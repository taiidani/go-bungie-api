// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Definitions_DestinyVendorAcceptedItemDefinition struct {
    // AcceptedInventoryBucketHash.
    //
    // The "source" bucket for a transfer. When a user wants to transfer an item, the appropriate DestinyVendorDefinition's acceptedItems property is evaluated, looking for an entry where acceptedInventoryBucketHash matches the bucket that the item being transferred is currently located. If it exists, the item will be transferred into whatever bucket is defined by destinationInventoryBucketHash.
    AcceptedInventoryBucketHash uint32 `json:"acceptedInventoryBucketHash"`

    // DestinationInventoryBucketHash.
    //
    // This is the bucket where the item being transferred will be put, given that it was being transferred *from* the bucket defined in acceptedInventoryBucketHash.
    DestinationInventoryBucketHash uint32 `json:"destinationInventoryBucketHash"`
}
