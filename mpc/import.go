package mpc

import "errors"

func ImportEnclave(options ...ImportOption) (Enclave, error) {
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

func WithInitialShares(valKeyshare Message, userKeyshare Message) ImportOption {
	return func(opts importOptions) importOptions {
		opts.valKeyshare = valKeyshare
		opts.userKeyshare = userKeyshare
		return opts
	}
}

func WithEnclaveBytes(enclaveBytes []byte) ImportOption {
	return func(opts importOptions) importOptions {
		opts.enclaveBytes = enclaveBytes
		return opts
	}
}

func (opts importOptions) Apply() (Enclave, error) {
	if len(opts.enclaveBytes) > 0 {
		return restoreEnclave(opts.enclaveBytes)
	}
	if opts.valKeyshare != nil && opts.userKeyshare != nil {
		return buildEnclave(opts.valKeyshare, opts.userKeyshare)
	}
	return nil, errors.New("invalid shares")
}

func buildEnclave(valShare, userShare Message) (Enclave, error) {
	pubPoint, err := getAlicePubPoint(valShare)
	if err != nil {
		return nil, err
	}
	return &enclave{
		PubPoint:  pubPoint,
		ValShare:  valShare,
		UserShare: userShare,
		nonce:     randNonce(),
	}, nil
}

func restoreEnclave(data []byte) (Enclave, error) {
	keyclave := &enclave{}
	err := keyclave.Unmarshal(data)
	if err != nil {
		return nil, err
	}
	return keyclave, nil
}
