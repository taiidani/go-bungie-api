// This file was generated by https://github.com/taiidani/go-bungie-api/generate.
// DO NOT EDIT THIS FILE.
package api

type GroupsV2_GroupEditHistory struct {
    // MottoEditors.
    //
    // 
    MottoEditors int64 `json:"mottoEditors"`

    // Motto.
    //
    // 
    Motto string `json:"motto"`

    // NameEditors.
    //
    // 
    NameEditors int64 `json:"nameEditors"`

    // ClanCallsignEditors.
    //
    // 
    ClanCallsignEditors int64 `json:"clanCallsignEditors"`

    // GroupEditors.
    //
    // 
    GroupEditors []User_UserInfoCard `json:"groupEditors"`

    // GroupId.
    //
    // 
    GroupId int64 `json:"groupId"`

    // Name.
    //
    // 
    Name string `json:"name"`

    // AboutEditors.
    //
    // 
    AboutEditors int64 `json:"aboutEditors"`

    // ClanCallsign.
    //
    // 
    ClanCallsign string `json:"clanCallsign"`

    // About.
    //
    // 
    About string `json:"about"`

    // EditDate.
    //
    // 
    EditDate string `json:"editDate"`
}
