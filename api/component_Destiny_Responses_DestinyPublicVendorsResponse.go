// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Responses_DestinyPublicVendorsResponse struct {
    // Vendors.
    //
    // The base properties of the vendor. These are keyed by the Vendor Hash, so you will get one Vendor Component per vendor returned.
    //
    // COMPONENT TYPE: Vendors
    Vendors any `json:"vendors"`

    // Categories.
    //
    // Categories that the vendor has available, and references to the sales therein. These are keyed by the Vendor Hash, so you will get one Categories Component per vendor returned.
    //
    // COMPONENT TYPE: VendorCategories
    Categories any `json:"categories"`

    // Sales.
    //
    // Sales, keyed by the vendorItemIndex of the item being sold. These are keyed by the Vendor Hash, so you will get one Sale Item Set Component per vendor returned.
    //
    // Note that within the Sale Item Set component, the sales are themselves keyed by the vendorSaleIndex, so you can relate it to the corrent sale item definition within the Vendor's definition.
    //
    // COMPONENT TYPE: VendorSales
    Sales any `json:"sales"`

    // StringVariables.
    //
    // A set of string variable values by hash for a public vendors context.
    //
    // COMPONENT TYPE: StringVariables
    StringVariables any `json:"stringVariables"`

    // VendorGroups.
    //
    // For Vendors being returned, this will give you the information you need to group them and order them in the same way that the Bungie Companion app performs grouping. It will automatically be returned if you request the Vendors component.
    //
    // COMPONENT TYPE: Vendors
    VendorGroups any `json:"vendorGroups"`
}
