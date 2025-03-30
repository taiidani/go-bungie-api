// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Responses_InventoryChangedResponse struct {
    // AddedInventoryItems.
    //
    // Items that appeared in the inventory possibly as a result of an action.
    AddedInventoryItems []Destiny_Entities_Items_DestinyItemComponent `json:"addedInventoryItems"`

    // RemovedInventoryItems.
    //
    // Items that disappeared from the inventory possibly as a result of an action.
    RemovedInventoryItems []Destiny_Entities_Items_DestinyItemComponent `json:"removedInventoryItems"`
}
