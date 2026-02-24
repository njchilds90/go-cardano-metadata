# Contributing to go-cardano-metadata

Thank you for your interest in contributing!

## Getting Started

1. Fork the repository on GitHub
2. Clone your fork: `git clone https://github.com/YOUR_USERNAME/go-cardano-metadata.git`
3. Create a feature branch: `git checkout -b feature/my-feature`
4. Make your changes with tests
5. Run: `go test -race ./...`
6. Run: `go vet ./...`
7. Push and open a Pull Request against `main`

## Code Standards

- All exported symbols must have GoDoc comments
- New features require table-driven tests with good coverage
- Zero external runtime dependencies — keep it that way
- Follow standard Go idioms (gofmt, go vet clean)
- Structured errors only — never raw `errors.New("...")` for public API errors

## CIP Compliance

When adding support for new Cardano CIPs, please link to the official CIP document in your PR description and GoDoc.

## Reporting Issues

Please open a GitHub Issue with:
- Go version (`go version`)
- Minimal reproduction case
- Expected vs actual behavior

## License

By contributing, you agree your contributions will be licensed under MIT.
```

---

## 4. Release & Verification Instructions
```


Zero-dependency Go library for Cardano NFT metadata — CIP-25 and CIP-68 compliant.

### What's Included
- CIP-25 metadata builder (label 721) with automatic 64-byte string chunking
- CIP-68 reference datum metadata (NFT/FT/Royalty)
- Structured validation with machine-readable error types
- Pure functions, no global state, no external dependencies
- Full GoDoc + GitHub Actions CI

### Install
```
go get github.com/njchilds90/go-cardano-metadata@v1.0.0
```

   e. Check: ✅ "Set as the latest release"
   f. Click: "Publish release"

4. TRIGGER pkg.go.dev INDEXING
   - Visit this URL in your browser (triggers the proxy fetch):
     https://sum.golang.org/lookup/github.com/njchilds90/go-cardano-metadata@v1.0.0
   - Then visit:
     https://pkg.go.dev/github.com/njchilds90/go-cardano-metadata@v1.0.0

   Within 5–10 minutes, the full GoDoc will be live and searchable.

5. VERIFY
   ✅ https://pkg.go.dev/github.com/njchilds90/go-cardano-metadata
   ✅ https://goreportcard.com/report/github.com/njchilds90/go-cardano-metadata
   ✅ https://github.com/njchilds90/go-cardano-metadata/actions (CI green)

SEMANTIC VERSIONING GUIDANCE
==============================
- v1.0.x — Patch: bug fixes, doc improvements, no API changes
- v1.x.0 — Minor: new exported functions/types (backward compatible)
  e.g. v1.1.0: Add CBOR serialization helpers, asset name hex encoding
  e.g. v1.2.0: Add CIP-27 (royalty standard) support
  e.g. v1.3.0: Add CIP-60 (music metadata) support
- v2.0.0 — Major: breaking API changes (avoid as long as possible)
