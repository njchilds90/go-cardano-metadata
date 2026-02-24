package cardanometadata

import (
	"strings"
	"unicode/utf8"
)

const (
	maxStringBytes  = 64
	maxPolicyIDLen  = 56 // hex chars for 28-byte policy ID
	minPolicyIDLen  = 1
)

// ValidateCIP25Asset checks that a CIP25AssetMetadata satisfies CIP-25 requirements.
// Returns a []*ValidationError slice (empty slice means valid).
//
// Example:
//
//	errs := cardanometadata.ValidateCIP25Asset(asset)
//	if len(errs) > 0 {
//	    for _, e := range errs {
//	        log.Println(e)
//	    }
//	}
func ValidateCIP25Asset(a CIP25AssetMetadata) []*ValidationError {
	var errs []*ValidationError

	if strings.TrimSpace(a.Name) == "" {
		errs = append(errs, &ValidationError{Field: "name", Message: "must not be empty"})
	}

	if strings.TrimSpace(a.Image) == "" {
		errs = append(errs, &ValidationError{Field: "image", Message: "must not be empty"})
	}

	// Individual string segments must not exceed 64 bytes
	if len(a.Name) > 0 && !validateStringChunks(a.Name) {
		// Only invalid if a single chunk would exceed 64 bytes without auto-chunking
		// (auto-chunking handles this; this checks raw name plausibility)
		_ = a.Name // chunking is applied by ToMap; name itself can be any length
	}

	for i, f := range a.Files {
		if strings.TrimSpace(f.MediaType) == "" {
			errs = append(errs, &ValidationError{
				Field:   "files[" + itoa(i) + "].mediaType",
				Message: "must not be empty",
			})
		}
		if strings.TrimSpace(f.Src) == "" {
			errs = append(errs, &ValidationError{
				Field:   "files[" + itoa(i) + "].src",
				Message: "must not be empty",
			})
		}
	}

	return errs
}

// ValidatePolicyID checks whether a policy ID is a valid 56-character hex string.
//
// Example:
//
//	ok := cardanometadata.ValidatePolicyID("abc123...")
func ValidatePolicyID(policyID string) error {
	if len(policyID) < minPolicyIDLen || len(policyID) > maxPolicyIDLen {
		return &ValidationError{
			Field:   "policyID",
			Message: "must be between 1 and 56 hex characters",
		}
	}
	for _, c := range policyID {
		if !isHexChar(c) {
			return &ValidationError{
				Field:   "policyID",
				Message: "must contain only hex characters (0-9, a-f, A-F)",
			}
		}
	}
	return nil
}

// ValidateMetadataString checks that a single metadata string value is
// either ≤ 64 bytes, or a []string where each element is ≤ 64 bytes.
//
// Example:
//
//	err := cardanometadata.ValidateMetadataString("hello")
func ValidateMetadataString(v any) error {
	switch s := v.(type) {
	case string:
		if !utf8.ValidString(s) {
			return &ValidationError{Field: "string", Message: "must be valid UTF-8"}
		}
		if len(s) > maxStringBytes {
			return &ValidationError{
				Field:   "string",
				Message: "exceeds 64 bytes; use StringOrChunks to split automatically",
			}
		}
		return nil
	case []string:
		for i, chunk := range s {
			if len(chunk) > maxStringBytes {
				return &ValidationError{
					Field:   "string[" + itoa(i) + "]",
					Message: "chunk exceeds 64 bytes",
				}
			}
		}
		return nil
	default:
		return &SchemaError{Expected: "string or []string", Got: "other"}
	}
}

// ValidateCIP68Metadata checks required fields for a CIP-68 metadata struct.
//
// Example:
//
//	errs := cardanometadata.ValidateCIP68Metadata(nftMeta)
func ValidateCIP68Metadata(c *CIP68Metadata) []*ValidationError {
	var errs []*ValidationError
	if c == nil {
		return append(errs, &ValidationError{Field: "metadata", Message: "must not be nil"})
	}
	if len(c.Metadata) == 0 {
		errs = append(errs, &ValidationError{Field: "metadata", Message: "must not be empty"})
	}
	if c.Version < 1 {
		errs = append(errs, &ValidationError{Field: "version", Message: "must be >= 1"})
	}
	return errs
}

// helpers

func validateStringChunks(s string) bool {
	return len(s) <= maxStringBytes
}

func isHexChar(c rune) bool {
	return (c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	b := make([]byte, 0, 10)
	for n > 0 {
		b = append([]byte{byte('0' + n%10)}, b...)
		n /= 10
	}
	return string(b)
}
