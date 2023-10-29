package zabbixgosdk

// import (
// 	"fmt"
// 	"regexp"
// 	"strings"
// )

// // UserMacroService create a new service to access macro related methods and functions.
// type UserMacroService struct {
// 	Client *ApiClient
// }

// UserMacroType define the available macro types.
type UserMacroType string

const (
	UserMacroText   UserMacroType = "0"
	UserMacroSecret UserMacroType = "1"
	UserMacroVault  UserMacroType = "2"
)

// // HostMacro properties.
// // Some properties are read-only, which means they are only accessible after creation
// // and should not be passed as arguments in other methods.
// type HostMacro struct {
// 	// ReadOnly
// 	HostMacroId string        `json:"hostmacroid,omitempty"`
// 	HostId      string        `json:"hostid,omitempty"`
// 	Macro       string        `json:"macro"`
// 	Value       string        `json:"value"`
// 	Type        UserMacroType `json:"type,omitempty"`
// 	Description string        `json:"description,omitempty"`
// }

// // GlobalMacro properties.
// // Some properties are read-only, which means they are only accessible after creation
// // and should not be passed as arguments in other methods.
// type GlobalMacro struct {
// 	// ReadOnly
// 	GlobalMacroId string        `json:"globalmacroid,omitempty"`
// 	Macro         string        `json:"macro"`
// 	Value         string        `json:"value"`
// 	Type          UserMacroType `json:"type,omitempty"`
// 	Description   string        `json:"description,omitempty"`
// }

// // HostMacroResponse define the server response format for HostMacro methods.
// type HostMacroResponse struct {
// 	HostMacroIds []string `json:"hostmacroids"`
// }

// // GlobalMacroResponse define the server response format for GlobalMacro methods.
// type GlobalMacroResponse struct {
// 	GlobalMacroIds []string `json:"globalmacroids"`
// }

// // ValidateMacro is used to validate HostMacro format compliance.
// func (h *HostMacro) ValidateMacro() error {
// 	re := regexp.MustCompile(`^{\$[A-Z0-9_.]*}$`)
// 	if !re.Match([]byte(h.Macro)) {
// 		return fmt.Errorf("the following macro '%s' does not complies with the required format.\n- Start with : {$\n- End with : }\n- Contain only : A-Z // 0-9 // _ //", h.Macro)
// 	}

// 	return nil
// }

// // ValidateMacro is used to validate GlobalMacro format compliance.
// func (g *GlobalMacro) ValidateMacro() error {
// 	re := regexp.MustCompile(`^{\$[A-Z0-9_.]*}$`)
// 	if !re.Match([]byte(g.Macro)) {
// 		return fmt.Errorf("the following macro '%s' does not complies with the required format.\n- Start with : {$\n- End with : }\n- Contain only : A-Z // 0-9 // _ //", g.Macro)
// 	}

// 	return nil
// }

// // Create is used to create a new HostMacro.
// func (u *UserMacroService) Create(h *HostMacro) (*HostMacroResponse, error) {
// 	if h.HostId == "" {
// 		return nil, fmt.Errorf("missing required field 'HostId' in the given object.\nObject passed : %v", h)
// 	}

// 	// Prevent user from sending an Id when creating an HostMacro
// 	h.HostMacroId = ""

// 	if err := h.ValidateMacro(); err != nil {
// 		return nil, err
// 	}

// 	req := u.Client.NewRequest("usermacro.create", h)

// 	res, err := u.Client.Post(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	r := HostMacroResponse{}
// 	err = u.Client.ConvertResponse(*res, &r)
// 	if err != nil {
// 		escaped_macro := strings.Replace(h.Macro, "$", `\$`, -1)
// 		if u.Client.ResourceAlreadyExist("Macro", escaped_macro, res.Error) {
// 			fmt.Println(res.Error.Data)
// 		} else {
// 			return nil, err
// 		}
// 	}

// 	return &r, nil
// }

// // CreateGlobal is used to create a new GlobalMacro.
// func (u *UserMacroService) CreateGlobal(g *GlobalMacro) (*GlobalMacroResponse, error) {
// 	if err := g.ValidateMacro(); err != nil {
// 		return nil, err
// 	}

// 	// Prevent user from sending an Id when creating a GlobalMacro
// 	g.GlobalMacroId = ""

// 	req := u.Client.NewRequest("usermacro.createglobal", g)

// 	res, err := u.Client.Post(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	r := GlobalMacroResponse{}
// 	err = u.Client.ConvertResponse(*res, &r)
// 	if err != nil {
// 		escaped_macro := strings.Replace(g.Macro, "$", `\$`, -1)
// 		if u.Client.ResourceAlreadyExist("Macro", escaped_macro, res.Error) {
// 			fmt.Println(res.Error.Data)
// 		} else {
// 			return nil, err
// 		}
// 	}

