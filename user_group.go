package zabbixgosdk

// import (
// 	"fmt"
// )

// // UserGroupDebugMode define whether debug mode should be enabled or not
// type UserGroupDebugMode string

// // UserGroupGuiAccess define the type of frontend authentification for users in the group
// type UserGroupGuiAccess string

// // UserGroupStatus define if the group should be enabled or not
// type UserGroupStatus string

// // UserGroupPermissionType define if the available type of permission level
// type UserGroupPermissionType string

// const (
// 	// Default
// 	UserGroupDebugDisabled UserGroupDebugMode = "0"
// 	UserGroupDebugEnabled  UserGroupDebugMode = "1"

// 	// Default
// 	UserGroupSystemDefault UserGroupGuiAccess = "0"
// 	UserGroupInternal      UserGroupGuiAccess = "1"
// 	UserGroupLDAP          UserGroupGuiAccess = "2"
// 	UserGroupDisable       UserGroupGuiAccess = "3"

// 	// Default
// 	UserGroupStatusEnabled  UserGroupStatus = "0"
// 	UserGroupStatusDisabled UserGroupStatus = "1"

// 	UserGroupAccessDenied    UserGroupPermissionType = "0"
// 	UserGroupAccessReadOnly  UserGroupPermissionType = "2"
// 	UserGroupAccessReadWrite UserGroupPermissionType = "3"
// )

// // UserGroupService create a new service to access userGroup related methods and functions.
// type UserGroupService struct {
// 	Client *ApiClient
// }

// // UserGroup properties.
// // Some properties are read-only, which means they are only accessible after creation
// // and should not be passed as arguments in other methods.
// type UserGroup struct {
// 	// ReadOnly
// 	UsrGrpId        string             `json:"usrgrpid"`
// 	Name            string             `json:"name"`
// 	DebugMode       UserGroupDebugMode `json:"debug_mode"`
// 	GuiAccess       UserGroupGuiAccess `json:"gui_access"`
// 	UsersStatus     UserGroupStatus    `json:"users_status"`
// 	UserDirectoryId string             `json:"userdirectoryid"`
// }

// // UserGroupPermission define a permission assignable to a UserGroup
// type UserGroupPermission struct {
// 	Id         string                  `json:"id"`
// 	Permission UserGroupPermissionType `json:"permission"`
// }

// // UserGroupPermission define a tag assignable to a UserGroup
// type UserGroupTagPermission struct {
// 	GroupId string `json:"groupid"`
// 	Tag     string `json:"tag"`
// 	Value   string `json:"value"`
// }

// // UserGroupResponse define the server response format for UserGroup methods
// type UserGroupResponse struct {
// 	Usrgrpids []string `json:"usrgrpids"`
// }

// type UserGroupCreateParameters struct {
// 	Name            string                    `json:"name"`
// 	DebugMode       UserGroupDebugMode        `json:"debug_mode,omitempty"`
// 	GuiAccess       UserGroupGuiAccess        `json:"gui_access,omitempty"`
// 	UsersStatus     UserGroupStatus           `json:"users_status,omitempty"`
// 	UserDirectoryId string                    `json:"userdirectoryid,omitempty"`
// 	Rights          []*UserGroupPermission    `json:"rights,omitempty"`
// 	TagFilters      []*UserGroupTagPermission `json:"tag_filters,omitempty"`
// 	// TODO - Waiting for implemenation of Users
// 	// Users                []map[string]string      `json:"users,omitempty"`
// }

// // Create is used to create a new UserGroup.
// func (u *UserGroupService) Create(p *UserGroupCreateParameters) (*UserGroupResponse, error) {
// 	req := u.Client.NewRequest("usergroup.create", p)

// 	res, err := u.Client.Post(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	g := UserGroupResponse{}
// 	err = u.Client.ConvertResponse(*res, &g)
// 	if err != nil {
// 		if u.Client.ResourceAlreadyExist("User group", p.Name, res.Error) {
// 			fmt.Println(res.Error.Data)
// 		} else {
// 			return nil, err
// 		}
// 	}

// 	return &g, nil
// }

