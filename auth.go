package main

import "fmt"

// ApiUser defined the informations needed to authenticate against the Zabbix API
// and retrieve a API Token.
type ApiUser struct {
	User string
	Pwd  string
}

// AuthService create a new service to access authentification related methods and functions.
type AuthService struct {
	Client *ApiClient
	// Required only if you intend to use ZabbixService Authenticate method.
	User *ApiUser
}

// Request define the body format to interact with the API.
// The 'auth.login' method accept a different format than the other
type AuthRequest struct {
	Jsonrpc string            `json:"jsonrpc"`
	Method  string            `json:"method"`
	Params  map[string]string `json:"params"`
	Id      int               `json:"id"`
	Auth    *string           `json:"auth"`
}

// NewAuthRequest build a new API Request (for authentification only)
func (s *AuthService) NewAuthRequest(user string, password string) AuthRequest {
	body := AuthRequest{
		Jsonrpc: "2.0",
		Method:  "user.login",
		Params: map[string]string{
			"user":     user,
			"password": password,
		},
		Id:   1,
		Auth: nil,
	}

	return body
}

// GetCredentials execute an AuthRequest using the given user[name] // password and return a Response.
func (s *AuthService) GetCredentials(user string, password string) (*Response, error) {
	params := s.NewAuthRequest(user, password)

	res, err := s.Client.Post(params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Authenticate is a wrapper that while execute the GetCredentials method and update the Token property for the current ZabbixService.
// When updating the Token, all subservices while see their Token property updated as well.
func (s *ZabbixService) Authenticate() error {
	if s.Auth.User == nil {
		return fmt.Errorf("Missing authentification informations :\n- user\n- password")
	}

	if s.Auth.User.User == "" || s.Auth.User.Pwd == "" {
		return fmt.Errorf("Missing authentification informations :\n- user\n- password")
	}

	res, err := s.Auth.GetCredentials(s.Auth.User.User, s.Auth.User.Pwd)
	if err != nil {
		return err
	}

	// When using user.login, the server while return a Response such as :
	// {
	// 	"jsonrpc": string,
	// 	"result": string,
	// 	"id": int
	// }
	var token string
	err = s.Auth.Client.ConvertResponse(*res, &token)
	if err != nil {
		return err
	}

	if token == "" {
		return s.Auth.Client.Error(*res)
	}

	s.SetToken(token)

	return nil
}
