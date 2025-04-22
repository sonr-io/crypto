.PHONY: proto-gen proto-lint proto-format proto-check proto-clean proto-publish build

PROTO_DIR = ./proto

build:
	@earthly +signer

# Generate protobuf files
proto-gen:
	@echo "Generating Protobuf files..."
	@mkdir -p ucan/types/v1
	@cd $(PROTO_DIR) && buf generate
	@find proto/ucan/types/v1/crypto/ucan/v1/ -name "*.pb.go" -exec mv {} ucan/types/v1/ \;
	@find proto/ucan/types/v1/crypto/ucan/v1/ -name "*.pb.validate.go" -exec mv {} ucan/types/v1/ \;
	@rm -rf proto/ucan

# Lint protobuf files
proto-lint:
	@echo "Linting Protobuf files..."
	@cd $(PROTO_DIR) && buf lint

# Format protobuf files
proto-format:
	@echo "Formatting Protobuf files..."
	@find $(PROTO_DIR) -name "*.proto" -exec clang-format -i {} \;

# Check protobuf files
proto-check: proto-lint
	@echo "Checking Protobuf files..."
	@cd $(PROTO_DIR) && buf breaking --against "https://github.com/onsonr/crypto.git#branch=main"

# Clean generated files
proto-clean:
	@echo "Cleaning generated files..."
	@rm -rf ucan/types/v1/crypto

# Publish protobuf files to buf.build
proto-publish:
	@echo "Publishing Protobuf files to buf.build..."
	@cd $(PROTO_DIR) && buf push

plugins:
	@$(MAKE) -C mpc/enclave all
	@$(MAKE) -C mpc/signer all
	@$(MAKE) -C mpc/verifier all

all: proto-gen
