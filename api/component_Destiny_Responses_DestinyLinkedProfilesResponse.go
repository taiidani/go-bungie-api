// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Responses_DestinyLinkedProfilesResponse struct {
    // BnetMembership.
    //
    // If the requested membership had a linked Bungie.Net membership ID, this is the basic information about that BNet account.
    //
    // I know, Tetron; I know this is mixing UserServices concerns with DestinyServices concerns. But it's so damn convenient! https://www.youtube.com/watch?v=X5R-bB-gKVI
    BnetMembership any `json:"bnetMembership"`

    // Profiles.
    //
    // Any Destiny account for whom we could successfully pull characters will be returned here, as the Platform-level summary of user data. (no character data, no Destiny account data other than the Membership ID and Type so you can make further queries)
    Profiles []Destiny_Responses_DestinyProfileUserInfoCard `json:"profiles"`

    // ProfilesWithErrors.
    //
    // This is brief summary info for profiles that we believe have valid Destiny info, but who failed to return data for some other reason and thus we know that subsequent calls for their info will also fail.
    ProfilesWithErrors []Destiny_Responses_DestinyErrorProfile `json:"profilesWithErrors"`
}
