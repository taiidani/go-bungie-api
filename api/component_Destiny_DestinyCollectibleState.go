// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

const (
    // Destiny_DestinyCollectibleStateNone represents the None enum in the Destiny.DestinyCollectibleState component.
    Destiny_DestinyCollectibleStateNone int32 = 0

    // Destiny_DestinyCollectibleStateNotAcquired represents the NotAcquired enum in the Destiny.DestinyCollectibleState component.
    //
    // If this flag is set, you have not yet obtained this collectible.
    Destiny_DestinyCollectibleStateNotAcquired int32 = 1

    // Destiny_DestinyCollectibleStateObscured represents the Obscured enum in the Destiny.DestinyCollectibleState component.
    //
    // If this flag is set, the item is "obscured" to you: you can/should use the alternate item hash found in DestinyCollectibleDefinition.stateInfo.obscuredOverrideItemHash when displaying this collectible instead of the default display info.
    Destiny_DestinyCollectibleStateObscured int32 = 2

    // Destiny_DestinyCollectibleStateInvisible represents the Invisible enum in the Destiny.DestinyCollectibleState component.
    //
    // If this flag is set, the collectible should not be shown to the user.
    //
    // Please do consider honoring this flag. It is used - for example - to hide items that a person didn't get from the Eververse. I can't prevent these from being returned in definitions, because some people may have acquired them and thus they should show up: but I would hate for people to start feeling some variant of a Collector's Remorse about these items, and thus increasing their purchasing based on that compulsion. That would be a very unfortunate outcome, and one that I wouldn't like to see happen. So please, whether or not I'm your mom, consider honoring this flag and don't show people invisible collectibles.
    Destiny_DestinyCollectibleStateInvisible int32 = 4

    // Destiny_DestinyCollectibleStateCannotAffordMaterialRequirements represents the CannotAffordMaterialRequirements enum in the Destiny.DestinyCollectibleState component.
    //
    // If this flag is set, the collectible requires payment for creating an instance of the item, and you are lacking in currency. Bring the benjamins next time. Or spinmetal. Whatever.
    Destiny_DestinyCollectibleStateCannotAffordMaterialRequirements int32 = 8

    // Destiny_DestinyCollectibleStateInventorySpaceUnavailable represents the InventorySpaceUnavailable enum in the Destiny.DestinyCollectibleState component.
    //
    // If this flag is set, you can't pull this item out of your collection because there's no room left in your inventory.
    Destiny_DestinyCollectibleStateInventorySpaceUnavailable int32 = 16

    // Destiny_DestinyCollectibleStateUniquenessViolation represents the UniquenessViolation enum in the Destiny.DestinyCollectibleState component.
    //
    // If this flag is set, you already have one of these items and can't have a second one.
    Destiny_DestinyCollectibleStateUniquenessViolation int32 = 32

    // Destiny_DestinyCollectibleStatePurchaseDisabled represents the PurchaseDisabled enum in the Destiny.DestinyCollectibleState component.
    //
    // If this flag is set, the ability to pull this item out of your collection has been disabled.
    Destiny_DestinyCollectibleStatePurchaseDisabled int32 = 64
)
