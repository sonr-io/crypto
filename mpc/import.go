package mpc

type ImportOptions struct {
	valKeyshare  AliceOut
	userKeyshare BobOut
	enclaveBytes []byte
}
