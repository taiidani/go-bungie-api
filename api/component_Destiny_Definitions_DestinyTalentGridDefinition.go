// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Definitions_DestinyTalentGridDefinition struct {
    // ExclusiveSets.
    //
    // Talent Nodes can exist in "exclusive sets": these are sets of nodes in which only a single node in the set can be activated at any given time. Activating a node in this set will automatically deactivate the other nodes in the set (referred to as a "Swap").
    //
    // If a node in the exclusive set has already been activated, the game will not charge you materials to activate another node in the set, even if you have never activated it before, because you already paid the cost to activate one node in the set.
    //
    // Not to be confused with Exclusive Groups. (how the heck do we NOT get confused by that? Jeez) See the groups property for information about that only-tangentially-related concept.
    ExclusiveSets []Destiny_Definitions_DestinyTalentNodeExclusiveSetDefinition `json:"exclusiveSets"`

    // GridLevelPerColumn.
    //
    // The meaning of this has been lost in the sands of time: it still exists as a property, but appears to be unused in the modern UI of talent grids. It used to imply that each visual "column" of talent nodes required identical progression levels in order to be activated. Returning this value in case it is still useful to someone? Perhaps it's just a bit of interesting history.
    GridLevelPerColumn int32 `json:"gridLevelPerColumn"`

    // Groups.
    //
    // Talent Nodes can have "Exclusive Groups". These are not to be confused with Exclusive Sets (see exclusiveSets property).
    //
    // Look at the definition of DestinyTalentExclusiveGroup for more information and how they work. These groups are keyed by the "groupHash" from DestinyTalentExclusiveGroup.
    Groups map[uint32]Destiny_Definitions_DestinyTalentExclusiveGroup `json:"groups"`

    // Hash.
    //
    // The unique identifier for this entity. Guaranteed to be unique for the type of entity, but not globally.
    //
    // When entities refer to each other in Destiny content, it is this hash that they are referring to.
    Hash uint32 `json:"hash"`

    // IndependentNodeIndexes.
    //
    // This is a quick reference to the indexes of nodes that are not part of exclusive sets. Handy for knowing which talent nodes can only be activated directly, rather than via swapping.
    IndependentNodeIndexes []int32 `json:"independentNodeIndexes"`

    // Index.
    //
    // The index of the entity as it was found in the investment tables.
    Index int32 `json:"index"`

    // MaxGridLevel.
    //
    // The maximum possible level of the Talent Grid: at this level, any nodes are allowed to be activated.
    MaxGridLevel int32 `json:"maxGridLevel"`

    // NodeCategories.
    //
    // BNet wants to show talent nodes grouped by similar purpose with localized titles. This is the ordered list of those categories: if you want to show nodes by category, you can iterate over this list, render the displayProperties for the category as the title, and then iterate over the talent nodes referenced by the category to show the related nodes.
    //
    // Note that this is different from Exclusive Groups or Sets, because these categories also incorporate "Independent" nodes that belong to neither sets nor groups. These are purely for visual grouping of nodes rather than functional grouping.
    NodeCategories []Destiny_Definitions_DestinyTalentNodeCategory `json:"nodeCategories"`

    // Nodes.
    //
    // The list of Talent Nodes on the Grid (recall that Nodes themselves are really just locations in the UI to show whatever their current Step is. You will only know the current step for a node by retrieving instanced data through platform calls to the API that return DestinyItemTalentGridComponent).
    Nodes []Destiny_Definitions_DestinyTalentNodeDefinition `json:"nodes"`

    // ProgressionHash.
    //
    // The hash identifier of the Progression (DestinyProgressionDefinition) that drives whether and when Talent Nodes can be activated on the Grid. Items will have instances of this Progression, and will gain experience that will eventually cause the grid to increase in level. As the grid's level increases, it will cross the threshold where nodes can be activated. See DestinyTalentGridStepDefinition's activation requirements for more information.
    ProgressionHash uint32 `json:"progressionHash"`

    // Redacted.
    //
    // If this is true, then there is an entity with this identifier/type combination, but BNet is not yet allowed to show it. Sorry!
    Redacted bool `json:"redacted"`
}
