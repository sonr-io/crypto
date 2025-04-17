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

	opts := Options{}
	for _, opt := range options {
		opts = opt(opts)
	}
	return opts.Apply()
}

// Options is a struct that holds the import options
type Options struct {
	valKeyshare   Message
	userKeyshare  Message
	enclaveBytes  []byte
	initialShares bool
	curve         CurveName
}

// ImportOption is a function that modifies the import options
type ImportOption func(Options) Options

// WithInitialShares creates an option to import an enclave from validator and user keyshares.
func WithInitialShares(valKeyshare Message, userKeyshare Message, curve CurveName) ImportOption {
	return func(opts Options) Options {
		opts.valKeyshare = valKeyshare
		opts.userKeyshare = userKeyshare
		opts.initialShares = true
		opts.curve = curve
		return opts
	}
}

// WithEnclaveJSON creates an option to import an enclave from serialized bytes.
func WithEnclaveJSON(enclaveBytes []byte) ImportOption {
	return func(opts Options) Options {
		opts.enclaveBytes = enclaveBytes
		opts.initialShares = false
		return opts
	}
}

func (opts Options) Apply() (Enclave, error) {
	// First try to restore from enclave bytes if provided
	if !opts.initialShares {
		if len(opts.enclaveBytes) == 0 {
			return nil, errors.New("enclave bytes cannot be empty")
		}
		return RestoreEnclave(opts.enclaveBytes)
	} else {
		// Then try to build from keyshares
		if opts.valKeyshare == nil {
			return nil, errors.New("validator share cannot be nil")
		}
		if opts.userKeyshare == nil {
			return nil, errors.New("user share cannot be nil")
		}
		return BuildEnclave(opts.valKeyshare, opts.userKeyshare, opts)
	}
}

// BuildEnclave creates a new enclave from validator and user keyshares.
func BuildEnclave(valShare, userShare Message, options Options) (Enclave, error) {
	if valShare == nil {
		return nil, errors.New("validator share cannot be nil")
	}
	if userShare == nil {
		return nil, errors.New("user share cannot be nil")
	}

	pubPoint, err := GetAlicePublicPoint(valShare)
	if err != nil {
		return nil, fmt.Errorf("failed to get public point: %w", err)
	}
	return &EnclaveData{
		PubPoint:  pubPoint,
		PubBytes:  pubPoint.ToAffineCompressed(),
		ValShare:  valShare,
		UserShare: userShare,
		Nonce:     randNonce(),
		Curve:     options.curve,
	}, nil
}

// RestoreEnclave deserializes an enclave from its binary representation.
func RestoreEnclave(data []byte) (Enclave, error) {
	if len(data) == 0 {
		return nil, errors.New("enclave bytes cannot be empty")
	}

	keyclave := &EnclaveData{}
	err := keyclave.Deserialize(data)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal enclave: %w", err)
	}

	return keyclave, nil
}
