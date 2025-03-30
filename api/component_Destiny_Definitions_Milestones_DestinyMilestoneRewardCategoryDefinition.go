// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Definitions_Milestones_DestinyMilestoneRewardCategoryDefinition struct {
    // DisplayProperties.
    //
    // Hopefully this is obvious by now.
    DisplayProperties any `json:"displayProperties"`

    // Order.
    //
    // If you want to use BNet's recommended order for rendering categories programmatically, use this value and compare it to other categories to determine the order in which they should be rendered. I don't feel great about putting this here, I won't lie.
    Order int32 `json:"order"`

    // RewardEntries.
    //
    // If this milestone can provide rewards, this will define the sets of rewards that can be earned, the conditions under which they can be acquired, internal data that we'll use at runtime to determine whether you've already earned or redeemed this set of rewards, and the category that this reward should be placed under.
    RewardEntries any `json:"rewardEntries"`

    // CategoryHash.
    //
    // Identifies the reward category. Only guaranteed unique within this specific component!
    CategoryHash uint32 `json:"categoryHash"`

    // CategoryIdentifier.
    //
    // The string identifier for the category, if you want to use it for some end. Guaranteed unique within the specific component.
    CategoryIdentifier string `json:"categoryIdentifier"`
}
