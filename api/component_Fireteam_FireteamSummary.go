// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Fireteam_FireteamSummary struct {
    // ActivityType.
    //
    // 
    ActivityType int32 `json:"activityType"`

    // AlternateSlotCount.
    //
    // 
    AlternateSlotCount *int32 `json:"alternateSlotCount"`

    // AvailableAlternateSlotCount.
    //
    // 
    AvailableAlternateSlotCount int32 `json:"availableAlternateSlotCount"`

    // AvailablePlayerSlotCount.
    //
    // 
    AvailablePlayerSlotCount int32 `json:"availablePlayerSlotCount"`

    // DateCreated.
    //
    // 
    DateCreated string `json:"dateCreated"`

    // DateModified.
    //
    // 
    DateModified *string `json:"dateModified"`

    // DatePlayerModified.
    //
    // 
    DatePlayerModified string `json:"datePlayerModified"`

    // FireteamId.
    //
    // 
    FireteamId int64 `json:"fireteamId"`

    // GroupId.
    //
    // 
    GroupId int64 `json:"groupId"`

    // IsImmediate.
    //
    // 
    IsImmediate bool `json:"isImmediate"`

    // IsPublic.
    //
    // 
    IsPublic bool `json:"isPublic"`

    // IsValid.
    //
    // 
    IsValid bool `json:"isValid"`

    // Locale.
    //
    // 
    Locale string `json:"locale"`

    // OwnerCurrentGuardianRankSnapshot.
    //
    // 
    OwnerCurrentGuardianRankSnapshot int32 `json:"ownerCurrentGuardianRankSnapshot"`

    // OwnerHighestLifetimeGuardianRankSnapshot.
    //
    // 
    OwnerHighestLifetimeGuardianRankSnapshot int32 `json:"ownerHighestLifetimeGuardianRankSnapshot"`

    // OwnerMembershipId.
    //
    // 
    OwnerMembershipId int64 `json:"ownerMembershipId"`

    // OwnerTotalCommendationScoreSnapshot.
    //
    // 
    OwnerTotalCommendationScoreSnapshot int32 `json:"ownerTotalCommendationScoreSnapshot"`

    // Platform.
    //
    // 
    Platform byte `json:"platform"`

    // PlayerSlotCount.
    //
    // 
    PlayerSlotCount int32 `json:"playerSlotCount"`

    // ScheduledTime.
    //
    // 
    ScheduledTime *string `json:"scheduledTime"`

    // Title.
    //
    // 
    Title string `json:"title"`

    // TitleBeforeModeration.
    //
    // 
    TitleBeforeModeration string `json:"titleBeforeModeration"`
}
