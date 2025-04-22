.PHONY: build test

all: build publish

tidy:
	@gum spin --show-error --title "[ENCLAVE] Running go mod tidy..." -- sh -c "go mod tidy"
	@gum log --level info --time kitchen "[ENCLAVE] Completed go mod tidy successfully."

build: tidy
	@gum spin --show-error --title "[ENCLAVE] Running go build..." -- sh -c "GOOS=wasip1 GOARCH=wasm go build -buildmode=c-shared -o enclave.wasm main.go"
	@gum log --level info --time kitchen "[ENCLAVE] Completed go build successfully."

test: build
	@extism call enclave.wasm generate --wasi >> enclave.json

publish: build
	@gum spin --show-error --title "[ENCLAVE] Uploading enclave.wasm to r2" -- sh -c "rclone copy ./enclave.wasm r2:cdn/bin/"
	@gum log --level info --time kitchen "[ENCLAVE] Completed rclone upload successfully."


