// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Definitions_DestinyTalentExclusiveGroup struct {
    // OpposingGroupHashes.
    //
    // A quick reference of Groups whose nodes will be deactivated if any node in this group is activated.
    OpposingGroupHashes []any `json:"opposingGroupHashes"`

    // OpposingNodeHashes.
    //
    // A quick reference of Nodes that will be deactivated if any node in this group is activated, by their Talent Node hashes. (See DestinyTalentNodeDefinition.nodeHash)
    OpposingNodeHashes []any `json:"opposingNodeHashes"`

    // GroupHash.
    //
    // The identifier for this exclusive group. Only guaranteed unique within the talent grid, not globally.
    GroupHash uint32 `json:"groupHash"`

    // LoreHash.
    //
    // If this group has an associated piece of lore to show next to it, this will be the identifier for that DestinyLoreDefinition.
    LoreHash uint32 `json:"loreHash"`

    // NodeHashes.
    //
    // A quick reference of the talent nodes that are part of this group, by their Talent Node hashes. (See DestinyTalentNodeDefinition.nodeHash)
    NodeHashes []any `json:"nodeHashes"`
}
