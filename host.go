package main

type HostService struct {
	Client *ApiClient
}

type HostInventoryMode string
type HostIpmiAuthType string
type HostIpmiPrivilege string
type HostStatus string
type HostTlsMode string
type HostInventory string
type HostEvalType string
type HostTags string

const (
	Host_Disabled  HostInventoryMode = "-1"
	Host_Manual    HostInventoryMode = "0"
	Host_Automatic HostInventoryMode = "1"

	Host_Default  HostIpmiAuthType = "-1"
	Host_None     HostIpmiAuthType = "0"
	Host_MD2      HostIpmiAuthType = "1"
	Host_MD5      HostIpmiAuthType = "2"
	Host_Straight HostIpmiAuthType = "4"
	Host_OEM      HostIpmiAuthType = "5"
	Host_RMCP     HostIpmiAuthType = "6"

	Callback HostIpmiPrivilege = "1"
	User     HostIpmiPrivilege = "2"
	Operator HostIpmiPrivilege = "3"
	Admin    HostIpmiPrivilege = "4"
	OEM      HostIpmiPrivilege = "5"

	Monitored   HostStatus = "0"
	Unmonitored HostStatus = "1"

	NoEncryption HostTlsMode = "1"
	PSK          HostTlsMode = "2"
	Certificate  HostTlsMode = "4"

	alias             HostInventory = "alias"
	asset_tag         HostInventory = "asset_tag"
	chassis           HostInventory = "chassis"
	contact           HostInventory = "contact"
	contract_number   HostInventory = "contract_number"
	date_hw_decomm    HostInventory = "date_hw_decomm"
	date_hw_expiry    HostInventory = "date_hw_expiry"
	date_hw_install   HostInventory = "date_hw_install"
	date_hw_purchase  HostInventory = "date_hw_purchase"
	deployment_status HostInventory = "deployment_status"
	hardware          HostInventory = "hardware"
	hardware_full     HostInventory = "hardware_full"
	host_netmask      HostInventory = "host_netmask"
	host_networks     HostInventory = "host_networks"
	host_router       HostInventory = "host_router"
	hw_arch           HostInventory = "hw_arch"
	installer_name    HostInventory = "installer_name"
	location          HostInventory = "location"
	location_lat      HostInventory = "location_lat"
	location_lon      HostInventory = "location_lon"
	macaddress_a      HostInventory = "macaddress_a"
	macaddress_b      HostInventory = "macaddress_b"
	model             HostInventory = "model"
	name              HostInventory = "name"
	notes             HostInventory = "notes"
	oob_ip            HostInventory = "oob_ip"
	oob_netmask       HostInventory = "oob_netmask"
	oob_router        HostInventory = "oob_router"
	os                HostInventory = "os"
	os_full           HostInventory = "os_full"
	os_short          HostInventory = "os_short"
	poc_1_cell        HostInventory = "poc_1_cell"
	poc_1_email       HostInventory = "poc_1_email"
	poc_1_name        HostInventory = "poc_1_name"
	poc_1_notes       HostInventory = "poc_1_notes"
	poc_1_phone_a     HostInventory = "poc_1_phone_a"
	poc_1_phone_b     HostInventory = "poc_1_phone_b"
	poc_1_screen      HostInventory = "poc_1_screen"
	poc_2_cell        HostInventory = "poc_2_cell"
	poc_2_email       HostInventory = "poc_2_email"
	poc_2_name        HostInventory = "poc_2_name"
	poc_2_notes       HostInventory = "poc_2_notes"
	poc_2_phone_a     HostInventory = "poc_2_phone_a"
	poc_2_phone_b     HostInventory = "poc_2_phone_b"
	poc_2_screen      HostInventory = "poc_2_screen"
	serialno_a        HostInventory = "serialno_a"
	serialno_b        HostInventory = "serialno_b"
	site_address_a    HostInventory = "site_address_a"
	site_address_b    HostInventory = "site_address_b"
	site_address_c    HostInventory = "site_address_c"
	site_city         HostInventory = "site_city"
	site_country      HostInventory = "site_country"
	site_notes        HostInventory = "site_notes"
	site_rack         HostInventory = "site_rack"
	site_state        HostInventory = "site_state"
	site_zip          HostInventory = "site_zip"
	software          HostInventory = "software"
	software_app_a    HostInventory = "software_app_a"
	software_app_b    HostInventory = "software_app_b"
	software_app_c    HostInventory = "software_app_c"
	software_app_d    HostInventory = "software_app_d"
	software_app_e    HostInventory = "software_app_e"
	software_full     HostInventory = "software_full"
	tag               HostInventory = "tag"
	Type              HostInventory = "type"
	type_full         HostInventory = "type_full"
	url_a             HostInventory = "url_a"
	url_b             HostInventory = "url_b"
	url_c             HostInventory = "url_c"
	vendor            HostInventory = "vendor"

	And_Or HostEvalType = "0"
	Or     HostEvalType = "2"

	Contains   HostTags = "0"
	Equals     HostTags = "1"
	Not_like   HostTags = "2"
	Not_equal  HostTags = "3"
	Exists     HostTags = "4"
	Not_exists HostTags = "5"
)

