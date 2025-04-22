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

type VerifyRequest struct {
	PubKey  []byte `json:"pub_key"`
	Message []byte `json:"message"`
	Sig     []byte `json:"sig"`
}

type VerifyResponse struct {
	Valid bool `json:"valid"`
	Error string
}

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
