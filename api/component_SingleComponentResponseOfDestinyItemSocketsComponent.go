// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type SingleComponentResponseOfDestinyItemSocketsComponent struct {
    // Disabled.
    //
    // If true, this component is disabled.
    Disabled bool `json:"disabled"`

    // Privacy.
    //
    // 
    Privacy int32 `json:"privacy"`

    // Data.
    //
    // Instanced items can have sockets, which are slots on the item where plugs can be inserted.
    //
    // Sockets are a bit complex: be sure to examine the documentation on the DestinyInventoryItemDefinition's "socket" block and elsewhere on these objects for more details.
    Data Destiny_Entities_Items_DestinyItemSocketsComponent `json:"data"`
}
