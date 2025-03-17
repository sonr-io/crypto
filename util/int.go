package util

import (
	"math/big"
	"strconv"
)

func ConvertToUint64(v string) uint64 {
	i, _ := strconv.ParseUint(v, 10, 64)
	return i
}

func ConvertToBigInt(v string) *big.Int {
	b := new(big.Int)
	b.SetString(v, 10)
	return b
}

func ConvertToUint32(v string) uint32 {
	i, _ := strconv.Atoi(v)
	return uint32(i)
}
