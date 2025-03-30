// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Definitions_DestinyDisplayCategoryDefinition struct {
    // DisplayStyleHash.
    //
    // An indicator of how the category will be displayed in the UI. It's up to you to do something cool or interesting in response to this, or just to treat it as a normal category.
    DisplayStyleHash uint32 `json:"displayStyleHash"`

    // ProgressionHash.
    //
    // If it exists, this is the hash identifier of a DestinyProgressionDefinition that represents the progression to show on this display category.
    //
    // Specific categories can now have thier own distinct progression, apparently. So that's cool.
    ProgressionHash uint32 `json:"progressionHash"`

    // SortOrder.
    //
    // If this category sorts items in a nonstandard way, this will be the way we sort.
    SortOrder int32 `json:"sortOrder"`

    // DisplayProperties.
    //
    // Many Destiny*Definition contracts - the "first order" entities of Destiny that have their own tables in the Manifest Database - also have displayable information. This is the base class for that display information.
    DisplayProperties Destiny_Definitions_Common_DestinyDisplayPropertiesDefinition `json:"displayProperties"`

    // DisplayStyleIdentifier.
    //
    // An indicator of how the category will be displayed in the UI. It's up to you to do something cool or interesting in response to this, or just to treat it as a normal category.
    DisplayStyleIdentifier string `json:"displayStyleIdentifier"`

    // Identifier.
    //
    // A string identifier for the display category.
    Identifier string `json:"identifier"`

    // Index.
    //
    // 
    Index int32 `json:"index"`

    // DisplayCategoryHash.
    //
    // 
    DisplayCategoryHash uint32 `json:"displayCategoryHash"`

    // DisplayInBanner.
    //
    // If true, this category should be displayed in the "Banner" section of the vendor's UI.
    DisplayInBanner bool `json:"displayInBanner"`
}
