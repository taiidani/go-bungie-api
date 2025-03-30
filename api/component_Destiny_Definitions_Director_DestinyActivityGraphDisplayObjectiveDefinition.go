// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Definitions_Director_DestinyActivityGraphDisplayObjectiveDefinition struct {
    // ObjectiveHash.
    //
    // The objective being shown on the map.
    ObjectiveHash uint32 `json:"objectiveHash"`

    // Id.
    //
    // $NOTE $amola 2017-01-19 This field is apparently something that CUI uses to manually wire up objectives to display info. I am unsure how it works.
    Id uint32 `json:"id"`
}
