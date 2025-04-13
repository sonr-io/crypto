package main

import (
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
	e, err := mpc.ImportEnclave(input)
	if err != nil {
		pdk.Log(pdk.LogError, err.Error())
		return 1
	}
	msg := pdk.GetString(0)
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
	e, err := mpc.ImportEnclave(input)
	if err != nil {
		pdk.Log(pdk.LogError, err.Error())
		return 1
	}
	sig := pdk.GetString(0)
	msg := pdk.GetString(1)
	err = e.Verify(sig, msg)
	if err != nil {
		pdk.Log(pdk.LogError, err.Error())
		return 1
	}
	pdk.Log(pdk.LogInfo, "Signature verified")
	return 0
}

//go:wasmexport refresh_enclave
func refreshEnclave() int32 {
	input := pdk.Input()
	e, err := mpc.ImportEnclave(input)
	if err != nil {
		pdk.Log(pdk.LogError, err.Error())
		return 1
	}
	err = e.Refresh()
	if err != nil {
		pdk.Log(pdk.LogError, err.Error())
		return 1
	}
	pdk.Log(pdk.LogInfo, "Enclave refreshed")
	return 0
}
