// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Definitions_DestinyActivityModeDefinition struct {
    // ActivityModeCategory.
    //
    // The type of play being performed in broad terms (PVP, PVE)
    ActivityModeCategory int32 `json:"activityModeCategory"`

    // ActivityModeMappings.
    //
    // If this exists, the mode has specific Activities (referred to by the Key) that should instead map to other Activity Modes when they are played. This was useful in D1 for Private Matches, where we wanted to have Private Matches as an activity mode while still referring to the specific mode being played.
    ActivityModeMappings any `json:"activityModeMappings"`

    // Display.
    //
    // If FALSE, we want to ignore this type when we're showing activity modes in BNet UI. It will still be returned in case 3rd parties want to use it for any purpose.
    Display bool `json:"display"`

    // DisplayProperties.
    //
    // Many Destiny*Definition contracts - the "first order" entities of Destiny that have their own tables in the Manifest Database - also have displayable information. This is the base class for that display information.
    DisplayProperties Destiny_Definitions_Common_DestinyDisplayPropertiesDefinition `json:"displayProperties"`

    // FriendlyName.
    //
    // A Friendly identifier you can use for referring to this Activity Mode. We really only used this in our URLs, so... you know, take that for whatever it's worth.
    FriendlyName string `json:"friendlyName"`

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

    // IsAggregateMode.
    //
    // If true, this mode is an aggregation of other, more specific modes rather than being a mode in itself. This includes modes that group Features/Events rather than Gameplay, such as Trials of The Nine: Trials of the Nine being an Event that is interesting to see aggregate data for, but when you play the activities within Trials of the Nine they are more specific activity modes such as Clash.
    IsAggregateMode bool `json:"isAggregateMode"`

    // IsTeamBased.
    //
    // If True, this mode has oppositional teams fighting against each other rather than "Free-For-All" or Co-operative modes of play.
    //
    // Note that Aggregate modes are never marked as team based, even if they happen to be team based at the moment. At any time, an aggregate whose subordinates are only team based could be changed so that one or more aren't team based, and then this boolean won't make much sense (the aggregation would become "sometimes team based"). Let's not deal with that right now.
    IsTeamBased bool `json:"isTeamBased"`

    // ModeType.
    //
    // The Enumeration value for this Activity Mode. Pass this identifier into Stats endpoints to get aggregate stats for this mode.
    ModeType int32 `json:"modeType"`

    // Order.
    //
    // The relative ordering of activity modes.
    Order int32 `json:"order"`

    // ParentHashes.
    //
    // The hash identifiers of the DestinyActivityModeDefinitions that represent all of the "parent" modes for this mode. For instance, the Nightfall Mode is also a member of AllStrikes and AllPvE.
    ParentHashes []uint32 `json:"parentHashes"`

    // PgcrImage.
    //
    // If this activity mode has a related PGCR image, this will be the path to said image.
    PgcrImage string `json:"pgcrImage"`

    // Redacted.
    //
    // If this is true, then there is an entity with this identifier/type combination, but BNet is not yet allowed to show it. Sorry!
    Redacted bool `json:"redacted"`
}
