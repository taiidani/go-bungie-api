// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type FireteamFinder_DestinyFireteamFinderLobbyResponse struct {
    // ListingId.
    //
    // 
    ListingId int64 `json:"listingId"`

    // LobbyId.
    //
    // 
    LobbyId int64 `json:"lobbyId"`

    // Owner.
    //
    // 
    Owner FireteamFinder_DestinyFireteamFinderPlayerId `json:"owner"`

    // Players.
    //
    // 
    Players []FireteamFinder_DestinyFireteamFinderLobbyPlayer `json:"players"`

    // Revision.
    //
    // 
    Revision int32 `json:"revision"`

    // Settings.
    //
    // 
    Settings FireteamFinder_DestinyFireteamFinderLobbySettings `json:"settings"`

    // State.
    //
    // 
    State int32 `json:"state"`

    // CreatedDateTime.
    //
    // 
    CreatedDateTime string `json:"createdDateTime"`
}
