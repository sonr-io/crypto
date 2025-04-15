package mpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestKeyShareGeneration(t *testing.T) {
	t.Run("Generate Valid Enclave", func(t *testing.T) {
		// Generate enclave
		enclave, err := NewEnclave()
		require.NoError(t, err)
		require.NotNil(t, enclave)

		// Validate enclave contents
		assert.True(t, enclave.IsValid())
	})

	t.Run("Export and Import", func(t *testing.T) {
		// Generate original enclave
		original, err := NewEnclave()
		require.NoError(t, err)

		// Test key for encryption/decryption (32 bytes)
		testKey := []byte("test-key-12345678-test-key-123456")

		// Test Export/Import
		t.Run("Full Enclave", func(t *testing.T) {
			// Export enclave
			data, err := original.Encrypt(testKey)
			require.NoError(t, err)
			require.NotEmpty(t, data)

			// Create new empty enclave
			newEnclave, err := NewEnclave()
			require.NoError(t, err)

			// Verify the imported enclave works by signing
			testData := []byte("test message")
			sig, err := newEnclave.Sign(testData)
			require.NoError(t, err)
			valid, err := newEnclave.Verify(testData, sig)
			require.NoError(t, err)
			assert.True(t, valid)
		})
	})

	t.Run("Encrypt and Decrypt", func(t *testing.T) {
		// Generate enclave
		enclave, err := NewEnclave()
		require.NoError(t, err)
		
		// Get the enclave data
		keyclave, ok := enclave.(*EnclaveData)
		require.True(t, ok)
		
		// Create test data
		testKey := []byte("test-key-12345678-test-key-123456")
		
		// Test encryption
		encrypted, err := keyclave.Encrypt(testKey)
		require.NoError(t, err)
		require.NotEmpty(t, encrypted)
		
		// Test decryption
		decrypted, err := keyclave.Decrypt(testKey, encrypted)
		require.NoError(t, err)
		require.NotEmpty(t, decrypted)
		
		// Verify decrypted data can be used to restore the enclave
		restoredEnclave := &EnclaveData{}
		err = restoredEnclave.Deserialize(decrypted)
		require.NoError(t, err)
		
		// Ensure restored enclave is valid
		assert.True(t, restoredEnclave.IsValid())
		
		// Verify both enclaves have the same public key
		assert.True(t, keyclave.PubPoint.Equal(restoredEnclave.PubPoint))
		
		// Test decryption with wrong key (should fail)
		wrongKey := []byte("wrong-key-12345678-wrong-key-123456")
		_, err = keyclave.Decrypt(wrongKey, encrypted)
		assert.Error(t, err, "Decryption with wrong key should fail")
		
		// Test full round-trip encryption/decryption cycle
		t.Run("Round-trip Encryption/Decryption", func(t *testing.T) {
			// Generate original data to encrypt
			originalData, err := keyclave.Serialize()
			require.NoError(t, err)
			
			// Encrypt the data
			encrypted, err := keyclave.Encrypt(testKey)
			require.NoError(t, err)
			
			// Decrypt the data
			decrypted, err := keyclave.Decrypt(testKey, encrypted)
			require.NoError(t, err)
			
			// Verify the decrypted data matches the original
			assert.Equal(t, originalData, decrypted, "Decrypted data should match original")
		})
	})
}

func TestEnclaveOperations(t *testing.T) {
	t.Run("Signing and Verification", func(t *testing.T) {
		// Generate valid enclave
		enclave, err := NewEnclave()
		require.NoError(t, err)

		// Test signing
		testData := []byte("test message")
		signature, err := enclave.Sign(testData)
		require.NoError(t, err)
		require.NotNil(t, signature)

		// Verify the signature
		valid, err := enclave.Verify(testData, signature)
		require.NoError(t, err)
		assert.True(t, valid)

		// Test invalid data verification
		invalidData := []byte("wrong message")
		valid, err = enclave.Verify(invalidData, signature)
		require.NoError(t, err)
		assert.False(t, valid)
	})

	t.Run("Refresh Operation", func(t *testing.T) {
		enclave, err := NewEnclave()
		require.NoError(t, err)

		// Test refresh
		refreshedEnclave, err := enclave.Refresh()
		require.NoError(t, err)
		require.NotNil(t, refreshedEnclave)

		// Verify refreshed enclave is valid
		assert.True(t, refreshedEnclave.IsValid())
	})
}

func TestEnclaveSerialization(t *testing.T) {
	t.Run("Marshal and Unmarshal", func(t *testing.T) {
		// Generate original enclave
		original, err := NewEnclave()
		require.NoError(t, err)
		require.NotNil(t, original)

		// Marshal
		keyclave, ok := original.(*EnclaveData)
		require.True(t, ok)

		data, err := keyclave.Serialize()
		require.NoError(t, err)
		require.NotEmpty(t, data)

		// Unmarshal
		restored := &EnclaveData{}
		err = restored.Deserialize(data)
		require.NoError(t, err)

		// Verify restored enclave
		assert.True(t, keyclave.PubPoint.Equal(restored.PubPoint))
		assert.True(t, restored.IsValid())
	})
}
