// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Tokens_RewardAvailabilityModel struct {
    // RedemptionEndDate.
    //
    // 
    RedemptionEndDate string `json:"RedemptionEndDate"`

    // DecryptedToken.
    //
    // 
    DecryptedToken string `json:"DecryptedToken"`

    // HasExistingCode.
    //
    // 
    HasExistingCode bool `json:"HasExistingCode"`

    // HasOffer.
    //
    // 
    HasOffer bool `json:"HasOffer"`

    // IsOffer.
    //
    // 
    IsOffer bool `json:"IsOffer"`

    // OfferApplied.
    //
    // 
    OfferApplied bool `json:"OfferApplied"`

    // CollectibleDefinitions.
    //
    // 
    CollectibleDefinitions []Tokens_CollectibleDefinitions `json:"CollectibleDefinitions"`

    // GameEarnByDate.
    //
    // 
    GameEarnByDate string `json:"GameEarnByDate"`

    // RecordDefinitions.
    //
    // 
    RecordDefinitions []Destiny_Definitions_Records_DestinyRecordDefinition `json:"RecordDefinitions"`

    // IsLoyaltyReward.
    //
    // 
    IsLoyaltyReward bool `json:"IsLoyaltyReward"`

    // ShopifyEndDate.
    //
    // 
    ShopifyEndDate string `json:"ShopifyEndDate"`
}
