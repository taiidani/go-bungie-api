// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type User_EmailViewDefinition struct {
    // ViewSettings.
    //
    // The ordered list of settings to show in this view.
    ViewSettings []User_EmailViewDefinitionSetting `json:"viewSettings"`

    // Name.
    //
    // The identifier for this view.
    Name string `json:"name"`
}
