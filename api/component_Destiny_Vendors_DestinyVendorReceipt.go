// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Vendors_DestinyVendorReceipt struct {
    // CurrencyPaid.
    //
    // The amount paid for the item, in terms of items that were consumed in the purchase and their quantity.
    CurrencyPaid []Destiny_DestinyItemQuantity `json:"currencyPaid"`

    // ExpiresOn.
    //
    // The date at which this receipt is rendered invalid.
    ExpiresOn string `json:"expiresOn"`

    // ItemReceived.
    //
    // The item that was received, and its quantity.
    ItemReceived any `json:"itemReceived"`

    // LicenseUnlockHash.
    //
    // The unlock flag used to determine whether you still have the purchased item.
    LicenseUnlockHash uint32 `json:"licenseUnlockHash"`

    // PurchasedByCharacterId.
    //
    // The ID of the character who made the purchase.
    PurchasedByCharacterId int64 `json:"purchasedByCharacterId"`

    // RefundPolicy.
    //
    // Whether you can get a refund, and what happens in order for the refund to be received. See the DestinyVendorItemRefundPolicy enum for details.
    RefundPolicy int32 `json:"refundPolicy"`

    // SequenceNumber.
    //
    // The identifier of this receipt.
    SequenceNumber int32 `json:"sequenceNumber"`

    // TimeToExpiration.
    //
    // The seconds since epoch at which this receipt is rendered invalid.
    TimeToExpiration int64 `json:"timeToExpiration"`
}
