// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Definitions_DestinyDamageTypeDefinition struct {
    // Color.
    //
    // A color associated with the damage type. The displayProperties icon is tinted with a color close to this.
    Color any `json:"color"`

    // DisplayProperties.
    //
    // The description of the damage type, icon etc...
    DisplayProperties any `json:"displayProperties"`

    // EnumValue.
    //
    // We have an enumeration for damage types for quick reference. This is the current definition's damage type enum value.
    EnumValue int32 `json:"enumValue"`

    // Hash.
    //
    // The unique identifier for this entity. Guaranteed to be unique for the type of entity, but not globally.
    //
    // When entities refer to each other in Destiny content, it is this hash that they are referring to.
    Hash uint32 `json:"hash"`

    // Index.
    //
    // The index of the entity as it was found in the investment tables.
    Index int32 `json:"index"`

    // Redacted.
    //
    // If this is true, then there is an entity with this identifier/type combination, but BNet is not yet allowed to show it. Sorry!
    Redacted bool `json:"redacted"`

    // ShowIcon.
    //
    // If TRUE, the game shows this damage type's icon. Otherwise, it doesn't. Whether you show it or not is up to you.
    ShowIcon bool `json:"showIcon"`

    // TransparentIconPath.
    //
    // A variant of the icon that is transparent and colorless.
    TransparentIconPath string `json:"transparentIconPath"`
}
