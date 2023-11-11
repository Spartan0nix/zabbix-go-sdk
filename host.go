package zabbixgosdk

// HostService create a new service to access host related methods and functions.
type HostService struct {
	client *apiClient
}

// HostInventoryMode define the available mode of inventory for an host.
type HostInventoryMode string

// HostStatus define the available status for an host.
type HostStatus string

// HostTlsMode define the available tls modes for an host.
type HostTlsMode string

// HostInventory define the available inventory alias key for an host.
type HostInventory string

// HostInventory define the id associated with an alias key for an host.
type HostInventoryId string

// HostTagOperator define the available evaluation operators when searching hosts using tags.
type HostTagOperator string

const (
	HostDisabled  HostInventoryMode = "-1"
	HostManual    HostInventoryMode = "0"
	HostAutomatic HostInventoryMode = "1"

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

	HostTagContains  HostTagOperator = "0"
	HostTagEquals    HostTagOperator = "1"
	HostTagNotlike   HostTagOperator = "2"
	HostTagNotequal  HostTagOperator = "3"
	HostTagExists    HostTagOperator = "4"
	HostTagNotExists HostTagOperator = "5"
)

// Host properties.
type Host struct {
	Id                string            `json:"hostid,omitempty"`
	Host              string            `json:"host,omitempty"`
	Name              string            `json:"name,omitempty"`
	Description       string            `json:"description,omitempty"`
	Flags             OriginType        `json:"flags,omitempty"`
	InventoryMode     HostInventoryMode `json:"inventory_mode,omitempty"`
	IpmiAuthType      IpmiAuthType      `json:"ipmi_authtype,omitempty"`
	IpmiPassword      string            `json:"ipmi_password,omitempty"`
	IpmiPrivilege     IpmiPrivilege     `json:"ipmi_privilege,omitempty"`
	IpmiUsername      string            `json:"ipmi_username,omitempty"`
	MaintenanceFrom   string            `json:"maintenance_from,omitempty"`
	MaintenanceStatus MaintenanceStatus `json:"maintenance_status,omitempty"`
	MaintenanceType   MaintenanceType   `json:"maintenance_type,omitempty"`
	MaintenanceId     string            `json:"maintenanceid,omitempty"`
	ProxyHostId       string            `json:"proxy_hostid,omitempty"`
	Status            HostStatus        `json:"status,omitempty"`
	TlsConnect        HostTlsMode       `json:"tls_connect,omitempty"`
	TlsAccept         HostTlsMode       `json:"tls_accept,omitempty"`
	TlsIssuer         string            `json:"tls_issuer,omitempty"`
	TlsSubject        string            `json:"tls_subject,omitempty"`
	TlsPskIdentity    string            `json:"tls_psk_identity,omitempty"`
	TlsPsk            string            `json:"tls_psk,omitempty"`
	LastAccess        string            `json:"lastaccess,omitempty"`
}

// HostTag define a tag assignable to an Host.
type HostTag struct {
	Tag   string `json:"tag"`
	Value string `json:"value,omitempty"`
}

// HostId define the format for certain methods that required a reference to an Host using it's hostId.
type HostId struct {
	Id string `json:"hostid"`
}

// HostIds define the response format for most hostgroup methods.
// Ids of the affected hosts are returned as a list of string mapped under the 'hostids' property.
type HostIds struct {
	Ids []string `json:"hostids"`
}

// HostDiscovery define the properties when "selectHostDiscovery" is used with the host.get method.
type HostDiscovery struct {
	Host         string `json:"host"`
	HostId       string `json:"hostid"`
	ParentHostId string `json:"parent_hostid"`
	ParentItemId string `json:"parent_itemid"`
	LastCheck    string `json:"lastcheck"`
	TsDelete     string `json:"ts_delete"`
}

