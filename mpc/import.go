package mpc

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/json"
	"errors"
)

const gcmNonceSize = 12 // Standard GCM nonce size

// Encrypt encrypts data using AES-GCM with the correct nonce size
func encrypt(plaintext, key []byte) ([]byte, error) {
	// Hash the key to ensure it's the right length for AES
	keyHash := sha256.Sum256(key)
	
	// Create a new AES cipher with the provided key
	block, err := aes.NewCipher(keyHash[:])
	if err != nil {
		return nil, err
	}
	
	// Create a new GCM cipher
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	
	// Use only the first gcmNonceSize bytes of the package nonce
	if len(nonce) < gcmNonceSize {
		return nil, errors.New("nonce length must be at least 12 bytes")
	}
	gcmNonce := nonce[:gcmNonceSize]
	
	// Encrypt the data
	return gcm.Seal(nil, gcmNonce, plaintext, nil), nil
}

// Decrypt decrypts data using AES-GCM with the correct nonce size
func decrypt(ciphertext, key []byte) ([]byte, error) {
	// Hash the key to ensure it's the right length for AES
	keyHash := sha256.Sum256(key)
	
	// Create a new AES cipher with the provided key
	block, err := aes.NewCipher(keyHash[:])
	if err != nil {
		return nil, err
	}
	
	// Create a new GCM cipher
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	
	// Use only the first gcmNonceSize bytes of the package nonce
	if len(nonce) < gcmNonceSize {
		return nil, errors.New("nonce length must be at least 12 bytes")
	}
	gcmNonce := nonce[:gcmNonceSize]
	
	// Decrypt the data
	return gcm.Open(nil, gcmNonce, ciphertext, nil)
}

// Export encrypts the enclave data using the provided key
func (k *keyEnclave) Export(key []byte) ([]byte, error) {
	// Serialize the enclave
	data, err := k.Serialize()
	if err != nil {
		return nil, err
	}
	
	// Encrypt the data
	return encrypt(data, key)
}

// Import decrypts the enclave data using the provided key
func (k *keyEnclave) Import(encryptedData []byte, key []byte) error {
	// Decrypt the data
	data, err := decrypt(encryptedData, key)
	if err != nil {
		return err
	}
	
	// Unmarshal the data
	return k.Unmarshal(data)
}

// Serialize serializes the keyEnclave to JSON
func (k *keyEnclave) Serialize() ([]byte, error) {
	return json.Marshal(k)
}

type ImportOptions struct {
	valKeyshare  AliceOut
	userKeyshare BobOut
	enclaveBytes []byte
}
