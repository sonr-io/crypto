package mpc

type SignRequest struct {
	Data []byte `json:"data"`
}

type SignResponse struct {
	Signature []byte `json:"signature"`
	Data      []byte `json:"data"`
	Err       error  `json:"err"`
}

type VerifyRequest struct {
	Signature []byte `json:"signature"`
	Data      []byte `json:"data"`
}

type VerifyResponse struct {
	Valid bool  `json:"valid"`
	Err   error `json:"err"`
}
