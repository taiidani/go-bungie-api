// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Definitions_DestinyPlugItemCraftingRequirements struct {
    // MaterialRequirementHashes.
    //
    // 
    MaterialRequirementHashes []uint32 `json:"materialRequirementHashes"`

    // RequiredLevel.
    //
    // If the plug has a known level requirement, it'll be available here.
    RequiredLevel *int32 `json:"requiredLevel"`

    // UnlockRequirements.
    //
    // 
    UnlockRequirements []Destiny_Definitions_DestinyPlugItemCraftingUnlockRequirement `json:"unlockRequirements"`
}
