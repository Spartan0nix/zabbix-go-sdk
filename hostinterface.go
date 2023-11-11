package zabbixgosdk

// import (
// 	"encoding/json"
// 	"fmt"
// )

// // HostInterfaceService create a new service to access hostinterface related methods and functions.
// type HostInterfaceService struct {
// 	Client *ApiClient
// }

// HostInterfaceMain define if an interface should be configured as the default interface.
type HostInterfaceMain string

// HostInterfaceType define the available types of interface.
type HostInterfaceType string

// HostInterfaceConnection define if DNS or IP should be used to connect with an host.
type HostInterfaceConnection string

// HostInterfaceSnmpVersion define the available version of SNMP.
type HostInterfaceSnmpVersion string

// HostInterfaceSnmpBulk define if SNMP bulk requests should be used.
type HostInterfaceSnmpBulk string

// HostInterfaceSecurityLevel define the available types of SnmpV3 security level.
type HostInterfaceSecurityLevel string

// HostInterfaceAuthProtocol define the available types of SnmpV3 authentification protocol.
type HostInterfaceAuthProtocol string

// HostInterfacePrivProtocol define the available types of SnmpV3 encryption protocol.
type HostInterfacePrivProtocol string

// HostInterfaceAvailability define if an interface is available.
type HostInterfaceAvailability string

const (
	HostInterfaceNotDefault HostInterfaceMain = "0"
	HostInterfaceDefault    HostInterfaceMain = "1"

	HostInterfaceAgent HostInterfaceType = "1"
	HostInterfaceSNMP  HostInterfaceType = "2"
	HostInterfaceIPMI  HostInterfaceType = "3"
	HostInterfaceJMX   HostInterfaceType = "4"

	HostInterfaceUseDns HostInterfaceConnection = "0"
	HostInterfaceUseIp  HostInterfaceConnection = "1"

	HostInterfaceSnmpV1  HostInterfaceSnmpVersion = "1"
	HostInterfaceSnmpV2c HostInterfaceSnmpVersion = "2"
	HostInterfaceSnmpV3  HostInterfaceSnmpVersion = "3"

	HostInterfaceNoSnmpBulk  HostInterfaceSnmpBulk = "0"
	HostInterfaceUseSnmpBulk HostInterfaceSnmpBulk = "1"

	HostInterfaceNoAuthNoPriv HostInterfaceSecurityLevel = "0"
	HostInterfaceAuthNoPriv   HostInterfaceSecurityLevel = "1"
	HostInterfaceAuthPriv     HostInterfaceSecurityLevel = "2"

	HostInterfaceMD5    HostInterfaceAuthProtocol = "0"
	HostInterfaceSHA1   HostInterfaceAuthProtocol = "1"
	HostInterfaceSHA224 HostInterfaceAuthProtocol = "2"
	HostInterfaceSHA256 HostInterfaceAuthProtocol = "3"
	HostInterfaceSHA384 HostInterfaceAuthProtocol = "4"
	HostInterfaceSHA512 HostInterfaceAuthProtocol = "5"

	HostInterfaceDES     HostInterfacePrivProtocol = "0"
	HostInterfaceAES128  HostInterfacePrivProtocol = "1"
	HostInterfaceAES192  HostInterfacePrivProtocol = "2"
	HostInterfaceAES256  HostInterfacePrivProtocol = "3"
	HostInterfaceAES192C HostInterfacePrivProtocol = "4"
	HostInterfaceAES256C HostInterfacePrivProtocol = "5"

	HostInterfaceUnknown     HostInterfaceAvailability = "0"
	HostInterfaceAvailable   HostInterfaceAvailability = "1"
	HostInterfaceUnavailable HostInterfaceAvailability = "2"
)

