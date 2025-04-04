// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Definitions_Items_DestinyDerivedItemDefinition struct {
    // IconPath.
    //
    // An icon for the item.
    IconPath string `json:"iconPath"`

    // ItemDescription.
    //
    // A brief description of the item.
    ItemDescription string `json:"itemDescription"`

    // ItemDetail.
    //
    // Additional details about the derived item, in addition to the description.
    ItemDetail string `json:"itemDetail"`

    // ItemHash.
    //
    // The hash for the DestinyInventoryItemDefinition of this derived item, if there is one. Sometimes we are given this information as a manual override, in which case there won't be an actual DestinyInventoryItemDefinition for what we display, but you can still show the strings from this object itself.
    ItemHash *uint32 `json:"itemHash"`

    // ItemName.
    //
    // The name of the derived item.
    ItemName string `json:"itemName"`

    // VendorItemIndex.
    //
    // If the item was derived from a "Preview Vendor", this will be an index into the DestinyVendorDefinition's itemList property. Otherwise, -1.
    VendorItemIndex int32 `json:"vendorItemIndex"`
}