// HostCreateParameters define the properties used to create a new Host.
// Properties using the 'omitempty' json parameters are optional.
type HostCreateParameters struct {
	Host           string                           `json:"host"`
	Groups         []*HostGroupId                   `json:"groups"`
	Name           string                           `json:"name,omitempty"`
	Description    string                           `json:"description,omitempty"`
	InventoryMode  HostInventoryMode                `json:"inventory_mode,omitempty"`
	IpmiAuthType   IpmiAuthType                     `json:"ipmi_authtype,omitempty"`
	IpmiPassword   string                           `json:"ipmi_password,omitempty"`
	IpmiPrivilege  IpmiPrivilege                    `json:"ipmi_privilege,omitempty"`
	IpmiUsername   string                           `json:"ipmi_username,omitempty"`
	ProxyHostId    string                           `json:"proxy_hostid,omitempty"`
	Status         HostStatus                       `json:"status,omitempty"`
	TlsConnect     HostTlsMode                      `json:"tls_connect,omitempty"`
	TlsAccept      HostTlsMode                      `json:"tls_accept,omitempty"`
	TlsIssuer      string                           `json:"tls_issuer,omitempty"`
	TlsSubject     string                           `json:"tls_subject,omitempty"`
	TlsPskIdentity string                           `json:"tls_psk_identity,omitempty"`
	TlsPsk         string                           `json:"tls_psk,omitempty"`
	Interfaces     []*HostInterfaceCreateParameters `json:"interfaces,omitempty"`
	Tags           []*HostTag                       `json:"tags,omitempty"`
	Templates      []*TemplateId                    `json:"templates,omitempty"`
	Macros         []*HostMacroCreateParamaters     `json:"macros,omitempty"`
	Inventory      map[HostInventory]string         `json:"inventory,omitempty"`
}

// Create is used to create an Host.
func (h *HostService) Create(p *HostCreateParameters) (*Response[*HostIds], error) {
	r := Response[*HostIds]{}
	err := h.client.Post("host.create", p, &r)
	if err != nil {
		return nil, err
	}

	err = r.Validate()
	return &r, err
}

// HostGetParameters define the properties used to search Hosts.
// Properties using the 'omitempty' json parameters are optional.
type HostGetParameters struct {
	GroupIds                      []string          `json:"groupids,omitempty"`
	DserviceIds                   []string          `json:"dserviceids,omitempty"`
	GraphIds                      []string          `json:"graphids,omitempty"`
	HostIds                       []string          `json:"hostids,omitempty"`
	HttpTestIds                   []string          `json:"httptestids,omitempty"`
	InterfaceIds                  []string          `json:"interfaceids,omitempty"`
	ItemIds                       []string          `json:"itemids,omitempty"`
	MaintenanceIds                []string          `json:"maintenanceids,omitempty"`
	MonitoredHosts                bool              `json:"monitored_hosts,omitempty"`
	ProxyHosts                    bool              `json:"proxy_hosts,omitempty"`
	ProxyIds                      []string          `json:"proxyids,omitempty"`
	TemplatedHosts                bool              `json:"templated_hosts,omitempty"`
	TemplateIds                   []string          `json:"templateids,omitempty"`
	TriggerIds                    []string          `json:"triggerids,omitempty"`
	WithItems                     bool              `json:"with_items,omitempty"`
	WithItemPrototypes            bool              `json:"with_item_prototypes,omitempty"`
	WithSimpleGraphItemPrototypes bool              `json:"with_simple_graph_item_prototypes,omitempty"`
	WithGraphs                    bool              `json:"with_graphs,omitempty"`
	WithGraphPrototypes           bool              `json:"with_graph_prototypes,omitempty"`
	WithHttpTests                 bool              `json:"with_httptests,omitempty"`
	WithMonitoredHttpTests        bool              `json:"with_monitored_httptests,omitempty"`
	WithMonitoredItems            bool              `json:"with_monitored_items,omitempty"`
	WithSimpleGraphItems          bool              `json:"with_simple_graph_items,omitempty"`
	WithTriggers                  bool              `json:"with_triggers,omitempty"`
	WithProblemsSuppressed        bool              `json:"withProblemsSuppressed,omitempty"`
	EvalType                      EvalType          `json:"evaltype,omitempty"`
	Severities                    []TriggerSeverity `json:"severities,omitempty"`
	Tags                          []*SearchTag      `json:"tags,omitempty"`
	InheritedTags                 bool              `json:"inheritedTags,omitempty"`
	SelectDiscoveries             interface{}       `json:"selectDiscoveries,omitempty"`
	SelectDiscoveryRule           interface{}       `json:"selectDiscoveryRule,omitempty"`
	SelectGraphs                  interface{}       `json:"selectGraphs,omitempty"`
	SelectGroups                  interface{}       `json:"selectGroups,omitempty"`
	SelectHostDiscovery           interface{}       `json:"selectHostDiscovery,omitempty"`
	SelectHttpTests               interface{}       `json:"selectHttpTests,omitempty"`
	SelectInterfaces              interface{}       `json:"selectInterfaces,omitempty"`
	SelectInventory               interface{}       `json:"selectInventory,omitempty"`
	SelectItems                   interface{}       `json:"selectItems,omitempty"`
	SelectMacros                  interface{}       `json:"selectMacros,omitempty"`
	SelectParentTemplates         interface{}       `json:"selectParentTemplates,omitempty"`
	SelectDashboards              interface{}       `json:"selectDashboards,omitempty"`
	SelectTags                    interface{}       `json:"selectTags,omitempty"`
	SelectInheritedTags           interface{}       `json:"selectInheritedTags,omitempty"`
	SelectTriggers                interface{}       `json:"selectTriggers,omitempty"`
	SelectValueMaps               interface{}       `json:"selectValueMaps,omitempty"`
	LimitSelects                  string            `json:"limitSelects,omitempty"`
	SearchInventory               interface{}       `json:"searchInventory,omitempty"`
	CommonGetParameters
}

