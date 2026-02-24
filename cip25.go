package cardanometadata

import (
	"fmt"
)

// CIP25AssetMetadata represents the per-asset metadata object under CIP-25.
// See: https://github.com/cardano-foundation/CIPs/tree/master/CIP-0025
type CIP25AssetMetadata struct {
	// Name is the human-readable asset name (required).
	Name string `json:"name"`
	// Image is the primary image URI (required, max 64 bytes or chunked).
	Image string `json:"image"`
	// MediaType is the MIME type of the image (optional, e.g. "image/png").
	MediaType string `json:"mediaType,omitempty"`
	// Description is an optional description string (chunked if > 64 bytes).
	Description string `json:"description,omitempty"`
	// Files lists additional associated files (optional).
	Files []CIP25File `json:"files,omitempty"`
	// Extra holds any additional user-defined fields.
	Extra MetadataMap `json:"-"`
}

// CIP25File represents a file entry in CIP-25 metadata.
type CIP25File struct {
	Name      string `json:"name"`
	MediaType string `json:"mediaType"`
	Src       string `json:"src"`
}

// ToMap converts CIP25AssetMetadata into a MetadataMap suitable for on-chain encoding.
// Strings longer than 64 bytes are automatically chunked per CIP-25.
//
// Example:
//
//	asset := cardanometadata.CIP25AssetMetadata{
//	    Name:  "MyNFT #1",
//	    Image: "ipfs://Qm...",
//	}
//	m := asset.ToMap()
func (a CIP25AssetMetadata) ToMap() MetadataMap {
	m := MetadataMap{
		"name":  StringOrChunks(a.Name, 64),
		"image": StringOrChunks(a.Image, 64),
	}
	if a.MediaType != "" {
		m["mediaType"] = a.MediaType
	}
	if a.Description != "" {
		m["description"] = StringOrChunks(a.Description, 64)
	}
	if len(a.Files) > 0 {
		files := make([]MetadataMap, len(a.Files))
		for i, f := range a.Files {
			files[i] = MetadataMap{
				"name":      f.Name,
				"mediaType": f.MediaType,
				"src":       StringOrChunks(f.Src, 64),
			}
		}
		m["files"] = files
	}
	for k, v := range a.Extra {
		m[k] = v
	}
	return m
}

// CIP25PolicyMetadata maps asset names (hex or UTF-8) to their metadata.
type CIP25PolicyMetadata map[string]CIP25AssetMetadata

// CIP25Metadata is the top-level CIP-25 metadata object keyed by label 721.
type CIP25Metadata struct {
	// Policies maps policy IDs to their per-asset metadata.
	Policies map[string]CIP25PolicyMetadata
	// Version is "1.0" or "2.0" per CIP-25 spec (default "1.0").
	Version string
}

// NewCIP25Metadata creates a CIP25Metadata with version "1.0".
//
// Example:
//
//	meta := cardanometadata.NewCIP25Metadata()
func NewCIP25Metadata() *CIP25Metadata {
	return &CIP25Metadata{
		Policies: make(map[string]CIP25PolicyMetadata),
		Version:  "1.0",
	}
}

// AddAsset adds an asset's metadata under the given policyID and assetName.
//
// Example:
//
//	meta.AddAsset("abc123policy", "MyNFT1", cardanometadata.CIP25AssetMetadata{
//	    Name:  "My NFT #1",
//	    Image: "ipfs://Qm...",
//	})
func (c *CIP25Metadata) AddAsset(policyID, assetName string, asset CIP25AssetMetadata) {
	if _, ok := c.Policies[policyID]; !ok {
		c.Policies[policyID] = make(CIP25PolicyMetadata)
	}
	c.Policies[policyID][assetName] = asset
}

// ToMap produces the final MetadataMap keyed under label 721 as required by CIP-25.
//
// Example:
//
//	m := meta.ToMap()
//	// m[721]["abc123policy"]["MyNFT1"] = {...}
func (c *CIP25Metadata) ToMap() MetadataMap {
	version := c.Version
	if version == "" {
		version = "1.0"
	}
	policies := MetadataMap{}
	for policyID, assets := range c.Policies {
		assetMap := MetadataMap{}
		for assetName, asset := range assets {
			assetMap[assetName] = asset.ToMap()
		}
		policies[policyID] = assetMap
	}
	policies["version"] = version
	return MetadataMap{
		fmt.Sprintf("%d", LabelCIP25): policies,
	}
}