// HostInterfaceDetail properties for SNMP interface.
// Properties using the 'omitempty' json parameters are optional.
type HostInterfaceDetail struct {
	Version        HostInterfaceSnmpVersion   `json:"version"`
	Bulk           HostInterfaceSnmpBulk      `json:"bulk,omitempty"`
	Community      string                     `json:"community,omitempty"`
	SecurityName   string                     `json:"securityname,omitempty"`
	ContextName    string                     `json:"contextname,omitempty"`
	SecurityLevel  HostInterfaceSecurityLevel `json:"securitylevel,omitempty"`
	AuthPassPhrase string                     `json:"authpassphrase,omitempty"`
	PrivPassPhrase string                     `json:"privpassphrase,omitempty"`
	AuthProtocol   HostInterfaceAuthProtocol  `json:"authprotocol,omitempty"`
	PrivProtocol   HostInterfacePrivProtocol  `json:"privprotocol,omitempty"`
}

// HostInterface properties.
type HostInterface struct {
	InterfaceId  string                    `json:"interfaceid,omitempty"`
	HostId       string                    `json:"hostid"`
	Ip           string                    `json:"ip"`
	Dns          string                    `json:"dns"`
	Main         HostInterfaceMain         `json:"main"`
	Port         string                    `json:"port"`
	Type         HostInterfaceType         `json:"type"`
	UseIp        HostInterfaceConnection   `json:"useip"`
	Details      HostInterfaceDetail       `json:"details,omitempty"`
	Available    HostInterfaceAvailability `json:"available,omitempty"`
	DisableUntil string                    `json:"disable_until,omitempty"`
	Error        string                    `json:"error,omitempty"`
	ErrorsFrom   string                    `json:"errors_from,omitempty"`
}

// // HostInterfaceResponse define the server response format for HostInterface methods.
// type HostInterfaceResponse struct {
// 	InterfaceIds []json.Number `json:"interfaceids"`
// }

// // convertToHostInterface is used to convert a hostInterfaceGetResponse in an HostInterface struc
// func (h *hostInterfaceGetResponse) convertToHostInterface() (*HostInterface, error) {
// 	host := &HostInterface{
// 		HostId:       h.HostId,
// 		Ip:           h.Ip,
// 		Dns:          h.Dns,
// 		Main:         h.Main,
// 		Port:         h.Port,
// 		Type:         h.Type,
// 		UseIp:        h.UseIp,
// 		Available:    h.Available,
// 		DisableUntil: h.DisableUntil,
// 		Error:        h.Error,
// 		ErrorsFrom:   h.ErrorsFrom,
// 		InterfaceId:  h.InterfaceId,
// 	}

// 	if string(h.Details) != "[]" {
// 		details := HostInterfaceDetail{}
// 		if err := json.Unmarshal(h.Details, &details); err != nil {
// 			return nil, err
// 		}

// 		host.Details = &details
// 	}

// 	return host, nil
// }

// // ValidateSNMP help to verify that the 'Details' property is set when configuring SNMP interface.
// func (h *HostInterfaceCreateParameters) ValidateSNMP() error {
// 	if h.Type == HostInterfaceSNMP {
// 		if h.Details == nil {
// 			return fmt.Errorf("missing required field 'details' for hostInterface of type SNMP.\nObject passed : %v", h)
// 		}
// 	}
// 	return nil
// }

// HostInterfaceCreateParameters define the parameters use to create a new interface on an Host.
type HostInterfaceCreateParameters struct {
	Dns     string                  `json:"dns"`
	Ip      string                  `json:"ip"`
	Main    HostInterfaceMain       `json:"main"`
	Port    string                  `json:"port"`
	Type    HostInterfaceType       `json:"type"`
	UseIp   HostInterfaceConnection `json:"useip"`
	Details *HostInterfaceDetail    `json:"details,omitempty"`
}

// type HostInterfaceCreateParameters struct {
// 	HostId  string               `json:"hostid"`
// 	Ip      string               `json:"ip"`
// 	Dns     string               `json:"dns"`
// 	Main    HostInterfaceMain    `json:"main"`
// 	Port    string               `json:"port"`
// 	Type    HostInterfaceType    `json:"type"`
// 	UseIp   HostInterfaceUseIp   `json:"useip"`
// 	Details *HostInterfaceDetail `json:"details,omitempty"`
// }

// // Create is used to create a new HostInterface.
// func (h *HostInterfaceService) Create(p *HostInterfaceCreateParameters) (*HostInterfaceResponse, error) {
// 	// Validate configuration when creating SNMP interface.
// 	if err := p.ValidateSNMP(); err != nil {
// 		return nil, err
// 	}

