// Package cardanometadata provides zero-dependency utilities for creating,
// validating, and serializing Cardano NFT metadata compliant with CIP-25 and CIP-68.
//
// It is designed for use by both human developers and AI agents building
// automated NFT minting pipelines on Cardano.
package cardanometadata

import (
	"encoding/json"
	"fmt"
)

// Version is the library version.
const Version = "1.0.0"

// Standard CIP metadata label numbers.
const (
	LabelCIP25 = 721  // NFT metadata label (CIP-25)
	LabelCIP68 = 100  // Reference token label (CIP-68)
	LabelCIP68RFT = 333 // Royalty token label (CIP-68)
	LabelCIP68NFT = 222 // NFT token label (CIP-68)
	LabelCIP68FT  = 333 // Fungible token label (CIP-68)
)

// MetadataMap is an ordered representation of Cardano on-chain metadata.
// Keys are strings; values can be strings, numbers, booleans, slices, or nested MetadataMaps.
type MetadataMap map[string]any

// MarshalJSON encodes MetadataMap as a JSON object.
func (m MetadataMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any(m))
}

// Error types for structured, machine-readable errors.

// ValidationError is returned when metadata fails CIP compliance checks.
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error: field=%q message=%q", e.Field, e.Message)
}

// SchemaError is returned when the metadata structure is malformed.
type SchemaError struct {
	Expected string
	Got      string
}

func (e *SchemaError) Error() string {
	return fmt.Sprintf("schema error: expected=%q got=%q", e.Expected, e.Got)
}

// chunkString splits s into chunks of max length n, per CIP-25 string limits.
func chunkString(s string, n int) []string {
	if len(s) <= n {
		return []string{s}
	}
	var chunks []string
	for len(s) > 0 {
		end := n
		if end > len(s) {
			end = len(s)
		}
		chunks = append(chunks, s[:end])
		s = s[end:]
	}
	return chunks
}

// StringOrChunks returns a string as-is if <= maxLen, otherwise returns []string chunks.
// Cardano metadata strings are limited to 64 bytes each.
//
// Example:
//
//	val := cardanometadata.StringOrChunks("short", 64)
//	// returns "short"
//
//	val2 := cardanometadata.StringOrChunks(strings.Repeat("a", 100), 64)
//	// returns []string{"aaa...64 chars...", "aaa...36 chars..."}
func StringOrChunks(s string, maxLen int) any {
	if maxLen <= 0 {
		maxLen = 64
	}
	if len(s) <= maxLen {
		return s
	}
	return chunkString(s, maxLen)
}
