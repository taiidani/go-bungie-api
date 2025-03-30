// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type SingleComponentResponseOfDestinyVendorGroupComponent struct {
    // Privacy.
    //
    // 
    Privacy int32 `json:"privacy"`

    // Data.
    //
    // This component returns references to all of the Vendors in the response, grouped by categorizations that Bungie has deemed to be interesting, in the order in which both the groups and the vendors within that group should be rendered.
    Data Destiny_Components_Vendors_DestinyVendorGroupComponent `json:"data"`

    // Disabled.
    //
    // If true, this component is disabled.
    Disabled bool `json:"disabled"`
}
