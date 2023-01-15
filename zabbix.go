package zabbixgosdk

import "net/http"

// ZabbixService define the base service to interact with API methods.
type ZabbixService struct {
	Auth          *AuthService
	UserGroup     *UserGroupService
	UserMacro     *UserMacroService
	HostGroup     *HostGroupService
	Template      *TemplateService
	HostInterface *HostInterfaceService
	Host          *HostService
}

// NewZabbixService create a new ZabbixService.
// During the creation process, each sub-service http.Client is initialize.
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
		HostInterface: &HostInterfaceService{
			Client: c,
		},
		Host: &HostService{
			Client: c,
		},
	}
}

// SetUrl is used to set the Url property of each sub-service for the current ZabbiService with the given url.
func (s ZabbixService) SetUrl(url string) {
	s.Auth.Client.Url = url
	s.UserGroup.Client.Url = url
	s.UserMacro.Client.Url = url
	s.HostGroup.Client.Url = url
	s.HostInterface.Client.Url = url
	s.Host.Client.Url = url
}

// SetUser is used to set the User property of the AuthService for the current ZabbixService.
func (s ZabbixService) SetUser(user *ApiUser) {
	s.Auth.User = user
}

// SetToken is used to set the Token property of each sub-service for the current ZabbiService with the given token.
func (s ZabbixService) SetToken(token string) {
	s.Auth.Client.Token = token
	s.UserGroup.Client.Token = token
	s.UserMacro.Client.Token = token
	s.HostGroup.Client.Token = token
	s.HostInterface.Client.Token = token
	s.Host.Client.Token = token
}

func (s ZabbixService) TestConnectivity() {

}
