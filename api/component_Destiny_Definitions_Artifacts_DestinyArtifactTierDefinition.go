// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Definitions_Artifacts_DestinyArtifactTierDefinition struct {
    // DisplayTitle.
    //
    // The human readable title of this tier, if any.
    DisplayTitle string `json:"displayTitle"`

    // Items.
    //
    // The items that can be earned within this tier.
    Items []Destiny_Definitions_Artifacts_DestinyArtifactTierItemDefinition `json:"items"`

    // MinimumUnlockPointsUsedRequirement.
    //
    // The minimum number of "unlock points" that you must have used before you can unlock items from this tier.
    MinimumUnlockPointsUsedRequirement int32 `json:"minimumUnlockPointsUsedRequirement"`

    // ProgressRequirementMessage.
    //
    // A string representing the localized minimum requirement text for this Tier, if any.
    ProgressRequirementMessage string `json:"progressRequirementMessage"`

    // TierHash.
    //
    // An identifier, unique within the Artifact, for this specific tier.
    TierHash uint32 `json:"tierHash"`
}
