.PHONY: build

all: build publish

tidy:
	@gum spin --show-error --title "[SIGNER] Running go mod tidy..." -- sh -c "go mod tidy"
	@gum log --level info --time kitchen "[SIGNER] Completed go mod tidy successfully."

build: tidy
	@gum spin --show-error --title "[SIGNER] Running tinygo build..." -- sh -c "tinygo build -o signer.wasm -target wasip1 -buildmode=c-shared main.go"
	@gum log --level info --time kitchen "[SIGNER] Completed tinygo build successfully."

publish: build
	@gum spin --show-error --title "[SIGNER] Uploading signer.wasm to r2" -- sh -c "rclone copy ./signer.wasm r2:cdn/bin/"
	@gum log --level info --time kitchen "[SIGNER] Completed rclone upload successfully."


