// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Components_Craftables_DestinyCraftableSocketPlugComponent struct {
    // FailedRequirementIndexes.
    //
    // Index into the unlock requirements to display failure descriptions
    FailedRequirementIndexes []int32 `json:"failedRequirementIndexes"`

    // PlugItemHash.
    //
    // 
    PlugItemHash uint32 `json:"plugItemHash"`
}
