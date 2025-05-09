// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type SingleComponentResponseOfDestinyCharacterProgressionComponent struct {
    // Data.
    //
    // This component returns anything that could be considered "Progression" on a user: data where the user is gaining levels, reputation, completions, rewards, etc...
    Data Destiny_Entities_Characters_DestinyCharacterProgressionComponent `json:"data"`

    // Disabled.
    //
    // If true, this component is disabled.
    Disabled *bool `json:"disabled"`

    // Privacy.
    //
    // 
    Privacy int32 `json:"privacy"`
}
