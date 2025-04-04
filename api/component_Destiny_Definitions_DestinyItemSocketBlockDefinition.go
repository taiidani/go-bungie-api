// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Definitions_DestinyItemSocketBlockDefinition struct {
    // Detail.
    //
    // This was supposed to be a string that would give per-item details about sockets. In practice, it turns out that all this ever has is the localized word "details". ... that's lame, but perhaps it will become something cool in the future.
    Detail string `json:"detail"`

    // IntrinsicSockets.
    //
    // Each intrinsic (or immutable/permanent) socket on an item is defined here, along with the plug that is permanently affixed to the socket.
    IntrinsicSockets []Destiny_Definitions_DestinyItemIntrinsicSocketEntryDefinition `json:"intrinsicSockets"`

    // SocketCategories.
    //
    // A convenience property, that refers to the sockets in the "sockets" property, pre-grouped by category and ordered in the manner that they should be grouped in the UI. You could form this yourself with the existing data, but why would you want to? Enjoy life man.
    SocketCategories []Destiny_Definitions_DestinyItemSocketCategoryDefinition `json:"socketCategories"`

    // SocketEntries.
    //
    // Each non-intrinsic (or mutable) socket on an item is defined here. Check inside for more info.
    SocketEntries []Destiny_Definitions_DestinyItemSocketEntryDefinition `json:"socketEntries"`
}
