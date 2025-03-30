// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type GroupsV2_GroupEditAction struct {
    // Callsign.
    //
    // 
    Callsign string `json:"callsign"`

    // IsPublic.
    //
    // 
    IsPublic bool `json:"isPublic"`

    // AvatarImageIndex.
    //
    // 
    AvatarImageIndex int32 `json:"avatarImageIndex"`

    // IsPublicTopicAdminOnly.
    //
    // 
    IsPublicTopicAdminOnly bool `json:"isPublicTopicAdminOnly"`

    // Tags.
    //
    // 
    Tags string `json:"tags"`

    // DefaultPublicity.
    //
    // 
    DefaultPublicity int32 `json:"defaultPublicity"`

    // Locale.
    //
    // 
    Locale string `json:"locale"`

    // MembershipOption.
    //
    // 
    MembershipOption int32 `json:"membershipOption"`

    // Theme.
    //
    // 
    Theme string `json:"theme"`

    // AllowChat.
    //
    // 
    AllowChat bool `json:"allowChat"`

    // ChatSecurity.
    //
    // 
    ChatSecurity int32 `json:"chatSecurity"`

    // EnableInvitationMessagingForAdmins.
    //
    // 
    EnableInvitationMessagingForAdmins bool `json:"enableInvitationMessagingForAdmins"`

    // Homepage.
    //
    // 
    Homepage int32 `json:"homepage"`

    // Motto.
    //
    // 
    Motto string `json:"motto"`

    // About.
    //
    // 
    About string `json:"about"`

    // Name.
    //
    // 
    Name string `json:"name"`
}
