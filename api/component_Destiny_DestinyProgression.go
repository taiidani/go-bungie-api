// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_DestinyProgression struct {
    // CurrentProgress.
    //
    // This is the total amount of progress obtained overall for this progression (for instance, the total amount of Character Level experience earned)
    CurrentProgress int32 `json:"currentProgress"`

    // CurrentResetCount.
    //
    // The number of resets of this progression you've executed this season, if applicable to this progression.
    CurrentResetCount *int32 `json:"currentResetCount"`

    // DailyLimit.
    //
    // If this progression has a daily limit, this is that limit.
    DailyLimit int32 `json:"dailyLimit"`

    // DailyProgress.
    //
    // The amount of progress earned today for this progression.
    DailyProgress int32 `json:"dailyProgress"`

    // Level.
    //
    // This is the level of the progression (for instance, the Character Level).
    Level int32 `json:"level"`

    // LevelCap.
    //
    // This is the maximum possible level you can achieve for this progression (for example, the maximum character level obtainable)
    LevelCap int32 `json:"levelCap"`

    // NextLevelAt.
    //
    // The total amount of progression (i.e. "Experience") needed in order to reach the next level.
    NextLevelAt int32 `json:"nextLevelAt"`

    // ProgressToNextLevel.
    //
    // The amount of progression (i.e. "Experience") needed to reach the next level of this Progression. Jeez, progression is such an overloaded word.
    ProgressToNextLevel int32 `json:"progressToNextLevel"`

    // ProgressionHash.
    //
    // The hash identifier of the Progression in question. Use it to look up the DestinyProgressionDefinition in static data.
    ProgressionHash uint32 `json:"progressionHash"`

    // RewardItemSocketOverrideStates.
    //
    // Information about items stats and states that have socket overrides, if there is any data for it.
    RewardItemSocketOverrideStates any `json:"rewardItemSocketOverrideStates"`

    // RewardItemStates.
    //
    // Information about historical rewards for this progression, if there is any data for it.
    RewardItemStates []int32 `json:"rewardItemStates"`

    // SeasonResets.
    //
    // Information about historical resets of this progression, if there is any data for it.
    SeasonResets []Destiny_DestinyProgressionResetEntry `json:"seasonResets"`

    // StepIndex.
    //
    // Progressions define their levels in "steps". Since the last step may be repeatable, the user may be at a higher level than the actual Step achieved in the progression. Not necessarily useful, but potentially interesting for those cruising the API. Relate this to the "steps" property of the DestinyProgression to see which step the user is on, if you care about that. (Note that this is Content Version dependent since it refers to indexes.)
    StepIndex int32 `json:"stepIndex"`

    // WeeklyLimit.
    //
    // If this progression has a weekly limit, this is that limit.
    WeeklyLimit int32 `json:"weeklyLimit"`

    // WeeklyProgress.
    //
    // The amount of progress earned toward this progression in the current week.
    WeeklyProgress int32 `json:"weeklyProgress"`
}
