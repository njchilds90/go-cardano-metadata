package cardanometadata

import (
	"strings"
	"testing"
)

func TestValidateCIP25Asset_Valid(t *testing.T) {
	a := CIP25AssetMetadata{Name: "NFT", Image: "ipfs://Qm"}
	errs := ValidateCIP25Asset(a)
	if len(errs) != 0 {
		t.Errorf("expected no errors, got %v", errs)
	}
}

func TestValidateCIP25Asset_EmptyName(t *testing.T) {
	a := CIP25AssetMetadata{Name: "", Image: "ipfs://Qm"}
	errs := ValidateCIP25Asset(a)
	if len(errs) == 0 {
		t.Error("expected error for empty name")
	}
}

func TestValidateCIP25Asset_EmptyImage(t *testing.T) {
	a := CIP25AssetMetadata{Name: "NFT", Image: ""}
	errs := ValidateCIP25Asset(a)
	if len(errs) == 0 {
		t.Error("expected error for empty image")
	}
}

func TestValidateCIP25Asset_FileMissingMediaType(t *testing.T) {
	a := CIP25AssetMetadata{
		Name:  "NFT",
		Image: "ipfs://Qm",
		Files: []CIP25File{{Src: "ipfs://art"}},
	}
	errs := ValidateCIP25Asset(a)
	if len(errs) == 0 {
		t.Error("expected error for file missing mediaType")
	}
}

func TestValidatePolicyID_Valid(t *testing.T) {
	tests := []struct {
		id    string
		valid bool
	}{
		{"abc123", true},
		{strings.Repeat("a", 56), true},
		{"", false},
		{strings.Repeat("a", 57), false},
		{"xyz!!!", false},
	}
	for _, tt := range tests {
		err := ValidatePolicyID(tt.id)
		if tt.valid && err != nil {
			t.Errorf("policyID %q: expected valid, got %v", tt.id, err)
		}
		if !tt.valid && err == nil {
			t.Errorf("policyID %q: expected invalid, got nil", tt.id)
		}
	}
}

func TestValidateMetadataString_Short(t *testing.T) {
	if err := ValidateMetadataString("hello"); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestValidateMetadataString_TooLong(t *testing.T) {
	long := strings.Repeat("a", 65)
	if err := ValidateMetadataString(long); err == nil {
		t.Error("expected error for 65-byte string")
	}
}

func TestValidateMetadataString_Chunks(t *testing.T) {
	chunks := []string{"hello", "world"}
	if err := ValidateMetadataString(chunks); err != nil {
		t.Errorf("unexpected error for valid chunks: %v", err)
	}
}

func TestValidateMetadataString_ChunkTooLong(t *testing.T) {
	chunks := []string{strings.Repeat("x", 65)}
	if err := ValidateMetadataString(chunks); err == nil {
		t.Error("expected error for oversized chunk")
	}
}

func TestValidateMetadataString_WrongType(t *testing.T) {
	if err := ValidateMetadataString(42); err == nil {
		t.Error("expected schema error for int input")
	}
}

func TestValidateCIP68Metadata_Valid(t *testing.T) {
	m := NewCIP68NFTMetadata(MetadataMap{"name": "x"})
	errs := ValidateCIP68Metadata(m)
	if len(errs) != 0 {
		t.Errorf("expected no errors, got %v", errs)
	}
}

func TestValidateCIP68Metadata_Nil(t *testing.T) {
	errs := ValidateCIP68Metadata(nil)
	if len(errs) == 0 {
		t.Error("expected error for nil metadata")
	}
}

func TestValidateCIP68Metadata_Empty(t *testing.T) {
	m := &CIP68Metadata{Metadata: MetadataMap{}, Version: 1}
	errs := ValidateCIP68Metadata(m)
	if len(errs) == 0 {
		t.Error("expected error for empty metadata map")
	}
}
