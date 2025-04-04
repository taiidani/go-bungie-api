// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Milestones_DestinyMilestoneActivityCompletionStatus struct {
    // Completed.
    //
    // If the activity has been "completed", that information will be returned here.
    Completed bool `json:"completed"`

    // Phases.
    //
    // If the Activity has discrete "phases" that we can track, that info will be here. Otherwise, this value will be NULL. Note that this is a list and not a dictionary: the order implies the ascending order of phases or progression in this activity.
    Phases []Destiny_Milestones_DestinyMilestoneActivityPhase `json:"phases"`
}
