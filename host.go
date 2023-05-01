package zabbixgosdk

// HostService create a new service to access host related methods and functions.
type HostService struct {
	Client *ApiClient
}

// HostFlag define the origin of the host.
type HostFlag string

// HostInventoryMode define the available inventory modes.
type HostInventoryMode string

// HostIpmiAuthType define the available ipmi auth modes.
type HostIpmiAuthType string

// HostIpmiPrivilege define the available ipmi privilege modes.
type HostIpmiPrivilege string

// HostMaintenanceStatus define the status of maintenance on the host.
type HostMaintenanceStatus string

// HostMaintenanceType define the type of maintenance of the host.
type HostMaintenanceType string

// HostStatus define the available host status.
type HostStatus string

// HostTlsMode define the available TLS modes.
type HostTlsMode string

// HostInventory define the available inventory modes.
type HostInventory string

// HostEvalType define the available evaluation operators.
type HostEvalType string

// HostTags define the available evaluation operators when searching hosts with tags.
type HostTags string

const (
	HostPlain      HostFlag = "0"
	HostDiscovered HostFlag = "4"

	HostDisabled  HostInventoryMode = "-1"
	HostManual    HostInventoryMode = "0"
	HostAutomatic HostInventoryMode = "1"

	HostDefault  HostIpmiAuthType = "-1"
	HostNone     HostIpmiAuthType = "0"
	HostMD2      HostIpmiAuthType = "1"
	HostMD5      HostIpmiAuthType = "2"
	HostStraight HostIpmiAuthType = "4"
	HostIpmiOEM  HostIpmiAuthType = "5"
	HostRMCP     HostIpmiAuthType = "6"

	HostCallback HostIpmiPrivilege = "1"
	HostUser     HostIpmiPrivilege = "2"
	HostOperator HostIpmiPrivilege = "3"
	HostAdmin    HostIpmiPrivilege = "4"
	HostOEM      HostIpmiPrivilege = "5"

	HostNoMaintenance       HostMaintenanceStatus = "0"
	HostMaintenanceInEffect HostMaintenanceStatus = "1"

	HostMaintenanceWithDataCollection    HostMaintenanceType = "0"
	HostMaintenanceWithoutDataCollection HostMaintenanceType = "1"

	HostMonitored   HostStatus = "0"
	HostUnmonitored HostStatus = "1"

	HostNoEncryption HostTlsMode = "1"
	HostPSK          HostTlsMode = "2"
	HostCertificate  HostTlsMode = "4"

	Alias            HostInventory = "alias"
	AssetTag         HostInventory = "asset_tag"
	Chassis          HostInventory = "chassis"
	Contact          HostInventory = "contact"
	ContractNumber   HostInventory = "contract_number"
	DateHwDecomm     HostInventory = "date_hw_decomm"
	DateHwExpiry     HostInventory = "date_hw_expiry"
	DateHwInstall    HostInventory = "date_hw_install"
	DateHwPurchase   HostInventory = "date_hw_purchase"
	DeploymentStatus HostInventory = "deployment_status"
	Hardware         HostInventory = "hardware"
	HardwareFull     HostInventory = "hardware_full"
	HostNetmask      HostInventory = "host_netmask"
	HostNetworks     HostInventory = "host_networks"
	HostRouter       HostInventory = "host_router"
	HwArch           HostInventory = "hw_arch"
	InstallerName    HostInventory = "installer_name"
	Location         HostInventory = "location"
	LocationLat      HostInventory = "location_lat"
	LocationLon      HostInventory = "location_lon"
	MacAddressA      HostInventory = "macaddress_a"
	MacAddressB      HostInventory = "macaddress_b"
	Model            HostInventory = "model"
	Name             HostInventory = "name"
	Notes            HostInventory = "notes"
	OobIp            HostInventory = "oob_ip"
	OobNetmask       HostInventory = "oob_netmask"
	OobRouter        HostInventory = "oob_router"
	Os               HostInventory = "os"
	OsFull           HostInventory = "os_full"
	OsShort          HostInventory = "os_short"
	Poc1Cell         HostInventory = "poc_1_cell"
	Poc1Email        HostInventory = "poc_1_email"
	Poc1Name         HostInventory = "poc_1_name"
	Poc1Notes        HostInventory = "poc_1_notes"
	Poc1PhoneA       HostInventory = "poc_1_phone_a"
	Poc1PhoneB       HostInventory = "poc_1_phone_b"
	Poc1Screen       HostInventory = "poc_1_screen"
	Poc2Cell         HostInventory = "poc_2_cell"
	Poc2Email        HostInventory = "poc_2_email"
	Poc2Name         HostInventory = "poc_2_name"
	Poc2Notes        HostInventory = "poc_2_notes"
	Poc2PhoneA       HostInventory = "poc_2_phone_a"
	Poc2PhoneB       HostInventory = "poc_2_phone_b"
	Poc2Screen       HostInventory = "poc_2_screen"
	SerialNoA        HostInventory = "serialno_a"
	SerialNoB        HostInventory = "serialno_b"
	SiteAddressA     HostInventory = "site_address_a"
	SiteAddressB     HostInventory = "site_address_b"
	SiteAddressC     HostInventory = "site_address_c"
	SiteCity         HostInventory = "site_city"
	SiteCountry      HostInventory = "site_country"
	SitNotes         HostInventory = "site_notes"
	SiteRack         HostInventory = "site_rack"
	SiteState        HostInventory = "site_state"
	SiteZip          HostInventory = "site_zip"
	Software         HostInventory = "software"
	SoftwareAppA     HostInventory = "software_app_a"
	SoftwareAppB     HostInventory = "software_app_b"
	SoftwareAppC     HostInventory = "software_app_c"
	SoftwareAppD     HostInventory = "software_app_d"
	SoftwareAppE     HostInventory = "software_app_e"
	SoftwareFull     HostInventory = "software_full"
	Tag              HostInventory = "tag"
	Type             HostInventory = "type"
	TypeFull         HostInventory = "type_full"
	UrlA             HostInventory = "url_a"
	UrlB             HostInventory = "url_b"
	UrlC             HostInventory = "url_c"
	Vendor           HostInventory = "vendor"

	HostEvalAndOr HostEvalType = "0"
	HostEvalOr    HostEvalType = "2"

	HostTagContains  HostTags = "0"
	HostTagEquals    HostTags = "1"
	HostTagNotlike   HostTags = "2"
	HostTagNotequal  HostTags = "3"
	HostTagExists    HostTags = "4"
	HostTagNotExists HostTags = "5"
)

