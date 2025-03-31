// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_HistoricalStats_DestinyLeaderboardEntry struct {
    // CharacterId.
    //
    // ID of the player's best character for the reported stat.
    CharacterId int64 `json:"characterId"`

    // Player.
    //
    // Identity details of the player
    Player any `json:"player"`

    // Rank.
    //
    // Where this player ranks on the leaderboard. A value of 1 is the top rank.
    Rank int32 `json:"rank"`

    // Value.
    //
    // Value of the stat for this player
    Value any `json:"value"`
}
