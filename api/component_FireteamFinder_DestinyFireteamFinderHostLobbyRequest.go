// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type FireteamFinder_DestinyFireteamFinderHostLobbyRequest struct {
    // ActivityGraphHash.
    //
    // 
    ActivityGraphHash uint32 `json:"activityGraphHash"`

    // ActivityHash.
    //
    // 
    ActivityHash uint32 `json:"activityHash"`

    // ClanId.
    //
    // 
    ClanId int64 `json:"clanId"`

    // ListingValues.
    //
    // 
    ListingValues []FireteamFinder_DestinyFireteamFinderListingValue `json:"listingValues"`

    // MaxPlayerCount.
    //
    // 
    MaxPlayerCount int32 `json:"maxPlayerCount"`

    // OnlinePlayersOnly.
    //
    // 
    OnlinePlayersOnly bool `json:"onlinePlayersOnly"`

    // PrivacyScope.
    //
    // 
    PrivacyScope int32 `json:"privacyScope"`

    // ScheduledDateTime.
    //
    // 
    ScheduledDateTime string `json:"scheduledDateTime"`
}
