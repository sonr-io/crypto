package random

import "crypto/rand"

const (
	nonceSize = 32
)

// GenerateNonce generates a cryptographically secure random nonce of the specified length
func GenerateNonce() []byte {
	nonce := make([]byte, nonceSize)
	_, err := rand.Read(nonce)
	if err != nil {
		panic(err)
	}
	return nonce
}
