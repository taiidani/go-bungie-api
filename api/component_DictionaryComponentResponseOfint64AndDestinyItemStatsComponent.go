// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type DictionaryComponentResponseOfint64AndDestinyItemStatsComponent struct {
    // Data.
    //
    // 
    Data map[int64]Destiny_Entities_Items_DestinyItemStatsComponent `json:"data"`

    // Disabled.
    //
    // If true, this component is disabled.
    Disabled *bool `json:"disabled"`

    // Privacy.
    //
    // 
    Privacy int32 `json:"privacy"`
}
