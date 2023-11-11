package zabbixgosdk

// TemplateService create a new service to access template related methods and functions.
type TemplateService struct {
	client *apiClient
}

// Template properties.
type Template struct {
	Id                string            `json:"templateid,omitempty"`
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
	Uuid              string            `json:"uuid,omitempty"`
}

// TemplateId define the format for certain methods that required a reference to a Template using it's templateid.
type TemplateId struct {
	Id string `json:"templateid"`
}

// TemplateTag define a Tag assignable to a Template.
type TemplateTag struct {
	Tag   string `json:"tag,omitempty"`
	Value string `json:"value,omitempty"`
}

// TemplateValueMap define a ValueMap assignable to a Template.
type TemplateValueMap struct {
	Id       string          `json:"valuemapid,omitempty"`
	Name     string          `json:"name,omitempty"`
	Mappings []*ValueMapping `json:"mappings,omitempty"`
	Uuid     string          `json:"uuid,omitempty"`
}

// TemplateIds define the response format for most template methods.
// Ids of the affected Templates are returned as a list of string mapped under the 'templateids' property.
type TemplateIds struct {
	Ids []string `json:"templateids"`
}

// TemplateCreateParameters define the properties used to create a new Template.
// Properties using the 'omitempty' json parameters are optionals.
type TemplateCreateParameters struct {
	Host        string                       `json:"host"`
	Groups      []*HostGroupId               `json:"groups"`
	Description string                       `json:"description,omitempty"`
	Name        string                       `json:"name,omitempty"`
	Tags        []*TemplateTag               `json:"tags,omitempty"`
	Templates   []*TemplateId                `json:"templates,omitempty"`
	Macros      []*HostMacroCreateParamaters `json:",omitempty"`
}

// Create is used to create a new Template.
func (t *TemplateService) Create(p *TemplateCreateParameters) (*Response[*TemplateIds], error) {
	r := Response[*TemplateIds]{}
	err := t.client.Post("template.create", p, &r)
	if err != nil {
		return nil, err
	}

	err = r.Validate()
	return &r, err
}

// TemplateGetParameters define the properties used to search Templates.
// Properties using the 'omitempty' json parameters are optional.
type TemplateGetParameters struct {
	TemplateIds           []string     `json:"templateids,omitempty"`
	GroupIds              []string     `json:"groupids,omitempty"`
	ParentTemplateIds     []string     `json:"parentTemplateids,omitempty"`
	HostIds               []string     `json:"hostids,omitempty"`
	GraphIds              []string     `json:"graphids,omitempty"`
	ItemIds               []string     `json:"itemids,omitempty"`
	TriggerIds            []string     `json:"triggerids,omitempty"`
	WithItems             bool         `json:"with_items,omitempty"`
	WithTriggers          bool         `json:"with_triggers,omitempty"`
	WithGraphs            bool         `json:"with_graphs,omitempty"`
	WithHttpTests         bool         `json:"with_httptests,omitempty"`
	EvalType              EvalType     `json:"evaltype,omitempty"`
	Tags                  []*SearchTag `json:"tags,omitempty"`
	SelectGroups          interface{}  `json:"selectGroups,omitempty"`
	SelectTags            interface{}  `json:"selectTags,omitempty"`
	SelectHosts           interface{}  `json:"selectHosts,omitempty"`
	SelectTemplates       interface{}  `json:"selectTemplates,omitempty"`
	SelectParentTemplates interface{}  `json:"selectParentTemplates,omitempty"`
	SelectHttpTests       interface{}  `json:"selectHttpTests,omitempty"`
	SelectItems           interface{}  `json:"selectItems,omitempty"`
	SelectDiscoveries     interface{}  `json:"selectDiscoveries,omitempty"`
	SelectTriggers        interface{}  `json:"selectTriggers,omitempty"`
	SelectGraphs          interface{}  `json:"selectGraphs,omitempty"`
	SelectMacros          interface{}  `json:"selectMacros,omitempty"`
	SelectDashboards      interface{}  `json:"selectDashboards,omitempty"`
	SelectValueMaps       interface{}  `json:"selectValueMaps,omitempty"`
	LimitSelects          string       `json:"limitSelects,omitempty"`
	CommonGetParameters
}

// TemplateGetResponse define the format of the Result field for the Response struct.
type TemplateGetResponse struct {
	Template
	ProxyAddress     string               `json:"proxy_address,omitempty"`
	AutoCompress     string               `json:"auto_compress,omitempty"`
	CustomInterfaces string               `json:"custom_interfaces,omitempty"`
	Uuid             string               `json:"uuid,omitempty"`
	HostGroups       []*HostGroup         `json:"groups,omitempty"`
	Tags             []*TemplateTag       `json:"tags,omitempty"`
	Hosts            []*Host              `json:"hosts,omitempty"`
	Templates        []*Template          `json:"templates,omitempty"`
	ParentTemplates  []*Template          `json:"parentTemplates,omitempty"`
	HttpTests        []*WebScenario       `json:"httpTests,omitempty"`
	Items            []*Item              `json:"items,omitempty"`
	Discoveries      []*LowLevelDiscovery `json:"discoveries,omitempty"`
	Triggers         []*Trigger           `json:"triggers,omitempty"`
	Graphs           []*Graph             `json:"graphs,omitempty"`
	Macros           []*HostMacro         `json:"macros,omitempty"`
	Dashboards       []*Dashboard         `json:"dashboards,omitempty"`
	ValueMaps        []*TemplateValueMap  `json:"valuemaps,omitempty"`
}

