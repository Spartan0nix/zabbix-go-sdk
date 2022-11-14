package main

import "fmt"

type ApiUser struct {
	User string
	Pwd  string
}

type AuthService struct {
	Client *ApiClient
	User   *ApiUser
}

type AuthRequest struct {
	Jsonrpc string            `json:"jsonrpc"`
	Method  string            `json:"method"`
	Params  map[string]string `json:"params"`
	Id      int               `json:"id"`
	Auth    *string           `json:"auth"`
}

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

func (s *AuthService) GetCredentials(user string, password string) (*Response, error) {
	params := s.NewAuthRequest(user, password)

	res, err := s.Client.Post(params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

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
