// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type GroupsV2_GroupEditAction struct {
    // IsPublicTopicAdminOnly.
    //
    // 
    IsPublicTopicAdminOnly bool `json:"isPublicTopicAdminOnly"`

    // Theme.
    //
    // 
    Theme string `json:"theme"`

    // AllowChat.
    //
    // 
    AllowChat bool `json:"allowChat"`

    // MembershipOption.
    //
    // 
    MembershipOption int32 `json:"membershipOption"`

    // Tags.
    //
    // 
    Tags string `json:"tags"`

    // Homepage.
    //
    // 
    Homepage int32 `json:"homepage"`

    // EnableInvitationMessagingForAdmins.
    //
    // 
    EnableInvitationMessagingForAdmins bool `json:"enableInvitationMessagingForAdmins"`

    // Locale.
    //
    // 
    Locale string `json:"locale"`

    // About.
    //
    // 
    About string `json:"about"`

    // AvatarImageIndex.
    //
    // 
    AvatarImageIndex int32 `json:"avatarImageIndex"`

    // DefaultPublicity.
    //
    // 
    DefaultPublicity int32 `json:"defaultPublicity"`

    // Name.
    //
    // 
    Name string `json:"name"`

    // Callsign.
    //
    // 
    Callsign string `json:"callsign"`

    // Motto.
    //
    // 
    Motto string `json:"motto"`

    // IsPublic.
    //
    // 
    IsPublic bool `json:"isPublic"`

    // ChatSecurity.
    //
    // 
    ChatSecurity int32 `json:"chatSecurity"`
}
