// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Definitions_DestinyItemMetricBlockDefinition struct {
    // AvailableMetricCategoryNodeHashes.
    //
    // Hash identifiers for any DestinyPresentationNodeDefinition entry that can be used to list available metrics. Any metric listed directly below these nodes, or in any of these nodes' children will be made available for selection.
    AvailableMetricCategoryNodeHashes []uint32 `json:"availableMetricCategoryNodeHashes"`
}
