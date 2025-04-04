// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Definitions_DestinyItemSetBlockEntryDefinition struct {
    // ItemHash.
    //
    // This is the hash identifier for a DestinyInventoryItemDefinition representing this quest step.
    ItemHash uint32 `json:"itemHash"`

    // TrackingValue.
    //
    // Used for tracking which step a user reached. These values will be populated in the user's internal state, which we expose externally as a more usable DestinyQuestStatus object. If this item has been obtained, this value will be set in trackingUnlockValueHash.
    TrackingValue int32 `json:"trackingValue"`
}
