// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type User_EmailSubscriptionDefinition struct {
    // Localization.
    //
    // A dictionary of localized text for the EMail Opt-in setting, keyed by the locale.
    Localization any `json:"localization"`

    // Name.
    //
    // The unique identifier for this subscription.
    Name string `json:"name"`

    // Value.
    //
    // The bitflag value for this subscription. Should be a unique power of two value.
    Value int64 `json:"value"`
}
