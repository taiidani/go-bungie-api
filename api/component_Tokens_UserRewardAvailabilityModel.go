// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Tokens_UserRewardAvailabilityModel struct {
    // IsAvailableForUser.
    //
    // 
    IsAvailableForUser bool `json:"IsAvailableForUser"`

    // IsUnlockedForUser.
    //
    // 
    IsUnlockedForUser bool `json:"IsUnlockedForUser"`

    // AvailabilityModel.
    //
    // 
    AvailabilityModel Tokens_RewardAvailabilityModel `json:"AvailabilityModel"`
}
