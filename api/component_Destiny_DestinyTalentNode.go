// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_DestinyTalentNode struct {
    // IsActivated.
    //
    // If true, the node is activated: it's current step then provides its benefits.
    IsActivated bool `json:"isActivated"`

    // NodeIndex.
    //
    // The index of the Talent Node being referred to (an index into DestinyTalentGridDefinition.nodes[]). CONTENT VERSION DEPENDENT.
    NodeIndex int32 `json:"nodeIndex"`

    // State.
    //
    // An DestinyTalentNodeState enum value indicating the node's state: whether it can be activated or swapped, and why not if neither can be performed.
    State int32 `json:"state"`

    // StepIndex.
    //
    // The currently relevant Step for the node. It is this step that has rendering data for the node and the benefits that are provided if the node is activated. (the actual rules for benefits provided are extremely complicated in theory, but with how Talent Grids are being used in Destiny 2 you don't have to worry about a lot of those old Destiny 1 rules.) This is an index into: DestinyTalentGridDefinition.nodes[nodeIndex].steps[stepIndex]
    StepIndex int32 `json:"stepIndex"`

    // ActivationGridLevel.
    //
    // The progression level required on the Talent Grid in order to be able to activate this talent node. Talent Grids have their own Progression - similar to Character Level, but in this case it is experience related to the item itself.
    ActivationGridLevel int32 `json:"activationGridLevel"`

    // Hidden.
    //
    // Whether or not the talent node is actually visible in the game's UI. Whether you want to show it in your own UI is up to you! I'm not gonna tell you who to sock it to.
    Hidden bool `json:"hidden"`

    // NodeStatsBlock.
    //
    // This property has some history. A talent grid can provide stats on both the item it's related to and the character equipping the item. This returns data about those stat bonuses.
    NodeStatsBlock any `json:"nodeStatsBlock"`

    // ProgressPercent.
    //
    // If you want to show a progress bar or circle for how close this talent node is to being activate-able, this is the percentage to show. It follows the node's underlying rules about when the progress bar should first show up, and when it should be filled.
    ProgressPercent float32 `json:"progressPercent"`

    // MaterialsToUpgrade.
    //
    // If the node has material requirements to be activated, this is the list of those requirements.
    MaterialsToUpgrade []Destiny_Definitions_DestinyMaterialRequirement `json:"materialsToUpgrade"`

    // NodeHash.
    //
    // The hash of the Talent Node being referred to (in DestinyTalentGridDefinition.nodes). Deceptively CONTENT VERSION DEPENDENT. We have no guarantee of the hash's immutability between content versions.
    NodeHash uint32 `json:"nodeHash"`
}
