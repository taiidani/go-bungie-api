// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type DictionaryComponentResponseOfuint32AndPublicDestinyVendorSaleItemSetComponent struct {
    // Data.
    //
    // 
    Data map[uint32]Destiny_Responses_PublicDestinyVendorSaleItemSetComponent `json:"data"`

    // Disabled.
    //
    // If true, this component is disabled.
    Disabled *bool `json:"disabled"`

    // Privacy.
    //
    // 
    Privacy int32 `json:"privacy"`
}