// Host properties.
// Some properties are read-only, which means they are only accessible after creation
// and should not be passed as arguments in other methods.
type Host struct {
	// ReadOnly
	HostId      string `json:"hostid"`
	Host        string `json:"host"`
	Name        string `json:"name"`
	Description string `json:"description"`
	// ReadOnly
	Flags         HostFlag          `json:"flags"`
	InventoryMode HostInventoryMode `json:"inventory_mode"`
	IpmiAuthType  HostIpmiAuthType  `json:"ipmi_authtype"`
	IpmiPassword  string            `json:"ipmi_password"`
	IpmiPrivilege HostIpmiPrivilege `json:"ipmi_privilege"`
	IpmiUsername  string            `json:"ipmi_username"`
	// ReadOnly
	MaintenanceFrom string `json:"maintenance_from"`
	// ReadOnly
	MaintenanceStatus HostMaintenanceStatus `json:"maintenance_status"`
	// ReadOnly
	MaintenanceType HostMaintenanceType `json:"maintenance_type"`
	// ReadOnly
	MaintenanceId  string      `json:"maintenanceid"`
	ProxyHostId    string      `json:"proxy_hostid"`
	Status         HostStatus  `json:"status"`
	TlsConnect     HostTlsMode `json:"tls_connect"`
	TlsAccept      HostTlsMode `json:"tls_accept"`
	TlsIssuer      string      `json:"tls_issuer"`
	TlsSubject     string      `json:"tls_subject"`
	TlsPskIdentity string      `json:"tls_psk_identity"`
	TlsPsk         string      `json:"tls_psk"`
	LastAccess     string      `json:"lastaccess"`
}

