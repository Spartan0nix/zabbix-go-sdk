package main

import (
	"encoding/json"
	"fmt"
)

type HostInterfaceService struct {
	Client *ApiClient
}

type HostInterfaceMain string
type HostInterfaceType string
type HostInterfaceVersion string
type HostInterfaceSecurityLevel string
type HostInterfaceAuthProtocol string
type HostInterfacePrivProtocol string

const (
	NotDefault   HostInterfaceMain          = "0"
	Default      HostInterfaceMain          = "1"
	Agent        HostInterfaceType          = "1"
	SNMP         HostInterfaceType          = "2"
	IPMI         HostInterfaceType          = "3"
	JMX          HostInterfaceType          = "4"
	SNMPv1       HostInterfaceVersion       = "1"
	SNMPv2c      HostInterfaceVersion       = "2"
	SNMPv3       HostInterfaceVersion       = "3"
	noAuthNoPriv HostInterfaceSecurityLevel = "0"
	authNoPriv   HostInterfaceSecurityLevel = "1"
	authPriv     HostInterfaceSecurityLevel = "3"
	MD5          HostInterfaceAuthProtocol  = "0"
	SHA1         HostInterfaceAuthProtocol  = "1"
	SHA224       HostInterfaceAuthProtocol  = "2"
	SHA256       HostInterfaceAuthProtocol  = "3"
	SHA384       HostInterfaceAuthProtocol  = "4"
	SHA512       HostInterfaceAuthProtocol  = "5"
	DES          HostInterfacePrivProtocol  = "0"
	AES128       HostInterfacePrivProtocol  = "1"
	AES192       HostInterfacePrivProtocol  = "2"
	AES256       HostInterfacePrivProtocol  = "3"
	AES192C      HostInterfacePrivProtocol  = "4"
	AES256C      HostInterfacePrivProtocol  = "5"
)

type HostInterfaceDetail struct {
	Version        HostInterfaceVersion       `json:"version"`
	Bulk           string                     `json:"bulk,omitempty"`
	Community      string                     `json:"community,omitempty"`
	Securityname   string                     `json:"securityname,omitempty"`
	Securitylevel  HostInterfaceSecurityLevel `json:"securitylevel,omitempty"`
	Authpassphrase string                     `json:"authpassphrase,omitempty"`
	Privpassphrase string                     `json:"privpassphrase,omitempty"`
	Authprotocol   HostInterfaceAuthProtocol  `json:"authprotocol,omitempty"`
	Privprotocol   HostInterfacePrivProtocol  `json:"privprotocol,omitempty"`
	Contextname    string                     `json:"contextname,omitempty"`
}

type HostInterface struct {
	Hostid        string               `json:"hostid"`
	Ip            string               `json:"ip"`
	Dns           string               `json:"dns"`
	Main          HostInterfaceMain    `json:"main"`
	Port          string               `json:"port"`
	Type          HostInterfaceType    `json:"type"`
	Useip         string               `json:"useip"`
	Details       *HostInterfaceDetail `json:"details,omitempty"`
	Available     string               `json:"available,omitempty"`
	Disable_until string               `json:"disable_until,omitempty"`
	Error         string               `json:"error,omitempty"`
	Errors_from   string               `json:"errors_from,omitempty"`
	Interfaceid   string               `json:"interfaceid,omitempty"`
}

type HostInterfaceResponse struct {
	InterfaceIds []json.Number `json:"interfaceids"`
}

func (h *HostInterface) ValidateSNMP() error {
	if h.Type == SNMP {
		if h.Details == nil {
			return fmt.Errorf("Missing required field 'details' for hostInterface of type SNMP.\nObject passed : %v", h)
		}
	}
	return nil
}

