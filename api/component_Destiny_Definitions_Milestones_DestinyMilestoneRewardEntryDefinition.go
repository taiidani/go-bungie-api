// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Definitions_Milestones_DestinyMilestoneRewardEntryDefinition struct {
    // DisplayProperties.
    //
    // For us to bother returning this info, we should be able to return some kind of information about why these rewards are grouped together. This is ideally that information. Look at how confident I am that this will always remain true.
    DisplayProperties any `json:"displayProperties"`

    // Items.
    //
    // The items you will get as rewards, and how much of it you'll get.
    Items []Destiny_DestinyItemQuantity `json:"items"`

    // Order.
    //
    // If you want to follow BNet's ordering of these rewards, use this number within a given category to order the rewards. Yeah, I know. I feel dirty too.
    Order int32 `json:"order"`

    // RewardEntryHash.
    //
    // The identifier for this reward entry. Runtime data will refer to reward entries by this hash. Only guaranteed unique within the specific Milestone.
    RewardEntryHash uint32 `json:"rewardEntryHash"`

    // RewardEntryIdentifier.
    //
    // The string identifier, if you care about it. Only guaranteed unique within the specific Milestone.
    RewardEntryIdentifier string `json:"rewardEntryIdentifier"`

    // VendorHash.
    //
    // If this reward is redeemed at a Vendor, this is the hash of the Vendor to go to in order to redeem the reward. Use this hash to look up the DestinyVendorDefinition.
    VendorHash *uint32 `json:"vendorHash"`
}