type Host struct {
	// ReadOnly
	Hostid      string `json:"hostid"`
	Host        string `json:"host"`
	Name        string `json:"name"`
	Description string `json:"description"`
	// ReadOnly
	Flags          string            `json:"flags"`
	Inventory_mode HostInventoryMode `json:"inventory_mode"`
	Ipmi_authtype  HostIpmiAuthType  `json:"ipmi_authtype"`
	Ipmi_password  string            `json:"ipmi_password"`
	Ipmi_privilege HostIpmiPrivilege `json:"ipmi_privilege"`
	Ipmi_username  string            `json:"ipmi_username"`
	// ReadOnly
	Maintenance_from string `json:"maintenance_from"`
	// ReadOnly
	Maintenance_status string `json:"maintenance_status"`
	// ReadOnly
	Maintenance_type string `json:"maintenance_type"`
	// ReadOnly
	Maintenanceid    string      `json:"maintenanceid"`
	Proxy_hostid     string      `json:"proxy_hostid"`
	Status           HostStatus  `json:"status"`
	Tls_connect      HostTlsMode `json:"tls_connect"`
	Tls_accept       HostTlsMode `json:"tls_accept"`
	Tls_issuer       string      `json:"tls_issuer"`
	Tls_subject      string      `json:"tls_subject"`
	Tls_psk_identity string      `json:"tls_psk_identity"`
	Tls_psk          string      `json:"tls_psk"`
}

type HostId struct {
	Hostid string `json:"hostid"`
}

