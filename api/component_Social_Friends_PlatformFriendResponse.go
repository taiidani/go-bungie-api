// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Social_Friends_PlatformFriendResponse struct {
    // CurrentPage.
    //
    // 
    CurrentPage int32 `json:"currentPage"`

    // HasMore.
    //
    // 
    HasMore bool `json:"hasMore"`

    // ItemsPerPage.
    //
    // 
    ItemsPerPage int32 `json:"itemsPerPage"`

    // PlatformFriends.
    //
    // 
    PlatformFriends []Social_Friends_PlatformFriend `json:"platformFriends"`
}
