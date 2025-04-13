package bech32

import (
	"errors"
	"strings"
)

// Charset is the Bech32 character set for encoding
const charset = "qpzry9x8gf2tvdw0s3jn54khce6mua7l"

// CharsetRev is a mapping of the charset for decoding
var charsetRev = map[byte]int{
	'q': 0, 'p': 1, 'z': 2, 'r': 3, 'y': 4, '9': 5, 'x': 6, '8': 7,
	'g': 8, 'f': 9, '2': 10, 't': 11, 'v': 12, 'd': 13, 'w': 14, '0': 15,
	's': 16, '3': 17, 'j': 18, 'n': 19, '5': 20, '4': 21, 'k': 22, 'h': 23,
	'c': 24, 'e': 25, '6': 26, 'm': 27, 'u': 28, 'a': 29, '7': 30, 'l': 31,
}

// Errors
var (
	ErrInvalidLength       = errors.New("invalid bech32 string length")
	ErrInvalidCharacter    = errors.New("invalid character in bech32 string")
	ErrInvalidChecksum     = errors.New("invalid bech32 checksum")
	ErrInvalidSeparator    = errors.New("invalid separator index")
	ErrInvalidPrefix       = errors.New("invalid bech32 prefix")
	ErrInvalidHRPCharacter = errors.New("invalid character in hrp")
)

// Encode encodes a byte array to a bech32 string with the given prefix
func Encode(prefix string, data []byte) (string, error) {
	// Check if the prefix is valid
	if err := validateHRP(prefix); err != nil {
		return "", err
	}

	// Convert data bytes to 5-bit values
	converted := convertBits(data, 8, 5, true)

	// Create checksum
	checksum := createChecksum(prefix, converted)

	// Combine data and checksum
	combined := append(converted, checksum...)

	// Encode to bech32 characters
	result := prefix + "1"
	for _, b := range combined {
		result += string(charset[b])
	}

	return result, nil
}

// Decode decodes a bech32 string and returns the prefix and data
func Decode(bech string) (string, []byte, error) {
	// Check length
	if len(bech) < 8 || len(bech) > 90 {
		return "", nil, ErrInvalidLength
	}

	// Convert to lowercase
	bech = strings.ToLower(bech)

	// Find separator
	pos := strings.LastIndex(bech, "1")
	if pos < 1 || pos+7 > len(bech) {
		return "", nil, ErrInvalidSeparator
	}

	// Extract prefix and data part
	prefix := bech[:pos]
	data := bech[pos+1:]

	// Check prefix
	if err := validateHRP(prefix); err != nil {
		return "", nil, err
	}

	// Decode data
	decoded := make([]byte, 0, len(data))
	for i := 0; i < len(data); i++ {
		value, ok := charsetRev[data[i]]
		if !ok {
			return "", nil, ErrInvalidCharacter
		}
		decoded = append(decoded, byte(value))
	}

	// Verify checksum
	if !verifyChecksum(prefix, decoded) {
		return "", nil, ErrInvalidChecksum
	}

	// Extract data without checksum (last 6 bytes are checksum)
	dataWithoutChecksum := decoded[:len(decoded)-6]

	// Convert from 5-bit to 8-bit
	result := convertBits(dataWithoutChecksum, 5, 8, false)

	return prefix, result, nil
}

// validateHRP checks if the prefix is valid
func validateHRP(hrp string) error {
	if len(hrp) < 1 || len(hrp) > 83 {
		return ErrInvalidPrefix
	}

	for _, c := range hrp {
		if c < 33 || c > 126 {
			return ErrInvalidHRPCharacter
		}
	}

	return nil
}

// convertBits converts from one bit width to another
func convertBits(data []byte, fromBits, toBits uint8, pad bool) []byte {
	acc := uint32(0)
	bits := uint8(0)
	ret := make([]byte, 0, len(data)*int(fromBits)/int(toBits)+1)
	maxv := uint32((1 << toBits) - 1)

	for _, value := range data {
		acc = (acc << fromBits) | uint32(value)
		bits += fromBits

		for bits >= toBits {
			bits -= toBits
			ret = append(ret, byte((acc>>bits)&maxv))
		}
	}

	if pad && bits > 0 {
		ret = append(ret, byte((acc<<(toBits-bits))&maxv))
	}

	return ret
}

// polymod calculates the checksum
func polymod(values []byte) uint32 {
	chk := uint32(1)
	for _, v := range values {
		top := chk >> 25
		chk = (chk&0x1ffffff)<<5 ^ uint32(v)
		for i := 0; i < 5; i++ {
			if ((top >> i) & 1) == 1 {
				chk ^= 0x3b6a57b2 << uint(i*5)
			}
		}
	}
	return chk
}

// hrpExpand expands the human-readable part for checksum calculation
func hrpExpand(hrp string) []byte {
	result := make([]byte, len(hrp)*2+1)
	for i, c := range hrp {
		result[i] = byte(c >> 5)
		result[i+len(hrp)+1] = byte(c & 31)
	}
	result[len(hrp)] = 0
	return result
}

// createChecksum creates a checksum for the given prefix and data
func createChecksum(hrp string, data []byte) []byte {
	values := append(hrpExpand(hrp), data...)
	values = append(values, []byte{0, 0, 0, 0, 0, 0}...)
	mod := polymod(values) ^ 1
	ret := make([]byte, 6)
	for i := 0; i < 6; i++ {
		ret[i] = byte((mod >> uint(5*(5-i))) & 31)
	}
	return ret
}

// verifyChecksum verifies the checksum for the given prefix and data
func verifyChecksum(hrp string, data []byte) bool {
	return polymod(append(hrpExpand(hrp), data...)) == 1
}
