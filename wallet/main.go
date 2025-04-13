package main

import (
	"fmt"

	"github.com/extism/go-pdk"
	"github.com/sonr-io/crypto/mpc"
)

//go:wasmexport new_enclave
func newEnclave() int32 {
	input := pdk.Input()
	e, err := mpc.GenEnclave(input)
	if err != nil {
		pdk.OutputString(err.Error())
		return 1
	}
	resp := fmt.Sprintf("Enclave: %v", e.IsValid())
	pdk.Log(pdk.LogInfo, resp)
	pdk.OutputString(resp)
	return 0
}
