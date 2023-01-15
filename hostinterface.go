package zabbixgosdk

import (
	"encoding/json"
	"fmt"
)

// HostInterfaceService create a new service to access hostinterface related methods and functions.
type HostInterfaceService struct {
	Client *ApiClient
}

// HostInterfaceMain define if an interface should be used as default or not.
type HostInterfaceMain string

// HostInterfaceType define the available types of interface.
type HostInterfaceType string

// HostInterfaceVersion define the available SNMP versions.
type HostInterfaceVersion string

// HostInterfaceSecurityLevel define the available SNMPv3 security levels.
type HostInterfaceSecurityLevel string

// HostInterfaceAuthProtocol define the available SNMPv3 authentification protocols.
type HostInterfaceAuthProtocol string

// HostInterfacePrivProtocol define the available SNMPv3 encryption (priv) protocols.
type HostInterfacePrivProtocol string

const (
	NotDefault HostInterfaceMain = "0"
	Default    HostInterfaceMain = "1"

	Agent HostInterfaceType = "1"
	SNMP  HostInterfaceType = "2"
	IPMI  HostInterfaceType = "3"
	JMX   HostInterfaceType = "4"

	SNMPv1  HostInterfaceVersion = "1"
	SNMPv2c HostInterfaceVersion = "2"
	SNMPv3  HostInterfaceVersion = "3"

	NoAuthNoPriv HostInterfaceSecurityLevel = "0"
	AuthNoPriv   HostInterfaceSecurityLevel = "1"
	AuthPriv     HostInterfaceSecurityLevel = "3"

	MD5     HostInterfaceAuthProtocol = "0"
	SHA1    HostInterfaceAuthProtocol = "1"
	SHA224  HostInterfaceAuthProtocol = "2"
	SHA256  HostInterfaceAuthProtocol = "3"
	SHA384  HostInterfaceAuthProtocol = "4"
	SHA512  HostInterfaceAuthProtocol = "5"
	DES     HostInterfacePrivProtocol = "0"
	AES128  HostInterfacePrivProtocol = "1"
	AES192  HostInterfacePrivProtocol = "2"
	AES256  HostInterfacePrivProtocol = "3"
	AES192C HostInterfacePrivProtocol = "4"
	AES256C HostInterfacePrivProtocol = "5"
)

// HostInterface details properties for SNMP interface.
// Properties using the 'omitempty' json parameters are optional
type HostInterfaceDetail struct {
	Version HostInterfaceVersion `json:"version"`
	// Whether to use bulk SNMP requests.
	//
	// 0 = Don't use bulk requests.
	//
	// 1 = Use bulk requests (default).
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

// HostInterface properties.
// Some properties are read-only, which means they are only accessible after creation
// and should not be passed as arguments in other methods.
type HostInterface struct {
	Hostid  string               `json:"hostid,omitempty"`
	Ip      string               `json:"ip"`
	Dns     string               `json:"dns"`
	Main    HostInterfaceMain    `json:"main"`
	Port    string               `json:"port"`
	Type    HostInterfaceType    `json:"type"`
	Useip   string               `json:"useip"`
	Details *HostInterfaceDetail `json:"details,omitempty"`
	// ReadOnly
	Available string `json:"available,omitempty"`
	// ReadOnly
	Disable_until string `json:"disable_until,omitempty"`
	// ReadOnly
	Error string `json:"error,omitempty"`
	// ReadOnly
	Errors_from string `json:"errors_from,omitempty"`
	// ReadOnly
	Interfaceid string `json:"interfaceid,omitempty"`
}

// HostInterfaceResponse define the server response format for HostInterface methods.
type HostInterfaceResponse struct {
	InterfaceIds []json.Number `json:"interfaceids"`
}

// convertToHostInterface is used to convert a hostInterfaceGetResponse in an HostInterface struc
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

// ValidateSNMP help to verify that the 'Details' property is set when configuring SNMP interface.
func (h *HostInterface) ValidateSNMP() error {
	if h.Type == SNMP {
		if h.Details == nil {
			return fmt.Errorf("missing required field 'details' for hostInterface of type SNMP.\nObject passed : %v", h)
		}
	}
	return nil
}

// Create is used to create a new HostInterface.
func (h *HostInterfaceService) Create(p *HostInterface) (*HostInterfaceResponse, error) {
	// Validate configuration when creating SNMP interface.
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

// Delete is used to delete one or multiples HostInterfaces.
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

// HostInterfaceGetParameters define the properties used to search HostInterface(s)
// Properties using the 'omitempty' json parameters are optional
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

// hostInterfaceGetResponse define the server response format for the Get method.
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

// Get is used to retrieve one or multiples HostInterfaces matching the given criteria(s).
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

// HostInterfaceMassProperties define the HostInterface properties used for the Mass method.
// Properties using the 'omitempty' json parameters are optional.
type HostInterfaceMassProperties struct {
	Ip      string               `json:"ip"`
	Dns     string               `json:"dns"`
	Main    HostInterfaceMain    `json:"main"`
	Port    string               `json:"port"`
	Type    HostInterfaceType    `json:"type"`
	Useip   string               `json:"useip"`
	Details *HostInterfaceDetail `json:"details,omitempty"`
}

// HostInterfaceMassAddParameters define the properties used for the MassAdd method.
type HostInterfaceMassAddParameters struct {
	Hosts      []*HostId                      `json:"hosts"`
	Interfaces []*HostInterfaceMassProperties `json:"interfaces"`
}

// HostInterfaceMassAddResponse define the server response format for the MassAdd method.
type HostInterfaceMassAddResponse struct {
	Response *HostInterfaceResponse `json:"interfaceids"`
}

// MassAdd is used to massively add HostInterfaces to existing Hosts.
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

// HostInterfaceMassRemoveParameters define the properties used for the MassRemove method.
type HostInterfaceMassRemoveParameters struct {
	HostIds    []string                       `json:"hostids"`
	Interfaces []*HostInterfaceMassProperties `json:"interfaces"`
}

// MassRemove is used to massively remove HostInterfaces from existing Hosts.
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

// HostInterfaceReplaceParameters define the properties used for the ReplaceHostInterfaces method.
type HostInterfaceReplaceParameters struct {
	Host       string                         `json:"hostid"`
	Interfaces []*HostInterfaceMassProperties `json:"interfaces"`
}

// ReplaceHostInterfaces is used to massively replace HostInterfaces from existing Hosts.
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

// HostInterfaceUpdateParameters define the properties needed for the Update method.
// Properties using the 'omitempty' json parameters are optional.
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

// Update is used to update or overwrite HostInterfaces from an existing Hosts.
func (h *HostInterfaceService) Update(p *HostInterfaceUpdateParameters) (*HostInterfaceResponse, error) {
	if p.Interfaceid == "" {
		return nil, fmt.Errorf("property 'Interfaceid' must be set in order to update an host interface.\nObject passed : %v", p)
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
