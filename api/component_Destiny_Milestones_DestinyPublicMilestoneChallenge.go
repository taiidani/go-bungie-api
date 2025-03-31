// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Milestones_DestinyPublicMilestoneChallenge struct {
    // ActivityHash.
    //
    // IF the Objective is related to a specific Activity, this will be that activity's hash. Use it to look up the DestinyActivityDefinition for additional data to show.
    ActivityHash *uint32 `json:"activityHash"`

    // ObjectiveHash.
    //
    // The objective for the Challenge, which should have human-readable data about what needs to be done to accomplish the objective. Use this hash to look up the DestinyObjectiveDefinition.
    ObjectiveHash uint32 `json:"objectiveHash"`
}
