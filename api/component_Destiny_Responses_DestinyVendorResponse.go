// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Responses_DestinyVendorResponse struct {
    // Categories.
    //
    // Categories that the vendor has available, and references to the sales therein.
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
    // Item components, keyed by the vendorItemIndex of the active sale items.
    //
    // COMPONENT TYPE: [See inside the DestinyVendorItemComponentSet contract for component types.]
    ItemComponents any `json:"itemComponents"`

    // Sales.
    //
    // Sales, keyed by the vendorItemIndex of the item being sold.
    //
    // COMPONENT TYPE: VendorSales
    Sales any `json:"sales"`

    // StringVariables.
    //
    // A map of string variable values by hash for this character context.
    //
    // COMPONENT TYPE: StringVariables
    StringVariables any `json:"stringVariables"`

    // Vendor.
    //
    // The base properties of the vendor.
    //
    // COMPONENT TYPE: Vendors
    Vendor any `json:"vendor"`
}
