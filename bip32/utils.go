package bip32

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math/big"

	"github.com/btcsuite/btcd/btcec/v2"
	"golang.org/x/crypto/ripemd160"
)

var (
	// Use btcsuite's implementation of secp256k1
	curve       = btcec.S256()
	curveParams = curve.Params()

	// Errors
	ErrInvalidPrivateKey = errors.New("invalid private key")
	ErrInvalidPublicKey  = errors.New("invalid public key")
)

// Base58 encoding alphabet
const base58Alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

//
// Hashes
//

func hashSha256(data []byte) ([]byte, error) {
	hasher := sha256.New()
	_, err := hasher.Write(data)
	if err != nil {
		return nil, err
	}
	return hasher.Sum(nil), nil
}

func hashDoubleSha256(data []byte) ([]byte, error) {
	hash1, err := hashSha256(data)
	if err != nil {
		return nil, err
	}

	hash2, err := hashSha256(hash1)
	if err != nil {
		return nil, err
	}
	return hash2, nil
}

func hashRipeMD160(data []byte) ([]byte, error) {
	hasher := ripemd160.New()
	_, err := io.WriteString(hasher, string(data))
	if err != nil {
		return nil, err
	}
	return hasher.Sum(nil), nil
}

func hash160(data []byte) ([]byte, error) {
	hash1, err := hashSha256(data)
	if err != nil {
		return nil, err
	}

	hash2, err := hashRipeMD160(hash1)
	if err != nil {
		return nil, err
	}

	return hash2, nil
}

//
// Encoding
//

func checksum(data []byte) ([]byte, error) {
	hash, err := hashDoubleSha256(data)
	if err != nil {
		return nil, err
	}

	return hash[:4], nil
}

func addChecksumToBytes(data []byte) ([]byte, error) {
	checksum, err := checksum(data)
	if err != nil {
		return nil, err
	}
	return append(data, checksum...), nil
}

// Base58Encode encodes a byte slice to a base58 string
func base58Encode(data []byte) string {
	// Convert the byte slice to a big integer
	x := new(big.Int).SetBytes(data)

	// Initialize the result
	result := ""

	// Base58 encoding
	base := big.NewInt(58)
	zero := big.NewInt(0)
	mod := new(big.Int)

	// Count leading zeros in the data
	zeroCount := 0
	for i := 0; i < len(data) && data[i] == 0; i++ {
		zeroCount++
	}

	// Convert to base58
	for x.Cmp(zero) > 0 {
		x.DivMod(x, base, mod)
		result = string(base58Alphabet[mod.Int64()]) + result
	}

	// Add leading '1's for each leading zero byte
	for i := 0; i < zeroCount; i++ {
		result = "1" + result
	}

	return result
}

// Base58Decode decodes a base58 string to a byte slice
func base58Decode(data string) ([]byte, error) {
	// Initialize the result
	result := big.NewInt(0)
	base := big.NewInt(58)

	// Count leading '1's
	zeroCount := 0
	for i := 0; i < len(data) && data[i] == '1'; i++ {
		zeroCount++
	}

	// Convert from base58
	for i := zeroCount; i < len(data); i++ {
		charIndex := bytes.IndexByte([]byte(base58Alphabet), data[i])
		if charIndex == -1 {
			return nil, errors.New("invalid base58 character")
		}

		result.Mul(result, base)
		result.Add(result, big.NewInt(int64(charIndex)))
	}

	// Convert big int to bytes
	resultBytes := result.Bytes()

	// Add leading zeros
	if zeroCount > 0 {
		zeros := make([]byte, zeroCount)
		resultBytes = append(zeros, resultBytes...)
	}

	return resultBytes, nil
}

// Keys
func publicKeyForPrivateKey(key []byte) []byte {
	// Create a private key from the bytes
	privKey, _ := btcec.PrivKeyFromBytes(key)

	// Get the public key
	pubKey := privKey.PubKey()

	// Return the compressed public key
	return pubKey.SerializeCompressed()
}