// hostGetResponseInterface is an intermediate format used during the host.get method to prevent issues while converting JSON to struct.
// For SNMP interface, the details field is returned as map[string]interface{} while for other interfaces the field is returned as []interfaces{}.
// The ParseDetails function can be used to return an *HostInterfaceDetail struct for the current hostGetResponseInterface.
type hostGetResponseInterface struct {
	InterfaceId  string                    `json:"interfaceid,omitempty"`
	HostId       string                    `json:"hostid"`
	Ip           string                    `json:"ip"`
	Dns          string                    `json:"dns"`
	Main         HostInterfaceMain         `json:"main"`
	Port         string                    `json:"port"`
	Type         HostInterfaceType         `json:"type"`
	UseIp        HostInterfaceConnection   `json:"useip"`
	Details      interface{}               `json:"details,omitempty"`
	Available    HostInterfaceAvailability `json:"available,omitempty"`
	DisableUntil string                    `json:"disable_until,omitempty"`
	Error        string                    `json:"error,omitempty"`
	ErrorsFrom   string                    `json:"errors_from,omitempty"`
}

// ParseDetails is used to parse the details field for the current hostGetResponseInterface.
func (i *hostGetResponseInterface) ParseDetails() *HostInterfaceDetail {
	detail := HostInterfaceDetail{}
	m, ok := i.Details.(map[string]interface{})
	// Field can be casted as a map[string]interface{}
	if ok {
		// Each value need to be extracted individually to prevent convertion error.

		if val, ok := m["version"]; ok {
			detail.Version = HostInterfaceSnmpVersion(val.(string))
		}

		if val, ok := m["bulk"]; ok {
			detail.Bulk = HostInterfaceSnmpBulk(val.(string))
		}

		if val, ok := m["community"]; ok {
			detail.Community = val.(string)
		}

		if val, ok := m["securityname"]; ok {
			detail.SecurityName = val.(string)
		}

		if val, ok := m["contextname"]; ok {
			detail.ContextName = val.(string)
		}

		if val, ok := m["securitylevel"]; ok {
			detail.SecurityLevel = HostInterfaceSecurityLevel(val.(string))
		}

		if val, ok := m["authpassphrase"]; ok {
			detail.AuthPassPhrase = val.(string)
		}

		if val, ok := m["privpassphrase"]; ok {
			detail.PrivPassPhrase = val.(string)
		}

		if val, ok := m["authprotocol"]; ok {
			detail.AuthProtocol = HostInterfaceAuthProtocol(val.(string))
		}

		if val, ok := m["privprotocol"]; ok {
			detail.PrivProtocol = HostInterfacePrivProtocol(val.(string))
		}
	}

	return &detail
}

// hostGetResponseInventory is an intermediate format used during the host.get method to prevent issues while converting JSON to struct.
// Instead of returned a map[HostInventory]string, entries are returned as a map[HostInventory]interface{}.
// The ParseInventory function can be used to return an map[HostInventory]string{} for the current map[HostInventory]interface{}.
type hostGetResponseInventory map[HostInventory]interface{}

