# Sonr Crypto Library Guidelines

## Build Commands
- `make`: Run default build target (proto-gen)
- `make proto-gen`: Generate protobuf files
- `make proto-lint`: Lint protobuf files

## Test Commands
- `go test ./...`: Run all tests
- `go test ./path/to/package`: Test specific package
- `go test ./path/to/package -run TestName`: Run specific test
- `go test -v -cover ./...`: Verbose test output with coverage
- `go test -short ./...`: Run tests with short flag (faster)

## Code Style
- **Formatting**: Use standard Go formatting with `go fmt`
- **Imports**: Group standard library, third-party, and project imports
- **Error Handling**: Return meaningful errors, use package-specific error variables
- **Documentation**: Include godoc comments for exported types and functions
- **Testing**: Use testify/require for assertions, tables for multiple test cases
- **Naming**: Use camelCase for private and PascalCase for exported names
- **Crypto**: Prefer constant-time operations and avoid memory leaks

## Package Structure
Keep crypto implementations in domain-specific packages with separate test files.
Use interfaces for abstraction, especially for curve implementations.