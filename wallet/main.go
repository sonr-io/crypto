package main

import (
	"fmt"

	"github.com/extism/go-pdk"
	"github.com/sonr-io/crypto/mpc"
)

//go:wasmexport new_enclave
func newEnclave() int32 {
	input := pdk.Input()
	e, err := mpc.NewEnclave()
	if err != nil {
		pdk.Log(pdk.LogError, err.Error())
		return 1
	}
	pdk.Log(pdk.LogInfo, "Enclave created")
	bz, err := e.Export(input)
	if err != nil {
		pdk.Log(pdk.LogError, err.Error())
		return 1
	}
	pdk.Log(pdk.LogInfo, "Enclave export successful")
	pdk.OutputJSON(bz)
	return 0
}

//go:wasmexport sign_message
func signMessage() int32 {
	input := pdk.Input()
	e, err := mpc.ImportEnclave(mpc.WithEnclaveBytes(input))
	if err != nil {
		pdk.Log(pdk.LogError, err.Error())
		return 1
	}
	sig, err := e.Sign(msg)
	if err != nil {
		pdk.Log(pdk.LogError, err.Error())
		return 1
	}
	pdk.Log(pdk.LogInfo, "Signature successful")
	pdk.OutputJSON(sig)
	return 0
}

//go:wasmexport verify_signature
func verifySignature() int32 {
	input := pdk.Input()
	e, err := mpc.ImportEnclave(mpc.WithEnclaveBytes(input))
	if err != nil {
		pdk.Log(pdk.LogError, err.Error())
		return 1
	}
	ok, err := e.Verify(sig, msg)
	if err != nil {
		pdk.Log(pdk.LogError, err.Error())
		return 1
	}
	pdk.Log(pdk.LogInfo, fmt.Sprintf("Signature verified: %v", ok))
	return 0
}
