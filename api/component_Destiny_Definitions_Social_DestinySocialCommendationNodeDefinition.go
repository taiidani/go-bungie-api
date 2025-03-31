// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Definitions_Social_DestinySocialCommendationNodeDefinition struct {
    // ChildCommendationHashes.
    //
    // A list of hashes that map to child commendations.
    ChildCommendationHashes []uint32 `json:"childCommendationHashes"`

    // ChildCommendationNodeHashes.
    //
    // A list of hashes that map to child commendation nodes. Only the root commendations node is expected to have child nodes.
    ChildCommendationNodeHashes []uint32 `json:"childCommendationNodeHashes"`

    // Color.
    //
    // The color associated with this group of commendations.
    Color any `json:"color"`

    // DisplayProperties.
    //
    // Many Destiny*Definition contracts - the "first order" entities of Destiny that have their own tables in the Manifest Database - also have displayable information. This is the base class for that display information.
    DisplayProperties Destiny_Definitions_Common_DestinyDisplayPropertiesDefinition `json:"displayProperties"`

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

    // ParentCommendationNodeHash.
    //
    // 
    ParentCommendationNodeHash uint32 `json:"parentCommendationNodeHash"`

    // Redacted.
    //
    // If this is true, then there is an entity with this identifier/type combination, but BNet is not yet allowed to show it. Sorry!
    Redacted bool `json:"redacted"`

    // TintedIcon.
    //
    // A version of the displayProperties icon tinted with the color of this node.
    TintedIcon string `json:"tintedIcon"`
}