// 	req := h.Client.NewRequest("hostinterface.create", p)

// 	res, err := h.Client.Post(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	r := HostInterfaceResponse{}
// 	err = h.Client.ConvertResponse(*res, &r)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &r, nil
// }

// // Delete is used to delete one or multiples HostInterfaces.
// func (h *HostInterfaceService) Delete(ids []string) (*HostInterfaceResponse, error) {
// 	req := h.Client.NewRequest("hostinterface.delete", ids)

// 	res, err := h.Client.Post(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	r := HostInterfaceResponse{}
// 	err = h.Client.ConvertResponse(*res, &r)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &r, nil
// }

// // HostInterfaceGetParameters define the properties used to search HostInterface(s)
// // Properties using the 'omitempty' json parameters are optional
// type HostInterfaceGetParameters struct {
// 	HostIds                []string    `json:"hostids,omitempty"`
// 	InterfaceIds           []string    `json:"interfaceids,omitempty"`
// 	ItemIds                []string    `json:"itemids,omitempty"`
// 	TriggerIds             []string    `json:"triggerids,omitempty"`
// 	SelectItems            interface{} `json:"selectItems,omitempty"`
// 	SelectHosts            interface{} `json:"selectHosts,omitempty"`
// 	LimitSelects           string      `json:"limitSelects,omitempty"`
// 	SortField              []string    `json:"sortfield,omitempty"`
// 	CountOutput            bool        `json:"countOutput,omitempty"`
// 	Editable               bool        `json:"editable,omitempty"`
// 	ExcludeSearch          bool        `json:"excludeSearch,omitempty"`
// 	Filter                 interface{} `json:"filter,omitempty"`
// 	Limit                  string      `json:"limit,omitempty"`
// 	Output                 interface{} `json:"output,omitempty"`
// 	PreserveKeys           bool        `json:"preservekeys,omitempty"`
// 	Search                 interface{} `json:"search,omitempty"`
// 	SearchByAny            bool        `json:"searchByAny,omitempty"`
// 	SearchWildcardsEnabled bool        `json:"searchWildcardsEnabled,omitempty"`
// 	SortOrder              []string    `json:"sortorder,omitempty"`
// 	StartSearch            bool        `json:"startSearch,omitempty"`
// }

// // hostInterfaceGetResponse define the server response format for the Get method.
// type hostInterfaceGetResponse struct {
// 	HostId       string                    `json:"hostid"`
// 	Ip           string                    `json:"ip"`
// 	Dns          string                    `json:"dns"`
// 	Main         HostInterfaceMain         `json:"main"`
// 	Port         string                    `json:"port"`
// 	Type         HostInterfaceType         `json:"type"`
// 	UseIp        HostInterfaceUseIp        `json:"useip"`
// 	Details      json.RawMessage           `json:"details,omitempty"`
// 	Available    HostInterfaceAvailability `json:"available,omitempty"`
// 	DisableUntil string                    `json:"disable_until,omitempty"`
// 	Error        string                    `json:"error,omitempty"`
// 	ErrorsFrom   string                    `json:"errors_from,omitempty"`
// 	InterfaceId  string                    `json:"interfaceid,omitempty"`
// }

// // Get is used to retrieve one or multiples HostInterfaces matching the given criteria(s).
// func (h *HostInterfaceService) Get(p *HostInterfaceGetParameters) ([]*HostInterface, error) {
// 	req := h.Client.NewRequest("hostinterface.get", p)

// 	res, err := h.Client.Post(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	tmp_response := make([]*hostInterfaceGetResponse, 0)
// 	err = h.Client.ConvertResponse(*res, &tmp_response)
// 	if err != nil {
// 		return nil, err
// 	}

// 	r := make([]*HostInterface, 0)
// 	for _, el := range tmp_response {
// 		h, err := el.convertToHostInterface()
// 		if err != nil {
// 			return nil, err
// 		}

// 		r = append(r, h)
// 	}

// 	return r, nil
// }

