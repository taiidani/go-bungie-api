// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Definitions_Milestones_DestinyMilestoneChallengeActivityDefinition struct {
    // ActivityGraphNodes.
    //
    // If the activity and its challenge is visible on any of these nodes, it will be returned.
    ActivityGraphNodes []Destiny_Definitions_Milestones_DestinyMilestoneChallengeActivityGraphNodeEntry `json:"activityGraphNodes"`

    // ActivityHash.
    //
    // The activity for which this challenge is active.
    ActivityHash uint32 `json:"activityHash"`

    // Challenges.
    //
    // 
    Challenges []Destiny_Definitions_Milestones_DestinyMilestoneChallengeDefinition `json:"challenges"`

    // Phases.
    //
    // Phases related to this activity, if there are any.
    //
    // These will be listed in the order in which they will appear in the actual activity.
    Phases []Destiny_Definitions_Milestones_DestinyMilestoneChallengeActivityPhase `json:"phases"`
}
