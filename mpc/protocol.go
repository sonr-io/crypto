package mpc

import (
	"github.com/sonr-io/crypto/core/curves"
	"github.com/sonr-io/crypto/core/protocol"
	"github.com/sonr-io/crypto/tecdsa/dklsv1"
)

// NewEnclave generates a new MPC keyshare
func NewEnclave(nonce []byte) (Enclave, error) {
	curve := curves.K256()
	valKs := dklsv1.NewAliceDkg(curve, protocol.Version1)
	userKs := dklsv1.NewBobDkg(curve, protocol.Version1)
	aErr, bErr := RunProtocol(userKs, valKs)
	if err := checkIteratedErrors(aErr, bErr); err != nil {
		return nil, err
	}
	valRes, err := valKs.Result(protocol.Version1)
	if err != nil {
		return nil, err
	}
	userRes, err := userKs.Result(protocol.Version1)
	if err != nil {
		return nil, err
	}
	return createEnclave(valRes, userRes)
}

// ExecuteSigning runs the MPC signing protocol
func ExecuteSigning(signFuncVal SignFunc, signFuncUser SignFunc) ([]byte, error) {
	aErr, bErr := RunProtocol(signFuncVal, signFuncUser)
	if err := checkIteratedErrors(aErr, bErr); err != nil {
		return nil, err
	}
	out, err := signFuncUser.Result(protocol.Version1)
	if err != nil {
		return nil, err
	}
	s, err := dklsv1.DecodeSignature(out)
	if err != nil {
		return nil, err
	}
	sig, err := SerializeSignature(s)
	if err != nil {
		return nil, err
	}
	return sig, nil
}

// ExecuteRefresh runs the MPC refresh protocol
func ExecuteRefresh(refreshFuncVal RefreshFunc, refreshFuncUser RefreshFunc, nonce []byte) (Enclave, error) {
	aErr, bErr := RunProtocol(refreshFuncVal, refreshFuncUser)
	if err := checkIteratedErrors(aErr, bErr); err != nil {
		return nil, err
	}
	valRefreshResult, err := refreshFuncVal.Result(protocol.Version1)
	if err != nil {
		return nil, err
	}
	userRefreshResult, err := refreshFuncUser.Result(protocol.Version1)
	if err != nil {
		return nil, err
	}
	return createEnclave(valRefreshResult, userRefreshResult)
}

// For DKG bob starts first. For refresh and sign, Alice starts first.
func RunProtocol(firstParty protocol.Iterator, secondParty protocol.Iterator) (error, error) {
	var (
		message *protocol.Message
		aErr    error
		bErr    error
	)

	for aErr != protocol.ErrProtocolFinished || bErr != protocol.ErrProtocolFinished {
		// Crank each protocol forward one iteration
		message, bErr = firstParty.Next(message)
		if bErr != nil && bErr != protocol.ErrProtocolFinished {
			return nil, bErr
		}

		message, aErr = secondParty.Next(message)
		if aErr != nil && aErr != protocol.ErrProtocolFinished {
			return aErr, nil
		}
	}
	return aErr, bErr
}
