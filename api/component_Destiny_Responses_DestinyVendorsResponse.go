// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Responses_DestinyVendorsResponse struct {
    // Categories.
    //
    // Categories that the vendor has available, and references to the sales therein. These are keyed by the Vendor Hash, so you will get one Categories Component per vendor returned.
    //
    // COMPONENT TYPE: VendorCategories
    Categories any `json:"categories"`

    // CurrencyLookups.
    //
    // A "lookup" convenience component that can be used to quickly check if the character has access to items that can be used for purchasing.
    //
    // COMPONENT TYPE: CurrencyLookups
    CurrencyLookups any `json:"currencyLookups"`

    // ItemComponents.
    //
    // The set of item detail components, one set of item components per Vendor. These are keyed by the Vendor Hash, so you will get one Item Component Set per vendor returned.
    //
    // The components contained inside are themselves keyed by the vendorSaleIndex, and will have whatever item-level components you requested (Sockets, Stats, Instance data etc...) per item being sold by the vendor.
    ItemComponents map[uint32]DestinyVendorItemComponentSetOfint32 `json:"itemComponents"`

    // Sales.
    //
    // Sales, keyed by the vendorItemIndex of the item being sold. These are keyed by the Vendor Hash, so you will get one Sale Item Set Component per vendor returned.
    //
    // Note that within the Sale Item Set component, the sales are themselves keyed by the vendorSaleIndex, so you can relate it to the current sale item definition within the Vendor's definition.
    //
    // COMPONENT TYPE: VendorSales
    Sales any `json:"sales"`

    // StringVariables.
    //
    // A map of string variable values by hash for this character context.
    //
    // COMPONENT TYPE: StringVariables
    StringVariables any `json:"stringVariables"`

    // VendorGroups.
    //
    // For Vendors being returned, this will give you the information you need to group them and order them in the same way that the Bungie Companion app performs grouping. It will automatically be returned if you request the Vendors component.
    //
    // COMPONENT TYPE: Vendors
    VendorGroups any `json:"vendorGroups"`

    // Vendors.
    //
    // The base properties of the vendor. These are keyed by the Vendor Hash, so you will get one Vendor Component per vendor returned.
    //
    // COMPONENT TYPE: Vendors
    Vendors any `json:"vendors"`
}
