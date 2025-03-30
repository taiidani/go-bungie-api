// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Forum_PostResponse struct {
    // LatestReplyAuthorId.
    //
    // 
    LatestReplyAuthorId int64 `json:"latestReplyAuthorId"`

    // Popularity.
    //
    // 
    Popularity int32 `json:"popularity"`

    // UrlMediaType.
    //
    // 
    UrlMediaType int32 `json:"urlMediaType"`

    // UserHasRated.
    //
    // 
    UserHasRated bool `json:"userHasRated"`

    // UserRating.
    //
    // 
    UserRating int32 `json:"userRating"`

    // IsPinned.
    //
    // 
    IsPinned bool `json:"IsPinned"`

    // IgnoreStatus.
    //
    // 
    IgnoreStatus Ignores_IgnoreResponse `json:"ignoreStatus"`

    // LatestReplyPostId.
    //
    // 
    LatestReplyPostId int64 `json:"latestReplyPostId"`

    // Thumbnail.
    //
    // 
    Thumbnail string `json:"thumbnail"`

    // IsActive.
    //
    // 
    IsActive bool `json:"isActive"`

    // IsAnnouncement.
    //
    // 
    IsAnnouncement bool `json:"isAnnouncement"`

    // Locale.
    //
    // 
    Locale string `json:"locale"`

    // UserHasMutedPost.
    //
    // 
    UserHasMutedPost bool `json:"userHasMutedPost"`

    // LastReplyTimestamp.
    //
    // 
    LastReplyTimestamp string `json:"lastReplyTimestamp"`
}
