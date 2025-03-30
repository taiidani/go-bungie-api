// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Definitions_Loadouts_DestinyLoadoutConstantsDefinition struct {
    // LoadoutNameHashes.
    //
    // A list of the loadout name hashes in index order, for convenience.
    LoadoutNameHashes []uint32 `json:"loadoutNameHashes"`

    // LoadoutPreviewFilterOutSocketCategoryHashes.
    //
    // A list of the socket category hashes to be filtered out of loadout item preview displays.
    LoadoutPreviewFilterOutSocketCategoryHashes []uint32 `json:"loadoutPreviewFilterOutSocketCategoryHashes"`

    // Redacted.
    //
    // If this is true, then there is an entity with this identifier/type combination, but BNet is not yet allowed to show it. Sorry!
    Redacted bool `json:"redacted"`

    // WhiteIconImagePath.
    //
    // This is the same icon as the one in the display properties, offered here as well with a more descriptive name.
    WhiteIconImagePath string `json:"whiteIconImagePath"`

    // Index.
    //
    // The index of the entity as it was found in the investment tables.
    Index int32 `json:"index"`

    // LoadoutColorHashes.
    //
    // A list of the loadout color hashes in index order, for convenience.
    LoadoutColorHashes []uint32 `json:"loadoutColorHashes"`

    // LoadoutCountPerCharacter.
    //
    // The maximum number of loadouts available to each character. The loadouts component API response can return fewer loadouts than this, as more loadouts are unlocked by reaching higher Guardian Ranks.
    LoadoutCountPerCharacter int32 `json:"loadoutCountPerCharacter"`

    // LoadoutPreviewFilterOutSocketTypeHashes.
    //
    // A list of the socket type hashes to be filtered out of loadout item preview displays.
    LoadoutPreviewFilterOutSocketTypeHashes []uint32 `json:"loadoutPreviewFilterOutSocketTypeHashes"`

    // BlackIconImagePath.
    //
    // This is a color-inverted version of the whiteIconImagePath.
    BlackIconImagePath string `json:"blackIconImagePath"`

    // DisplayProperties.
    //
    // Many Destiny*Definition contracts - the "first order" entities of Destiny that have their own tables in the Manifest Database - also have displayable information. This is the base class for that display information.
    DisplayProperties Destiny_Definitions_Common_DestinyDisplayPropertiesDefinition `json:"displayProperties"`

    // LoadoutIconHashes.
    //
    // A list of the loadout icon hashes in index order, for convenience.
    LoadoutIconHashes []uint32 `json:"loadoutIconHashes"`

    // Hash.
    //
    // The unique identifier for this entity. Guaranteed to be unique for the type of entity, but not globally.
    //
    // When entities refer to each other in Destiny content, it is this hash that they are referring to.
    Hash uint32 `json:"hash"`
}
