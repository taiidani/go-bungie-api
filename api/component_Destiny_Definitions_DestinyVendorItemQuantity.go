// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Definitions_DestinyVendorItemQuantity struct {
    // ItemHash.
    //
    // The hash identifier for the item in question. Use it to look up the item's DestinyInventoryItemDefinition.
    ItemHash uint32 `json:"itemHash"`

    // ItemInstanceId.
    //
    // If this quantity is referring to a specific instance of an item, this will have the item's instance ID. Normally, this will be null.
    ItemInstanceId int64 `json:"itemInstanceId"`

    // Quantity.
    //
    // The amount of the item needed/available depending on the context of where DestinyItemQuantity is being used.
    Quantity int32 `json:"quantity"`

    // HasConditionalVisibility.
    //
    // Indicates that this item quantity may be conditionally shown or hidden, based on various sources of state. For example: server flags, account state, or character progress.
    HasConditionalVisibility bool `json:"hasConditionalVisibility"`
}
