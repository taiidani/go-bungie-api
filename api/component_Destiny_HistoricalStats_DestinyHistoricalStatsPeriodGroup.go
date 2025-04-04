// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_HistoricalStats_DestinyHistoricalStatsPeriodGroup struct {
    // ActivityDetails.
    //
    // If the period group is for a specific activity, this property will be set.
    ActivityDetails any `json:"activityDetails"`

    // Period.
    //
    // Period for the group. If the stat periodType is day, then this will have a specific day. If the type is monthly, then this value will be the first day of the applicable month. This value is not set when the periodType is 'all time'.
    Period string `json:"period"`

    // Values.
    //
    // Collection of stats for the period.
    Values any `json:"values"`
}
