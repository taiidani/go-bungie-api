// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Forum_PollResponse struct {
    // Results.
    //
    // 
    Results []Forum_PollResult `json:"results"`

    // TopicId.
    //
    // 
    TopicId int64 `json:"topicId"`

    // TotalVotes.
    //
    // 
    TotalVotes int32 `json:"totalVotes"`
}
