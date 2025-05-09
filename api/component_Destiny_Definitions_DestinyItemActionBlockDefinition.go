// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Definitions_DestinyItemActionBlockDefinition struct {
    // ActionTypeLabel.
    //
    // The internal identifier for the action.
    ActionTypeLabel string `json:"actionTypeLabel"`

    // ConsumeEntireStack.
    //
    // If true, the entire stack is deleted when the action completes.
    ConsumeEntireStack bool `json:"consumeEntireStack"`

    // DeleteOnAction.
    //
    // If true, the item is deleted when the action completes.
    DeleteOnAction bool `json:"deleteOnAction"`

    // IsPositive.
    //
    // The content has this property, however it's not entirely clear how it is used.
    IsPositive bool `json:"isPositive"`

    // OverlayIcon.
    //
    // The icon associated with the overlay screen for the action, if any.
    OverlayIcon string `json:"overlayIcon"`

    // OverlayScreenName.
    //
    // If the action has an overlay screen associated with it, this is the name of that screen. Unfortunately, we cannot return the screen's data itself.
    OverlayScreenName string `json:"overlayScreenName"`

    // ProgressionRewards.
    //
    // If performing this action earns you Progression, this is the list of progressions and values granted for those progressions by performing this action.
    ProgressionRewards []Destiny_Definitions_DestinyProgressionRewardDefinition `json:"progressionRewards"`

    // RequiredCooldownHash.
    //
    // The identifier hash for the Cooldown associated with this action. We have not pulled this data yet for you to have more data to use for cooldowns.
    RequiredCooldownHash uint32 `json:"requiredCooldownHash"`

    // RequiredCooldownSeconds.
    //
    // The number of seconds to delay before allowing this action to be performed again.
    RequiredCooldownSeconds int32 `json:"requiredCooldownSeconds"`

    // RequiredItems.
    //
    // If the action requires other items to exist or be destroyed, this is the list of those items and requirements.
    RequiredItems []Destiny_Definitions_DestinyItemActionRequiredItemDefinition `json:"requiredItems"`

    // RequiredLocation.
    //
    // Theoretically, an item could have a localized string for a hint about the location in which the action should be performed. In practice, no items yet have this property.
    RequiredLocation string `json:"requiredLocation"`

    // UseOnAcquire.
    //
    // If true, this action will be performed as soon as you earn this item. Some rewards work this way, providing you a single item to pick up from a reward-granting vendor in-game and then immediately consuming itself to provide you multiple items.
    UseOnAcquire bool `json:"useOnAcquire"`

    // VerbDescription.
    //
    // Localized text describing the action being performed.
    VerbDescription string `json:"verbDescription"`

    // VerbName.
    //
    // Localized text for the verb of the action being performed.
    VerbName string `json:"verbName"`
}
