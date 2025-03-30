// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Components_Presentation_DestinyPresentationNodeComponent struct {
    // ProgressValue.
    //
    // How much of the presentation node is considered to be completed so far by the given character/profile.
    ProgressValue int32 `json:"progressValue"`

    // RecordCategoryScore.
    //
    // If available, this is the current score for the record category that this node represents.
    RecordCategoryScore int32 `json:"recordCategoryScore"`

    // State.
    //
    // 
    State int32 `json:"state"`

    // CompletionValue.
    //
    // The value at which the presentation node is considered to be completed.
    CompletionValue int32 `json:"completionValue"`

    // Objective.
    //
    // An optional property: presentation nodes MAY have objectives, which can be used to infer more human readable data about the progress. However, progressValue and completionValue ought to be considered the canonical values for progress on Progression Nodes.
    Objective any `json:"objective"`
}
