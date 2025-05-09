// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Definitions_DestinyNodeActivationRequirement struct {
    // GridLevel.
    //
    // The Progression level on the Talent Grid required to activate this node.
    //
    // See DestinyTalentGridDefinition.progressionHash for the related Progression, and read DestinyProgressionDefinition's documentation to learn more about Progressions.
    GridLevel int32 `json:"gridLevel"`

    // MaterialRequirementHashes.
    //
    // The list of hash identifiers for material requirement sets: materials that are required for the node to be activated. See DestinyMaterialRequirementSetDefinition for more information about material requirements.
    //
    // In this case, only a single DestinyMaterialRequirementSetDefinition will be chosen from this list, and we won't know which one will be chosen until an instance of the item is created.
    MaterialRequirementHashes []uint32 `json:"materialRequirementHashes"`
}
