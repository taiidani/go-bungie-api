// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Forum_ForumRecruitmentDetail struct {
    // ConversationId.
    //
    // 
    ConversationId int64 `json:"conversationId"`

    // MicrophoneRequired.
    //
    // 
    MicrophoneRequired bool `json:"microphoneRequired"`

    // Tone.
    //
    // 
    Tone byte `json:"tone"`

    // PlayerSlotsTotal.
    //
    // 
    PlayerSlotsTotal int32 `json:"playerSlotsTotal"`

    // TopicId.
    //
    // 
    TopicId int64 `json:"topicId"`

    // Fireteam.
    //
    // 
    Fireteam []User_GeneralUser `json:"Fireteam"`

    // KickedPlayerIds.
    //
    // 
    KickedPlayerIds []any `json:"kickedPlayerIds"`

    // Intensity.
    //
    // 
    Intensity byte `json:"intensity"`

    // PlayerSlotsRemaining.
    //
    // 
    PlayerSlotsRemaining int32 `json:"playerSlotsRemaining"`

    // Approved.
    //
    // 
    Approved bool `json:"approved"`
}
