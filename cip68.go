package cardanometadata

import "fmt"

// CIP68Metadata represents a CIP-68 datum metadata structure.
// CIP-68 stores metadata in a reference UTxO datum rather than transaction metadata.
// See: https://github.com/cardano-foundation/CIPs/tree/master/CIP-0068
type CIP68Metadata struct {
	// Metadata holds the key-value pairs for the asset.
	Metadata MetadataMap
	// Version is the CIP-68 version integer (currently 1).
	Version int
	// Extra is reserved for forward-compatibility (ignored fields).
	Extra any
}

// NewCIP68NFTMetadata creates a CIP-68 NFT (label 222) metadata object.
//
// Example:
//
//	m := cardanometadata.NewCIP68NFTMetadata(cardanometadata.MetadataMap{
//	    "name":  "My NFT",
//	    "image": "ipfs://Qm...",
//	})
func NewCIP68NFTMetadata(fields MetadataMap) *CIP68Metadata {
	return &CIP68Metadata{
		Metadata: fields,
		Version:  1,
	}
}

// NewCIP68FTMetadata creates a CIP-68 fungible token (label 333) metadata object.
//
// Example:
//
//	m := cardanometadata.NewCIP68FTMetadata(cardanometadata.MetadataMap{
//	    "name":     "MyToken",
//	    "ticker":   "MTK",
//	    "decimals": 6,
//	})
func NewCIP68FTMetadata(fields MetadataMap) *CIP68Metadata {
	return &CIP68Metadata{
		Metadata: fields,
		Version:  1,
	}
}

// Set adds or updates a field in the metadata.
//
// Example:
//
//	m.Set("ticker", "MTK")
func (c *CIP68Metadata) Set(key string, value any) {
	if c.Metadata == nil {
		c.Metadata = make(MetadataMap)
	}
	c.Metadata[key] = value
}

// Get retrieves a field value by key. Returns nil if not found.
//
// Example:
//
//	name := m.Get("name")
func (c *CIP68Metadata) Get(key string) any {
	return c.Metadata[key]
}

// ToMap returns the metadata as a MetadataMap with version and extra fields,
// suitable for CBOR encoding as a Plutus datum.
//
// The structure is: [metadata_map, version, extra]
// represented here as a MetadataMap for JSON/agent-friendly output.
//
// Example:
//
//	m := nft.ToMap()
func (c *CIP68Metadata) ToMap() MetadataMap {
	v := c.Version
	if v == 0 {
		v = 1
	}
	return MetadataMap{
		"metadata": c.Metadata,
		"version":  fmt.Sprintf("%d", v),
		"extra":    c.Extra,
	}
}

// CIP68RoyaltyMetadata represents a CIP-68 royalty token (label 444) structure.
type CIP68RoyaltyMetadata struct {
	// Rate is the royalty rate as a decimal string, e.g. "0.05" for 5%.
	Rate string
	// Addr is the royalty recipient address.
	Addr string
}

// ToMap returns royalty metadata as a MetadataMap.
//
// Example:
//
//	r := cardanometadata.CIP68RoyaltyMetadata{Rate: "0.05", Addr: "addr1..."}
//	m := r.ToMap()
func (r CIP68RoyaltyMetadata) ToMap() MetadataMap {
	return MetadataMap{
		"rate": r.Rate,
		"addr": r.Addr,
	}
}
