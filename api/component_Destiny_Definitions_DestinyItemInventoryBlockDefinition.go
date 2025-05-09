// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Definitions_DestinyItemInventoryBlockDefinition struct {
    // BucketTypeHash.
    //
    // The hash identifier for the DestinyInventoryBucketDefinition to which this item belongs. I should have named this "bucketHash", but too many things refer to it now. Sigh.
    BucketTypeHash uint32 `json:"bucketTypeHash"`

    // ExpirationTooltip.
    //
    // The tooltip message to show, if any, when the item expires.
    ExpirationTooltip string `json:"expirationTooltip"`

    // ExpiredInActivityMessage.
    //
    // If the item expires while playing in an activity, we show a different message.
    ExpiredInActivityMessage string `json:"expiredInActivityMessage"`

    // ExpiredInOrbitMessage.
    //
    // If the item expires in orbit, we show a... more different message. ("Consummate V's, consummate!")
    ExpiredInOrbitMessage string `json:"expiredInOrbitMessage"`

    // IsInstanceItem.
    //
    // If TRUE, this item is instanced. Otherwise, it is a generic item that merely has a quantity in a stack (like Glimmer).
    IsInstanceItem bool `json:"isInstanceItem"`

    // MaxStackSize.
    //
    // The maximum quantity of this item that can exist in a stack.
    MaxStackSize int32 `json:"maxStackSize"`

    // RecipeItemHash.
    //
    // A reference to the associated crafting 'recipe' item definition, if this item can be crafted.
    RecipeItemHash *uint32 `json:"recipeItemHash"`

    // RecoveryBucketTypeHash.
    //
    // If the item is picked up by the lost loot queue, this is the hash identifier for the DestinyInventoryBucketDefinition into which it will be placed. Again, I should have named this recoveryBucketHash instead.
    RecoveryBucketTypeHash uint32 `json:"recoveryBucketTypeHash"`

    // StackUniqueLabel.
    //
    // If this string is populated, you can't have more than one stack with this label in a given inventory. Note that this is different from the equipping block's unique label, which is used for equipping uniqueness.
    StackUniqueLabel string `json:"stackUniqueLabel"`

    // SuppressExpirationWhenObjectivesComplete.
    //
    // 
    SuppressExpirationWhenObjectivesComplete bool `json:"suppressExpirationWhenObjectivesComplete"`

    // TierType.
    //
    // The enumeration matching the tier type of the item to known values, again for convenience sake.
    TierType int32 `json:"tierType"`

    // TierTypeHash.
    //
    // The hash identifier for the Tier Type of the item, use to look up its DestinyItemTierTypeDefinition if you need to show localized data for the item's tier.
    TierTypeHash uint32 `json:"tierTypeHash"`

    // TierTypeName.
    //
    // The localized name of the tier type, which is a useful shortcut so you don't have to look up the definition every time. However, it's mostly a holdover from days before we had a DestinyItemTierTypeDefinition to refer to.
    TierTypeName string `json:"tierTypeName"`
}
