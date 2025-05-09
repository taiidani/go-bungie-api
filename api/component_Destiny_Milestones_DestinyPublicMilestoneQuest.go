// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Milestones_DestinyPublicMilestoneQuest struct {
    // Activity.
    //
    // A milestone need not have an active activity, but if there is one it will be returned here, along with any variant and additional information.
    Activity any `json:"activity"`

    // Challenges.
    //
    // For the given quest there could be 0-to-Many challenges: mini quests that you can perform in the course of doing this quest, that may grant you rewards and benefits.
    Challenges []Destiny_Milestones_DestinyPublicMilestoneChallenge `json:"challenges"`

    // QuestItemHash.
    //
    // Quests are defined as Items in content. As such, this is the hash identifier of the DestinyInventoryItemDefinition that represents this quest. It will have pointers to all of the steps in the quest, and display information for the quest (title, description, icon etc) Individual steps will be referred to in the Quest item's DestinyInventoryItemDefinition.setData property, and themselves are Items with their own renderable data.
    QuestItemHash uint32 `json:"questItemHash"`
}
