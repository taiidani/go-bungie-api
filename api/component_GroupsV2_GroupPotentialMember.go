// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type GroupsV2_GroupPotentialMember struct {
    // DestinyUserInfo.
    //
    // 
    DestinyUserInfo any `json:"destinyUserInfo"`

    // GroupId.
    //
    // 
    GroupId int64 `json:"groupId"`

    // JoinDate.
    //
    // 
    JoinDate string `json:"joinDate"`

    // PotentialStatus.
    //
    // 
    PotentialStatus int32 `json:"potentialStatus"`

    // BungieNetUserInfo.
    //
    // This contract supplies basic information commonly used to display a minimal amount of information about a user. Take care to not add more properties here unless the property applies in all (or at least the majority) of the situations where UserInfoCard is used. Avoid adding game specific or platform specific details here. In cases where UserInfoCard is a subset of the data needed in a contract, use UserInfoCard as a property of other contracts.
    BungieNetUserInfo any `json:"bungieNetUserInfo"`
}
