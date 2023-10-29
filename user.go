package zabbixgosdk

import "fmt"

// UserService create a new service to access user related methods and functions.
type UserService struct {
	client *apiClient
}

type UserAutoLogin uint8
type UserTheme string

const (
	UserAutoLoginEnabled UserAutoLogin = 0
	UserAutoLoginDisable UserAutoLogin = 1
	UserThemeDefault     UserTheme     = "default"
	UserThemeBlue        UserTheme     = "blue-theme"
	UserThemeDark        UserTheme     = "dark-theme"
)

type User struct {
	UserId        string        `json:"userid,omitempty"`
	Username      string        `json:"username"`
	RoleId        string        `json:"roleid"`
	AttemptClock  string        `json:"attempt_clock,omitempty"`
	AttemptFailed uint16        `json:"attempt_failed,omitempty"`
	AttemptIp     string        `json:"attempt_ip,omitempty"`
	AutoLogin     UserAutoLogin `json:"autologin,omitempty"`
	AutoLogout    string        `json:"autologout,omitempty"`
	Lang          string        `json:"lang,omitempty"`
	Name          string        `json:"name,omitempty"`
	Refresh       string        `json:"refresh,omitempty"`
	RowsPerPage   uint16        `json:"rows_per_page,omitempty"`
	Surname       string        `json:"surname,omitempty"`
	Theme         UserTheme     `json:"theme,omitempty"`
	Url           string        `json:"url,omitempty"`
	Timezone      string        `json:"timezone,omitempty"`
}

func (u *UserService) Login(username string, password string) (*Response[string], error) {
	// user.login required the auth field not to be set.
	type loginRequest struct {
		JsonRpc string            `json:"jsonrpc"`
		Method  string            `json:"method"`
		Params  map[string]string `json:"params"`
		Id      int               `json:"id"`
	}

	body := loginRequest{
		JsonRpc: "2.0",
		Method:  "user.login",
		Params: map[string]string{
			"user":     username,
			"password": password,
		},
		Id: 1,
	}

	req, err := u.client.NewCustomRequest(&body)
	if err != nil {
		return nil, err
	}

	res, err := u.client.ExecuteRequest(req)
	if err != nil {
		return nil, err
	}

	r := Response[string]{}
	err = u.client.ConvertResponse(res, &r)
	if err != nil {
		return nil, err
	}

	err = r.Validate()
	return &r, err
}

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	UserData bool   `json:"userData"`
}

func (u *UserService) GetUserInformations(p *UserLoginRequest) (*Response[*User], error) {
	r := Response[*User]{}
	err := u.client.Post("user.login", p, &r)
	if err != nil {
		return nil, err
	}

	err = r.Validate()
	return &r, err
}

func (u *UserService) Logout() error {
	r := Response[bool]{}
	err := u.client.Post("user.logout", map[string]string{}, &r)
	if !r.Result {
		return fmt.Errorf("failed to logout of the Zabbix API")
	}

	return err
}
