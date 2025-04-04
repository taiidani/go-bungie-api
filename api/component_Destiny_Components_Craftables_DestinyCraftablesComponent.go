// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Components_Craftables_DestinyCraftablesComponent struct {
    // Craftables.
    //
    // A map of craftable item hashes to craftable item state components.
    Craftables map[uint32]Destiny_Components_Craftables_DestinyCraftableComponent `json:"craftables"`

    // CraftingRootNodeHash.
    //
    // The hash for the root presentation node definition of craftable item categories.
    CraftingRootNodeHash uint32 `json:"craftingRootNodeHash"`
}
