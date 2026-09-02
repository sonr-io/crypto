# Sonr Crypto

Cryptographic primitives used by Sonr: elliptic-curve arithmetic, threshold signatures,
multi-party computation, zero-knowledge proofs, and the identity layer built on top of them.

**Documentation: https://sonr-io.github.io/crypto**

```bash
go get github.com/sonr-io/crypto
```

Requires Go 1.24.7 or newer.

> [!WARNING]
> This library has no public security audit, and several packages contain stubs or known defects.
> Read the [security notes](https://sonr-io.github.io/crypto/reference/security) before depending on
> any of it.

## Packages

### Foundations

- [`core/curves`](./core/curves): the `Curve` / `Point` / `Scalar` abstraction every other package is
  written against, with secp256k1, NIST P-256, Ed25519, Pallas, BLS12-381, and BLS12-377.
- [`core`](./core): modular arithmetic, hash-to-field, HMAC commitments, and safe-prime generation.
- [`core/protocol`](./core/protocol): the message and iterator types that drive multi-round protocols.

### Symmetric and key derivation

- [`aead`](./aead): AES-256-GCM with generated, prepended nonces.
- [`daed`](./daed): AES-SIV-CMAC deterministic AEAD (RFC 5297).
- [`argon2`](./argon2): Argon2id password hashing and key derivation.
- [`subtle`](./subtle): HKDF, X25519, and hash/curve name helpers.
- [`salt`](./salt), [`password`](./password), [`secure`](./secure): salt handling, password policy,
  and memory-hygiene helpers.

### Signatures

- [`signatures/bls`](./signatures/bls): BLS signatures with aggregation, proof of possession, and
  threshold key generation.
- [`signatures/bbs`](./signatures/bbs): BBS+ signatures for selective disclosure and blind signing.
- [`signatures/schnorr`](./signatures/schnorr): Mina (Pallas/Poseidon) and NEM (Ed25519-Keccak) schemes.
- [`ecdsa`](./ecdsa): low-S canonicalization and RFC 6979 deterministic signing.
- [`vrf`](./vrf): verifiable random function over Edwards25519.

### Threshold cryptography and MPC

- [`sharing`](./sharing): Shamir, Feldman, and Pedersen secret sharing.
- [`dkg`](./dkg): FROST and Gennaro distributed key generation.
- [`tecdsa`](./tecdsa): two-party threshold ECDSA (DKLs18) with key refresh.
- [`ted25519`](./ted25519): threshold Ed25519 and FROST threshold Schnorr signing.
- [`ot`](./ot): base oblivious transfer and correlated OT extension — building blocks for the above.
- [`mpc`](./mpc): threshold ECDSA packaged as a single enclave value: keygen, sign, verify, refresh.

### Zero-knowledge

- [`zkp/schnorr`](./zkp/schnorr): non-interactive proof of knowledge of a discrete log.
- [`accumulator`](./accumulator): pairing-based accumulator with constant-size membership proofs.
- [`bulletproof`](./bulletproof): inner-product argument and range proofs.
- [`paillier`](./paillier): additively homomorphic encryption with a square-free modulus proof.

### Identity

- [`keys`](./keys): `did:key` encoding and decoding over libp2p public keys.
- [`ucan`](./ucan): UCAN capability tokens, attenuation, and verification.
- [`ecies`](./ecies): ECIES payload encryption over secp256k1.
- [`wasm`](./wasm): Ed25519 signing and hash pinning for WebAssembly modules.

## Documentation

The site under `docs/` is built with [Blume](https://useblume.dev) and deployed to GitHub Pages by
`.github/workflows/docs.yml`.

```bash
npm ci
npm run dev    # local preview with hot reload
npm run build  # static output to dist/
```

## License

This project is licensed under the [Apache 2.0 License](./LICENSE).
