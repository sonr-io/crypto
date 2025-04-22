package main

import (
	"github.com/extism/go-pdk"
	"github.com/sonr-io/crypto/mpc"
)

func main() {}

//go:wasmexport generate
func generate() int32 {
	e, err := mpc.NewEnclave()
	if err != nil {
		pdk.Log(pdk.LogError, err.Error())
		return 1
	}
	pdk.Log(pdk.LogInfo, "Generated enclave successfully")
	pdk.OutputJSON(e.GetData())
	return 0
}
