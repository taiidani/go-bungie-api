// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Components_Inventory_DestinyMaterialRequirementSetState struct {
    // MaterialRequirementSetHash.
    //
    // The hash identifier of the material requirement set. Use it to look up the DestinyMaterialRequirementSetDefinition.
    MaterialRequirementSetHash uint32 `json:"materialRequirementSetHash"`

    // MaterialRequirementStates.
    //
    // The dynamic state values for individual material requirements.
    MaterialRequirementStates []Destiny_Components_Inventory_DestinyMaterialRequirementState `json:"materialRequirementStates"`
}
