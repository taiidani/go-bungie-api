// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Definitions_DestinyBubbleDefinition struct {
    // DisplayProperties.
    //
    // The display properties of this bubble, so you don't have to look them up in a separate list anymore.
    DisplayProperties any `json:"displayProperties"`

    // Hash.
    //
    // The identifier for the bubble: only guaranteed to be unique within the Destination.
    Hash uint32 `json:"hash"`
}
