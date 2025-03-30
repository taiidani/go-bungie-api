// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Definitions_DestinyItemTranslationBlockDefinition struct {
    // CustomDyes.
    //
    // 
    CustomDyes []Destiny_DyeReference `json:"customDyes"`

    // DefaultDyes.
    //
    // 
    DefaultDyes []Destiny_DyeReference `json:"defaultDyes"`

    // HasGeometry.
    //
    // 
    HasGeometry bool `json:"hasGeometry"`

    // LockedDyes.
    //
    // 
    LockedDyes []Destiny_DyeReference `json:"lockedDyes"`

    // WeaponPatternHash.
    //
    // 
    WeaponPatternHash uint32 `json:"weaponPatternHash"`

    // WeaponPatternIdentifier.
    //
    // 
    WeaponPatternIdentifier string `json:"weaponPatternIdentifier"`

    // Arrangements.
    //
    // 
    Arrangements []Destiny_Definitions_DestinyGearArtArrangementReference `json:"arrangements"`
}
