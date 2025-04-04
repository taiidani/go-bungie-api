// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Components_Records_DestinyRecordComponent struct {
    // CompletedCount.
    //
    // If available, this is the number of times this record has been completed. For example, the number of times a seal title has been gilded.
    CompletedCount *int32 `json:"completedCount"`

    // IntervalObjectives.
    //
    // 
    IntervalObjectives []Destiny_Quests_DestinyObjectiveProgress `json:"intervalObjectives"`

    // IntervalsRedeemedCount.
    //
    // 
    IntervalsRedeemedCount int32 `json:"intervalsRedeemedCount"`

    // Objectives.
    //
    // 
    Objectives []Destiny_Quests_DestinyObjectiveProgress `json:"objectives"`

    // RewardVisibilty.
    //
    // If available, a list that describes which reward rewards should be shown (true) or hidden (false). This property is for regular record rewards, and not for interval objective rewards.
    RewardVisibilty []bool `json:"rewardVisibilty"`

    // State.
    //
    // 
    State int32 `json:"state"`
}
