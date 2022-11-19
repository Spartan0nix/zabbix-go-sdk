package main

import (
	"fmt"
	"regexp"
	"strings"
)

type UserMacroService struct {
	Client *ApiClient
}

type MacroType string

const (
	Text   MacroType = "0"
	Secret MacroType = "1"
	Vault  MacroType = "2"
)

type GlobalMacro struct {
	Id          string    `json:"globalmacroid,omitempty"`
	Macro       string    `json:"macro"`
	Value       string    `json:"value"`
	Type        MacroType `json:"type,omitempty"`
	Description string    `json:"description,omitempty"`
}

type HostMacro struct {
	Id          string    `json:"hostmacroid,omitempty"`
	Hostid      string    `json:"hostid"`
	Macro       string    `json:"macro"`
	Value       string    `json:"value"`
	Type        MacroType `json:"type,omitempty"`
	Description string    `json:"description,omitempty"`
}

type HostMacroResponse struct {
	Hostmacroids []string `json:"hostmacroids"`
}

type GlobalMacroResponse struct {
	Globalmacroids []string `json:"globalmacroids"`
}

type GetRequestBaseParameters struct {
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

type UserMacroGetParameters struct {
	Globalmacro          bool        `json:"globalmacro,omitempty"`
	Globalmacroids       []string    `json:"globalmacroids,omitempty"`
	Groupids             []string    `json:"groupids,omitempty"`
	Hostids              []string    `json:"hostids,omitempty"`
	Hostmacroids         []string    `json:"hostmacroids,omitempty"`
	Inherited            bool        `json:"inherited,omitempty"`
	SelectHostGroups     interface{} `json:"selectHostGroups,omitempty"`
	SelectHosts          interface{} `json:"selectHosts,omitempty"`
	SelectTemplateGroups interface{} `json:"selectTemplateGroups,omitempty"`
	SelectTemplates      interface{} `json:"selectTemplates,omitempty"`
	Sortfield            []string    `json:"sortfield,omitempty"`
	BaseParameters       *GetRequestBaseParameters
}

func (u *UserMacroService) GetHostMacro(p *UserMacroGetParameters) ([]HostMacro, error) {
	p.Globalmacro = false

	req := u.Client.NewRequest("usermacro.get", p)

	res, err := u.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := make([]HostMacro, 0)
	err = u.Client.ConvertResponse(*res, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (u *UserMacroService) GetGlobalMacro(p *UserMacroGetParameters) ([]GlobalMacro, error) {
	p.Globalmacro = true

	req := u.Client.NewRequest("usermacro.get", p)

	res, err := u.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := make([]GlobalMacro, 0)
	err = u.Client.ConvertResponse(*res, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (u *UserMacroService) Create(h HostMacro) (*HostMacroResponse, error) {
	re := regexp.MustCompile(`^{\$.*}$`)
	if !re.Match([]byte(h.Macro)) {
		return nil, fmt.Errorf("The following macro '%s' does not complies with the required format.\nFormat : {$...your data..}", h.Macro)
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

func (u *UserMacroService) CreateGlobal(g GlobalMacro) (*GlobalMacroResponse, error) {
	re := regexp.MustCompile(`^{\$.*}$`)
	if !re.Match([]byte(g.Macro)) {
		return nil, fmt.Errorf("The following macro '%s' does not complies with the required format.\nFormat : {$...your data..}", g.Macro)
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
