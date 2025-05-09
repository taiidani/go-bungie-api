// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Definitions_Social_DestinySocialCommendationDefinition struct {
    // ActivityGivingLimit.
    //
    // 
    ActivityGivingLimit int32 `json:"activityGivingLimit"`

    // CardImagePath.
    //
    // 
    CardImagePath string `json:"cardImagePath"`

    // Color.
    //
    // Represents a color whose RGBA values are all represented as values between 0 and 255.
    Color Destiny_Misc_DestinyColor `json:"color"`

    // DisplayActivities.
    //
    // The display properties for the the activities that this commendation is available in.
    DisplayActivities []Destiny_Definitions_Common_DestinyDisplayPropertiesDefinition `json:"displayActivities"`

    // DisplayPriority.
    //
    // 
    DisplayPriority int32 `json:"displayPriority"`

    // DisplayProperties.
    //
    // Many Destiny*Definition contracts - the "first order" entities of Destiny that have their own tables in the Manifest Database - also have displayable information. This is the base class for that display information.
    DisplayProperties Destiny_Definitions_Common_DestinyDisplayPropertiesDefinition `json:"displayProperties"`

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

    // ParentCommendationNodeHash.
    //
    // 
    ParentCommendationNodeHash uint32 `json:"parentCommendationNodeHash"`

    // Redacted.
    //
    // If this is true, then there is an entity with this identifier/type combination, but BNet is not yet allowed to show it. Sorry!
    Redacted bool `json:"redacted"`
}
