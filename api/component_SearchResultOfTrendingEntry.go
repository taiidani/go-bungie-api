// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type SearchResultOfTrendingEntry struct {
    // HasMore.
    //
    // 
    HasMore bool `json:"hasMore"`

    // Query.
    //
    // 
    Query Queries_PagedQuery `json:"query"`

    // ReplacementContinuationToken.
    //
    // 
    ReplacementContinuationToken string `json:"replacementContinuationToken"`

    // Results.
    //
    // 
    Results []Trending_TrendingEntry `json:"results"`

    // TotalResults.
    //
    // 
    TotalResults int32 `json:"totalResults"`

    // UseTotalResults.
    //
    // If useTotalResults is true, then totalResults represents an accurate count.
    //
    // If False, it does not, and may be estimated/only the size of the current page.
    //
    // Either way, you should probably always only trust hasMore.
    //
    // This is a long-held historical throwback to when we used to do paging with known total results. Those queries toasted our database, and we were left to hastily alter our endpoints and create backward- compatible shims, of which useTotalResults is one.
    UseTotalResults bool `json:"useTotalResults"`
}
