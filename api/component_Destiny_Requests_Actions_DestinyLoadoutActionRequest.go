// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Requests_Actions_DestinyLoadoutActionRequest struct {
    // CharacterId.
    //
    // 
    CharacterId int64 `json:"characterId"`

    // LoadoutIndex.
    //
    // The index of the loadout for this action request.
    LoadoutIndex int32 `json:"loadoutIndex"`

    // MembershipType.
    //
    // 
    MembershipType int32 `json:"membershipType"`
}
