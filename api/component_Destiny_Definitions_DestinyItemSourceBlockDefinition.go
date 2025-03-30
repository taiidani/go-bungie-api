// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Definitions_DestinyItemSourceBlockDefinition struct {
    // SourceHashes.
    //
    // The list of hash identifiers for Reward Sources that hint where the item can be found (DestinyRewardSourceDefinition).
    SourceHashes []any `json:"sourceHashes"`

    // Sources.
    //
    // A collection of details about the stats that were computed for the ways we found that the item could be spawned.
    Sources []Destiny_Definitions_Sources_DestinyItemSourceDefinition `json:"sources"`

    // VendorSources.
    //
    // A denormalized reference back to vendors that potentially sell this item.
    VendorSources []Destiny_Definitions_DestinyItemVendorSourceReference `json:"vendorSources"`

    // Exclusive.
    //
    // If we found that this item is exclusive to a specific platform, this will be set to the BungieMembershipType enumeration that matches that platform.
    Exclusive int32 `json:"exclusive"`
}
