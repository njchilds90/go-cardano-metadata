# go-cardano-metadata

[![CI](https://github.com/njchilds90/go-cardano-metadata/actions/workflows/ci.yml/badge.svg)](https://github.com/njchilds90/go-cardano-metadata/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/njchilds90/go-cardano-metadata.svg)](https://pkg.go.dev/github.com/njchilds90/go-cardano-metadata)
[![Go Report Card](https://goreportcard.com/badge/github.com/njchilds90/go-cardano-metadata)](https://goreportcard.com/report/github.com/njchilds90/go-cardano-metadata)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Version](https://img.shields.io/github/v/tag/njchilds90/go-cardano-metadata?label=version)](https://github.com/njchilds90/go-cardano-metadata/releases)

**Zero-dependency Go library for Cardano NFT metadata — CIP-25 and CIP-68 compliant.**

Built for human developers and AI agents that need to generate, validate, and serialize Cardano-native metadata without pulling in heavy blockchain SDKs.

---

## Features

- ✅ **CIP-25** NFT metadata generation (label 721)
- ✅ **CIP-68** reference datum metadata (NFT label 222, FT label 333, royalty label 444)
- ✅ Automatic 64-byte string chunking (CIP-25 compliant)
- ✅ Structured, machine-readable validation errors
- ✅ Pure functions — no global state, no side effects
- ✅ Zero external dependencies
- ✅ Full GoDoc coverage with examples

---

## Installation
```bash
go get github.com/njchilds90/go-cardano-metadata@v1.0.0
```

---

## Quick Start

### CIP-25 NFT Metadata
```go
package main

import (
    "encoding/json"
    "fmt"

    cardanometadata "github.com/njchilds90/go-cardano-metadata"
)

func main() {
    meta := cardanometadata.NewCIP25Metadata()

    meta.AddAsset("your_policy_id_hex", "MyNFT001", cardanometadata.CIP25AssetMetadata{
        Name:        "My NFT #1",
        Image:       "ipfs://QmYourImageHashHere",
        MediaType:   "image/png",
        Description: "A foundational Cardano NFT",
        Files: []cardanometadata.CIP25File{
            {Name: "Full Resolution", MediaType: "image/png", Src: "ipfs://QmFullRes"},
        },
        Extra: cardanometadata.MetadataMap{
            "rarity": "legendary",
            "traits": []string{"golden", "animated"},
        },
    })

    m := meta.ToMap()
    b, _ := json.MarshalIndent(m, "", "  ")
    fmt.Println(string(b))
}
```

**Output:**
```json
{
  "721": {
    "your_policy_id_hex": {
      "MyNFT001": {
        "description": "A foundational Cardano NFT",
        "files": [{"mediaType": "image/png", "name": "Full Resolution", "src": "ipfs://QmFullRes"}],
        "image": "ipfs://QmYourImageHashHere",
        "mediaType": "image/png",
        "name": "My NFT #1",
        "rarity": "legendary",
        "traits": ["golden", "animated"]
      }
    },
    "version": "1.0"
  }
}
```

---

### CIP-68 NFT Metadata (Reference Datum)
```go
nft := cardanometadata.NewCIP68NFTMetadata(cardanometadata.MetadataMap{
    "name":        "My CIP-68 NFT",
    "image":       "ipfs://QmHash",
    "description": "Stored in reference UTxO",
})
nft.Set("rarity", "rare")

out := nft.ToMap()
b, _ := json.MarshalIndent(out, "", "  ")
fmt.Println(string(b))
```

---

### CIP-68 Fungible Token
```go
ft := cardanometadata.NewCIP68FTMetadata(cardanometadata.MetadataMap{
    "name":     "My Token",
    "ticker":   "MTK",
    "decimals": 6,
    "url":      "https://mytoken.io",
    "logo":     "ipfs://QmLogo",
})
```

---

### Royalty Metadata
```go
royalty := cardanometadata.CIP68RoyaltyMetadata{
    Rate: "0.05", // 5%
    Addr: "addr1q...",
}
m := royalty.ToMap()
```

---

### Validation
```go
asset := cardanometadata.CIP25AssetMetadata{
    Name:  "",         // invalid — empty
    Image: "ipfs://Qm",
}

errs := cardanometadata.ValidateCIP25Asset(asset)
for _, e := range errs {
    fmt.Println(e) // validation error: field="name" message="must not be empty"
}

// Validate a policy ID
if err := cardanometadata.ValidatePolicyID("not-a-valid-id!"); err != nil {
    fmt.Println(err)
}

// Validate a metadata string for length compliance
if err := cardanometadata.ValidateMetadataString("a short string"); err != nil {
    fmt.Println(err)
}
```

---

### Auto String Chunking

CIP-25 requires that no metadata string exceed 64 bytes. Use `StringOrChunks` to handle this automatically:
```go
uri := "ipfs://" + strings.Repeat("Q", 80) // 87 bytes total

val := cardanometadata.StringOrChunks(uri, 64)
// val is []string{"ipfs://QQQQQ...64 chars", "QQQ...23 chars"}
```

`ToMap()` on any asset calls this automatically for `image`, `name`, `description`, and file `src` fields.

---

## CIP Compliance Notes

| Feature | Standard | Status |
|---------|----------|--------|
| NFT metadata structure | CIP-25 | ✅ v1.0 |
| 64-byte string chunking | CIP-25 | ✅ v1.0 |
| Reference datum structure | CIP-68 | ✅ v1.0 |
| FT/NFT/Royalty token types | CIP-68 | ✅ v1.0 |
| Policy ID validation | CIP-25 | ✅ v1.0 |
| CBOR serialization | CIP-25/68 | 🔜 v1.1 (bring your own CBOR) |
| Asset name hex encoding | CIP-25 | 🔜 v1.1 |

> For CBOR encoding of the produced `MetadataMap`, pair this library with a CBOR encoder such as [`fxamacker/cbor`](https://github.com/fxamacker/cbor) (optional, external).

---

## AI Agent Usage

This library is designed to be called directly by AI agents in NFT minting pipelines:
```go
// Agent receives structured config; produces compliant metadata deterministically
func buildMetadata(config AgentConfig) (map[string]any, []error) {
    meta := cardanometadata.NewCIP25Metadata()

    asset := cardanometadata.CIP25AssetMetadata{
        Name:      config.Name,
        Image:     config.IPFSURI,
        MediaType: config.MIMEType,
        Extra:     cardanometadata.MetadataMap(config.Traits),
    }

    errs := cardanometadata.ValidateCIP25Asset(asset)
    if len(errs) != 0 {
        var out []error
        for _, e := range errs { out = append(out, e) }
        return nil, out
    }

    meta.AddAsset(config.PolicyID, config.AssetName, asset)
    return meta.ToMap(), nil
}
```

---

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md).

---

## Changelog

See [CHANGELOG.md](CHANGELOG.md).

---

## License

MIT — see [LICENSE](LICENSE).
