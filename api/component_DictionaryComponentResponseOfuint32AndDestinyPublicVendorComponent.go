// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type DictionaryComponentResponseOfuint32AndDestinyPublicVendorComponent struct {
    // Data.
    //
    // 
    Data map[uint32]Destiny_Components_Vendors_DestinyPublicVendorComponent `json:"data"`

    // Disabled.
    //
    // If true, this component is disabled.
    Disabled *bool `json:"disabled"`

    // Privacy.
    //
    // 
    Privacy int32 `json:"privacy"`
}
