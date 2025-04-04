// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Responses_DestinyCharacterResponse struct {
    // Activities.
    //
    // Activity data - info about current activities available to the player.
    //
    // COMPONENT TYPE: CharacterActivities
    Activities any `json:"activities"`

    // Character.
    //
    // Base information about the character in question.
    //
    // COMPONENT TYPE: Characters
    Character any `json:"character"`

    // Collectibles.
    //
    // COMPONENT TYPE: Collectibles
    Collectibles any `json:"collectibles"`

    // CurrencyLookups.
    //
    // A "lookup" convenience component that can be used to quickly check if the character has access to items that can be used for purchasing.
    //
    // COMPONENT TYPE: CurrencyLookups
    CurrencyLookups any `json:"currencyLookups"`

    // Equipment.
    //
    // Equipped items on the character.
    //
    // COMPONENT TYPE: CharacterEquipment
    Equipment any `json:"equipment"`

    // Inventory.
    //
    // The character-level non-equipped inventory items.
    //
    // COMPONENT TYPE: CharacterInventories
    Inventory any `json:"inventory"`

    // ItemComponents.
    //
    // The set of components belonging to the player's instanced items.
    //
    // COMPONENT TYPE: [See inside the DestinyItemComponentSet contract for component types.]
    ItemComponents any `json:"itemComponents"`

    // Kiosks.
    //
    // Items available from Kiosks that are available to this specific character. 
    //
    // COMPONENT TYPE: Kiosks
    Kiosks any `json:"kiosks"`

    // Loadouts.
    //
    // The loadouts available to the character.
    //
    // COMPONENT TYPE: CharacterLoadouts
    Loadouts any `json:"loadouts"`

    // PlugSets.
    //
    // When sockets refer to reusable Plug Sets (see DestinyPlugSetDefinition for more info), this is the set of plugs and their states that are scoped to this character.
    //
    // This comes back with ItemSockets, as it is needed for a complete picture of the sockets on requested items.
    //
    // COMPONENT TYPE: ItemSockets
    PlugSets any `json:"plugSets"`

    // PresentationNodes.
    //
    // COMPONENT TYPE: PresentationNodes
    PresentationNodes any `json:"presentationNodes"`

    // Progressions.
    //
    // Character progression data, including Milestones.
    //
    // COMPONENT TYPE: CharacterProgressions
    Progressions any `json:"progressions"`

    // Records.
    //
    // COMPONENT TYPE: Records
    Records any `json:"records"`

    // RenderData.
    //
    // Character rendering data - a minimal set of information about equipment and dyes used for rendering.
    //
    // COMPONENT TYPE: CharacterRenderData
    RenderData any `json:"renderData"`

    // UninstancedItemComponents.
    //
    // The set of components belonging to the player's UNinstanced items. Because apparently now those too can have information relevant to the character's state.
    //
    // COMPONENT TYPE: [See inside the DestinyItemComponentSet contract for component types.]
    UninstancedItemComponents any `json:"uninstancedItemComponents"`
}
