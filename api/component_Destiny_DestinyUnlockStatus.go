// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_DestinyUnlockStatus struct {
    // IsSet.
    //
    // Whether the unlock flag is set.
    IsSet bool `json:"isSet"`

    // UnlockHash.
    //
    // The hash identifier for the Unlock Flag. Use to lookup DestinyUnlockDefinition for static data. Not all unlocks have human readable data - in fact, most don't. But when they do, it can be very useful to show. Even if they don't have human readable data, you might be able to infer the meaning of an unlock flag with a bit of experimentation...
    UnlockHash uint32 `json:"unlockHash"`
}
