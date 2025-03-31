// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Definitions_Collectibles_DestinyCollectibleDefinition struct {
    // AcquisitionInfo.
    //
    // 
    AcquisitionInfo Destiny_Definitions_Collectibles_DestinyCollectibleAcquisitionBlock `json:"acquisitionInfo"`

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

    // ItemHash.
    //
    // 
    ItemHash uint32 `json:"itemHash"`

    // ParentNodeHashes.
    //
    // A quick reference to presentation nodes that have this node as a child. Presentation nodes can be parented under multiple parents.
    ParentNodeHashes []uint32 `json:"parentNodeHashes"`

    // PresentationInfo.
    //
    // 
    PresentationInfo Destiny_Definitions_Presentation_DestinyPresentationChildBlock `json:"presentationInfo"`

    // PresentationNodeType.
    //
    // 
    PresentationNodeType int32 `json:"presentationNodeType"`

    // Redacted.
    //
    // If this is true, then there is an entity with this identifier/type combination, but BNet is not yet allowed to show it. Sorry!
    Redacted bool `json:"redacted"`

    // Scope.
    //
    // Indicates whether the state of this Collectible is determined on a per-character or on an account-wide basis.
    Scope int32 `json:"scope"`

    // SourceHash.
    //
    // This is a hash identifier we are building on the BNet side in an attempt to let people group collectibles by similar sources.
    //
    // I can't promise that it's going to be 100% accurate, but if the designers were consistent in assigning the same source strings to items with the same sources, it *ought to* be. No promises though.
    //
    // This hash also doesn't relate to an actual definition, just to note: we've got nothing useful other than the source string for this data.
    SourceHash *uint32 `json:"sourceHash"`

    // SourceString.
    //
    // A human readable string for a hint about how to acquire the item.
    SourceString string `json:"sourceString"`

    // StateInfo.
    //
    // 
    StateInfo Destiny_Definitions_Collectibles_DestinyCollectibleStateBlock `json:"stateInfo"`

    // TraitHashes.
    //
    // 
    TraitHashes []uint32 `json:"traitHashes"`

    // TraitIds.
    //
    // 
    TraitIds []string `json:"traitIds"`
}
