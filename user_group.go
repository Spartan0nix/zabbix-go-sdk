package main

import (
	"fmt"
)

type UserGroupService struct {
	Client *ApiClient
}

type UserGroup struct {
	Id              string `json:"usrgrpid"`
	Name            string `json:"name"`
	Debug_mode      string `json:"debug_mode"`
	Gui_access      string `json:"gui_access"`
	Users_status    string `json:"users_status"`
	Userdirectoryid string `json:"userdirectoryid"`
}

type UserGroupPermission struct {
	Id         int `json:"id"`
	Permission int `json:"permission"`
}

type UserGroupTagPermission struct {
	GroupId string `json:"groupid"`
	Tag     string `json:"tag"`
	Value   string `json:"value"`
}

func (u *UserGroupService) List() ([]UserGroup, error) {
	params := map[string]interface{}{
		"output": "extend",
	}

	req := u.Client.NewRequest("usergroup.get", params)

	res, err := u.Client.Post(req)
	if err != nil {
		return nil, err
	}

	g := make([]UserGroup, 0)
	err = u.Client.ConvertResponse(*res, &g)
	if err != nil {
		return nil, err
	}

	return g, nil
}

type UserGroupGetParameters struct {
	Status                    int         `json:"status,omitempty"`
	Userids                   []string    `json:"userids,omitempty"`
	Usrgrpids                 []string    `json:"usrgrpids,omitempty"`
	SelectTagFilters          interface{} `json:"selectTagFilters,omitempty"`
	SelectUsers               interface{} `json:"selectUsers,omitempty"`
	SelectHostGroupRights     interface{} `json:"selectHostGroupRights,omitempty"`
	SelectTemplateGroupRights interface{} `json:"selectTemplateGroupRights,omitempty"`
	LimitSelects              int         `json:"limitSelects,omitempty"`
	Sortfield                 []string    `json:"sortfield,omitempty"`
	CountOutput               bool        `json:"countOutput,omitempty"`
	Editable                  bool        `json:"editable,omitempty"`
	ExcludeSearch             bool        `json:"excludeSearch,omitempty"`
	Filter                    interface{} `json:"filter,omitempty"`
	Limit                     int         `json:"limit,omitempty"`
	Output                    interface{} `json:"output,omitempty"`
	Preservekeys              bool        `json:"preservekeys,omitempty"`
	Search                    interface{} `json:"search,omitempty"`
	SearchByAny               bool        `json:"searchByAny,omitempty"`
	SearchWildcardsEnabled    bool        `json:"searchWildcardsEnabled,omitempty"`
	Sortorder                 []string    `json:"sortorder,omitempty"`
	StartSearch               bool        `json:"startSearch,omitempty"`
}

func (u *UserGroupService) Get(params *UserGroupGetParameters) ([]UserGroup, error) {
	req := u.Client.NewRequest("usergroup.get", params)

	res, err := u.Client.Post(req)
	if err != nil {
		return nil, err
	}

	g := make([]UserGroup, 0)
	err = u.Client.ConvertResponse(*res, &g)
	if err != nil {
		return nil, err
	}

	return g, nil
}

type UserGroupCreateParameters struct {
	Name            string                   `json:"name"`
	Debug_mode      int                      `json:"debug_mode,omitempty"`
	Gui_access      int                      `json:"gui_access,omitempty"`
	Users_status    int                      `json:"users_status,omitempty"`
	Userdirectoryid string                   `json:"userdirectoryid,omitempty"`
	Rights          []UserGroupPermission    `json:"rights,omitempty"`
	Tag_filters     []UserGroupTagPermission `json:"tag_filters,omitempty"`
	// TODO - Waiting for implemenation of Users
	// Users                []map[string]string      `json:"users,omitempty"`
}

type UserGroupResponse struct {
	Usrgrpids []string `json:"usrgrpids"`
}

func (u *UserGroupService) Create(params *UserGroupCreateParameters) (*UserGroupResponse, error) {
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

type UserGroupUpdateParameters struct {
	Id              string                   `json:"usrgrpid"`
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

func (u *UserGroupService) Update(params *UserGroupUpdateParameters) (*UserGroupResponse, error) {
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
