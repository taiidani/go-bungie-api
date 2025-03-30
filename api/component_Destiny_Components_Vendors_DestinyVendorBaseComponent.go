// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Components_Vendors_DestinyVendorBaseComponent struct {
    // VendorHash.
    //
    // The unique identifier for the vendor. Use it to look up their DestinyVendorDefinition.
    VendorHash uint32 `json:"vendorHash"`

    // Enabled.
    //
    // If True, the Vendor is currently accessible. 
    //
    // If False, they may not actually be visible in the world at the moment.
    Enabled bool `json:"enabled"`

    // NextRefreshDate.
    //
    // The date when this vendor's inventory will next rotate/refresh.
    //
    // Note that this is distinct from the date ranges that the vendor is visible/available in-game: this field indicates the specific time when the vendor's available items refresh and rotate, regardless of whether the vendor is actually available at that time. Unfortunately, these two values may be (and are, for the case of important vendors like Xur) different.
    //
    // Issue https://github.com/Bungie-net/api/issues/353 is tracking a fix to start providing visibility date ranges where possible in addition to this refresh date, so that all important dates for vendors are available for use.
    NextRefreshDate string `json:"nextRefreshDate"`
}
