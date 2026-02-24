package cardanometadata

import "testing"

func TestNewCIP68NFTMetadata_ToMap(t *testing.T) {
	m := NewCIP68NFTMetadata(MetadataMap{
		"name":  "My NFT",
		"image": "ipfs://Qm",
	})
	out := m.ToMap()
	meta, ok := out["metadata"].(MetadataMap)
	if !ok {
		t.Fatalf("expected MetadataMap for metadata key, got %T", out["metadata"])
	}
	if meta["name"] != "My NFT" {
		t.Errorf("name mismatch: %v", meta["name"])
	}
	if out["version"] != "1" {
		t.Errorf("expected version 1, got %v", out["version"])
	}
}

func TestCIP68Metadata_SetGet(t *testing.T) {
	m := NewCIP68NFTMetadata(MetadataMap{})
	m.Set("ticker", "MTK")
	if m.Get("ticker") != "MTK" {
		t.Errorf("Get after Set failed: %v", m.Get("ticker"))
	}
}

func TestCIP68Metadata_GetMissing(t *testing.T) {
	m := NewCIP68NFTMetadata(MetadataMap{})
	if v := m.Get("nonexistent"); v != nil {
		t.Errorf("expected nil for missing key, got %v", v)
	}
}

func TestNewCIP68FTMetadata(t *testing.T) {
	m := NewCIP68FTMetadata(MetadataMap{
		"name":     "MyToken",
		"decimals": 6,
	})
	if m.Version != 1 {
		t.Errorf("expected version 1, got %d", m.Version)
	}
	out := m.ToMap()
	if out["version"] != "1" {
		t.Errorf("expected version string '1', got %v", out["version"])
	}
}

func TestCIP68RoyaltyMetadata_ToMap(t *testing.T) {
	r := CIP68RoyaltyMetadata{Rate: "0.05", Addr: "addr1qtest"}
	m := r.ToMap()
	if m["rate"] != "0.05" {
		t.Errorf("rate mismatch: %v", m["rate"])
	}
	if m["addr"] != "addr1qtest" {
		t.Errorf("addr mismatch: %v", m["addr"])
	}
}