// HostId define a representation for certain methods that only requires the 'hostid' property.
type HostId struct {
	HostId string `json:"hostid"`
}

// HostTag define a tag assignable to an Host
type HostTag struct {
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

// HostReponse define the server response format for Host methods.
type HostReponse struct {
	HostIds []string `json:"hostids"`
}

// HostCreateParameters define the properties needed to create a new Host
// Properties using the 'omitempty' json parameters are optional
type HostCreateParameters struct {
	Host           string                   `json:"host"`
	Groups         []*HostGroupId           `json:"groups"`
	Name           string                   `json:"name,omitempty"`
	Description    string                   `json:"description,omitempty"`
	InventoryMode  HostInventoryMode        `json:"inventory_mode,omitempty"`
	IpmiAuthType   HostIpmiAuthType         `json:"ipmi_authtype,omitempty"`
	IpmiPassword   string                   `json:"ipmi_password,omitempty"`
	IpmiPrivilege  HostIpmiPrivilege        `json:"ipmi_privilege,omitempty"`
	IpmiUsername   string                   `json:"ipmi_username,omitempty"`
	ProxyHostId    string                   `json:"proxy_hostid,omitempty"`
	Status         HostStatus               `json:"status,omitempty"`
	TlsConnect     HostTlsMode              `json:"tls_connect,omitempty"`
	TlsAccept      HostTlsMode              `json:"tls_accept,omitempty"`
	TlsIssuer      string                   `json:"tls_issuer,omitempty"`
	TlsSubject     string                   `json:"tls_subject,omitempty"`
	TlsPskIdentity string                   `json:"tls_psk_identity,omitempty"`
	TlsPsk         string                   `json:"tls_psk,omitempty"`
	Interfaces     []*HostInterface         `json:"interfaces,omitempty"`
	Tags           []*HostTag               `json:"tags,omitempty"`
	Templates      []*TemplateId            `json:"templates,omitempty"`
	Macros         []*HostMacro             `json:"macros,omitempty"`
	Inventory      map[HostInventory]string `json:"inventory,omitempty"`
}

// Create is used to create a new Host.
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

// Delete is used to delete one or multiples Hosts.
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

// HostGetParameters define the properties used to search Host(s)
// Properties using the 'omitempty' json parameters are optional
type HostGetParameters struct {
	GroupIds                      []string     `json:"groupids,omitempty"`
	DserviceIds                   []string     `json:"dserviceids,omitempty"`
	GraphIds                      []string     `json:"graphids,omitempty"`
	HostIds                       []string     `json:"hostids,omitempty"`
	HttpTestIds                   []string     `json:"httptestids,omitempty"`
	InterfaceIds                  []string     `json:"interfaceids,omitempty"`
	ItemIds                       []string     `json:"itemids,omitempty"`
	MaintenanceIds                []string     `json:"maintenanceids,omitempty"`
	MonitoredHosts                bool         `json:"monitored_hosts,omitempty"`
	ProxyHosts                    bool         `json:"proxy_hosts,omitempty"`
	ProxyIds                      []string     `json:"proxyids,omitempty"`
	TemplatedHosts                bool         `json:"templated_hosts,omitempty"`
	TemplateIds                   []string     `json:"templateids,omitempty"`
	TriggerIds                    []string     `json:"triggerids,omitempty"`
	WithItems                     bool         `json:"with_items,omitempty"`
	WithItemPrototypes            bool         `json:"with_item_prototypes,omitempty"`
	WithSimpleGraphItemPrototypes bool         `json:"with_simple_graph_item_prototypes,omitempty"`
	WithGraphs                    bool         `json:"with_graphs,omitempty"`
	WithGraphPrototypes           bool         `json:"with_graph_prototypes,omitempty"`
	WithHttpTests                 bool         `json:"with_httptests,omitempty"`
	WithMonitoredHttpTests        bool         `json:"with_monitored_httptests,omitempty"`
	WithMonitoredItems            bool         `json:"with_monitored_items,omitempty"`
	WithSimpleGraphItems          bool         `json:"with_simple_graph_items,omitempty"`
	WithTriggers                  bool         `json:"with_triggers,omitempty"`
	WithProblemsSuppressed        bool         `json:"withProblemsSuppressed,omitempty"`
	EvalType                      HostEvalType `json:"evaltype,omitempty"`
	Severities                    []string     `json:"severities,omitempty"`
	Tags                          HostTags     `json:"tags,omitempty"`
	InheritedTags                 bool         `json:"inheritedTags,omitempty"`
	SelectDiscoveries             interface{}  `json:"selectDiscoveries,omitempty"`
	SelectDiscoveryRule           interface{}  `json:"selectDiscoveryRule,omitempty"`
	SelectGraphs                  interface{}  `json:"selectGraphs,omitempty"`
	SelectGroups                  interface{}  `json:"selectGroups,omitempty"`
	SelectHostDiscovery           interface{}  `json:"selectHostDiscovery,omitempty"`
	SelectHttpTests               interface{}  `json:"selectHttpTests,omitempty"`
	SelectInterfaces              interface{}  `json:"selectInterfaces,omitempty"`
	SelectInventory               interface{}  `json:"selectInventory,omitempty"`
	SelectItems                   interface{}  `json:"selectItems,omitempty"`
	SelectMacros                  interface{}  `json:"selectMacros,omitempty"`
	SelectParentTemplates         interface{}  `json:"selectParentTemplates,omitempty"`
	SelectDashboards              interface{}  `json:"selectDashboards,omitempty"`
	SelectTags                    interface{}  `json:"selectTags,omitempty"`
	SelectInheritedTags           interface{}  `json:"selectInheritedTags,omitempty"`
	SelectTriggers                interface{}  `json:"selectTriggers,omitempty"`
	SelectValueMaps               interface{}  `json:"selectValueMaps,omitempty"`
	LimitSelects                  string       `json:"limitSelects,omitempty"`
	SearchInventory               interface{}  `json:"searchInventory,omitempty"`
	CountOutput                   bool         `json:"countOutput,omitempty"`
	Editable                      bool         `json:"editable,omitempty"`
	ExcludeSearch                 bool         `json:"excludeSearch,omitempty"`
	Filter                        interface{}  `json:"filter,omitempty"`
	Limit                         string       `json:"limit,omitempty"`
	Output                        interface{}  `json:"output,omitempty"`
	PreserveKeys                  bool         `json:"preservekeys,omitempty"`
	Search                        interface{}  `json:"search,omitempty"`
	SearchByAny                   bool         `json:"searchByAny,omitempty"`
	SearchWildcardsEnabled        bool         `json:"searchWildcardsEnabled,omitempty"`
	SortField                     []string     `json:"sortfield,omitempty"`
	SortOrder                     []string     `json:"sortorder,omitempty"`
	StartSearch                   bool         `json:"startSearch,omitempty"`
}

// Get is used to retrieve one or multiples Hosts matching the given criteria(s).
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

// HostMassAddParameters define the properties used for the MassAdd method.
// Properties using the 'omitempty' json parameters are optional.
type HostMassAddParameters struct {
	Hosts      []*HostId        `json:"hosts"`
	Groups     []*HostGroupId   `json:"groups,omitempty"`
	Interfaces []*HostInterface `json:"interfaces,omitempty"`
	Macros     []*HostMacro     `json:"macros,omitempty"`
	Templates  []*TemplateId    `json:"templates,omitempty"`
}

// MassAdd is used to massively add properties to existing Hosts.
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

// HostMassRemoveParameters define the properties used for the MassRemove method.
// Properties using the 'omitempty' json parameters are optional.
type HostMassRemoveParameters struct {
	HostIds          []string         `json:"hostids"`
	GroupIds         []string         `json:"groupids,omitempty"`
	Interfaces       []*HostInterface `json:"interfaces,omitempty"`
	Macros           []string         `json:"macros,omitempty"`
	TemplateIds      []string         `json:"templateids,omitempty"`
	TemplateIdsClear []string         `json:"templateids_clear,omitempty"`
}

// MassRemove is used to massively remove properties from existing Hosts.
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

// HostMassUpdateParameters define the properties used for the MassUpdate method.
// Properties using the 'omitempty' json parameters are optional.
type HostMassUpdateParameters struct {
	Hosts          []*HostId                `json:"hosts"`
	Host           string                   `json:"host,omitempty"`
	Name           string                   `json:"name,omitempty"`
	Description    string                   `json:"description,omitempty"`
	InventoryMode  HostInventoryMode        `json:"inventory_mode,omitempty"`
	IpmiAuthType   HostIpmiAuthType         `json:"ipmi_authtype,omitempty"`
	IpmiPassword   string                   `json:"ipmi_password,omitempty"`
	IpmiPrivilege  HostIpmiPrivilege        `json:"ipmi_privilege,omitempty"`
	IpmiUsername   string                   `json:"ipmi_username,omitempty"`
	ProxyHostId    string                   `json:"proxy_hostid,omitempty"`
	Status         HostStatus               `json:"status,omitempty"`
	TlsConnect     HostTlsMode              `json:"tls_connect,omitempty"`
	TlsAccept      HostTlsMode              `json:"tls_accept,omitempty"`
	TlsIssuer      string                   `json:"tls_issuer,omitempty"`
	TlsSubject     string                   `json:"tls_subject,omitempty"`
	TlsPskIdentity string                   `json:"tls_psk_identity,omitempty"`
	TlsPsk         string                   `json:"tls_psk,omitempty"`
	Groups         []*HostGroupId           `json:"groups,omitempty"`
	Interfaces     []*HostInterface         `json:"interfaces,omitempty"`
	Inventory      map[HostInventory]string `json:"inventory,omitempty"`
	Macros         []*HostMacro             `json:"macros,omitempty"`
	Templates      []*TemplateId            `json:"templates,omitempty"`
	TemplatesClear []*TemplateId            `json:"templates_clear,omitempty"`
}

// MassUpdate is used to massively update or overwrite properties from existing Hosts.
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

// HostUpdateParameters define the properties needed for the Update method.
// Properties using the 'omitempty' json parameters are optional.
type HostUpdateParameters struct {
	HostId         string                   `json:"hostid"`
	Host           string                   `json:"host,omitempty"`
	Name           string                   `json:"name,omitempty"`
	Description    string                   `json:"description,omitempty"`
	InventoryMode  HostInventoryMode        `json:"inventory_mode,omitempty"`
	IpmiAuthType   HostIpmiAuthType         `json:"ipmi_authtype,omitempty"`
	IpmiPassword   string                   `json:"ipmi_password,omitempty"`
	IpmiPrivilege  HostIpmiPrivilege        `json:"ipmi_privilege,omitempty"`
	IpmiUsername   string                   `json:"ipmi_username,omitempty"`
	ProxyHostId    string                   `json:"proxy_hostid,omitempty"`
	Status         HostStatus               `json:"status,omitempty"`
	TlsConnect     HostTlsMode              `json:"tls_connect,omitempty"`
	TlsAccept      HostTlsMode              `json:"tls_accept,omitempty"`
	TlsIssuer      string                   `json:"tls_issuer,omitempty"`
	TlsSubject     string                   `json:"tls_subject,omitempty"`
	TlsPskIdentity string                   `json:"tls_psk_identity,omitempty"`
	TlsPsk         string                   `json:"tls_psk,omitempty"`
	Groups         []*HostGroupId           `json:"groups,omitempty"`
	Interfaces     []*HostInterface         `json:"interfaces,omitempty"`
	Tags           []*HostTag               `json:"tags,omitempty"`
	Inventory      map[HostInventory]string `json:"inventory,omitempty"`
	Macros         []*HostMacro             `json:"macros,omitempty"`
	Templates      []*TemplateId            `json:"templates,omitempty"`
	TemplatesClear []*TemplateId            `json:"templates_clear,omitempty"`
}

// Update is used to update or overwrite properties from an existing Host.
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
