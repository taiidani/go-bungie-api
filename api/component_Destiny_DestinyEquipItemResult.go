// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_DestinyEquipItemResult struct {
    // EquipStatus.
    //
    // A PlatformErrorCodes enum indicating whether it succeeded, and if it failed why.
    EquipStatus int32 `json:"equipStatus"`

    // ItemInstanceId.
    //
    // The instance ID of the item in question (all items that can be equipped must, but definition, be Instanced and thus have an Instance ID that you can use to refer to them)
    ItemInstanceId int64 `json:"itemInstanceId"`
}
