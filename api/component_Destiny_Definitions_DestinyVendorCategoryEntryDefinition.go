// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Definitions_DestinyVendorCategoryEntryDefinition struct {
    // BuyStringOverride.
    //
    // The localized string for making purchases from this category, if it is different from the vendor's string for purchasing.
    BuyStringOverride string `json:"buyStringOverride"`

    // CategoryHash.
    //
    // The hashed identifier for the category.
    CategoryHash uint32 `json:"categoryHash"`

    // CategoryIndex.
    //
    // The index of the category in the original category definitions for the vendor.
    CategoryIndex int32 `json:"categoryIndex"`

    // DisabledDescription.
    //
    // If the category is disabled, this is the localized description to show.
    DisabledDescription string `json:"disabledDescription"`

    // DisplayTitle.
    //
    // The localized title of the category.
    DisplayTitle string `json:"displayTitle"`

    // HideFromRegularPurchase.
    //
    // True if this category doesn't allow purchases.
    HideFromRegularPurchase bool `json:"hideFromRegularPurchase"`

    // HideIfNoCurrency.
    //
    // If you don't have the currency required to buy items from this category, should the items be hidden?
    HideIfNoCurrency bool `json:"hideIfNoCurrency"`

    // IsDisplayOnly.
    //
    // If true, this category only displays items: you can't purchase anything in them.
    IsDisplayOnly bool `json:"isDisplayOnly"`

    // IsPreview.
    //
    // Sometimes a category isn't actually used to sell items, but rather to preview them. This implies different UI (and manual placement of the category in the UI) in the game, and special treatment.
    IsPreview bool `json:"isPreview"`

    // Overlay.
    //
    // If this category has an overlay prompt that should appear, this contains the details of that prompt.
    Overlay any `json:"overlay"`

    // QuantityAvailable.
    //
    // The amount of items that will be available when this category is shown.
    QuantityAvailable int32 `json:"quantityAvailable"`

    // ResetIntervalMinutesOverride.
    //
    // 
    ResetIntervalMinutesOverride int32 `json:"resetIntervalMinutesOverride"`

    // ResetOffsetMinutesOverride.
    //
    // 
    ResetOffsetMinutesOverride int32 `json:"resetOffsetMinutesOverride"`

    // ShowUnavailableItems.
    //
    // If items aren't up for sale in this category, should we still show them (greyed out)?
    ShowUnavailableItems bool `json:"showUnavailableItems"`

    // SortValue.
    //
    // Used in sorting items in vendors... but there's a lot more to it. Just go with the order provided in the itemIndexes property on the DestinyVendorCategoryComponent instead, it should be more reliable than trying to recalculate it yourself.
    SortValue int32 `json:"sortValue"`

    // VendorItemIndexes.
    //
    // A shortcut for the vendor item indexes sold under this category. Saves us from some expensive reorganization at runtime.
    VendorItemIndexes []int32 `json:"vendorItemIndexes"`
}