// Get is used to retrieve one or multiple Templates matching the given criterias.
func (t *TemplateService) Get(p *TemplateGetParameters) (*Response[[]*TemplateGetResponse], error) {
	r := Response[[]*TemplateGetResponse]{}
	err := t.client.Post("template.get", p, &r)
	if err != nil {
		return nil, err
	}

	err = r.Validate()
	return &r, err
}

// TemplateUpdateParameters define the properties used to update a Template.
// Properties using the 'omitempty' json parameters are optional.
type TemplateUpdateParameters struct {
	Id            string                       `json:"templateid"`
	Host          string                       `json:"host,omitempty"`
	Description   string                       `json:"description,omitempty"`
	Name          string                       `json:"name,omitempty"`
	Groups        []*HostGroupId               `json:"groups,omitempty"`
	Tags          []*TemplateTag               `json:"tags,omitempty"`
	Macros        []*HostMacroUpdateParamaters `json:"macros,omitempty"`
	TemplateLink  []*TemplateId                `json:"templates,omitempty"`
	TemplateClear []*TemplateId                `json:"templates_clear,omitempty"`
}

// Update is used to update or replace properties for a Template.
func (t *TemplateService) Update(p *TemplateUpdateParameters) (*Response[*TemplateIds], error) {
	r := Response[*TemplateIds]{}
	err := t.client.Post("template.update", p, &r)
	if err != nil {
		return nil, err
	}

	err = r.Validate()
	return &r, err
}

// Delete is used to delete one or multiple Templates.
func (t *TemplateService) Delete(ids []string) (*Response[*TemplateIds], error) {
	r := Response[*TemplateIds]{}
	err := t.client.Post("template.delete", ids, &r)
	if err != nil {
		return nil, err
	}

	err = r.Validate()
	return &r, err
}

// TemplateMassAddParameters define the properties used for the MassAdd method.
// Properties using the 'omitempty' json parameters are optional.
type TemplateMassAddParameters struct {
	Templates     []*TemplateId                `json:"templates"`
	Groups        []*HostGroupId               `json:"groups,omitempty"`
	Macros        []*HostMacroCreateParamaters `json:"macros,omitempty"`
	TemplatesLink []*TemplateId                `json:"templates_link,omitempty"`
}

// MassAdd is used to massively add properties to a given list of Templates.
func (t *TemplateService) MassAdd(p *TemplateMassAddParameters) (*Response[*TemplateIds], error) {
	r := Response[*TemplateIds]{}
	err := t.client.Post("template.massadd", p, &r)
	if err != nil {
		return nil, err
	}

	err = r.Validate()
	return &r, err
}

// TemplateMassUpdateParameters define the properties used for the MassUpdate method.
// Properties using the 'omitempty' json parameters are optional.
type TemplateMassUpdateParameters struct {
	Templates      []*TemplateId                `json:"templates"`
	Groups         []*HostGroupId               `json:"groups,omitempty"`
	Macros         []*HostMacroUpdateParamaters `json:"macros,omitempty"`
	TemplateClear  []*TemplateId                `json:"templates_clear,omitempty"`
	TemplateUnlink []*TemplateId                `json:"templates_link,omitempty"`
}

// MassUpdate is used to massively update or replace properties for a given list of Templates.
func (t *TemplateService) MassUpdate(p *TemplateMassUpdateParameters) (*Response[*TemplateIds], error) {
	r := Response[*TemplateIds]{}
	err := t.client.Post("template.massupdate", p, &r)
	if err != nil {
		return nil, err
	}

	err = r.Validate()
	return &r, err
}

// TemplateMassRemoveParameters define the properties used for the MassRemove method.
// Properties using the 'omitempty' json parameters are optional.
type TemplateMassRemoveParameters struct {
	TemplateIds       []string `json:"templateids"`
	GroupIds          []string `json:"groupids,omitempty"`
	Macros            []string `json:"macros,omitempty"`
	TemplateIdsClear  []string `json:"templateids_clear,omitempty"`
	TemplateIdsUnlink []string `json:"templateids_link,omitempty"`
}

// MassRemove is used to massively remove properties for a given list of Templates.
func (t *TemplateService) MassRemove(p *TemplateMassRemoveParameters) (*Response[*TemplateIds], error) {
	r := Response[*TemplateIds]{}
	err := t.client.Post("template.massremove", p, &r)
	if err != nil {
		return nil, err
	}

	err = r.Validate()
	return &r, err
}
