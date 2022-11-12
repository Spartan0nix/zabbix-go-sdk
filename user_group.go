package main

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
	Status                    int         `json:"status"`
	Userids                   []string    `json:"userids"`
	Usrgrpids                 []string    `json:"usrgrpids"`
	SelectTagFilters          interface{} `json:"selectTagFilters"`
	SelectUsers               interface{} `json:"selectUsers"`
	SelectHostGroupRights     interface{} `json:"selectHostGroupRights"`
	SelectTemplateGroupRights interface{} `json:"selectTemplateGroupRights"`
	LimitSelects              int         `json:"limitSelects"`
	Sortfield                 []string    `json:"sortfield"`
	CountOutput               bool        `json:"countOutput"`
	Editable                  bool        `json:"editable"`
	ExcludeSearch             bool        `json:"excludeSearch"`
	Filter                    interface{} `json:"filter"`
	Limit                     int         `json:"limit"`
	Output                    interface{} `json:"output"`
	Preservekeys              bool        `json:"preservekeys"`
	Search                    interface{} `json:"search"`
	SearchByAny               bool        `json:"searchByAny"`
	SearchWildcardsEnabled    bool        `json:"searchWildcardsEnabled"`
	Sortorder                 []string    `json:"sortorder"`
	StartSearch               bool        `json:"startSearch"`
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
