// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type DictionaryComponentResponseOfint32AndDestinyItemInstanceComponent struct {
    // Data.
    //
    // 
    Data map[int32]Destiny_Entities_Items_DestinyItemInstanceComponent `json:"data"`

    // Disabled.
    //
    // If true, this component is disabled.
    Disabled *bool `json:"disabled"`

    // Privacy.
    //
    // 
    Privacy int32 `json:"privacy"`
}
