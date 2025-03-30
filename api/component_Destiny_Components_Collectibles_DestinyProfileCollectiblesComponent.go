// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Destiny_Components_Collectibles_DestinyProfileCollectiblesComponent struct {
    // NewnessFlaggedCollectibleHashes.
    //
    // The list of collectibles determined by the game as having been "recently" acquired.
    //
    // The game client itself actually controls this data, so I personally question whether anyone will get much use out of this: because we can't edit this value through the API. But in case anyone finds it useful, here it is.
    NewnessFlaggedCollectibleHashes []any `json:"newnessFlaggedCollectibleHashes"`

    // RecentCollectibleHashes.
    //
    // The list of collectibles determined by the game as having been "recently" acquired.
    RecentCollectibleHashes []any `json:"recentCollectibleHashes"`

    // Collectibles.
    //
    // 
    Collectibles any `json:"collectibles"`

    // CollectionBadgesRootNodeHash.
    //
    // The hash for the root presentation node definition of Collection Badges.
    CollectionBadgesRootNodeHash uint32 `json:"collectionBadgesRootNodeHash"`

    // CollectionCategoriesRootNodeHash.
    //
    // The hash for the root presentation node definition of Collection categories.
    CollectionCategoriesRootNodeHash uint32 `json:"collectionCategoriesRootNodeHash"`
}
