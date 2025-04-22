package main

import (
	"github.com/extism/go-pdk"
	"github.com/sonr-io/crypto/mpc"
)

type GenerateResponse struct {
	PubKeyHex string           `json:"pub_key_hex"`
	Enclave   *mpc.EnclaveData `json:"enclave"`
}

func main() {
	generate()
}

//go:wasmexport generate
func generate() int32 {
	e, err := mpc.NewEnclave()
	if err != nil {
		pdk.SetError(err)
		return 1
	}
	pdk.Log(pdk.LogInfo, "Generated enclave successfully")
	resp := &GenerateResponse{
		PubKeyHex: e.PubKeyHex(),
		Enclave:   e.GetData(),
	}
	pdk.OutputJSON(resp)
	return 0
}