// // Delete is used to delete one or multiples UserGroups.
// func (u *UserGroupService) Delete(ids []string) (*UserGroupResponse, error) {
// 	req := u.Client.NewRequest("usergroup.delete", ids)

// 	res, err := u.Client.Post(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	g := UserGroupResponse{}
// 	err = u.Client.ConvertResponse(*res, &g)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &g, nil
// }

// // UserGroupGetParameters define the properties used to search UserGroup(s)
// // Properties using the 'omitempty' json parameters are optional
// type UserGroupGetParameters struct {
// 	Status                 int         `json:"status,omitempty"`
// 	UserIds                []string    `json:"userids,omitempty"`
// 	UsrGrpIds              []string    `json:"usrgrpids,omitempty"`
// 	SelectTagFilters       interface{} `json:"selectTagFilters,omitempty"`
// 	SelectUsers            interface{} `json:"selectUsers,omitempty"`
// 	SelectRights           interface{} `json:"selectRights,omitempty"`
// 	LimitSelects           int         `json:"limitSelects,omitempty"`
// 	SortField              []string    `json:"sortfield,omitempty"`
// 	CountOutput            bool        `json:"countOutput,omitempty"`
// 	Editable               bool        `json:"editable,omitempty"`
// 	ExcludeSearch          bool        `json:"excludeSearch,omitempty"`
// 	Filter                 interface{} `json:"filter,omitempty"`
// 	Limit                  int         `json:"limit,omitempty"`
// 	Output                 interface{} `json:"output,omitempty"`
// 	PreserveKeys           bool        `json:"preservekeys,omitempty"`
// 	Search                 interface{} `json:"search,omitempty"`
// 	SearchByAny            bool        `json:"searchByAny,omitempty"`
// 	SearchWildcardsEnabled bool        `json:"searchWildcardsEnabled,omitempty"`
// 	Sortorder              []string    `json:"sortorder,omitempty"`
// 	StartSearch            bool        `json:"startSearch,omitempty"`
// }

// // List is used to retrieve all UserGroups.
// func (u *UserGroupService) List() ([]*UserGroup, error) {
// 	p := UserGroupGetParameters{
// 		Output: "extend",
// 	}

// 	req := u.Client.NewRequest("usergroup.get", p)

// 	res, err := u.Client.Post(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	g := make([]*UserGroup, 0)
// 	err = u.Client.ConvertResponse(*res, &g)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return g, nil
// }

// // Get is used to retrieve one or multiples UserGroups matching the given criteria(s).
// func (u *UserGroupService) Get(p *UserGroupGetParameters) ([]*UserGroup, error) {
// 	req := u.Client.NewRequest("usergroup.get", p)

// 	res, err := u.Client.Post(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	g := make([]*UserGroup, 0)
// 	err = u.Client.ConvertResponse(*res, &g)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return g, nil
// }

// type UserGroupUpdateParameters struct {
// 	UsrGrpId        string                    `json:"usrgrpid"`
// 	Name            string                    `json:"name,omitempty"`
// 	DebugMode       UserGroupDebugMode        `json:"debug_mode,omitempty"`
// 	GuiAccess       UserGroupGuiAccess        `json:"gui_access,omitempty"`
// 	UsersStatus     UserGroupStatus           `json:"users_status,omitempty"`
// 	UserDirectoryId string                    `json:"userdirectoryid,omitempty"`
// 	Rights          []*UserGroupPermission    `json:"rights,omitempty"`
// 	TagFilters      []*UserGroupTagPermission `json:"tag_filters,omitempty"`
// 	// TODO - Waiting for implemenation of Users
// 	// Users                []map[string]string      `json:"users,omitempty"`
// }

// // Update is used to update or overwrite properties from an existing UserGroup.
// func (u *UserGroupService) Update(p *UserGroupUpdateParameters) (*UserGroupResponse, error) {
// 	if p.UsrGrpId == "" {
// 		return nil, fmt.Errorf("missing required field 'Id' in the given object.\nObject passed : %v", p)
// 	}

// 	req := u.Client.NewRequest("usergroup.update", p)

// 	res, err := u.Client.Post(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	g := UserGroupResponse{}
// 	err = u.Client.ConvertResponse(*res, &g)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &g, nil
// }
