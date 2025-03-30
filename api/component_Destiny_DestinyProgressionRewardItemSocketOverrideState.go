// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_DestinyProgressionRewardItemSocketOverrideState struct {
    // ItemState.
    //
    // Information about the item state, specifically deepsight if there is any data for it
    ItemState int32 `json:"itemState"`

    // RewardItemStats.
    //
    // Information about the computed stats from socket and plug overrides for this progression, if there is any data for it.
    RewardItemStats any `json:"rewardItemStats"`
}
