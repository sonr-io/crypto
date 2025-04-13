package mpc

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	
	"github.com/sonr-io/crypto/core/protocol"
)

func TestImportEnclave(t *testing.T) {
	// Mock message values for testing
	mockValShare := &protocol.Message{Protocol: "test-val"}
	mockUserShare := &protocol.Message{Protocol: "test-user"}
	
	// Create a mock enclave for testing
	mockEnclave, err := buildEnclave(mockValShare, mockUserShare)
	require.NoError(t, err)
	
	// Serialize the enclave
	mockEnclaveBytes, err := mockEnclave.Serialize()
	require.NoError(t, err)
	
	tests := []struct {
		name        string
		options     []ImportOption
		expectError bool
		errorMsg    string
	}{
		{
			name:        "No options",
			options:     []ImportOption{},
			expectError: true,
			errorMsg:    "no import options provided",
		},
		{
			name: "With valid shares",
			options: []ImportOption{
				WithInitialShares(mockValShare, mockUserShare),
			},
			expectError: false,
		},
		{
			name: "With valid enclave bytes",
			options: []ImportOption{
				WithEnclaveBytes(mockEnclaveBytes),
			},
			expectError: false,
		},
		{
			name: "With both shares and bytes (bytes take priority)",
			options: []ImportOption{
				WithInitialShares(mockValShare, mockUserShare),
				WithEnclaveBytes(mockEnclaveBytes),
			},
			expectError: false,
		},
		{
			name: "With nil val share",
			options: []ImportOption{
				WithInitialShares(nil, mockUserShare),
			},
			expectError: true,
			errorMsg:    "validator share cannot be nil",
		},
		{
			name: "With nil user share",
			options: []ImportOption{
				WithInitialShares(mockValShare, nil),
			},
			expectError: true,
			errorMsg:    "user share cannot be nil",
		},
		{
			name: "With empty enclave bytes",
			options: []ImportOption{
				WithEnclaveBytes([]byte{}),
			},
			expectError: true,
			errorMsg:    "enclave bytes cannot be empty",
		},
		{
			name: "With invalid enclave bytes",
			options: []ImportOption{
				WithEnclaveBytes([]byte("invalid")),
			},
			expectError: true,
		},
	}
	
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			enclave, err := ImportEnclave(tc.options...)
			
			if tc.expectError {
				assert.Error(t, err)
				if tc.errorMsg != "" {
					assert.Contains(t, err.Error(), tc.errorMsg)
				}
				assert.Nil(t, enclave)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, enclave)
			}
		})
	}
}

// TestImportOptions tests the individual option functions
func TestImportOptions(t *testing.T) {
	// Create test values
	testValShare := &protocol.Message{Protocol: "test-val"}
	testUserShare := &protocol.Message{Protocol: "test-user"}
	testBytes := []byte("test bytes")
	
	// Test WithInitialShares option
	t.Run("WithInitialShares", func(t *testing.T) {
		opt := WithInitialShares(testValShare, testUserShare)
		opts := importOptions{}
		result := opt(opts)
		
		assert.Equal(t, testValShare, result.valKeyshare)
		assert.Equal(t, testUserShare, result.userKeyshare)
		assert.Nil(t, result.enclaveBytes)
	})
	
	// Test WithEnclaveBytes option
	t.Run("WithEnclaveBytes", func(t *testing.T) {
		opt := WithEnclaveBytes(testBytes)
		opts := importOptions{}
		result := opt(opts)
		
		assert.Equal(t, testBytes, result.enclaveBytes)
		assert.Nil(t, result.valKeyshare)
		assert.Nil(t, result.userKeyshare)
	})
	
	// Test option precedence (bytes takes priority over shares)
	t.Run("Option precedence", func(t *testing.T) {
		// Create a mock enclave
		mockEnclave, err := buildEnclave(testValShare, testUserShare)
		require.NoError(t, err)
		
		// Apply options in different orders
		enclave1, err := ImportEnclave(
			WithInitialShares(testValShare, testUserShare),
			WithEnclaveBytes(testBytes),
		)
		require.NoError(t, err)
		
		enclave2, err := ImportEnclave(
			WithEnclaveBytes(testBytes),
			WithInitialShares(testValShare, testUserShare),
		)
		require.NoError(t, err)
		
		// Both should prioritize bytes over shares
		_, ok1 := enclave1.(*enclave)
		_, ok2 := enclave2.(*enclave)
		assert.True(t, ok1)
		assert.True(t, ok2)
	})
}
