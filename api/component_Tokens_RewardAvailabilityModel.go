// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Tokens_RewardAvailabilityModel struct {
    // CollectibleDefinitions.
    //
    // 
    CollectibleDefinitions []Tokens_CollectibleDefinitions `json:"CollectibleDefinitions"`

    // DecryptedToken.
    //
    // 
    DecryptedToken string `json:"DecryptedToken"`

    // GameEarnByDate.
    //
    // 
    GameEarnByDate string `json:"GameEarnByDate"`

    // HasExistingCode.
    //
    // 
    HasExistingCode bool `json:"HasExistingCode"`

    // HasOffer.
    //
    // 
    HasOffer bool `json:"HasOffer"`

    // IsLoyaltyReward.
    //
    // 
    IsLoyaltyReward bool `json:"IsLoyaltyReward"`

    // IsOffer.
    //
    // 
    IsOffer bool `json:"IsOffer"`

    // OfferApplied.
    //
    // 
    OfferApplied bool `json:"OfferApplied"`

    // RecordDefinitions.
    //
    // 
    RecordDefinitions []Destiny_Definitions_Records_DestinyRecordDefinition `json:"RecordDefinitions"`

    // RedemptionEndDate.
    //
    // 
    RedemptionEndDate string `json:"RedemptionEndDate"`

    // ShopifyEndDate.
    //
    // 
    ShopifyEndDate *string `json:"ShopifyEndDate"`
}
