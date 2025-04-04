// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Definitions_Artifacts_DestinyArtifactDefinition struct {
    // DisplayProperties.
    //
    // Any basic display info we know about the Artifact. Currently sourced from a related inventory item, but the source of this data is subject to change.
    DisplayProperties any `json:"displayProperties"`

    // Hash.
    //
    // The unique identifier for this entity. Guaranteed to be unique for the type of entity, but not globally.
    //
    // When entities refer to each other in Destiny content, it is this hash that they are referring to.
    Hash uint32 `json:"hash"`

    // Index.
    //
    // The index of the entity as it was found in the investment tables.
    Index int32 `json:"index"`

    // Redacted.
    //
    // If this is true, then there is an entity with this identifier/type combination, but BNet is not yet allowed to show it. Sorry!
    Redacted bool `json:"redacted"`

    // Tiers.
    //
    // Any Tier/Rank data related to this artifact, listed in display order.  Currently sourced from a Vendor, but this source is subject to change.
    Tiers []Destiny_Definitions_Artifacts_DestinyArtifactTierDefinition `json:"tiers"`

    // TranslationBlock.
    //
    // Any Geometry/3D info we know about the Artifact. Currently sourced from a related inventory item's gearset information, but the source of this data is subject to change.
    TranslationBlock any `json:"translationBlock"`
}
