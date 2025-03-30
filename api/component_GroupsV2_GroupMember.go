// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type GroupsV2_GroupMember struct {
    // GroupId.
    //
    // 
    GroupId int64 `json:"groupId"`

    // IsOnline.
    //
    // 
    IsOnline bool `json:"isOnline"`

    // JoinDate.
    //
    // 
    JoinDate string `json:"joinDate"`

    // LastOnlineStatusChange.
    //
    // 
    LastOnlineStatusChange int64 `json:"lastOnlineStatusChange"`

    // MemberType.
    //
    // 
    MemberType int32 `json:"memberType"`

    // BungieNetUserInfo.
    //
    // This contract supplies basic information commonly used to display a minimal amount of information about a user. Take care to not add more properties here unless the property applies in all (or at least the majority) of the situations where UserInfoCard is used. Avoid adding game specific or platform specific details here. In cases where UserInfoCard is a subset of the data needed in a contract, use UserInfoCard as a property of other contracts.
    BungieNetUserInfo User_UserInfoCard `json:"bungieNetUserInfo"`

    // DestinyUserInfo.
    //
    // 
    DestinyUserInfo GroupsV2_GroupUserInfoCard `json:"destinyUserInfo"`
}
