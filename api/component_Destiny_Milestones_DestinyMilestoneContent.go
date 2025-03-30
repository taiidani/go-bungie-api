// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Milestones_DestinyMilestoneContent struct {
    // About.
    //
    // The "About this Milestone" text from the Firehose.
    About string `json:"about"`

    // ItemCategories.
    //
    // If DPS has defined items related to this Milestone, they can categorize those items in the Firehose. That data will then be returned as item categories here.
    ItemCategories []Destiny_Milestones_DestinyMilestoneContentItemCategory `json:"itemCategories"`

    // Status.
    //
    // The Current Status of the Milestone, as driven by the Firehose.
    Status string `json:"status"`

    // Tips.
    //
    // A list of tips, provided by the Firehose.
    Tips []string `json:"tips"`
}