// 	return &r, nil
// }

// // Delete is used to delete one or multiples HostMacro.
// func (u *UserMacroService) Delete(ids []string) (*HostMacroResponse, error) {
// 	req := u.Client.NewRequest("usermacro.delete", ids)

// 	res, err := u.Client.Post(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	r := HostMacroResponse{}
// 	err = u.Client.ConvertResponse(*res, &r)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &r, nil
// }

// // DeleteGlobal is used to delete one or multiples GlobalMacros.
// func (u *UserMacroService) DeleteGlobal(ids []string) (*GlobalMacroResponse, error) {
// 	req := u.Client.NewRequest("usermacro.deleteglobal", ids)

// 	res, err := u.Client.Post(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	r := GlobalMacroResponse{}
// 	err = u.Client.ConvertResponse(*res, &r)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &r, nil
// }

// // UserMacroGetParameters define the properties used to search Macro(s)
// // Properties using the 'omitempty' json parameters are optional
// type UserMacroGetParameters struct {
// 	GlobalMacro            bool        `json:"globalmacro,omitempty"`
// 	GlobalMacroIds         []string    `json:"globalmacroids,omitempty"`
// 	GroupIds               []string    `json:"groupids,omitempty"`
// 	HostIds                []string    `json:"hostids,omitempty"`
// 	HostMacroIds           []string    `json:"hostmacroids,omitempty"`
// 	Inherited              bool        `json:"inherited,omitempty"`
// 	SelectHostGroups       interface{} `json:"selectHostGroups,omitempty"`
// 	SelectHosts            interface{} `json:"selectHosts,omitempty"`
// 	SelectTemplateGroups   interface{} `json:"selectTemplateGroups,omitempty"`
// 	SelectTemplates        interface{} `json:"selectTemplates,omitempty"`
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
// 	SortOrder              []string    `json:"sortorder,omitempty"`
// 	StartSearch            bool        `json:"startSearch,omitempty"`
// }

// // Get is used to retrieve one or multiples HostMacros matching the given criteria(s).
// func (u *UserMacroService) Get(p *UserMacroGetParameters) ([]*HostMacro, error) {
// 	p.GlobalMacro = false

// 	req := u.Client.NewRequest("usermacro.get", p)

// 	res, err := u.Client.Post(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	r := make([]*HostMacro, 0)
// 	err = u.Client.ConvertResponse(*res, &r)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return r, nil
// }

// // GetGlobal is used to retrieve one or multiples GlobalMacros matching the given criteria(s).
// func (u *UserMacroService) GetGlobal(p *UserMacroGetParameters) ([]*GlobalMacro, error) {
// 	p.GlobalMacro = true

// 	req := u.Client.NewRequest("usermacro.get", p)

// 	res, err := u.Client.Post(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	r := make([]*GlobalMacro, 0)
// 	err = u.Client.ConvertResponse(*res, &r)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return r, nil
// }

// type HostMacroUpdateParameters struct {
// 	HostMacroId string        `json:"hostmacroid"`
// 	Macro       string        `json:"macro,omitempty"`
// 	Value       string        `json:"value,omitempty"`
// 	Type        UserMacroType `json:"type,omitempty"`
// 	Description string        `json:"description,omitempty"`
// }

// // Update is used to update or overwrite properties from an existing HostMacro.
// func (u *UserMacroService) Update(p *HostMacroUpdateParameters) (*HostMacroResponse, error) {
// 	if p.HostMacroId == "" {
// 		return nil, fmt.Errorf("missing required field 'HostMacroId' in the given object.\nObject passed : %v", p)
// 	}

// 	req := u.Client.NewRequest("usermacro.update", p)

// 	res, err := u.Client.Post(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	r := HostMacroResponse{}
// 	err = u.Client.ConvertResponse(*res, &r)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &r, nil
// }

// type GlobalMacroUpdateParameters struct {
// 	GlobalMacroId string        `json:"globalmacroid"`
// 	Macro         string        `json:"macro,omitempty"`
// 	Value         string        `json:"value,omitempty"`
// 	Type          UserMacroType `json:"type,omitempty"`
// 	Description   string        `json:"description,omitempty"`
// }

// // UpdateGlobal is used to update or overwrite properties from an existing GlobalMacro.
// func (u *UserMacroService) UpdateGlobal(p *GlobalMacroUpdateParameters) (*GlobalMacroResponse, error) {
// 	if p.GlobalMacroId == "" {
// 		return nil, fmt.Errorf("missing required field 'GlobalMacroId' in the given object.\nObject passed : %v", p)
// 	}

// 	req := u.Client.NewRequest("usermacro.updateglobal", p)

// 	res, err := u.Client.Post(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	r := GlobalMacroResponse{}
// 	err = u.Client.ConvertResponse(*res, &r)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &r, nil
// }
