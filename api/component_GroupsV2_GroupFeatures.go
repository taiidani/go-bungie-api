// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type GroupsV2_GroupFeatures struct {
    // Capabilities.
    //
    // 
    Capabilities int32 `json:"capabilities"`

    // HostGuidedGamePermissionOverride.
    //
    // Minimum Member Level allowed to host guided games
    //
    // Always Allowed: Founder, Acting Founder, Admin
    //
    // Allowed Overrides: None, Member, Beginner
    //
    // Default is Member for clans, None for groups, although this means nothing for groups.
    HostGuidedGamePermissionOverride int32 `json:"hostGuidedGamePermissionOverride"`

    // InvitePermissionOverride.
    //
    // Minimum Member Level allowed to invite new members to group
    //
    // Always Allowed: Founder, Acting Founder
    //
    // True means admins have this power, false means they don't
    //
    // Default is false for clans, true for groups.
    InvitePermissionOverride bool `json:"invitePermissionOverride"`

    // JoinLevel.
    //
    // Level to join a member at when accepting an invite, application, or joining an open clan
    //
    // Default is Beginner.
    JoinLevel int32 `json:"joinLevel"`

    // MaximumMembers.
    //
    // 
    MaximumMembers int32 `json:"maximumMembers"`

    // MaximumMembershipsOfGroupType.
    //
    // Maximum number of groups of this type a typical membership may join. For example, a user may join about 50 General groups with their Bungie.net account. They may join one clan per Destiny membership.
    MaximumMembershipsOfGroupType int32 `json:"maximumMembershipsOfGroupType"`

    // MembershipTypes.
    //
    // 
    MembershipTypes []int32 `json:"membershipTypes"`

    // UpdateBannerPermissionOverride.
    //
    // Minimum Member Level allowed to update banner
    //
    // Always Allowed: Founder, Acting Founder
    //
    // True means admins have this power, false means they don't
    //
    // Default is false for clans, true for groups.
    UpdateBannerPermissionOverride bool `json:"updateBannerPermissionOverride"`

    // UpdateCulturePermissionOverride.
    //
    // Minimum Member Level allowed to update group culture
    //
    // Always Allowed: Founder, Acting Founder
    //
    // True means admins have this power, false means they don't
    //
    // Default is false for clans, true for groups.
    UpdateCulturePermissionOverride bool `json:"updateCulturePermissionOverride"`
}
