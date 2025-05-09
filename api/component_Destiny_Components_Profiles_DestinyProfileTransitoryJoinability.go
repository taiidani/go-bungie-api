// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Components_Profiles_DestinyProfileTransitoryJoinability struct {
    // ClosedReasons.
    //
    // Reasons why a person can't join this person's fireteam.
    ClosedReasons int32 `json:"closedReasons"`

    // OpenSlots.
    //
    // The number of slots still available on this person's fireteam.
    OpenSlots int32 `json:"openSlots"`

    // PrivacySetting.
    //
    // Who the person is currently allowing invites from.
    PrivacySetting int32 `json:"privacySetting"`
}
