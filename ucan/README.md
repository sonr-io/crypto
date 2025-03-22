# UCAN Implementation Guide for Sonr Blockchain

This guide outlines how to implement the UCAN (User Controlled Authorization Network) specification in the Sonr blockchain modules.

## Project Structure

The UCAN implementation consists of:

1. **Protobuf Definitions** - In the `sonr-io/crypto` repository, published to buf.build
2. **Go Implementation** - Updated from v0.7.0 to v1.0.0-rc.1 
3. **Module Integration** - Integration with x/did, x/dwn, and x/svc modules

## Capability Hierarchies

Based on your diagram, here are the capability hierarchies for each module:

### DWN (Vault) Module

```
/
└── /vault
    ├── /vault/sign
    ├── /vault/refresh
    └── /vault/verify
```

### DID (Account) Module

```
/
└── /account
    ├── /account/execute
    ├── /account/deposit
    ├── /account/transfer
    ├── /account/derive
    ├── /account/link
    ├── /account/unlink
    └── /account/delegate
```

### SVC (Service) Module

```
/
└── /service
    ├── /service/register
    ├── /service/update
    └── /service/delete
```

### IBC Module

```
/
└── /ibc
    ├── /ibc/broadcast
    ├── /ibc/query
    └── /ibc/simulate
    └── /ibc/chain
        ├── /ibc/chain/nomic
        ├── /ibc/chain/evmos
        └── /ibc/chain/osmos
```

## Implementation Steps

1. **Update UCAN Go Implementation**
   - Update existing implementation in `ucan/` to match v1.0.0-rc.1
   - Implement proper envelope format with DAG-CBOR encoding
   - Implement the policy language for attenuation validation

2. **Publish Protobuf Definitions**
   - Upload protobuf definitions to buf.build
   - Create a new repository for UCAN implementation

3. **Integrate with Module Genesis**
   - Add capability hierarchies to each module's genesis state
   - Create initialization logic to load capabilities on startup

4. **Implement Middleware**
   - Create UCAN validation middleware for each module's message handlers
   - Implement token verification and capability checking

5. **Update Client Code**
   - Create client-side libraries for generating UCAN tokens
   - Implement delegation and invocation patterns

## Usage Examples

### Creating a UCAN Token

```go
// Create a new UCAN token
token, err := ucanSource.NewOriginToken(
    audienceDID,
    ucan.DWNCapability("/vault/sign"),
    []ucan.Fact{},
    time.Now(),
    time.Now().Add(24 * time.Hour),
)
```

### Validating a UCAN Token

```go
// In the message handler
func (s msgServer) SignMessage(ctx context.Context, msg *types.MsgSignMessage) (*types.MsgSignMessageResponse, error) {
    // Parse and verify the UCAN token
    token, err := s.ucanParser.ParseAndVerify(ctx, msg.Authorization)
    if err != nil {
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "invalid token")
    }
    
    // Check if token has the right capability
    if !hasCapability(token, "/vault/sign", msg.VaultId) {
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "missing capability")
    }
    
    // Proceed with the operation
    // ...
}
```

## Genesis Configuration

Add the capability hierarchies to the genesis state of each module. Example for the DWN module:

```json
{
  "params": {
    "attenuations": [
      {
        "resource": {
          "kind": "vault",
          "template": "{id}"
        },
        "capabilities": [
          {
            "name": "sign",
            "parent": "/vault",
            "description": "Sign a message using the vault"
          },
          {
            "name": "refresh",
            "parent": "/vault",
            "description": "Refresh vault access"
          },
          {
            "name": "verify",
            "parent": "/vault",
            "description": "Verify a signature from the vault"
          }
        ]
      }
    ]
  }
}
```

