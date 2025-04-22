.PHONY: build

all: build publish

tidy:
	@gum spin --show-error --title "[VERIFIER] Running go mod tidy..." -- sh -c "go mod tidy"
	@gum log --level info --time kitchen "[VERIFIER] Completed go mod tidy successfully."

build: tidy
	@gum spin --show-error --title "[VERIFIER] Running go build..." -- sh -c "GOOS=wasip1 GOARCH=wasm go build -buildmode=c-shared -o verifier.wasm main.go"
	@gum log --level info --time kitchen "[VERIFIER] Completed go build successfully."

test:
	@extism call verifier.wasm verify --input test.json

publish: build
	@gum spin --show-error --title "[VERIFIER] Uploading verifier.wasm to r2" -- sh -c "rclone copy ./verifier.wasm r2:cdn/bin/"
	@gum log --level info --time kitchen "[VERIFIER] Completed rclone upload successfully."


