# Changelog

All notable changes to this project will be documented in this file.
This project adheres to [Semantic Versioning](https://semver.org/).

## [1.0.0] - 2026-02-24

### Added
- `CIP25AssetMetadata` struct with `ToMap()` for CIP-25 compliant NFT metadata generation
- `CIP25Metadata` with `AddAsset()` and `ToMap()` producing label-721 structured output
- `CIP25File` for multi-file asset support
- `CIP68Metadata` with `NewCIP68NFTMetadata`, `NewCIP68FTMetadata`, `Set`, `Get`, `ToMap`
- `CIP68RoyaltyMetadata` with `ToMap()`
- `StringOrChunks` for automatic 64-byte CIP-25 string chunking
- `ValidateCIP25Asset` returning structured `[]*ValidationError`
- `ValidatePolicyID` for hex policy ID validation
- `ValidateMetadataString` for per-string byte compliance checking
- `ValidateCIP68Metadata` for CIP-68 struct validation
- Structured error types: `ValidationError`, `SchemaError`
- Zero external dependencies
- GitHub Actions CI (Go 1.21, 1.22) with race detector
- Full GoDoc examples on all exported functions
