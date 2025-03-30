// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Definitions_DestinyVendorDisplayPropertiesDefinition struct {
    // Name.
    //
    // 
    Name string `json:"name"`

    // SmallTransparentIcon.
    //
    // This is the icon used in parts of the game UI such as the vendor's waypoint.
    SmallTransparentIcon string `json:"smallTransparentIcon"`

    // Icon.
    //
    // Note that "icon" is sometimes misleading, and should be interpreted in the context of the entity. For instance, in Destiny 1 the DestinyRecordBookDefinition's icon was a big picture of a book.
    //
    // But usually, it will be a small square image that you can use as... well, an icon.
    //
    // They are currently represented as 96px x 96px images.
    Icon string `json:"icon"`

    // IconSequences.
    //
    // 
    IconSequences []Destiny_Definitions_Common_DestinyIconSequenceDefinition `json:"iconSequences"`

    // LargeIcon.
    //
    // I regret calling this a "large icon". It's more like a medium-sized image with a picture of the vendor's mug on it, trying their best to look cool. Not what one would call an icon.
    LargeIcon string `json:"largeIcon"`

    // OriginalIcon.
    //
    // If we replaced the icon with something more glitzy, this is the original icon that the vendor had according to the game's content. It may be more lame and/or have less razzle-dazzle. But who am I to tell you which icon to use.
    OriginalIcon string `json:"originalIcon"`

    // Description.
    //
    // 
    Description string `json:"description"`

    // HasIcon.
    //
    // 
    HasIcon bool `json:"hasIcon"`

    // HighResIcon.
    //
    // If this item has a high-res icon (at least for now, many things won't), then the path to that icon will be here.
    HighResIcon string `json:"highResIcon"`

    // LargeTransparentIcon.
    //
    // This is apparently the "Watermark". I am not certain offhand where this is actually used in the Game UI, but some people may find it useful.
    LargeTransparentIcon string `json:"largeTransparentIcon"`

    // RequirementsDisplay.
    //
    // Vendors, in addition to expected display property data, may also show some "common requirements" as statically defined definition data. This might be when a vendor accepts a single type of currency, or when the currency is unique to the vendor and the designers wanted to show that currency when you interact with the vendor.
    RequirementsDisplay []Destiny_Definitions_DestinyVendorRequirementDisplayEntryDefinition `json:"requirementsDisplay"`

    // Subtitle.
    //
    // 
    Subtitle string `json:"subtitle"`

    // MapIcon.
    //
    // This is the icon used in the map overview, when the vendor is located on the map.
    MapIcon string `json:"mapIcon"`
}
