package mpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sonr-io/crypto/core/curves"
	"github.com/sonr-io/crypto/core/protocol"
	"github.com/sonr-io/crypto/tecdsa/dklsv1"
)

func TestImportEnclave(t *testing.T) {
	// Mock message values for testing
	curve := curves.K256()
	valKs := dklsv1.NewAliceDkg(curve, protocol.Version1)
	userKs := dklsv1.NewBobDkg(curve, protocol.Version1)
	aErr, bErr := RunProtocol(userKs, valKs)
	if err := checkIteratedErrors(aErr, bErr); err != nil {
		require.NoError(t, err)
	}
	mockValShare, err := valKs.Result(protocol.Version1)
	if err != nil {
		require.NoError(t, err)
	}
	mockUserShare, err := userKs.Result(protocol.Version1)
	if err != nil {
	}

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
				WithEnclaveJSON(mockEnclaveBytes),
			},
			expectError: false,
		},
		{
			name: "With both shares and bytes (bytes take priority)",
			options: []ImportOption{
				WithInitialShares(mockValShare, mockUserShare),
				WithEnclaveJSON(mockEnclaveBytes),
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
				WithEnclaveJSON([]byte{}),
			},
			expectError: true,
			errorMsg:    "enclave bytes cannot be empty",
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
