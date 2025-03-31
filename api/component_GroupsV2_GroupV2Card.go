// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type GroupsV2_GroupV2Card struct {
    // About.
    //
    // 
    About string `json:"about"`

    // AvatarPath.
    //
    // 
    AvatarPath string `json:"avatarPath"`

    // Capabilities.
    //
    // 
    Capabilities int32 `json:"capabilities"`

    // ClanInfo.
    //
    // This contract contains clan-specific group information. It does not include any investment data.
    ClanInfo GroupsV2_GroupV2ClanInfo `json:"clanInfo"`

    // CreationDate.
    //
    // 
    CreationDate string `json:"creationDate"`

    // GroupId.
    //
    // 
    GroupId int64 `json:"groupId"`

    // GroupType.
    //
    // 
    GroupType int32 `json:"groupType"`

    // Locale.
    //
    // 
    Locale string `json:"locale"`

    // MemberCount.
    //
    // 
    MemberCount int32 `json:"memberCount"`

    // MembershipOption.
    //
    // 
    MembershipOption int32 `json:"membershipOption"`

    // Motto.
    //
    // 
    Motto string `json:"motto"`

    // Name.
    //
    // 
    Name string `json:"name"`

    // RemoteGroupId.
    //
    // 
    RemoteGroupId *int64 `json:"remoteGroupId"`

    // Theme.
    //
    // 
    Theme string `json:"theme"`
}
