package main

import (
	"github.com/extism/go-pdk"
	"github.com/sonr-io/crypto/mpc"
)

type SignRequest struct {
	Message []byte `json:"message"`
	Enclave []byte `json:"enclave"`
}

func main() {}

//
// //go:wasmexport new_enclave
// func newEnclave() int32 {
// 	e, err := mpc.NewEnclave()
// 	if err != nil {
// 		pdk.Log(pdk.LogError, err.Error())
// 		return 1
// 	}
// 	pdk.Log(pdk.LogInfo, "Enclave created")
// 	bz, err := e.Serialize()
// 	if err != nil {
// 		pdk.Log(pdk.LogError, err.Error())
// 		return 1
// 	}
// 	pdk.Log(pdk.LogInfo, "Enclave export successful")
// 	pdk.OutputJSON(bz)
// 	return 0
// }

//go:wasmexport sign_message
func signMessage() int32 {
	req := SignRequest{}
	err := pdk.InputJSON(req)
	if err != nil {
		pdk.Log(pdk.LogError, err.Error())
		return 1
	}
	e, err := mpc.ImportEnclave(mpc.WithEnclaveJSON(req.Enclave))
	if err != nil {
		pdk.Log(pdk.LogError, err.Error())
		return 1
	}
	sig, err := e.Sign(req.Message)
	if err != nil {
		pdk.Log(pdk.LogError, err.Error())
		return 1
	}
	pdk.Log(pdk.LogInfo, "Signature successful")
	pdk.OutputJSON(sig)
	return 0
}
