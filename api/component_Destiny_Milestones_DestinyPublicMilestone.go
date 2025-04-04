// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Milestones_DestinyPublicMilestone struct {
    // Activities.
    //
    // 
    Activities []Destiny_Milestones_DestinyPublicMilestoneChallengeActivity `json:"activities"`

    // AvailableQuests.
    //
    // A milestone not need have even a single quest, but if there are active quests they will be returned here.
    AvailableQuests []Destiny_Milestones_DestinyPublicMilestoneQuest `json:"availableQuests"`

    // EndDate.
    //
    // If known, this is the date when the Milestone will expire/recycle/end.
    EndDate *string `json:"endDate"`

    // MilestoneHash.
    //
    // The hash identifier for the milestone. Use it to look up the DestinyMilestoneDefinition for static data about the Milestone.
    MilestoneHash uint32 `json:"milestoneHash"`

    // Order.
    //
    // Used for ordering milestones in a display to match how we order them in BNet. May pull from static data, or possibly in the future from dynamic information.
    Order int32 `json:"order"`

    // StartDate.
    //
    // If known, this is the date when the Milestone started/became active.
    StartDate *string `json:"startDate"`

    // VendorHashes.
    //
    // Sometimes milestones - or activities active in milestones - will have relevant vendors. These are the vendors that are currently relevant.
    //
    // Deprecated, already, for the sake of the new "vendors" property that has more data. What was I thinking.
    VendorHashes []uint32 `json:"vendorHashes"`

    // Vendors.
    //
    // This is why we can't have nice things. This is the ordered list of vendors to be shown that relate to this milestone, potentially along with other interesting data.
    Vendors []Destiny_Milestones_DestinyPublicMilestoneVendor `json:"vendors"`
}
