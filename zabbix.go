package zabbixgosdk

import (
	"fmt"
	"net/http"
	"time"
)

// ZabbixService define the base service to interact with API methods.
type ZabbixService struct {
	client *apiClient
	User   *UserService
	// UserGroup *UserGroupService
	// UserMacro *UserMacroService
	HostGroup *HostGroupService
	Template  *TemplateService
	// HostInterface *HostInterfaceService
	// Host          *HostService
	// Service       *ZbxService
	// IconMap       *IconMapService
	// Image         *ImageService
	// Trigger       *TriggerService
	// Map           *MapService
}

// NewZabbixService create a new ZabbixService.
// During the creation process, each sub-service http.Client is initialize.
func NewZabbixService(url string) *ZabbixService {
	c := &apiClient{
		client: &http.Client{
			Timeout: time.Second * 30,
		},
		url: url,
	}

	s := &ZabbixService{
		client: c,
		User: &UserService{
			client: c,
		},
		// UserGroup: &UserGroupService{
		// 	Client: c,
		// },
		// UserMacro: &UserMacroService{
		// 	Client: c,
		// },
		HostGroup: &HostGroupService{
			client: c,
		},
		Template: &TemplateService{
			client: c,
		},
		// HostInterface: &HostInterfaceService{
		// 	Client: c,
		// },
		// Host: &HostService{
		// 	Client: c,
		// },
		// Service: &ZbxService{
		// 	Client: c,
		// },
		// IconMap: &IconMapService{
		// 	Client: c,
		// },
		// Image: &ImageService{
		// 	Client: c,
		// },
		// Trigger: &TriggerService{
		// 	Client: c,
		// },
		// Map: &MapService{
		// 	Client: c,
		// },
	}

	return s
}

// SetUrl is used to set the Url property of each sub-service for the current ZabbiService with the given url.
func (s *ZabbixService) SetUrl(url string) {
	s.client.url = url
	s.User.client.url = url
	// s.UserGroup.Client.Url = url
	// s.UserMacro.Client.Url = url
	s.HostGroup.client.url = url
	s.Template.client.url = url
	// s.HostInterface.Client.Url = url
	// s.Host.Client.Url = url
	// s.Service.Client.Url = url
	// s.IconMap.Client.Url = url
	// s.Image.Client.Url = url
	// s.Trigger.Client.Url = url
	// s.Map.Client.Url = url
}

// SetToken is used to set the Token property of each sub-service for the current ZabbiService with the given token.
func (s *ZabbixService) SetToken(token string) {
	s.client.token = token
	s.User.client.token = token
	// s.UserGroup.Client.Token = token
	// s.UserMacro.Client.Token = token
	s.HostGroup.client.token = token
	s.Template.client.token = token
	// s.HostInterface.Client.Token = token
	// s.Host.Client.Token = token
	// s.Service.Client.Token = token
	// s.IconMap.Client.Token = token
	// s.Image.Client.Token = token
	// s.Trigger.Client.Token = token
	// s.Map.Client.Token = token
}

// Authenticate is used to retrieve a new API token from the Zabbix server.
func (s *ZabbixService) Authenticate(username string, password string) error {
	res, err := s.User.Login(username, password)
	if err != nil {
		return err
	}

	s.SetToken(res.Result)
	return nil
}

// Logout is used to release the current API token.
func (s *ZabbixService) Logout() error {
	return s.User.Logout()
}

// GetApiVersion is used to retrieve the version of the Zabbix API from the server.
func (s *ZabbixService) GetApiVersion() (string, error) {
	// This method is available only to unauthenticated user (without the 'auth' field)
	type request struct {
		JsonRpc string   `json:"jsonrpc"`
		Method  string   `json:"method"`
		Params  []string `json:"params"`
		Id      int      `json:"id"`
	}

	// If the Url property is not set, do not throw an error.
	if s.client.url == "" {
		return "", fmt.Errorf("missing 'url' property. CheckConnectivity function will not be executed")
	}

	// Create the custom body
	body := request{
		JsonRpc: "2.0",
		Method:  "apiinfo.version",
		Params:  []string{},
		Id:      1,
	}

	req, err := s.client.NewCustomRequest(&body)
	if err != nil {
		return "", err
	}

	res, err := s.client.ExecuteRequest(req)
	if err != nil {
		return "", err
	}

	r := Response[string]{}
	err = s.client.ConvertResponse(res, &r)
	if err != nil {
		return "", err
	}

	if r.Result == "" {
		return "", fmt.Errorf("an empty response was returned by the API")
	}

	return r.Result, nil
}

// CheckConnectivity is used to validate the connectivity between the current client and the Zabbix API.
func (s *ZabbixService) CheckConnectivity() error {
	_, err := s.GetApiVersion()
	return err
}
