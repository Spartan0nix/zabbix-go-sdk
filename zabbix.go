package main

import "net/http"

type ZabbixService struct {
	Auth      *AuthService
	UserGroup *UserGroupService
	UserMacro *UserMacroService
	HostGroup *HostGroupService
	Template  *TemplateService
}

func NewZabbixService() *ZabbixService {
	c := &ApiClient{
		client: &http.Client{},
	}

	return &ZabbixService{
		Auth: &AuthService{
			Client: c,
		},
		UserGroup: &UserGroupService{
			Client: c,
		},
		UserMacro: &UserMacroService{
			Client: c,
		},
		HostGroup: &HostGroupService{
			Client: c,
		},
		Template: &TemplateService{
			Client: c,
		},
	}
}

func (s ZabbixService) SetUrl(url string) {
	s.Auth.Client.Url = url
	s.UserGroup.Client.Url = url
	s.UserMacro.Client.Url = url
	s.HostGroup.Client.Url = url
}

func (s ZabbixService) SetUser(user *ApiUser) {
	s.Auth.User = user
}

func (s ZabbixService) SetToken(token string) {
	s.Auth.Client.Token = token
	s.UserGroup.Client.Token = token
	s.UserMacro.Client.Token = token
	s.HostGroup.Client.Token = token
}