// ParseInventory is used to parse the current hostGetResponseInventory.
func (i hostGetResponseInventory) ParseInventory() map[HostInventory]string {
	inventory := map[HostInventory]string{}

	for key, value := range i {
		inventory[key] = value.(string)
	}

	return inventory
}

// HostGetResponse define the format of the Result field for the Response struct.
type HostGetResponse struct {
	Host
	TemplateId       string                      `json:"templateid,omitempty"`
	ProxyAddress     string                      `json:"proxy_address,omitempty"`
	AutoCompress     string                      `json:"auto_compress,omitempty"`
	CustomInterfaces string                      `json:"custom_interfaces,omitempty"`
	Uuid             string                      `json:"uuid,omitempty"`
	Groups           []*HostGroup                `json:"groups,omitempty"`
	Tags             []*TemplateTag              `json:"tags,omitempty"`
	ParentTemplates  []*Template                 `json:"parentTemplates,omitempty"`
	HttpTests        []*WebScenario              `json:"httpTests,omitempty"`
	Items            []*Item                     `json:"items,omitempty"`
	Discoveries      []*LowLevelDiscovery        `json:"discoveries,omitempty"`
	Triggers         []*Trigger                  `json:"triggers,omitempty"`
	Graphs           []*Graph                    `json:"graphs,omitempty"`
	Macros           []*HostMacro                `json:"macros,omitempty"`
	Dashboards       []*Dashboard                `json:"dashboards,omitempty"`
	ValueMaps        []*TemplateValueMap         `json:"valuemaps,omitempty"`
	Interfaces       []*hostGetResponseInterface `json:"interfaces,omitempty"`
	HostDiscoveries  []*HostDiscovery            `json:"hostDiscovery,omitempty"`
	// Inventory        map[HostInventory]string    `json:"inventory,omitempty"`
	Inventory     hostGetResponseInventory `json:"inventory,omitempty"`
	InheritedTags []*TemplateTag           `json:"inheritedTags,omitempty"`
}

// Get is used to retrieve one or multiple Hosts matching the given criterias.
func (h *HostService) Get(p *HostGetParameters) (*Response[[]*HostGetResponse], error) {
	r := Response[[]*HostGetResponse]{}
	err := h.client.Post("host.get", p, &r)
	if err != nil {
		return nil, err
	}

	err = r.Validate()
	return &r, err
}

// HostUpdateParameters define the properties used to update an Host.
// Properties using the 'omitempty' json parameters are optional.
type HostUpdateParameters struct {
	HostId         string                           `json:"hostid"`
	Host           string                           `json:"host,omitempty"`
	Name           string                           `json:"name,omitempty"`
	Description    string                           `json:"description,omitempty"`
	InventoryMode  HostInventoryMode                `json:"inventory_mode,omitempty"`
	IpmiAuthType   IpmiAuthType                     `json:"ipmi_authtype,omitempty"`
	IpmiPassword   string                           `json:"ipmi_password,omitempty"`
	IpmiPrivilege  IpmiPrivilege                    `json:"ipmi_privilege,omitempty"`
	IpmiUsername   string                           `json:"ipmi_username,omitempty"`
	ProxyHostId    string                           `json:"proxy_hostid,omitempty"`
	Status         HostStatus                       `json:"status,omitempty"`
	TlsConnect     HostTlsMode                      `json:"tls_connect,omitempty"`
	TlsAccept      HostTlsMode                      `json:"tls_accept,omitempty"`
	TlsIssuer      string                           `json:"tls_issuer,omitempty"`
	TlsSubject     string                           `json:"tls_subject,omitempty"`
	TlsPskIdentity string                           `json:"tls_psk_identity,omitempty"`
	TlsPsk         string                           `json:"tls_psk,omitempty"`
	Groups         []*HostGroupId                   `json:"groups,omitempty"`
	Interfaces     []*HostInterfaceCreateParameters `json:"interfaces,omitempty"`
	Tags           []*HostTag                       `json:"tags,omitempty"`
	Templates      []*TemplateId                    `json:"templates,omitempty"`
	TemplatesClear []*TemplateId                    `json:"templates_clear,omitempty"`
	Macros         []*HostMacroCreateParamaters     `json:"macros,omitempty"`
	Inventory      map[HostInventory]string         `json:"inventory,omitempty"`
}