// // HostInterfaceMassProperties define the HostInterface properties used for the Mass method.
// // Properties using the 'omitempty' json parameters are optional.
// type HostInterfaceMassProperties struct {
// 	Ip      string               `json:"ip"`
// 	Dns     string               `json:"dns"`
// 	Main    HostInterfaceMain    `json:"main"`
// 	Port    string               `json:"port"`
// 	Type    HostInterfaceType    `json:"type"`
// 	UseIp   HostInterfaceUseIp   `json:"useip"`
// 	Details *HostInterfaceDetail `json:"details,omitempty"`
// }

// // HostInterfaceMassAddParameters define the properties used for the MassAdd method.
// type HostInterfaceMassAddParameters struct {
// 	Hosts      []*HostId                      `json:"hosts"`
// 	Interfaces []*HostInterfaceMassProperties `json:"interfaces"`
// }

// // HostInterfaceMassAddResponse define the server response format for the MassAdd method.
// type HostInterfaceMassAddResponse struct {
// 	Response *HostInterfaceResponse `json:"interfaceids"`
// }

// // MassAdd is used to massively add HostInterfaces to existing Hosts.
// func (h *HostInterfaceService) MassAdd(p *HostInterfaceMassAddParameters) (*HostInterfaceResponse, error) {
// 	req := h.Client.NewRequest("hostinterface.massadd", p)

// 	res, err := h.Client.Post(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	r := HostInterfaceMassAddResponse{}
// 	err = h.Client.ConvertResponse(*res, &r)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return r.Response, nil
// }

// // HostInterfaceMassRemoveParameters define the properties used for the MassRemove method.
// type HostInterfaceMassRemoveParameters struct {
// 	HostIds    []string                       `json:"hostids"`
// 	Interfaces []*HostInterfaceMassProperties `json:"interfaces"`
// }

// // MassRemove is used to massively remove HostInterfaces from existing Hosts.
// func (h *HostInterfaceService) MassRemove(p *HostInterfaceMassRemoveParameters) (*HostInterfaceResponse, error) {
// 	req := h.Client.NewRequest("hostinterface.massremove", p)

// 	res, err := h.Client.Post(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	r := HostInterfaceResponse{}
// 	err = h.Client.ConvertResponse(*res, &r)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &r, nil
// }

// // HostInterfaceReplaceParameters define the properties used for the ReplaceHostInterfaces method.
// type HostInterfaceReplaceParameters struct {
// 	Host       string                         `json:"hostid"`
// 	Interfaces []*HostInterfaceMassProperties `json:"interfaces"`
// }

// // ReplaceHostInterfaces is used to massively replace HostInterfaces from existing Hosts.
// func (h *HostInterfaceService) ReplaceHostInterfaces(p *HostInterfaceReplaceParameters) (*HostInterfaceResponse, error) {
// 	req := h.Client.NewRequest("hostinterface.replacehostinterfaces", p)

// 	res, err := h.Client.Post(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	r := HostInterfaceResponse{}
// 	err = h.Client.ConvertResponse(*res, &r)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &r, nil
// }

// // HostInterfaceUpdateParameters define the properties needed for the Update method.
// // Properties using the 'omitempty' json parameters are optional.
// type HostInterfaceUpdateParameters struct {
// 	InterfaceId string               `json:"interfaceid"`
// 	Ip          string               `json:"ip,omitempty"`
// 	Dns         string               `json:"dns,omitempty"`
// 	Main        HostInterfaceMain    `json:"main,omitempty"`
// 	Port        string               `json:"port,omitempty"`
// 	Type        HostInterfaceType    `json:"type,omitempty"`
// 	UseIp       HostInterfaceUseIp   `json:"useip,omitempty"`
// 	Details     *HostInterfaceDetail `json:"details,omitempty"`
// }

// // Update is used to update or overwrite HostInterfaces from an existing Hosts.
// func (h *HostInterfaceService) Update(p *HostInterfaceUpdateParameters) (*HostInterfaceResponse, error) {
// 	if p.InterfaceId == "" {
// 		return nil, fmt.Errorf("property 'InterfaceId' must be set in order to update an host interface.\nObject passed : %v", p)
// 	}

// 	req := h.Client.NewRequest("hostinterface.update", p)

// 	res, err := h.Client.Post(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	r := HostInterfaceResponse{}
// 	err = h.Client.ConvertResponse(*res, &r)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &r, nil
// }
