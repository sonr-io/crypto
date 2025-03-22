module github.com/sonr-io/crypto

go 1.24.1

replace launchpad.net/gocheck => gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c

require (
	filippo.io/edwards25519 v1.1.0
	git.sr.ht/~sircmpwn/go-bare v0.0.0-20210406120253-ab86bc2846d9
	github.com/btcsuite/btcd/btcec/v2 v2.3.4
	github.com/bwesterb/go-ristretto v1.2.3
	github.com/consensys/gnark-crypto v0.16.0
	github.com/cosmos/btcutil v1.0.5
	github.com/cosmos/cosmos-sdk v0.50.12
	github.com/dustinxie/ecc v0.0.0-20210511000915-959544187564
	github.com/ecies/go/v2 v2.0.10
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/gtank/merlin v0.1.1
	github.com/holiman/uint256 v1.3.2 // indirect
	github.com/ipfs/go-cid v0.5.0
	github.com/libp2p/go-libp2p v0.41.0
	github.com/mr-tron/base58 v1.2.0
	github.com/multiformats/go-multibase v0.2.0
	github.com/multiformats/go-multihash v0.2.3
	github.com/multiformats/go-varint v0.0.7
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.10.0
	golang.org/x/crypto v0.36.0
	lukechampine.com/blake3 v1.4.0
)

require (
	github.com/bits-and-blooms/bitset v1.20.0 // indirect
	github.com/consensys/bavard v0.1.27 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.4.0 // indirect
	github.com/ethereum/go-ethereum v1.15.5
	github.com/klauspost/cpuid/v2 v2.2.10 // indirect
	github.com/mimoo/StrobeGo v0.0.0-20181016162300-f8f6d4d2b643 // indirect
	github.com/minio/sha256-simd v1.0.1 // indirect
	github.com/mmcloughlin/addchain v0.4.0 // indirect
	github.com/multiformats/go-base32 v0.1.0 // indirect
	github.com/multiformats/go-base36 v0.2.0 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/rogpeppe/go-internal v1.13.1 // indirect
	github.com/spaolacci/murmur3 v1.1.0 // indirect
	go.dedis.ch/kyber/v3 v3.1.0
	golang.org/x/sys v0.31.0 // indirect
	google.golang.org/protobuf v1.36.5 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	rsc.io/tmplfunc v0.0.3 // indirect
)

require (
	github.com/btcsuite/btcd v0.22.3
	github.com/btcsuite/btcutil v1.0.3-0.20201208143702-a53e38424cce
	github.com/okx/go-wallet-sdk/util v0.0.1
)
