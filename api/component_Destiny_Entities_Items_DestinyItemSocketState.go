// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Entities_Items_DestinyItemSocketState struct {
    // EnableFailIndexes.
    //
    // If a plug is inserted but not enabled, this will be populated with indexes into the plug item definition's plug.enabledRules property, so that you can show the reasons why it is not enabled.
    EnableFailIndexes []int32 `json:"enableFailIndexes"`

    // IsEnabled.
    //
    // Even if a plug is inserted, it doesn't mean it's enabled.
    //
    // This flag indicates whether the plug is active and providing its benefits.
    IsEnabled bool `json:"isEnabled"`

    // IsVisible.
    //
    // A plug may theoretically provide benefits but not be visible - for instance, some older items use a plug's damage type perk to modify their own damage type. These, though they are not visible, still affect the item. This field indicates that state.
    //
    // An invisible plug, while it provides benefits if it is Enabled, cannot be directly modified by the user.
    IsVisible bool `json:"isVisible"`

    // PlugHash.
    //
    // The currently active plug, if any.
    //
    // Note that, because all plugs are statically defined, its effect on stats and perks can be statically determined using the plug item's definition. The stats and perks can be taken at face value on the plug item as the stats and perks it will provide to the user/item.
    PlugHash *uint32 `json:"plugHash"`
}
