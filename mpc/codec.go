package mpc

import (
	"crypto/rand"

	"github.com/sonr-io/crypto/core/curves"
	"github.com/sonr-io/crypto/core/protocol"
	"github.com/sonr-io/crypto/tecdsa/dklsv1/dkg"
)

// ╭───────────────────────────────────────────────────────────╮
// │                    Exported Generics                      │
// ╰───────────────────────────────────────────────────────────╯

type (
	AliceOut    *dkg.AliceOutput
	BobOut      *dkg.BobOutput
	Point       curves.Point
	Role        string                         // Role is the type for the role
	Message     *protocol.Message              // Message is the protocol.Message that is used for MPC
	Signature   *curves.EcdsaSignature         // Signature is the type for the signature
	RefreshFunc interface{ protocol.Iterator } // RefreshFunc is the type for the refresh function
	SignFunc    interface{ protocol.Iterator } // SignFunc is the type for the sign function
)

const (
	RoleVal  = "validator"
	RoleUser = "user"
)

func randNonce() []byte {
	nonce := make([]byte, 12)
	rand.Read(nonce)
	return nonce
}

// Enclave defines the interface for key management operations
type Enclave interface {
	// DID() keys.DID                        // DID returns the DID of the keyEnclave
	Export(key []byte) ([]byte, error) // Export returns encrypted enclave data
	IsValid() bool                     // IsValid returns true if the keyEnclave is valid
	// PubKey() keys.PubKey                          // PubKey returns the public key of the keyEnclave
	Refresh() (Enclave, error)                    // Refresh returns a new keyEnclave
	Serialize() ([]byte, error)                   // Serialize returns the serialized keyEnclave
	Sign(data []byte) ([]byte, error)             // Sign returns the signature of the data
	Verify(data []byte, sig []byte) (bool, error) // Verify returns true if the signature is valid
}
