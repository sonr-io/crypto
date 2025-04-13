package mpc

import (
	"errors"
	"fmt"
)

// ImportEnclave creates an Enclave instance from various import options.
// It prioritizes enclave bytes over keyshares if both are provided.
func ImportEnclave(options ...ImportOption) (Enclave, error) {
	if len(options) == 0 {
		return nil, errors.New("no import options provided")
	}
	
	opts := importOptions{}
	for _, opt := range options {
		opts = opt(opts)
	}
	return opts.Apply()
}

type importOptions struct {
	valKeyshare  Message
	userKeyshare Message
	enclaveBytes []byte
}

type ImportOption func(importOptions) importOptions

// WithInitialShares creates an option to import an enclave from validator and user keyshares.
func WithInitialShares(valKeyshare Message, userKeyshare Message) ImportOption {
	return func(opts importOptions) importOptions {
		opts.valKeyshare = valKeyshare
		opts.userKeyshare = userKeyshare
		return opts
	}
}

// WithEnclaveBytes creates an option to import an enclave from serialized bytes.
func WithEnclaveBytes(enclaveBytes []byte) ImportOption {
	return func(opts importOptions) importOptions {
		opts.enclaveBytes = enclaveBytes
		return opts
	}
}

func (opts importOptions) Apply() (Enclave, error) {
	// First try to restore from enclave bytes if provided
	if len(opts.enclaveBytes) > 0 {
		if len(opts.enclaveBytes) < 10 { // Minimum size for valid serialized enclave
			return nil, errors.New("enclave bytes cannot be empty or too small")
		}
		return restoreEnclave(opts.enclaveBytes)
	}
	
	// Then try to build from keyshares
	if opts.valKeyshare != nil && opts.userKeyshare != nil {
		return buildEnclave(opts.valKeyshare, opts.userKeyshare)
	}
	
	// Report specific errors for missing components
	if opts.valKeyshare == nil && opts.userKeyshare != nil {
		return nil, errors.New("validator share cannot be nil")
	}
	if opts.valKeyshare != nil && opts.userKeyshare == nil {
		return nil, errors.New("user share cannot be nil")
	}
	
	return nil, errors.New("no valid import options provided")
}

// buildEnclave creates a new enclave from validator and user keyshares.
func buildEnclave(valShare, userShare Message) (Enclave, error) {
	if valShare == nil {
		return nil, errors.New("validator share cannot be nil")
	}
	if userShare == nil {
		return nil, errors.New("user share cannot be nil")
	}
	
	pubPoint, err := getAlicePubPoint(valShare)
	if err != nil {
		return nil, fmt.Errorf("failed to get public point: %w", err)
	}
	
	return &enclave{
		PubPoint:  pubPoint,
		ValShare:  valShare,
		UserShare: userShare,
		nonce:     randNonce(),
	}, nil
}

// restoreEnclave deserializes an enclave from its binary representation.
func restoreEnclave(data []byte) (Enclave, error) {
	if len(data) == 0 {
		return nil, errors.New("enclave bytes cannot be empty")
	}
	
	keyclave := &enclave{}
	err := keyclave.Unmarshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal enclave: %w", err)
	}
	
	return keyclave, nil
}