// Update is used to update or replace properties from an Host.
func (h *HostService) Update(p *HostUpdateParameters) (*Response[*HostIds], error) {
	r := Response[*HostIds]{}
	err := h.client.Post("host.update", p, &r)
	if err != nil {
		return nil, err
	}

	err = r.Validate()
	return &r, err
}

// Delete is used to delete one or multiple Hosts.
func (h *HostService) Delete(ids []string) (*Response[*HostIds], error) {
	r := Response[*HostIds]{}
	err := h.client.Post("host.delete", ids, &r)
	if err != nil {
		return nil, err
	}

	err = r.Validate()
	return &r, err
}

// HostMassAddParameters define the properties used for the MassAdd method.
// Properties using the 'omitempty' json parameters are optional.
type HostMassAddParameters struct {
	Hosts      []*HostId                        `json:"hosts"`
	Groups     []*HostGroupId                   `json:"groups,omitempty"`
	Interfaces []*HostInterfaceCreateParameters `json:"interfaces,omitempty"`
	Macros     []*HostMacroCreateParamaters     `json:"macros,omitempty"`
	Templates  []*TemplateId                    `json:"templates,omitempty"`
}

// MassAdd is used to massively add properties to a given list of Hosts.
func (h *HostService) MassAdd(p *HostMassAddParameters) (*Response[*HostIds], error) {
	r := Response[*HostIds]{}
	err := h.client.Post("host.massadd", p, &r)
	if err != nil {
		return nil, err
	}

	err = r.Validate()
	return &r, err
}

// HostMassUpdateParameters define the properties used for the MassUpdate method.
// Properties using the 'omitempty' json parameters are optional.
type HostMassUpdateParameters struct {
	Hosts          []*HostId                        `json:"hosts"`
	Host           string                           `json:"host,omitempty"`
	Name           string                           `json:"name,omitempty"`
	Description    string                           `json:"description,omitempty"`
	InventoryMode  HostInventoryMode                `json:"inventory_mode,omitempty"`
	IpmiAuthType   IpmiAuthType                     `json:"ipmi_authtype,omitempty"`
	IpmiPassword   string                           `json:"ipmi_password,omitempty"`
	IpmiPrivilege  IpmiPrivilege                    `json:"ipmi_privilege,omitempty"`
	IpmiUsername   string                           `json:"ipmi_username,omitempty"`
	ProxyHostId    string                           `json:"proxy_hostid,omitempty"`
	Status         HostStatus                       `json:"status,omitempty"`
	TlsConnect     HostTlsMode                      `json:"tls_connect,omitempty"`
	TlsAccept      HostTlsMode                      `json:"tls_accept,omitempty"`
	TlsIssuer      string                           `json:"tls_issuer,omitempty"`
	TlsSubject     string                           `json:"tls_subject,omitempty"`
	TlsPskIdentity string                           `json:"tls_psk_identity,omitempty"`
	TlsPsk         string                           `json:"tls_psk,omitempty"`
	Groups         []*HostGroupId                   `json:"groups,omitempty"`
	Interfaces     []*HostInterfaceCreateParameters `json:"interfaces,omitempty"`
	Inventory      map[HostInventory]string         `json:"inventory,omitempty"`
	Macros         []*HostMacroCreateParamaters     `json:"macros,omitempty"`
	Templates      []*TemplateId                    `json:"templates,omitempty"`
	TemplatesClear []*TemplateId                    `json:"templates_clear,omitempty"`
}

// MassUpdate is used to massively replace properties from a list of Hosts.
func (h *HostService) MassUpdate(p *HostMassUpdateParameters) (*Response[*HostIds], error) {
	r := Response[*HostIds]{}
	err := h.client.Post("host.massupdate", p, &r)
	if err != nil {
		return nil, err
	}

	err = r.Validate()
	return &r, err
}

// HostMassRemoveInterfaces define the properties required to remove interfaces using the MassRemove method.
type HostMassRemoveInterfaces struct {
	Ip   string `json:"ip"`
	Dns  string `json:"dns"`
	Port string `json:"port"`
}

// HostMassRemoveParameters define the properties used for the MassRemove method.
// Properties using the 'omitempty' json parameters are optional.
type HostMassRemoveParameters struct {
	HostIds          []string                    `json:"hostids"`
	GroupIds         []string                    `json:"groupids,omitempty"`
	Interfaces       []*HostMassRemoveInterfaces `json:"interfaces,omitempty"`
	Macros           []string                    `json:"macros,omitempty"`
	TemplateIds      []string                    `json:"templateids,omitempty"`
	TemplateIdsClear []string                    `json:"templateids_clear,omitempty"`
}

