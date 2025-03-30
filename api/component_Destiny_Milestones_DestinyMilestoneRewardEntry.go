// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Milestones_DestinyMilestoneRewardEntry struct {
    // RewardEntryHash.
    //
    // The identifier for the reward entry in question. It is important to look up the related DestinyMilestoneRewardEntryDefinition to get the static details about the reward, which you can do by looking up the milestone's DestinyMilestoneDefinition and examining the DestinyMilestoneDefinition.rewards[rewardCategoryHash].rewardEntries[rewardEntryHash] data.
    RewardEntryHash uint32 `json:"rewardEntryHash"`

    // Earned.
    //
    // If TRUE, the player has earned this reward.
    Earned bool `json:"earned"`

    // Redeemed.
    //
    // If TRUE, the player has redeemed/picked up/obtained this reward. Feel free to alias this to "gotTheShinyBauble" in your own codebase.
    Redeemed bool `json:"redeemed"`
}
