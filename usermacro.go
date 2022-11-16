package main

import (
	"fmt"
	"regexp"
	"strings"
)

type UserMacroService struct {
	Client *ApiClient
}

type MacroType int

const (
	Text MacroType = iota
	Secret
	Vault
)

type GlobalMacro struct {
	Id          string    `json:"globalmacroid"`
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

type HostMacroCreateParameters struct {
	Hostid      int       `json:"hostid"`
	Macro       string    `json:"macro"`
	Value       string    `json:"value"`
	Type        MacroType `json:"type,omitempty"`
	Description string    `json:"description,omitempty"`
}

func (u *UserMacroService) Create(h HostMacroCreateParameters) (*HostMacroResponse, error) {
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