// MassRemove is used to massively remove properties from a list of Hosts.
func (h *HostService) MassRemove(p *HostMassRemoveParameters) (*Response[*HostIds], error) {
	r := Response[*HostIds]{}
	err := h.client.Post("host.massremove", p, &r)
	if err != nil {
		return nil, err
	}

	err = r.Validate()
	return &r, err
}

// // GetId is used to retrieve the zabbix id associated with the curent HostInventory field
// func GetHostInventoryId(h HostInventory) HostInventoryId {
// 	switch h {
// 	case Alias:
// 		return "4"
// 	case AssetTag:
// 		return "11"
// 	case Chassis:
// 		return "28"
// 	case Contact:
// 		return "23"
// 	case ContractNumber:
// 		return "32"
// 	case DateHwDecomm:
// 		return "47"
// 	case DateHwExpiry:
// 		return "46"
// 	case DateHwInstall:
// 		return "45"
// 	case DateHwPurchase:
// 		return "44"
// 	case DeploymentStatus:
// 		return "34"
// 	case Hardware:
// 		return "14"
// 	case HardwareFull:
// 		return "15"
// 	case HostNetmask:
// 		return "39"
// 	case HostNetworks:
// 		return "38"
// 	case HostRouter:
// 		return "40"
// 	case HwArch:
// 		return "30"
// 	case InstallerName:
// 		return "33"
// 	case Location:
// 		return "24"
// 	case LocationLat:
// 		return "25"
// 	case LocationLon:
// 		return "26"
// 	case MacAddressA:
// 		return "12"
// 	case MacAddressB:
// 		return "13"
// 	case Model:
// 		return "29"
// 	case Name:
// 		return "3"
// 	case Notes:
// 		return "27"
// 	case OobIp:
// 		return "41"
// 	case OobNetmask:
// 		return "42"
// 	case OobRouter:
// 		return "43"
// 	case Os:
// 		return "5"
// 	case OsFull:
// 		return "6"
// 	case OsShort:
// 		return "7"
// 	case Poc1Cell:
// 		return "61"
// 	case Poc1Email:
// 		return "58"
// 	case Poc1Name:
// 		return "57"
// 	case Poc1Notes:
// 		return "63"
// 	case Poc1PhoneA:
// 		return "59"
// 	case Poc1PhoneB:
// 		return "60"
// 	case Poc1Screen:
// 		return "62"
// 	case Poc2Cell:
// 		return "68"
// 	case Poc2Email:
// 		return "65"
// 	case Poc2Name:
// 		return "64"
// 	case Poc2Notes:
// 		return "70"
// 	case Poc2PhoneA:
// 		return "66"
// 	case Poc2PhoneB:
// 		return "67"
// 	case Poc2Screen:
// 		return "69"
// 	case SerialNoA:
// 		return "8"
// 	case SerialNoB:
// 		return "9"
// 	case SiteAddressA:
// 		return "48"
// 	case SiteAddressB:
// 		return "49"
// 	case SiteAddressC:
// 		return "50"
// 	case SiteCity:
// 		return "51"
// 	case SiteCountry:
// 		return "53"
// 	case SitNotes:
// 		return "56"
// 	case SiteRack:
// 		return "55"
// 	case SiteState:
// 		return "52"
// 	case SiteZip:
// 		return "54"
// 	case Software:
// 		return "16"
// 	case SoftwareAppA:
// 		return "18"
// 	case SoftwareAppB:
// 		return "19"
// 	case SoftwareAppC:
// 		return "20"
// 	case SoftwareAppD:
// 		return "21"
// 	case SoftwareAppE:
// 		return "22"
// 	case SoftwareFull:
// 		return "17"
// 	case Tag:
// 		return "10"
// 	case Type:
// 		return "1"
// 	case TypeFull:
// 		return "2"
// 	case UrlA:
// 		return "35"
// 	case UrlB:
// 		return "36"
// 	case UrlC:
// 		return "37"
// 	case Vendor:
// 		return "31"
// 	default:
// 		return "0"
// 	}
// }
