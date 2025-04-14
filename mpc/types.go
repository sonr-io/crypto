package mpc

// SignRequest is the request struct for signing
type SignRequest struct {
	Data []byte `json:"data"`
}

// SignResponse is the response struct for signing
type SignResponse struct {
	Signature []byte `json:"signature"`
	Data      []byte `json:"data"`
	Err       error  `json:"err"`
}

// VerifyRequest is the request struct for verifying
type VerifyRequest struct {
	Signature []byte `json:"signature"`
	Data      []byte `json:"data"`
}

// VerifyResponse is the response struct for verifying
type VerifyResponse struct {
	Valid bool  `json:"valid"`
	Err   error `json:"err"`
}
