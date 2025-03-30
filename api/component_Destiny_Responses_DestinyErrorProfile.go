// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Responses_DestinyErrorProfile struct {
    // ErrorCode.
    //
    // The error that we encountered. You should be able to look up localized text to show to the user for these failures.
    ErrorCode int32 `json:"errorCode"`

    // InfoCard.
    //
    // Basic info about the account that failed. Don't expect anything other than membership ID, Membership Type, and displayName to be populated.
    InfoCard any `json:"infoCard"`
}
