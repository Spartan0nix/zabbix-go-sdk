package zabbixgosdk

import (
	"fmt"
	"regexp"
	"strings"
)

// UserMacroService create a new service to access macro related methods and functions.
type UserMacroService struct {
	Client *ApiClient
}

// MacroType define the available macro types.
type MacroType string

const (
	Text   MacroType = "0"
	Secret MacroType = "1"
	Vault  MacroType = "2"
)

// HostMacro properties.
// Some properties are read-only, which means they are only accessible after creation
// and should not be passed as arguments in other methods.
type HostMacro struct {
	// ReadOnly
	Id          string    `json:"hostmacroid,omitempty"`
	Hostid      string    `json:"hostid,omitempty"`
	Macro       string    `json:"macro"`
	Value       string    `json:"value"`
	Type        MacroType `json:"type,omitempty"`
	Description string    `json:"description,omitempty"`
}

// GlobalMacro properties.
// Some properties are read-only, which means they are only accessible after creation
// and should not be passed as arguments in other methods.
type GlobalMacro struct {
	// ReadOnly
	Id          string    `json:"globalmacroid,omitempty"`
	Macro       string    `json:"macro"`
	Value       string    `json:"value"`
	Type        MacroType `json:"type,omitempty"`
	Description string    `json:"description,omitempty"`
}

// HostMacroResponse define the server response format for HostMacro methods.
type HostMacroResponse struct {
	Hostmacroids []string `json:"hostmacroids"`
}

// GlobalMacroResponse define the server response format for GlobalMacro methods.
type GlobalMacroResponse struct {
	Globalmacroids []string `json:"globalmacroids"`
}

// ValidateMacro is used to validate HostMacro format compliance.
func (h *HostMacro) ValidateMacro() error {
	re := regexp.MustCompile(`^{\$[A-Z0-9_.]*}$`)
	if !re.Match([]byte(h.Macro)) {
		return fmt.Errorf("the following macro '%s' does not complies with the required format.\n- Start with : {$\n- End with : }\n- Contain only : A-Z // 0-9 // _ //", h.Macro)
	}

	return nil
}

// ValidateMacro is used to validate GlobalMacro format compliance.
func (g *GlobalMacro) ValidateMacro() error {
	re := regexp.MustCompile(`^{\$[A-Z0-9_.]*}$`)
	if !re.Match([]byte(g.Macro)) {
		return fmt.Errorf("the following macro '%s' does not complies with the required format.\n- Start with : {$\n- End with : }\n- Contain only : A-Z // 0-9 // _ //", g.Macro)
	}

	return nil
}

// Create is used to create a new HostMacro.
func (u *UserMacroService) Create(h *HostMacro) (*HostMacroResponse, error) {
	if h.Hostid == "" {
		return nil, fmt.Errorf("missing required field 'HostId' in the given object.\nObject passed : %v", h)
	}

	if err := h.ValidateMacro(); err != nil {
		return nil, err
	}

	req := u.Client.NewRequest("usermacro.create", h)

	res, err := u.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := HostMacroResponse{}
	err = u.Client.ConvertResponse(*res, &r)
	if err != nil {
		escaped_macro := strings.Replace(h.Macro, "$", `\$`, -1)
		if u.Client.ResourceAlreadyExist("Macro", escaped_macro, res.Error) {
			fmt.Println(res.Error.Data)
		} else {
			return nil, err
		}
	}

	return &r, nil
}

// CreateGlobal is used to create a new GlobalMacro.
func (u *UserMacroService) CreateGlobal(g *GlobalMacro) (*GlobalMacroResponse, error) {
	if err := g.ValidateMacro(); err != nil {
		return nil, err
	}

	req := u.Client.NewRequest("usermacro.createglobal", g)

	res, err := u.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := GlobalMacroResponse{}
	err = u.Client.ConvertResponse(*res, &r)
	if err != nil {
		escaped_macro := strings.Replace(g.Macro, "$", `\$`, -1)
		if u.Client.ResourceAlreadyExist("Macro", escaped_macro, res.Error) {
			fmt.Println(res.Error.Data)
		} else {
			return nil, err
		}
	}

	return &r, nil
}

// Delete is used to delete one or multiples HostMacro.
func (u *UserMacroService) Delete(ids []string) (*HostMacroResponse, error) {
	req := u.Client.NewRequest("usermacro.delete", ids)

	res, err := u.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := HostMacroResponse{}
	err = u.Client.ConvertResponse(*res, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// DeleteGlobal is used to delete one or multiples GlobalMacros.
func (u *UserMacroService) DeleteGlobal(ids []string) (*GlobalMacroResponse, error) {
	req := u.Client.NewRequest("usermacro.deleteglobal", ids)

	res, err := u.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := GlobalMacroResponse{}
	err = u.Client.ConvertResponse(*res, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// UserMacroGetParameters define the properties used to search Macro(s)
// Properties using the 'omitempty' json parameters are optional
type UserMacroGetParameters struct {
	Globalmacro            bool        `json:"globalmacro,omitempty"`
	Globalmacroids         []string    `json:"globalmacroids,omitempty"`
	Groupids               []string    `json:"groupids,omitempty"`
	Hostids                []string    `json:"hostids,omitempty"`
	Hostmacroids           []string    `json:"hostmacroids,omitempty"`
	Inherited              bool        `json:"inherited,omitempty"`
	SelectHostGroups       interface{} `json:"selectHostGroups,omitempty"`
	SelectHosts            interface{} `json:"selectHosts,omitempty"`
	SelectTemplateGroups   interface{} `json:"selectTemplateGroups,omitempty"`
	SelectTemplates        interface{} `json:"selectTemplates,omitempty"`
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

// Get is used to retrieve one or multiples HostMacros matching the given criteria(s).
func (u *UserMacroService) Get(p *UserMacroGetParameters) ([]*HostMacro, error) {
	p.Globalmacro = false

	req := u.Client.NewRequest("usermacro.get", p)

	res, err := u.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := make([]*HostMacro, 0)
	err = u.Client.ConvertResponse(*res, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// GetGlobal is used to retrieve one or multiples GlobalMacros matching the given criteria(s).
func (u *UserMacroService) GetGlobal(p *UserMacroGetParameters) ([]*GlobalMacro, error) {
	p.Globalmacro = true

	req := u.Client.NewRequest("usermacro.get", p)

	res, err := u.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := make([]*GlobalMacro, 0)
	err = u.Client.ConvertResponse(*res, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Update is used to update or overwrite properties from an existing HostMacro.
func (u *UserMacroService) Update(h *HostMacro) (*HostMacroResponse, error) {
	if h.Id == "" {
		return nil, fmt.Errorf("missing required field 'Id' in the given object.\nObject passed : %v", h)
	}

	h.Hostid = ""

	req := u.Client.NewRequest("usermacro.update", h)

	res, err := u.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := HostMacroResponse{}
	err = u.Client.ConvertResponse(*res, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// UpdateGlobal is used to update or overwrite properties from an existing GlobalMacro.
func (u *UserMacroService) UpdateGlobal(m *GlobalMacro) (*GlobalMacroResponse, error) {
	if m.Id == "" {
		return nil, fmt.Errorf("missing required field 'Id' in the given object.\nObject passed : %v", m)
	}

	req := u.Client.NewRequest("usermacro.updateglobal", m)

	res, err := u.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := GlobalMacroResponse{}
	err = u.Client.ConvertResponse(*res, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
