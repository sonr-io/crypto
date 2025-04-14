# Wallet Makefile
# -----------------------------------------------

.PHONY: build

build:
	@gum spin --show-error --title "[WALLET] Running tinygo build..." -- tinygo build -o signer.wasm -target wasip1 -buildmode=c-shared .
	@gum log --level info --time kitchen "[WALLET] Completed tinygo build successfully."
