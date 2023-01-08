package main

import (
	"fmt"
)

// UserGroupService create a new service to access userGroup related methods and functions.
type UserGroupService struct {
	Client *ApiClient
}

// UserGroup properties.
// Some properties are read-only, which means they are only accessible after creation
// and should not be passed as arguments in other methods.
type UserGroup struct {
	// ReadOnly
	Id              string `json:"usrgrpid,omitempty"`
	Name            string `json:"name"`
	Debug_mode      string `json:"debug_mode,omitempty"`
	Gui_access      string `json:"gui_access,omitempty"`
	Users_status    string `json:"users_status,omitempty"`
	Userdirectoryid string `json:"userdirectoryid,omitempty"`
}

// UserGroupPermission define a permission assignable to a UserGroup
type UserGroupPermission struct {
	Id         int `json:"id"`
	Permission int `json:"permission"`
}

// UserGroupPermission define a tag assignable to a UserGroup
type UserGroupTagPermission struct {
	GroupId string `json:"groupid"`
	Tag     string `json:"tag"`
	Value   string `json:"value"`
}

// UserGroupResponse define the server response format for UserGroup methods
type UserGroupResponse struct {
	Usrgrpids []string `json:"usrgrpids"`
}

// TODO : Refactor for Create and Update methods
type UserGroupExtendedParameters struct {
	Id              string                   `json:"usrgrpid,omitempty"`
	Name            string                   `json:"name,omitempty"`
	Debug_mode      int                      `json:"debug_mode,omitempty"`
	Gui_access      int                      `json:"gui_access,omitempty"`
	Users_status    int                      `json:"users_status,omitempty"`
	Userdirectoryid string                   `json:"userdirectoryid,omitempty"`
	Rights          []UserGroupPermission    `json:"rights,omitempty"`
	Tag_filters     []UserGroupTagPermission `json:"tag_filters,omitempty"`
	// TODO - Waiting for implemenation of Users
	// Users                []map[string]string      `json:"users,omitempty"`
}

// Create is used to create a new UserGroup.
func (u *UserGroupService) Create(params *UserGroupExtendedParameters) (*UserGroupResponse, error) {
	if params.Name == "" {
		return nil, fmt.Errorf("Missing required field 'Name' in the given object.\nObject passed : %v", params)
	}

	req := u.Client.NewRequest("usergroup.create", params)

	res, err := u.Client.Post(req)
	if err != nil {
		return nil, err
	}

	g := UserGroupResponse{}
	err = u.Client.ConvertResponse(*res, &g)
	if err != nil {
		if u.Client.ResourceAlreadyExist("User group", params.Name, res.Error) {
			fmt.Println(res.Error.Data)
		} else {
			return nil, err
		}
	}

	return &g, nil
}

// Delete is used to delete one or multiples UserGroups.
func (u *UserGroupService) Delete(ids []string) (*UserGroupResponse, error) {
	req := u.Client.NewRequest("usergroup.delete", ids)

	res, err := u.Client.Post(req)
	if err != nil {
		return nil, err
	}

	g := UserGroupResponse{}
	err = u.Client.ConvertResponse(*res, &g)
	if err != nil {
		return nil, err
	}

	return &g, nil
}

// UserGroupGetParameters define the properties used to search UserGroup(s)
// Properties using the 'omitempty' json parameters are optional
type UserGroupGetParameters struct {
	Status                 int         `json:"status,omitempty"`
	Userids                []string    `json:"userids,omitempty"`
	Usrgrpids              []string    `json:"usrgrpids,omitempty"`
	SelectTagFilters       interface{} `json:"selectTagFilters,omitempty"`
	SelectUsers            interface{} `json:"selectUsers,omitempty"`
	SelectRights           interface{} `json:"selectRights,omitempty"`
	LimitSelects           int         `json:"limitSelects,omitempty"`
	Sortfield              []string    `json:"sortfield,omitempty"`
	CountOutput            bool        `json:"countOutput,omitempty"`
	Editable               bool        `json:"editable,omitempty"`
	ExcludeSearch          bool        `json:"excludeSearch,omitempty"`
	Filter                 interface{} `json:"filter,omitempty"`
	Limit                  int         `json:"limit,omitempty"`
	Output                 interface{} `json:"output,omitempty"`
	Preservekeys           bool        `json:"preservekeys,omitempty"`
	Search                 interface{} `json:"search,omitempty"`
	SearchByAny            bool        `json:"searchByAny,omitempty"`
	SearchWildcardsEnabled bool        `json:"searchWildcardsEnabled,omitempty"`
	Sortorder              []string    `json:"sortorder,omitempty"`
	StartSearch            bool        `json:"startSearch,omitempty"`
}

// List is used to retrieve all UserGroups.
func (u *UserGroupService) List() ([]*UserGroup, error) {
	params := UserGroupGetParameters{
		Output: "extend",
	}

	req := u.Client.NewRequest("usergroup.get", params)

	res, err := u.Client.Post(req)
	if err != nil {
		return nil, err
	}

	g := make([]*UserGroup, 0)
	err = u.Client.ConvertResponse(*res, &g)
	if err != nil {
		return nil, err
	}

	return g, nil
}

// Get is used to retrieve one or multiples UserGroups matching the given criteria(s).
func (u *UserGroupService) Get(params *UserGroupGetParameters) ([]*UserGroup, error) {
	req := u.Client.NewRequest("usergroup.get", params)

	res, err := u.Client.Post(req)
	if err != nil {
		return nil, err
	}

	g := make([]*UserGroup, 0)
	err = u.Client.ConvertResponse(*res, &g)
	if err != nil {
		return nil, err
	}

	return g, nil
}

// Update is used to update or overwrite properties from an existing UserGroup.
func (u *UserGroupService) Update(params *UserGroupExtendedParameters) (*UserGroupResponse, error) {
	if params.Id == "" {
		return nil, fmt.Errorf("Missing required field 'Id' in the given object.\nObject passed : %v", params)
	}

	req := u.Client.NewRequest("usergroup.update", params)

	res, err := u.Client.Post(req)
	if err != nil {
		return nil, err
	}

	g := UserGroupResponse{}
	err = u.Client.ConvertResponse(*res, &g)
	if err != nil {
		return nil, err
	}

	return &g, nil
}
