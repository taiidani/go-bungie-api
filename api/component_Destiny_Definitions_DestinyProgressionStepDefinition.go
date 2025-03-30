// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Definitions_DestinyProgressionStepDefinition struct {
    // DisplayEffectType.
    //
    // This appears to be, when you "level up", whether a visual effect will display and on what entity. See DestinyProgressionStepDisplayEffect for slightly more info.
    DisplayEffectType int32 `json:"displayEffectType"`

    // Icon.
    //
    // If this progression step has a specific icon related to it, this is the icon to show.
    Icon string `json:"icon"`

    // ProgressTotal.
    //
    // The total amount of progression points/"experience" you will need to initially reach this step. If this is the last step and the progression is repeating indefinitely (DestinyProgressionDefinition.repeatLastStep), this will also be the progress needed to level it up further by repeating this step again.
    ProgressTotal int32 `json:"progressTotal"`

    // RewardItems.
    //
    // A listing of items rewarded as a result of reaching this level.
    RewardItems []Destiny_DestinyItemQuantity `json:"rewardItems"`

    // StepName.
    //
    // Very rarely, Progressions will have localized text describing the Level of the progression. This will be that localized text, if it exists. Otherwise, the standard appears to be to simply show the level numerically.
    StepName string `json:"stepName"`
}
