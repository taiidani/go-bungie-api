// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type User_UserSearchResponseDetail struct {
    // BungieGlobalDisplayName.
    //
    // 
    BungieGlobalDisplayName string `json:"bungieGlobalDisplayName"`

    // BungieGlobalDisplayNameCode.
    //
    // 
    BungieGlobalDisplayNameCode int16 `json:"bungieGlobalDisplayNameCode"`

    // BungieNetMembershipId.
    //
    // 
    BungieNetMembershipId int64 `json:"bungieNetMembershipId"`

    // DestinyMemberships.
    //
    // 
    DestinyMemberships []User_UserInfoCard `json:"destinyMemberships"`
}
