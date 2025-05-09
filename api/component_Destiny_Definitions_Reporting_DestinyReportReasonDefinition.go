// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Definitions_Reporting_DestinyReportReasonDefinition struct {
    // DisplayProperties.
    //
    // Many Destiny*Definition contracts - the "first order" entities of Destiny that have their own tables in the Manifest Database - also have displayable information. This is the base class for that display information.
    DisplayProperties Destiny_Definitions_Common_DestinyDisplayPropertiesDefinition `json:"displayProperties"`

    // ReasonHash.
    //
    // The identifier for the reason: they are only guaranteed unique under the Category in which they are found.
    ReasonHash uint32 `json:"reasonHash"`
}
