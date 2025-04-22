.PHONY: build

all: build publish

tidy:
	@gum spin --show-error --title "[ENCLAVE] Running go mod tidy..." -- sh -c "go mod tidy"
	@gum log --level info --time kitchen "[ENCLAVE] Completed go mod tidy successfully."

build: tidy
	@gum spin --show-error --title "[ENCLAVE] Running tinygo build..." -- sh -c "tinygo build -o enclave.wasm -target wasip1 -buildmode=c-shared ."
	@gum log --level info --time kitchen "[ENCLAVE] Completed tinygo build successfully."

publish: build
	@gum spin --show-error --title "[ENCLAVE] Uploading enclave.wasm to r2" -- sh -c "rclone copy ./enclave.wasm r2:cdn/bin/"
	@gum log --level info --time kitchen "[ENCLAVE] Completed rclone upload successfully."


