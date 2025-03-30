// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Definitions_Vendors_DestinyVendorLocationDefinition struct {
    // DestinationHash.
    //
    // The hash identifier for a Destination at which this vendor may be located. Each destination where a Vendor may exist will only ever have a single entry.
    DestinationHash uint32 `json:"destinationHash"`

    // BackgroundImagePath.
    //
    // The relative path to the background image representing this Vendor at this location, for use in a banner.
    BackgroundImagePath string `json:"backgroundImagePath"`
}
