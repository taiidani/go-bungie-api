// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type Common_Models_CoreSetting struct {
    // ChildSettings.
    //
    // 
    ChildSettings []Common_Models_CoreSetting `json:"childSettings"`

    // DisplayName.
    //
    // 
    DisplayName string `json:"displayName"`

    // Identifier.
    //
    // 
    Identifier string `json:"identifier"`

    // ImagePath.
    //
    // 
    ImagePath string `json:"imagePath"`

    // IsDefault.
    //
    // 
    IsDefault bool `json:"isDefault"`

    // Summary.
    //
    // 
    Summary string `json:"summary"`
}
