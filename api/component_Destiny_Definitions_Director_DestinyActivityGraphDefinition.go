// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Definitions_Director_DestinyActivityGraphDefinition struct {
    // ArtElements.
    //
    // Represents one-off/special UI elements that appear on the map.
    ArtElements []Destiny_Definitions_Director_DestinyActivityGraphArtElementDefinition `json:"artElements"`

    // Connections.
    //
    // Represents connections between graph nodes. However, it lacks context that we'd need to make good use of it.
    Connections []Destiny_Definitions_Director_DestinyActivityGraphConnectionDefinition `json:"connections"`

    // DisplayObjectives.
    //
    // Objectives can display on maps, and this is supposedly metadata for that. I have not had the time to analyze the details of what is useful within however: we could be missing important data to make this work. Expect this property to be expanded on later if possible.
    DisplayObjectives []Destiny_Definitions_Director_DestinyActivityGraphDisplayObjectiveDefinition `json:"displayObjectives"`

    // DisplayProgressions.
    //
    // Progressions can also display on maps, but similarly to displayObjectives we appear to lack some required information and context right now. We will have to look into it later and add more data if possible.
    DisplayProgressions []Destiny_Definitions_Director_DestinyActivityGraphDisplayProgressionDefinition `json:"displayProgressions"`

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

    // LinkedGraphs.
    //
    // Represents links between this Activity Graph and other ones.
    LinkedGraphs []Destiny_Definitions_Director_DestinyLinkedGraphDefinition `json:"linkedGraphs"`

    // Nodes.
    //
    // These represent the visual "nodes" on the map's view. These are the activities you can click on in the map.
    Nodes []Destiny_Definitions_Director_DestinyActivityGraphNodeDefinition `json:"nodes"`

    // Redacted.
    //
    // If this is true, then there is an entity with this identifier/type combination, but BNet is not yet allowed to show it. Sorry!
    Redacted bool `json:"redacted"`
}
