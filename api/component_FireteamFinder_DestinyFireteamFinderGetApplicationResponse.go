// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type FireteamFinder_DestinyFireteamFinderGetApplicationResponse struct {
    // ApplicantSet.
    //
    // 
    ApplicantSet FireteamFinder_DestinyFireteamFinderApplicantSet `json:"applicantSet"`

    // ApplicationId.
    //
    // 
    ApplicationId int64 `json:"applicationId"`

    // ApplicationType.
    //
    // 
    ApplicationType int32 `json:"applicationType"`

    // CreatedDateTime.
    //
    // 
    CreatedDateTime string `json:"createdDateTime"`

    // ListingId.
    //
    // 
    ListingId int64 `json:"listingId"`

    // ReferralToken.
    //
    // 
    ReferralToken int64 `json:"referralToken"`

    // Revision.
    //
    // 
    Revision int32 `json:"revision"`

    // State.
    //
    // 
    State int32 `json:"state"`

    // SubmitterId.
    //
    // 
    SubmitterId FireteamFinder_DestinyFireteamFinderPlayerId `json:"submitterId"`
}
