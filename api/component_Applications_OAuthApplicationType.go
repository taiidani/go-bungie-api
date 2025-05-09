// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

const (
    // Applications_OAuthApplicationTypeNone represents the None enum in the Applications.OAuthApplicationType component.
    Applications_OAuthApplicationTypeNone int32 = 0

    // Applications_OAuthApplicationTypeConfidential represents the Confidential enum in the Applications.OAuthApplicationType component.
    //
    // Indicates the application is server based and can keep its secrets from end users and other potential snoops.
    Applications_OAuthApplicationTypeConfidential int32 = 1

    // Applications_OAuthApplicationTypePublic represents the Public enum in the Applications.OAuthApplicationType component.
    //
    // Indicates the application runs in a public place, and it can't be trusted to keep a secret.
    Applications_OAuthApplicationTypePublic int32 = 2
)
