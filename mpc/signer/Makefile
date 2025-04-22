.PHONY: build

all: build publish

tidy:
	@gum spin --show-error --title "[SIGNER] Running go mod tidy..." -- sh -c "go mod tidy"
	@gum log --level info --time kitchen "[SIGNER] Completed go mod tidy successfully."

build: 
	@gum spin --show-error --title "[SIGNER] Running go build..." -- sh -c "GOOS=wasip1 GOARCH=wasm go build -buildmode=c-shared -o signer.wasm main.go"
	@gum log --level info --time kitchen "[SIGNER] Completed go build successfully."

publish: build
	@gum spin --show-error --title "[SIGNER] Uploading signer.wasm to r2" -- sh -c "rclone copy ./signer.wasm r2:cdn/bin/"
	@gum log --level info --time kitchen "[SIGNER] Completed rclone upload successfully."


