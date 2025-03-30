// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Definitions_Director_DestinyActivityGraphNodeDefinition struct {
    // OverrideDisplay.
    //
    // The node *may* have display properties that override the active Activity's display properties.
    OverrideDisplay any `json:"overrideDisplay"`

    // Position.
    //
    // The position on the map for this node.
    Position any `json:"position"`

    // States.
    //
    // Represents possible states that the graph node can be in. These are combined with some checking that happens in the game client and server to determine which state is actually active at any given time.
    States []Destiny_Definitions_Director_DestinyActivityGraphNodeStateEntry `json:"states"`

    // Activities.
    //
    // The node may have various possible activities that could be active for it, however only one may be active at a time. See the DestinyActivityGraphNodeActivityDefinition for details.
    Activities []Destiny_Definitions_Director_DestinyActivityGraphNodeActivityDefinition `json:"activities"`

    // FeaturingStates.
    //
    // The node may have various visual accents placed on it, or styles applied. These are the list of possible styles that the Node can have. The game iterates through each, looking for the first one that passes a check of the required game/character/account state in order to show that style, and then renders the node in that style.
    FeaturingStates []Destiny_Definitions_Director_DestinyActivityGraphNodeFeaturingStateDefinition `json:"featuringStates"`

    // NodeId.
    //
    // An identifier for the Activity Graph Node, only guaranteed to be unique within its parent Activity Graph.
    NodeId uint32 `json:"nodeId"`
}
