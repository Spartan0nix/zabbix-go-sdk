package main

import (
	"encoding/json"
	"fmt"
)

type AuthRequest struct {
	Jsonrpc string            `json:"jsonrpc"`
	Method  string            `json:"method"`
	Params  map[string]string `json:"params"`
	Id      int               `json:"id"`
	Auth    *string           `json:"auth"`
}

type AuthReponse struct {
	Jsonrpc string        `json:"jsonrpc"`
	Result  string        `json:"result"`
	Error   ResponseError `json:"error,omitempty"`
	Id      int           `json:"id"`
}

func (a *ApiClient) NewAuthRequest(user string, password string) (AuthRequest, error) {
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

	return body, nil
}

func (a *ApiClient) GetCredentials(user string, password string) (*AuthReponse, error) {
	params, err := a.NewAuthRequest(user, password)
	if err != nil {
		return nil, err
	}

	b, err := a.Post(params)
	if err != nil {
		return nil, err
	}

	res := AuthReponse{}
	err = json.Unmarshal(b, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (a *ApiClient) Authenticate() error {
	resp, err := a.GetCredentials(a.User.User, a.User.Pwd)
	if err != nil {
		return err
	}

	if resp.Result == "" {
		return fmt.Errorf("Code : %d\nMessage : %s\nData : %s", resp.Error.Code, resp.Error.Message, resp.Error.Data)
	}

	a.Token = resp.Result

	return nil
}
