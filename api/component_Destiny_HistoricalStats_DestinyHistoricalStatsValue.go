// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_HistoricalStats_DestinyHistoricalStatsValue struct {
    // ActivityId.
    //
    // When a stat represents the best, most, longest, fastest or some other personal best, the actual activity ID where that personal best was established is available on this property.
    ActivityId *int64 `json:"activityId"`

    // Basic.
    //
    // Basic stat value.
    Basic any `json:"basic"`

    // Pga.
    //
    // Per game average for the statistic, if applicable
    Pga any `json:"pga"`

    // StatId.
    //
    // Unique ID for this stat
    StatId string `json:"statId"`

    // Weighted.
    //
    // Weighted value of the stat if a weight greater than 1 has been assigned.
    Weighted any `json:"weighted"`
}