func (h *HostInterfaceService) Create(p *HostInterface) (*HostInterfaceResponse, error) {
	if err := p.ValidateSNMP(); err != nil {
		return nil, err
	}

	req := h.Client.NewRequest("hostinterface.create", p)

	res, err := h.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := HostInterfaceResponse{}
	err = h.Client.ConvertResponse(*res, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (h *HostInterfaceService) Delete(ids []string) (*HostInterfaceResponse, error) {
	req := h.Client.NewRequest("hostinterface.delete", ids)

	res, err := h.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := HostInterfaceResponse{}
	err = h.Client.ConvertResponse(*res, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

type HostInterfaceGetParameters struct {
	Hostids                []string    `json:"hostids,omitempty"`
	Interfaceids           []string    `json:"interfaceids,omitempty"`
	Itemids                []string    `json:"itemids,omitempty"`
	Triggerids             []string    `json:"triggerids,omitempty"`
	SelectItems            interface{} `json:"selectItems,omitempty"`
	SelectHosts            interface{} `json:"selectHosts,omitempty"`
	LimitSelects           int         `json:"limitSelects,omitempty"`
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

type hostInterfaceGetResponse struct {
	Hostid        string            `json:"hostid"`
	Ip            string            `json:"ip"`
	Dns           string            `json:"dns"`
	Main          HostInterfaceMain `json:"main"`
	Port          string            `json:"port"`
	Type          HostInterfaceType `json:"type"`
	Useip         string            `json:"useip"`
	Details       json.RawMessage   `json:"details,omitempty"`
	Available     string            `json:"available,omitempty"`
	Disable_until string            `json:"disable_until,omitempty"`
	Error         string            `json:"error,omitempty"`
	Errors_from   string            `json:"errors_from,omitempty"`
	Interfaceid   string            `json:"interfaceid,omitempty"`
}

func (h *hostInterfaceGetResponse) convertToHostInterface() (*HostInterface, error) {
	host := &HostInterface{
		Hostid:        h.Hostid,
		Ip:            h.Ip,
		Dns:           h.Dns,
		Main:          h.Main,
		Port:          h.Port,
		Type:          h.Type,
		Useip:         h.Useip,
		Available:     h.Available,
		Disable_until: h.Disable_until,
		Error:         h.Error,
		Errors_from:   h.Errors_from,
		Interfaceid:   h.Interfaceid,
	}

	if string(h.Details) != "[]" {
		detail := HostInterfaceDetail{}
		if err := json.Unmarshal(h.Details, &detail); err != nil {
			return nil, err
		}

		host.Details = &detail
	}

	return host, nil
}

func (h *HostInterfaceService) Get(p *HostInterfaceGetParameters) ([]*HostInterface, error) {
	req := h.Client.NewRequest("hostinterface.get", p)

	res, err := h.Client.Post(req)
	if err != nil {
		return nil, err
	}

	tmp_response := make([]*hostInterfaceGetResponse, 0)
	err = h.Client.ConvertResponse(*res, &tmp_response)
	if err != nil {
		return nil, err
	}

	r := make([]*HostInterface, 0)
	for _, el := range tmp_response {
		h, err := el.convertToHostInterface()
		if err != nil {
			return nil, err
		}

		r = append(r, h)
	}

	return r, nil
}

type HostInterfaceMassProperties struct {
	Ip      string               `json:"ip"`
	Dns     string               `json:"dns"`
	Main    HostInterfaceMain    `json:"main"`
	Port    string               `json:"port"`
	Type    HostInterfaceType    `json:"type"`
	Useip   string               `json:"useip"`
	Details *HostInterfaceDetail `json:"details,omitempty"`
}

type HostInterfaceHostId struct {
	Id string `json:"hostid"`
}

type HostInterfaceMassAddParameters struct {
	Hosts      []*HostInterfaceHostId         `json:"hosts"`
	Interfaces []*HostInterfaceMassProperties `json:"interfaces"`
}

type HostInterfaceMassAddResponse struct {
	Response *HostInterfaceResponse `json:"interfaceids"`
}

func (h *HostInterfaceService) MassAdd(p *HostInterfaceMassAddParameters) (*HostInterfaceResponse, error) {
	req := h.Client.NewRequest("hostinterface.massadd", p)

	res, err := h.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := HostInterfaceMassAddResponse{}
	err = h.Client.ConvertResponse(*res, &r)
	if err != nil {
		return nil, err
	}

	return r.Response, nil
}

type HostInterfaceMassRemoveParameters struct {
	HostIds    []string                       `json:"hostids"`
	Interfaces []*HostInterfaceMassProperties `json:"interfaces"`
}

func (h *HostInterfaceService) MassRemove(p *HostInterfaceMassRemoveParameters) (*HostInterfaceResponse, error) {
	req := h.Client.NewRequest("hostinterface.massremove", p)

	res, err := h.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := HostInterfaceResponse{}
	err = h.Client.ConvertResponse(*res, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

type HostInterfaceReplaceParameters struct {
	Host       string                         `json:"hostid"`
	Interfaces []*HostInterfaceMassProperties `json:"interfaces"`
}

func (h *HostInterfaceService) ReplaceHostInterfaces(p *HostInterfaceReplaceParameters) (*HostInterfaceResponse, error) {
	req := h.Client.NewRequest("hostinterface.replacehostinterfaces", p)

	res, err := h.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := HostInterfaceResponse{}
	err = h.Client.ConvertResponse(*res, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

type HostInterfaceUpdateParameters struct {
	Interfaceid string               `json:"interfaceid"`
	Ip          string               `json:"ip,omitempty"`
	Dns         string               `json:"dns,omitempty"`
	Main        HostInterfaceMain    `json:"main,omitempty"`
	Port        string               `json:"port,omitempty"`
	Type        HostInterfaceType    `json:"type,omitempty"`
	Useip       string               `json:"useip,omitempty"`
	Details     *HostInterfaceDetail `json:"details,omitempty"`
}

func (h *HostInterfaceService) Update(p *HostInterfaceUpdateParameters) (*HostInterfaceResponse, error) {
	if p.Interfaceid == "" {
		return nil, fmt.Errorf("Property 'Interfaceid' must be set in order to update an host interface.\nObject passed : %v", p)
	}

	req := h.Client.NewRequest("hostinterface.update", p)

	res, err := h.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := HostInterfaceResponse{}
	err = h.Client.ConvertResponse(*res, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
