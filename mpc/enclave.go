package mpc

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"

	"github.com/sonr-io/crypto/core/curves"
	"golang.org/x/crypto/sha3"
)

// enclave implements the Enclave interface
type enclave struct {
	PubPoint  curves.Point `json:"-"`
	PubBytes  []byte       `json:"pub_key"`
	ValShare  Message      `json:"val_share"`
	UserShare Message      `json:"user_share"`
	Nonce     []byte       `json:"nonce"`
}

// Export returns encrypted enclave data
func (k *enclave) Export(key []byte) ([]byte, error) {
	data, err := k.Serialize()
	if err != nil {
		return nil, fmt.Errorf("failed to serialize enclave: %w", err)
	}

	hashedKey := hashKey(key)
	block, err := aes.NewCipher(hashedKey)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	return aesgcm.Seal(nil, k.Nonce, data, nil), nil
}

// IsValid returns true if the keyEnclave is valid
func (k *enclave) IsValid() bool {
	return k.PubPoint != nil && k.ValShare != nil && k.UserShare != nil
}

// Refresh returns a new keyEnclave
func (k *enclave) Refresh() (Enclave, error) {
	refreshFuncVal, err := valRefreshFunc(k)
	if err != nil {
		return nil, err
	}
	refreshFuncUser, err := userRefreshFunc(k)
	if err != nil {
		return nil, err
	}
	return ExecuteRefresh(refreshFuncVal, refreshFuncUser, k.Nonce)
}

// Sign returns the signature of the data
func (k *enclave) Sign(data []byte) ([]byte, error) {
	userSign, err := userSignFunc(k, data)
	if err != nil {
		return nil, err
	}
	valSign, err := valSignFunc(k, data)
	if err != nil {
		return nil, err
	}
	return ExecuteSigning(valSign, userSign)
}

// Verify returns true if the signature is valid
func (k *enclave) Verify(data []byte, sig []byte) (bool, error) {
	edSig, err := deserializeSignature(sig)
	if err != nil {
		return false, err
	}
	ePub, err := getEcdsaPoint(k.PubPoint.ToAffineUncompressed())
	if err != nil {
		return false, err
	}
	pk := &ecdsa.PublicKey{
		Curve: ePub.Curve,
		X:     ePub.X,
		Y:     ePub.Y,
	}

	// Hash the message using SHA3-256
	hash := sha3.New256()
	hash.Write(data)
	digest := hash.Sum(nil)

	return ecdsa.Verify(pk, digest, edSig.R, edSig.S), nil
}

// Marshal returns the JSON encoding of keyEnclave
func (k *enclave) Serialize() ([]byte, error) {
	// Store compressed public point bytes before marshaling
	k.PubBytes = k.PubPoint.ToAffineCompressed()
	return json.Marshal(k)
}

// Deserialize parses the JSON-encoded data and stores the result
func (k *enclave) Deserialize(data []byte) error {
	if err := json.Unmarshal(data, k); err != nil {
		return err
	}
	// Reconstruct Point from bytes
	curve := curves.K256()
	point, err := curve.NewIdentityPoint().FromAffineCompressed(k.PubBytes)
	if err != nil {
		return err
	}
	k.PubPoint = point
	return nil
}
