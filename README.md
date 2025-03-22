# Crypto Library for Sonr

Sonr's advanced cryptography library, forked from Coinbase's Kryptology

## Quickstart

Use the latest version of this library:

```
go get github.com/sonr-io/crypto
```

Pin a specific version of this library:

```
go get github.com/sonr-io/crypto@v0.0.2
```

## Documentation

Documentation can be found at the Sonr developer portal.

To access the documentation of the local version, run `godoc -http=:6060` and open
the following URL in your browser:

http://localhost:6060/pkg/github.com/sonr-io/crypto/

## Developer Setup

**Prerequisites**: `golang 1.22+`, `make`

```
# Set up Go environment for private repositories
export GOPRIVATE=git.sonr.io
export GONOPROXY=git.sonr.io
export GONOSUMDB=git.sonr.io

# Clone and build
git clone https://github.com/sonr-io/crypto.git && cd crypto && make
```

## Components

The following is the list of primitives and protocols that are implemented in this repository.

### Curves

The curve abstraction code can be found at `pkg/core/curves/curve.go`

The curves that implement this abstraction are as follows:

- [BLS12377](pkg/core/curves/bls12377_curve.go)
- [BLS12381](pkg/core/curves/bls12381_curve.go)
- [Ed25519](pkg/core/curves/ed25519_curve.go)
- [Secp256k1](pkg/core/curves/k256_curve.go)
- [P256](pkg/core/curves/p256_curve.go)
- [Pallas](pkg/core/curves/pallas_curve.go)

### Protocols

The generic protocol interface `pkg/core/protocol/protocol.go`.
This abstraction is currently only used in DKLs18 implementation.

- [Cryptographic Accumulators](pkg/accumulator)
- [Bulletproof](pkg/bulletproof)
- Oblivious Transfer
  - [Verifiable Simplest OT](pkg/ot/base/simplest)
  - [KOS OT Extension](pkg/ot/extension/kos)
- Threshold ECDSA Signature
  - [DKLs18 - DKG and Signing](pkg/tecdsa/dkls/v1)
  - [GG20 - DKG](pkg/dkg/gennaro)
  - [GG20 - Signing](pkg/tecdsa/gg20)
- Threshold Schnorr Signature
  - [FROST threshold signature - DKG](pkg/dkg/frost)
  - [FROST threshold signature - Signing](pkg/ted25519/frost)
- [Paillier encryption system](pkg/paillier)
- Secret Sharing Schemes
  - [Shamir's secret sharing scheme](pkg/sharing/shamir.go)
  - [Pedersen](pkg/sharing/pedersen.go)
  - [Feldman](pkg/sharing/feldman.go)
- [Verifiable encryption](pkg/verenc)
- [ZKP Schnorr](pkg/zkp/schnorr)

### Sonr Enhancements

Sonr has actively maintained and enhanced this library with the following improvements:

- Updated to support Go 1.22+
- Fixed dependency issues and modernized package imports
- Enhanced performance for cryptographic operations
- Added additional curve support for blockchain compatibility
- Improved documentation and examples
- Integration with Sonr's wallet infrastructure

## Integration with Sonr Wallet SDK

This cryptography library serves as the foundation for the [Sonr Go Wallet SDK](https://git.sonr.io/pkg/coins), providing the cryptographic primitives needed for secure wallet operations across multiple blockchains.

## Contributing

Contributions to the Sonr cryptography library are welcome. Please follow these guidelines:

- [Versioning](https://blog.golang.org/publishing-go-modules): `vMajor.Minor.Patch`
  - Major revision indicates breaking API change or significant new features
  - Minor revision indicates no API breaking changes and may include significant new features or documentation
  - Patch indicates no API breaking changes and may include only fixes

## References

- [[GG20] _One Round Threshold ECDSA with Identifiable Abort._](https://eprint.iacr.org/2020/540.pdf)
- [[EL20] _Eliding RSA Group Membership Checks._](docs/rsa-membership.pdf)
- [[P99] _Public-Key Cryptosystems Based on Composite Degree Residuosity Classes._](http://citeseerx.ist.psu.edu/viewdoc/download?doi=10.1.1.112.4035&rep=rep1&type=pdf)

## License

This library is licensed under the [Apache License 2.0](LICENSE).
