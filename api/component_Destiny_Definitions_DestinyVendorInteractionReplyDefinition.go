// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Definitions_DestinyVendorInteractionReplyDefinition struct {
    // ItemRewardsSelection.
    //
    // The rewards granted upon responding to the vendor.
    ItemRewardsSelection int32 `json:"itemRewardsSelection"`

    // Reply.
    //
    // The localized text for the reply.
    Reply string `json:"reply"`

    // ReplyType.
    //
    // An enum indicating the type of reply being made.
    ReplyType int32 `json:"replyType"`
}
