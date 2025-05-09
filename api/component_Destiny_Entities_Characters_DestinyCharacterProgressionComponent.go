// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Entities_Characters_DestinyCharacterProgressionComponent struct {
    // Checklists.
    //
    // The set of checklists that can be examined for this specific character, keyed by the hash identifier of the Checklist (DestinyChecklistDefinition)
    //
    // For each checklist returned, its value is itself a Dictionary keyed by the checklist's hash identifier with the value being a boolean indicating if it's been discovered yet.
    Checklists map[uint32]map[uint32]bool `json:"checklists"`

    // Factions.
    //
    // A dictionary of all known Factions, keyed by the Faction's hash. It contains data about this character's status with the faction.
    Factions map[uint32]Destiny_Progression_DestinyFactionProgression `json:"factions"`

    // Milestones.
    //
    // Milestones are related to the simple progressions shown in the game, but return additional and hopefully helpful information for users about the specifics of the Milestone's status.
    Milestones map[uint32]Destiny_Milestones_DestinyMilestone `json:"milestones"`

    // Progressions.
    //
    // A Dictionary of all known progressions for the Character, keyed by the Progression's hash.
    //
    // Not all progressions have user-facing data, but those who do will have that data contained in the DestinyProgressionDefinition.
    Progressions map[uint32]Destiny_DestinyProgression `json:"progressions"`

    // Quests.
    //
    // If the user has any active quests, the quests' statuses will be returned here.
    //
    //  Note that quests have been largely supplanted by Milestones, but that doesn't mean that they won't make a comeback independent of milestones at some point.
    //
    //  (Fun fact: quests came back as I feared they would, but we never looped back to populate this... I'm going to put that in the backlog.)
    Quests []Destiny_Quests_DestinyQuestStatus `json:"quests"`

    // SeasonalArtifact.
    //
    // Data related to your progress on the current season's artifact that can vary per character.
    SeasonalArtifact any `json:"seasonalArtifact"`

    // UninstancedItemObjectives.
    //
    // Sometimes, you have items in your inventory that don't have instances, but still have Objective information. This provides you that objective information for uninstanced items. 
    //
    // This dictionary is keyed by the item's hash: which you can use to look up the name and description for the overall task(s) implied by the objective. The value is the list of objectives for this item, and their statuses.
    UninstancedItemObjectives map[uint32][]Destiny_Quests_DestinyObjectiveProgress `json:"uninstancedItemObjectives"`

    // UninstancedItemPerks.
    //
    // Sometimes, you have items in your inventory that don't have instances, but still have perks (for example: Trials passage cards). This gives you the perk information for uninstanced items.
    //
    // This dictionary is keyed by item hash, which you can use to look up the corresponding item definition. The value is the list of perks states for the item.
    UninstancedItemPerks map[uint32]Destiny_Entities_Items_DestinyItemPerksComponent `json:"uninstancedItemPerks"`
}
