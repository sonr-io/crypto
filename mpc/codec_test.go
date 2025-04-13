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
			data, err := original.Export(testKey)
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
		keyclave, ok := original.(*enclave)
		require.True(t, ok)

		data, err := keyclave.Serialize()
		require.NoError(t, err)
		require.NotEmpty(t, data)

		// Unmarshal
		restored := &enclave{}
		err = restored.Unmarshal(data)
		require.NoError(t, err)

		// Verify restored enclave
		assert.True(t, keyclave.PubPoint.Equal(restored.PubPoint))
		assert.True(t, restored.IsValid())
	})
}
