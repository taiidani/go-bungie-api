// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Components_Records_DestinyCharacterRecordsComponent struct {
    // FeaturedRecordHashes.
    //
    // 
    FeaturedRecordHashes []uint32 `json:"featuredRecordHashes"`

    // RecordCategoriesRootNodeHash.
    //
    // The hash for the root presentation node definition of Triumph categories.
    RecordCategoriesRootNodeHash uint32 `json:"recordCategoriesRootNodeHash"`

    // RecordSealsRootNodeHash.
    //
    // The hash for the root presentation node definition of Triumph Seals.
    RecordSealsRootNodeHash uint32 `json:"recordSealsRootNodeHash"`

    // Records.
    //
    // 
    Records map[uint32]Destiny_Components_Records_DestinyRecordComponent `json:"records"`
}
