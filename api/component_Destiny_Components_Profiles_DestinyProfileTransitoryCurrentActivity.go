// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Components_Profiles_DestinyProfileTransitoryCurrentActivity struct {
    // EndTime.
    //
    // If you're still in it but it "ended" (like when folks are dancing around the loot after they beat a boss), this is when the activity ended.
    EndTime *string `json:"endTime"`

    // HighestOpposingFactionScore.
    //
    // If you have human opponents, this is the highest opposing team's score.
    HighestOpposingFactionScore float32 `json:"highestOpposingFactionScore"`

    // NumberOfOpponents.
    //
    // This is how many human or poorly crafted aimbot opponents you have.
    NumberOfOpponents int32 `json:"numberOfOpponents"`

    // NumberOfPlayers.
    //
    // This is how many human or poorly crafted aimbots are on your team.
    NumberOfPlayers int32 `json:"numberOfPlayers"`

    // Score.
    //
    // This is what our non-authoritative source thought the score was.
    Score float32 `json:"score"`

    // StartTime.
    //
    // When the activity started.
    StartTime *string `json:"startTime"`
}
