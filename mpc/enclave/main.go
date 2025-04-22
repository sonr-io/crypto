package main

import (
	"context"

	"github.com/extism/go-pdk"
	"github.com/sonr-io/crypto/mpc"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

func main() {
	// Choose the context to use for function calls.
	ctx := context.Background()

	// Create a new WebAssembly Runtime.
	r := wazero.NewRuntime(ctx)
	defer r.Close(ctx) // This closes everything this Runtime created.

	// implement functions such as panic.
	wasi_snapshot_preview1.MustInstantiate(ctx, r)
}

//go:wasmexport generate
func generate() int32 {
	e, err := mpc.NewEnclave()
	if err != nil {
		pdk.Log(pdk.LogError, err.Error())
		return 1
	}
	pdk.Log(pdk.LogInfo, "Generated enclave successfully")
	dat := e.GetData()
	pdk.OutputJSON(dat)
	return 0
}