func addPublicKeys(key1 []byte, key2 []byte) []byte {
	pubKey1, err := btcec.ParsePubKey(key1)
	if err != nil {
		return nil
	}
	pubKey2, err := btcec.ParsePubKey(key2)
	if err != nil {
		return nil
	}

	// Get the curve from the btcec package
	curve := btcec.S256()

	// Get the big.Int coordinates
	x1, y1 := pubKey1.X(), pubKey1.Y()
	x2, y2 := pubKey2.X(), pubKey2.Y()

	// Add the points using the curve's Add method
	x3, y3 := curve.Add(x1, y1, x2, y2)

	// Convert the big.Int coordinates to FieldVal
	var fx3, fy3 btcec.FieldVal

	// Convert x3 to bytes and set in FieldVal
	x3Bytes := x3.Bytes()
	var x3Arr [32]byte
	// Pad if needed
	if len(x3Bytes) < 32 {
		copy(x3Arr[32-len(x3Bytes):], x3Bytes)
	} else {
		copy(x3Arr[:], x3Bytes)
	}
	fx3.SetBytes(&x3Arr)

	// Convert y3 to bytes and set in FieldVal
	y3Bytes := y3.Bytes()
	var y3Arr [32]byte
	// Pad if needed
	if len(y3Bytes) < 32 {
		copy(y3Arr[32-len(y3Bytes):], y3Bytes)
	} else {
		copy(y3Arr[:], y3Bytes)
	}
	fy3.SetBytes(&y3Arr)

	// Create a new public key using FieldVal values
	pubKey := btcec.NewPublicKey(&fx3, &fy3)
	return pubKey.SerializeCompressed()
}

func addPrivateKeys(key1 []byte, key2 []byte) []byte {
	var key1Int big.Int
	var key2Int big.Int
	key1Int.SetBytes(key1)
	key2Int.SetBytes(key2)

	key1Int.Add(&key1Int, &key2Int)
	key1Int.Mod(&key1Int, curve.N)

	b := key1Int.Bytes()
	if len(b) < 32 {
		extra := make([]byte, 32-len(b))
		b = append(extra, b...)
	}
	return b
}

func compressPublicKey(x, y *big.Int) []byte {
	// Convert big.Int to FieldVal
	var fx, fy btcec.FieldVal

	// Convert x to bytes and set in FieldVal
	xBytes := x.Bytes()
	var xArr [32]byte
	// Pad if needed
	if len(xBytes) < 32 {
		copy(xArr[32-len(xBytes):], xBytes)
	} else {
		copy(xArr[:], xBytes)
	}
	fx.SetBytes(&xArr)

	// Convert y to bytes and set in FieldVal
	yBytes := y.Bytes()
	var yArr [32]byte
	// Pad if needed
	if len(yBytes) < 32 {
		copy(yArr[32-len(yBytes):], yBytes)
	} else {
		copy(yArr[:], yBytes)
	}
	fy.SetBytes(&yArr)

	// Create a new public key using FieldVal values
	pubKey := btcec.NewPublicKey(&fx, &fy)

	// Serialize the public key in compressed format
	return pubKey.SerializeCompressed()
}

// expandPublicKey decompresses a public key
func expandPublicKey(key []byte) (*big.Int, *big.Int) {
	pubKey, err := btcec.ParsePubKey(key)
	if err != nil {
		return nil, nil
	}

	// The X() and Y() methods already return *big.Int values
	return pubKey.X(), pubKey.Y()
}

func validatePrivateKey(key []byte) error {
	if fmt.Sprintf("%x", key) == "0000000000000000000000000000000000000000000000000000000000000000" || // if the key is zero
		bytes.Compare(key, curve.N.Bytes()) >= 0 || // or is outside of the curve
		len(key) != 32 { // or is too short
		return ErrInvalidPrivateKey
	}

	return nil
}

func validateChildPublicKey(key []byte) error {
	_, err := btcec.ParsePubKey(key)
	if err != nil {
		return ErrInvalidPublicKey
	}
	return nil
}

// Numerical
func uint32Bytes(i uint32) []byte {
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, i)
	return bytes
}