type HostTag struct {
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

type HostCreateParameters struct {
	Host             string                   `json:"host"`
	Groups           []*HostGroupId           `json:"groups"`
	Name             string                   `json:"name,omitempty"`
	Description      string                   `json:"description,omitempty"`
	Inventory_mode   HostInventoryMode        `json:"inventory_mode,omitempty"`
	Ipmi_authtype    HostIpmiAuthType         `json:"ipmi_authtype,omitempty"`
	Ipmi_password    string                   `json:"ipmi_password,omitempty"`
	Ipmi_privilege   HostIpmiPrivilege        `json:"ipmi_privilege,omitempty"`
	Ipmi_username    string                   `json:"ipmi_username,omitempty"`
	Proxy_hostid     string                   `json:"proxy_hostid,omitempty"`
	Status           HostStatus               `json:"status,omitempty"`
	Tls_connect      HostTlsMode              `json:"tls_connect,omitempty"`
	Tls_accept       HostTlsMode              `json:"tls_accept,omitempty"`
	Tls_issuer       string                   `json:"tls_issuer,omitempty"`
	Tls_subject      string                   `json:"tls_subject,omitempty"`
	Tls_psk_identity string                   `json:"tls_psk_identity,omitempty"`
	Tls_psk          string                   `json:"tls_psk,omitempty"`
	Interfaces       []*HostInterface         `json:"interfaces,omitempty"`
	Tags             []*HostTag               `json:"tags,omitempty"`
	Templates        []*TemplateId            `json:"templates,omitempty"`
	Macros           []*HostMacro             `json:"macros,omitempty"`
	Inventory        map[HostInventory]string `json:"invetory,omitempty"`
}

type HostGetParameters struct {
	Groupids                          []string     `json:"groupids,omitempty"`
	Dserviceids                       []string     `json:"dserviceids,omitempty"`
	Graphids                          []string     `json:"graphids,omitempty"`
	Hostids                           []string     `json:"hostids,omitempty"`
	Httptestids                       []string     `json:"httptestids,omitempty"`
	Interfaceids                      []string     `json:"interfaceids,omitempty"`
	Itemids                           []string     `json:"itemids,omitempty"`
	Maintenanceids                    []string     `json:"maintenanceids,omitempty"`
	Monitored_hosts                   bool         `json:"monitored_hosts,omitempty"`
	Proxy_hosts                       bool         `json:"proxy_hosts,omitempty"`
	Proxyids                          []string     `json:"proxyids,omitempty"`
	Templated_hosts                   bool         `json:"templated_hosts,omitempty"`
	Templateids                       []string     `json:"templateids,omitempty"`
	Triggerids                        []string     `json:"triggerids,omitempty"`
	With_items                        bool         `json:"with_items,omitempty"`
	With_item_prototypes              bool         `json:"with_item_prototypes,omitempty"`
	With_simple_graph_item_prototypes bool         `json:"with_simple_graph_item_prototypes,omitempty"`
	With_graphs                       bool         `json:"with_graphs,omitempty"`
	With_graph_prototypes             bool         `json:"with_graph_prototypes,omitempty"`
	With_httptests                    bool         `json:"with_httptests,omitempty"`
	With_monitored_httptests          bool         `json:"with_monitored_httptests,omitempty"`
	With_monitored_items              bool         `json:"with_monitored_items,omitempty"`
	With_simple_graph_items           bool         `json:"with_simple_graph_items,omitempty"`
	With_triggers                     bool         `json:"with_triggers,omitempty"`
	WithProblemsSuppressed            bool         `json:"withProblemsSuppressed,omitempty"`
	Evaltype                          HostEvalType `json:"evaltype,omitempty"`
	Severities                        []string     `json:"severities,omitempty"`
	Tags                              HostTags     `json:"tags,omitempty"`
	InheritedTags                     bool         `json:"inheritedTags,omitempty"`
	SelectDiscoveries                 interface{}  `json:"selectDiscoveries,omitempty"`
	SelectDiscoveryRule               interface{}  `json:"selectDiscoveryRule,omitempty"`
	SelectGraphs                      interface{}  `json:"selectGraphs,omitempty"`
	SelectGroups                      interface{}  `json:"selectGroups,omitempty"`
	SelectHostDiscovery               interface{}  `json:"selectHostDiscovery,omitempty"`
	SelectHttpTests                   interface{}  `json:"selectHttpTests,omitempty"`
	SelectInterfaces                  interface{}  `json:"selectInterfaces,omitempty"`
	SelectInventory                   interface{}  `json:"selectInventory,omitempty"`
	SelectItems                       interface{}  `json:"selectItems,omitempty"`
	SelectMacros                      interface{}  `json:"selectMacros,omitempty"`
	SelectParentTemplates             interface{}  `json:"selectParentTemplates,omitempty"`
	SelectDashboards                  interface{}  `json:"selectDashboards,omitempty"`
	SelectTags                        interface{}  `json:"selectTags,omitempty"`
	SelectInheritedTags               interface{}  `json:"selectInheritedTags,omitempty"`
	SelectTriggers                    interface{}  `json:"selectTriggers,omitempty"`
	SelectValueMaps                   interface{}  `json:"selectValueMaps,omitempty"`
	Filter                            interface{}  `json:"filter,omitempty"`
	LimitSelects                      int          `json:"limitSelects,omitempty"`
	Search                            interface{}  `json:"search,omitempty"`
	SearchInventory                   interface{}  `json:"searchInventory,omitempty"`
	Sortfield                         []string     `json:"sortfield,omitempty"`
	CountOutput                       bool         `json:"countOutput,omitempty"`
	Editable                          bool         `json:"editable,omitempty"`
	ExcludeSearch                     bool         `json:"excludeSearch,omitempty"`
	Limit                             int          `json:"limit,omitempty"`
	Output                            interface{}  `json:"output,omitempty"`
	Preservekeys                      bool         `json:"preservekeys,omitempty"`
	SearchByAny                       bool         `json:"searchByAny,omitempty"`
	SearchWildcardsEnabled            bool         `json:"searchWildcardsEnabled,omitempty"`
	Sortorder                         []string     `json:"sortorder,omitempty"`
	StartSearch                       bool         `json:"startSearch,omitempty"`
}

type HostReponse struct {
	HostIds []string `json:"hostids"`
}

func (h *HostService) Create(p *HostCreateParameters) (*HostReponse, error) {
	req := h.Client.NewRequest("host.create", p)

	res, err := h.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := HostReponse{}
	err = h.Client.ConvertResponse(*res, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (h *HostService) Delete(ids []string) (*HostReponse, error) {
	req := h.Client.NewRequest("host.delete", ids)

	res, err := h.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := HostReponse{}
	err = h.Client.ConvertResponse(*res, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (h *HostService) Get(p *HostGetParameters) ([]*Host, error) {
	req := h.Client.NewRequest("host.get", p)

	res, err := h.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := make([]*Host, 0)
	err = h.Client.ConvertResponse(*res, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

type HostMassAddParameters struct {
	Hosts      []*HostId        `json:"hosts"`
	Groups     []*HostGroupId   `json:"groups,omitempty"`
	Interfaces []*HostInterface `json:"interfaces,omitempty"`
	Macros     []*HostMacro     `json:"macros,omitempty"`
	Templates  []*TemplateId    `json:"templates,omitempty"`
}

func (h *HostService) MassAdd(p *HostMassAddParameters) (*HostReponse, error) {
	req := h.Client.NewRequest("host.massadd", p)

	res, err := h.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := HostReponse{}
	err = h.Client.ConvertResponse(*res, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

type HostMassRemoveParameters struct {
	HostIds          []string         `json:"hostids"`
	GroupIds         []string         `json:"groupids,omitempty"`
	Interfaces       []*HostInterface `json:"interfaces,omitempty"`
	Macros           []string         `json:"macros,omitempty"`
	TemplateIds      []string         `json:"templateids,omitempty"`
	TemplateIdsClear []string         `json:"templateids_clear,omitempty"`
}

func (h *HostService) MassRemove(p *HostMassRemoveParameters) (*HostReponse, error) {
	req := h.Client.NewRequest("host.massremove", p)

	res, err := h.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := HostReponse{}
	err = h.Client.ConvertResponse(*res, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

type HostMassUpdateParameters struct {
	Hosts            []*HostId                `json:"hosts"`
	Host             string                   `json:"host,omitempty"`
	Name             string                   `json:"name,omitempty"`
	Description      string                   `json:"description,omitempty"`
	Inventory_mode   HostInventoryMode        `json:"inventory_mode,omitempty"`
	Ipmi_authtype    HostIpmiAuthType         `json:"ipmi_authtype,omitempty"`
	Ipmi_password    string                   `json:"ipmi_password,omitempty"`
	Ipmi_privilege   HostIpmiPrivilege        `json:"ipmi_privilege,omitempty"`
	Ipmi_username    string                   `json:"ipmi_username,omitempty"`
	Proxy_hostid     string                   `json:"proxy_hostid,omitempty"`
	Status           HostStatus               `json:"status,omitempty"`
	Tls_connect      HostTlsMode              `json:"tls_connect,omitempty"`
	Tls_accept       HostTlsMode              `json:"tls_accept,omitempty"`
	Tls_issuer       string                   `json:"tls_issuer,omitempty"`
	Tls_subject      string                   `json:"tls_subject,omitempty"`
	Tls_psk_identity string                   `json:"tls_psk_identity,omitempty"`
	Tls_psk          string                   `json:"tls_psk,omitempty"`
	Groups           []*HostGroupId           `json:"groups,omitempty"`
	Interfaces       []*HostInterface         `json:"interfaces,omitempty"`
	Inventory        map[HostInventory]string `json:"inventory,omitempty"`
	Macros           []string                 `json:"macros,omitempty"`
	Templates        []*TemplateId            `json:"templates,omitempty"`
	TemplatesClear   []*TemplateId            `json:"templates_clear,omitempty"`
}

func (h *HostService) MassUpdate(p *HostMassUpdateParameters) (*HostReponse, error) {
	req := h.Client.NewRequest("host.massupdate", p)

	res, err := h.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := HostReponse{}
	err = h.Client.ConvertResponse(*res, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

type HostUpdateParameters struct {
	Hostid           string                   `json:"hostid"`
	Host             string                   `json:"host,omitempty"`
	Name             string                   `json:"name,omitempty"`
	Description      string                   `json:"description,omitempty"`
	Inventory_mode   HostInventoryMode        `json:"inventory_mode,omitempty"`
	Ipmi_authtype    HostIpmiAuthType         `json:"ipmi_authtype,omitempty"`
	Ipmi_password    string                   `json:"ipmi_password,omitempty"`
	Ipmi_privilege   HostIpmiPrivilege        `json:"ipmi_privilege,omitempty"`
	Ipmi_username    string                   `json:"ipmi_username,omitempty"`
	Proxy_hostid     string                   `json:"proxy_hostid,omitempty"`
	Status           HostStatus               `json:"status,omitempty"`
	Tls_connect      HostTlsMode              `json:"tls_connect,omitempty"`
	Tls_accept       HostTlsMode              `json:"tls_accept,omitempty"`
	Tls_issuer       string                   `json:"tls_issuer,omitempty"`
	Tls_subject      string                   `json:"tls_subject,omitempty"`
	Tls_psk_identity string                   `json:"tls_psk_identity,omitempty"`
	Tls_psk          string                   `json:"tls_psk,omitempty"`
	Groups           []*HostGroupId           `json:"groups,omitempty"`
	Interfaces       []*HostInterface         `json:"interfaces,omitempty"`
	Tags             []*HostTag               `json:"tags,omitempty"`
	Inventory        map[HostInventory]string `json:"inventory,omitempty"`
	Macros           []string                 `json:"macros,omitempty"`
	Templates        []*TemplateId            `json:"templates,omitempty"`
	TemplatesClear   []*TemplateId            `json:"templates_clear,omitempty"`
}

func (h *HostService) Update(p *HostUpdateParameters) (*HostReponse, error) {
	req := h.Client.NewRequest("host.update", p)

	res, err := h.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := HostReponse{}
	err = h.Client.ConvertResponse(*res, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
