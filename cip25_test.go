package cardanometadata

import (
	"strings"
	"testing"
)

func TestCIP25AssetMetadata_ToMap_Basic(t *testing.T) {
	a := CIP25AssetMetadata{
		Name:  "Test NFT",
		Image: "ipfs://QmTest",
	}
	m := a.ToMap()
	if m["name"] != "Test NFT" {
		t.Errorf("name mismatch: %v", m["name"])
	}
	if m["image"] != "ipfs://QmTest" {
		t.Errorf("image mismatch: %v", m["image"])
	}
	if _, ok := m["description"]; ok {
		t.Error("description should be absent when empty")
	}
}

func TestCIP25AssetMetadata_ToMap_LongImage(t *testing.T) {
	longURI := "ipfs://" + strings.Repeat("Q", 70)
	a := CIP25AssetMetadata{
		Name:  "NFT",
		Image: longURI,
	}
	m := a.ToMap()
	if _, ok := m["image"].([]string); !ok {
		t.Errorf("expected []string for long image, got %T", m["image"])
	}
}

func TestCIP25AssetMetadata_ToMap_WithFiles(t *testing.T) {
	a := CIP25AssetMetadata{
		Name:  "NFT",
		Image: "ipfs://Qm",
		Files: []CIP25File{
			{Name: "artwork", MediaType: "image/png", Src: "ipfs://QmArt"},
		},
	}
	m := a.ToMap()
	files, ok := m["files"].([]MetadataMap)
	if !ok {
		t.Fatalf("expected []MetadataMap for files, got %T", m["files"])
	}
	if len(files) != 1 {
		t.Errorf("expected 1 file, got %d", len(files))
	}
	if files[0]["mediaType"] != "image/png" {
		t.Errorf("mediaType mismatch: %v", files[0]["mediaType"])
	}
}

func TestCIP25AssetMetadata_ToMap_Extra(t *testing.T) {
	a := CIP25AssetMetadata{
		Name:  "NFT",
		Image: "ipfs://Qm",
		Extra: MetadataMap{"rarity": "legendary"},
	}
	m := a.ToMap()
	if m["rarity"] != "legendary" {
		t.Errorf("extra field not propagated: %v", m["rarity"])
	}
}

func TestNewCIP25Metadata_AddAsset_ToMap(t *testing.T) {
	meta := NewCIP25Metadata()
	meta.AddAsset("policy1", "Asset1", CIP25AssetMetadata{
		Name:  "Asset One",
		Image: "ipfs://Qm1",
	})
	m := meta.ToMap()

	label, ok := m["721"]
	if !ok {
		t.Fatal("expected key '721' in map")
	}
	policies, ok := label.(MetadataMap)
	if !ok {
		t.Fatalf("expected MetadataMap under 721, got %T", label)
	}
	if _, ok := policies["policy1"]; !ok {
		t.Error("expected policy1 entry")
	}
}

func TestNewCIP25Metadata_DefaultVersion(t *testing.T) {
	meta := NewCIP25Metadata()
	m := meta.ToMap()
	policies := m["721"].(MetadataMap)
	if policies["version"] != "1.0" {
		t.Errorf("expected version 1.0, got %v", policies["version"])
	}
}

func TestCIP25Metadata_MultipleAssets(t *testing.T) {
	meta := NewCIP25Metadata()
	for i := 0; i < 5; i++ {
		meta.AddAsset("policyX", itoa(i), CIP25AssetMetadata{
			Name:  "NFT" + itoa(i),
			Image: "ipfs://Qm" + itoa(i),
		})
	}
	m := meta.ToMap()
	policies := m["721"].(MetadataMap)
	assets := policies["policyX"].(MetadataMap)
	if len(assets) != 5 {
		t.Errorf("expected 5 assets, got %d", len(assets))
	}
}
