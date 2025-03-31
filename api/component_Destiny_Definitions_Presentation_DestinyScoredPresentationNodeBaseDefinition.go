// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Definitions_Presentation_DestinyScoredPresentationNodeBaseDefinition struct {
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

    // MaxCategoryRecordScore.
    //
    // 
    MaxCategoryRecordScore int32 `json:"maxCategoryRecordScore"`

    // ParentNodeHashes.
    //
    // A quick reference to presentation nodes that have this node as a child. Presentation nodes can be parented under multiple parents.
    ParentNodeHashes []uint32 `json:"parentNodeHashes"`

    // PresentationNodeType.
    //
    // 
    PresentationNodeType int32 `json:"presentationNodeType"`

    // Redacted.
    //
    // If this is true, then there is an entity with this identifier/type combination, but BNet is not yet allowed to show it. Sorry!
    Redacted bool `json:"redacted"`

    // TraitHashes.
    //
    // 
    TraitHashes []uint32 `json:"traitHashes"`

    // TraitIds.
    //
    // 
    TraitIds []string `json:"traitIds"`
}
