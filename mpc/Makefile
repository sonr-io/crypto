.PHONY: build

all: enclave signer verifier

enclave:
	@$(MAKE) -C enclave all

signer:
	@$(MAKE) -C signer all

verifier:
	@$(MAKE) -C verifier all

test: all
	@extism call ./enclave/enclave.wasm generate --wasi >> enclave.json




