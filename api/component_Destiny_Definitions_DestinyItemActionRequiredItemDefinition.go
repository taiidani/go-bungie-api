// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Definitions_DestinyItemActionRequiredItemDefinition struct {
    // DeleteOnAction.
    //
    // If true, the item/quantity will be deleted from your inventory when the action is performed. Otherwise, you'll retain these required items after the action is complete.
    DeleteOnAction bool `json:"deleteOnAction"`

    // ItemHash.
    //
    // The hash identifier of the item you need to have. Use it to look up the DestinyInventoryItemDefinition for more info.
    ItemHash uint32 `json:"itemHash"`

    // Count.
    //
    // The minimum quantity of the item you have to have.
    Count int32 `json:"count"`
}
