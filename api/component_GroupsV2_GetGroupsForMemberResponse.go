// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type GroupsV2_GetGroupsForMemberResponse struct {
    // ReplacementContinuationToken.
    //
    // 
    ReplacementContinuationToken string `json:"replacementContinuationToken"`

    // Results.
    //
    // 
    Results []GroupsV2_GroupMembership `json:"results"`

    // TotalResults.
    //
    // 
    TotalResults int32 `json:"totalResults"`

    // UseTotalResults.
    //
    // If useTotalResults is true, then totalResults represents an accurate count.
    //
    // If False, it does not, and may be estimated/only the size of the current page.
    //
    // Either way, you should probably always only trust hasMore.
    //
    // This is a long-held historical throwback to when we used to do paging with known total results. Those queries toasted our database, and we were left to hastily alter our endpoints and create backward- compatible shims, of which useTotalResults is one.
    UseTotalResults bool `json:"useTotalResults"`

    // AreAllMembershipsInactive.
    //
    // A convenience property that indicates if every membership this user has that is a part of this group are part of an account that is considered inactive - for example, overridden accounts in Cross Save.
    //
    //  The key is the Group ID for the group being checked, and the value is true if the users' memberships for that group are all inactive.
    AreAllMembershipsInactive any `json:"areAllMembershipsInactive"`

    // HasMore.
    //
    // 
    HasMore bool `json:"hasMore"`

    // Query.
    //
    // 
    Query Queries_PagedQuery `json:"query"`
}
