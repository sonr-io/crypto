package main

import (
	"github.com/extism/go-pdk"
	"github.com/sonr-io/crypto/mpc"
)

type VerifyRequest struct {
	PubKey  []byte `json:"pub_key"`
	Message []byte `json:"message"`
	Sig     []byte `json:"sig"`
}

type VerifyResponse struct {
	Valid bool `json:"valid"`
	Error string
}

func main() {}

//go:wasmexport verify
func verify() int32 {
	req := VerifyRequest{}
	err := pdk.InputJSON(req)
	if err != nil {
		pdk.Log(pdk.LogError, err.Error())
		return 1
	}
	pdk.Log(pdk.LogInfo, "Deserialized request successfully")
	res := VerifyResponse{}
	valid, err := mpc.VerifyWithPubKey(req.PubKey, req.Message, req.Sig)
	if err != nil {
		res.Error = err.Error()
		res.Valid = false
	}
	pdk.Log(pdk.LogInfo, "Signature successful")
	res.Valid = valid
	pdk.OutputJSON(res)
	return 0
}
