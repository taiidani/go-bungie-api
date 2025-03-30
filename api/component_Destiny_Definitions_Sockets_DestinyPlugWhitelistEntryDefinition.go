// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Definitions_Sockets_DestinyPlugWhitelistEntryDefinition struct {
    // CategoryHash.
    //
    // The hash identifier of the Plug Category to compare against the plug item's plug.plugCategoryHash.
    //
    // Note that this does NOT relate to any Definition in itself, it is only used for comparison purposes.
    CategoryHash uint32 `json:"categoryHash"`

    // CategoryIdentifier.
    //
    // The string identifier for the category, which is here mostly for debug purposes.
    CategoryIdentifier string `json:"categoryIdentifier"`

    // ReinitializationPossiblePlugHashes.
    //
    // The list of all plug items (DestinyInventoryItemDefinition) that the socket may randomly be populated with when reinitialized.
    //
    // Which ones you should actually show are determined by the plug being inserted into the socket, and the socket’s type.
    //
    // When you inspect the plug that could go into a Masterwork Socket, look up the socket type of the socket being inspected and find the DestinySocketTypeDefinition.
    //
    // Then, look at the Plugs that can fit in that socket. Find the Whitelist in the DestinySocketTypeDefinition that matches the plug item’s categoryhash.
    //
    // That whitelist entry will potentially have a new “reinitializationPossiblePlugHashes” property.If it does, that means we know what it will roll if you try to insert this plug into this socket.
    ReinitializationPossiblePlugHashes []uint32 `json:"reinitializationPossiblePlugHashes"`
}
