// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Definitions_Director_DestinyActivityGraphNodeActivityDefinition struct {
    // ActivityHash.
    //
    // The activity that will be activated if the user clicks on this node. Controls all activity-related information displayed on the node if it is active (the text shown in the tooltip etc)
    ActivityHash uint32 `json:"activityHash"`

    // NodeActivityId.
    //
    // An identifier for this node activity. It is only guaranteed to be unique within the Activity Graph.
    NodeActivityId uint32 `json:"nodeActivityId"`
}
