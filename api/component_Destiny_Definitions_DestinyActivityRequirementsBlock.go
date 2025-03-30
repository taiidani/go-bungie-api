// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Definitions_DestinyActivityRequirementsBlock struct {
    // FireteamRequirementLabels.
    //
    // If being a fireteam member in this activity is gated, this is the gate being checked.
    FireteamRequirementLabels []Destiny_Definitions_DestinyActivityRequirementLabel `json:"fireteamRequirementLabels"`

    // LeaderRequirementLabels.
    //
    // If being a fireteam Leader in this activity is gated, this is the gate being checked.
    LeaderRequirementLabels []Destiny_Definitions_DestinyActivityRequirementLabel `json:"leaderRequirementLabels"`
}
