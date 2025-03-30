// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type SingleComponentResponseOfDestinyVendorReceiptsComponent struct {
    // Data.
    //
    // For now, this isn't used for much: it's a record of the recent refundable purchases that the user has made. In the future, it could be used for providing refunds/buyback via the API. Wouldn't that be fun?
    Data Destiny_Entities_Profiles_DestinyVendorReceiptsComponent `json:"data"`

    // Disabled.
    //
    // If true, this component is disabled.
    Disabled bool `json:"disabled"`

    // Privacy.
    //
    // 
    Privacy int32 `json:"privacy"`
}
