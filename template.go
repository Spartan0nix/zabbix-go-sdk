package zabbixgosdk

import "fmt"

// TemplateService create a new service to access template related methods and functions.
type TemplateService struct {
	Client *ApiClient
}

// Template properties.
// Some properties are read-only, which means they are only accessible after creation
// and should not be passed as arguments in other methods.
type Template struct {
	// ReadOnly
	Id          string `json:"templateid"`
	Host        string `json:"host"`
	Name        string `json:"name"`
	Description string `json:"description"`
	// ReadOnly
	Uuid string `json:"uuid"`
}

// TemplateId define a representation for certain methods that only requires the 'templateid' property.
type TemplateId struct {
	Id string `json:"templateid"`
}

// TemplateTag define a tag assignable to a Template
type TemplateTag struct {
	Name  string `json:"tag"`
	Value string `json:"value,omitempty"`
}

// TemplateResponse define the server response format for Template methods.
type TemplateResponse struct {
	TemplateIds []string `json:"templateids"`
}

// TemplateCreateParameters define the properties needed to create a new Template
// Properties using the 'omitempty' json parameters are optional
type TemplateCreateParameters struct {
	Host        string         `json:"host"`
	Groups      []*HostGroupId `json:"groups"`
	Name        string         `json:"name,omitempty"`
	Description string         `json:"description,omitempty"`
	Tags        []*TemplateTag `json:"tags,omitempty"`
	Templates   []*TemplateId  `json:"templates,omitempty"`
	Macros      []*HostMacro   `json:"macros,omitempty"`
}

// Create is used to create a new Template.
func (t *TemplateService) Create(p *TemplateCreateParameters) (*TemplateResponse, error) {
	req := t.Client.NewRequest("template.create", p)

	res, err := t.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := TemplateResponse{}
	err = t.Client.ConvertResponse(*res, &r)
	if err != nil {
		if t.Client.ResourceAlreadyExist("Template with host name", p.Host, res.Error) {
			fmt.Println(res.Error.Data)
		} else {
			return nil, err
		}
	}

	return &r, nil
}

// Delete is used to delete one or multiples Templates.
func (t *TemplateService) Delete(ids []string) (*TemplateResponse, error) {
	req := t.Client.NewRequest("template.delete", ids)

	res, err := t.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := TemplateResponse{}
	err = t.Client.ConvertResponse(*res, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// TemplateGetParameters define the properties used to search Template(s)
// Properties using the 'omitempty' json parameters are optional
type TemplateGetParameters struct {
	TemplateIds            []string    `json:"templateids,omitempty"`
	GroupIds               []string    `json:"groupids,omitempty"`
	ParentTemplateIds      []string    `json:"parentTemplateids,omitempty"`
	HostIds                []string    `json:"hostids,omitempty"`
	GraphIds               []string    `json:"graphids,omitempty"`
	ItemIds                []string    `json:"itemids,omitempty"`
	TriggerIds             []string    `json:"triggerids,omitempty"`
	WithItems              bool        `json:"with_items,omitempty"`
	WithTriggers           bool        `json:"with_triggers,omitempty"`
	WithGraphs             bool        `json:"with_graphs,omitempty"`
	WithHttpTests          bool        `json:"with_httptests,omitempty"`
	EvalType               string      `json:"evaltype,omitempty"`
	Tags                   []string    `json:"tags,omitempty"`
	SelectGroups           interface{} `json:"selectGroups,omitempty"`
	SelectTags             interface{} `json:"selectTags,omitempty"`
	SelectHosts            interface{} `json:"selectHosts,omitempty"`
	SelectTemplates        interface{} `json:"selectTemplates,omitempty"`
	SelectParentTemplates  interface{} `json:"selectParentTemplates,omitempty"`
	SelectHttpTests        interface{} `json:"selectHttpTests,omitempty"`
	SelectItems            interface{} `json:"selectItems,omitempty"`
	SelectDiscoveries      interface{} `json:"selectDiscoveries,omitempty"`
	SelectTriggers         interface{} `json:"selectTriggers,omitempty"`
	SelectGraphs           interface{} `json:"selectGraphs,omitempty"`
	SelectMacros           interface{} `json:"selectMacros,omitempty"`
	SelectDashboards       interface{} `json:"selectDashboards,omitempty"`
	SelectValueMaps        interface{} `json:"selectValueMaps,omitempty"`
	LimitSelects           int         `json:"limitSelects,omitempty"`
	SortField              []string    `json:"sortfield,omitempty"`
	CountOutput            bool        `json:"countOutput,omitempty"`
	Editable               bool        `json:"editable,omitempty"`
	ExcludeSearch          bool        `json:"excludeSearch,omitempty"`
	Filter                 interface{} `json:"filter,omitempty"`
	Limit                  int         `json:"limit,omitempty"`
	Output                 interface{} `json:"output,omitempty"`
	PreserveKeys           bool        `json:"preservekeys,omitempty"`
	Search                 interface{} `json:"search,omitempty"`
	SearchByAny            bool        `json:"searchByAny,omitempty"`
	SearchWildcardsEnabled bool        `json:"searchWildcardsEnabled,omitempty"`
	SortOrder              []string    `json:"sortorder,omitempty"`
	StartSearch            bool        `json:"startSearch,omitempty"`
}

// TemplateGetResponse define the server response format for Get method.
type TemplateGetResponse struct {
	TemplateId        string            `json:"templateid,omitempty"`
	Host              string            `json:"host"`
	Name              string            `json:"name"`
	Description       string            `json:"description"`
	ProxyHostid       string            `json:"proxy_hostid"`
	Status            HostStatus        `json:"status"`
	Available         string            `json:"available"`
	DisableUntil      string            `json:"disable_until"`
	Error             string            `json:"error"`
	ErrorsFrom        string            `json:"errors_from"`
	LastAccess        string            `json:"lastaccess"`
	Flags             string            `json:"flags"`
	SnmpDisableUntil  string            `json:"snmp_disable_until"`
	SnmpAvailable     string            `json:"snmp_available"`
	SnmpErrorsFrom    string            `json:"snmp_errors_from"`
	SnmpError         string            `json:"snmp_error"`
	IpmiAuthtype      HostIpmiAuthType  `json:"ipmi_authtype"`
	IpmiPrivilege     HostIpmiPrivilege `json:"ipmi_privilege"`
	IpmiUsername      string            `json:"ipmi_username"`
	IpmiPassword      string            `json:"ipmi_password"`
	IpmiAvailable     string            `json:"ipmi_available"`
	IpmiDisableUntil  string            `json:"ipmi_disable_until"`
	IpmiError         string            `json:"Ipmi_error"`
	IpmiErrorsFrom    string            `json:"ipmi_errors_from"`
	JmxAvailable      string            `json:"jmx_available"`
	JmxDisableUntil   string            `json:"jmx_disable_until"`
	JmxError          string            `json:"jmx_error"`
	JmxErrorsFrom     string            `json:"jmx_errors_from"`
	MaintenanceId     string            `json:"maintenanceid"`
	MaintenanceStatus string            `json:"maintenance_status"`
	MaintenanceType   string            `json:"maintenance_type"`
	MaintenanceFrom   string            `json:"maintenance_from"`
	TlsConnect        HostTlsMode       `json:"tls_connect"`
	TlsAccept         HostTlsMode       `json:"tls_accept"`
	TlsIssuer         string            `json:"tls_issuer"`
	TlsSubject        string            `json:"tls_subject"`
	TlsPskIdentity    string            `json:"tls_psk_identity"`
	TlsPsk            string            `json:"tls_psk"`
	Uuid              string            `json:"uuid"`
}

// List is used to retrieve all Templates.
func (t *TemplateService) List() ([]*TemplateGetResponse, error) {
	p := &TemplateGetParameters{
		Filter: map[string]string{},
	}

	req := t.Client.NewRequest("template.get", p)

	res, err := t.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := make([]*TemplateGetResponse, 0)
	err = t.Client.ConvertResponse(*res, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Get is used to retrieve one or multiples Templates matching the given criteria(s).
func (t *TemplateService) Get(p *TemplateGetParameters) ([]*TemplateGetResponse, error) {
	req := t.Client.NewRequest("template.get", p)

	res, err := t.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := make([]*TemplateGetResponse, 0)
	err = t.Client.ConvertResponse(*res, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// TemplateMassAddParameters define the properties used for the MassAdd method.
// Properties using the 'omitempty' json parameters are optional.
type TemplateMassAddParameters struct {
	Templates     []*TemplateId  `json:"templates"`
	Groups        []*HostGroupId `json:"groups,omitempty"`
	Macros        []*HostMacro   `json:"macros,omitempty"`
	TemplatesLink []*TemplateId  `json:"templates_link,omitempty"`
}

// MassAdd is used to massively add properties to existing Templates.
func (t *TemplateService) MassAdd(p *TemplateMassAddParameters) (*TemplateResponse, error) {
	req := t.Client.NewRequest("template.massadd", p)

	res, err := t.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := TemplateResponse{}
	err = t.Client.ConvertResponse(*res, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
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

// MassRemove is used to massively remove properties from existing Templates.
func (t *TemplateService) MassRemove(p *TemplateMassRemoveParameters) (*TemplateResponse, error) {
	req := t.Client.NewRequest("template.massremove", p)

	res, err := t.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := TemplateResponse{}
	err = t.Client.ConvertResponse(*res, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// TemplateMassUpdateParameters define the properties used for the MassUpdate method.
// Properties using the 'omitempty' json parameters are optional.
type TemplateMassUpdateParameters struct {
	Templates      []*TemplateId  `json:"templates"`
	Groups         []*HostGroupId `json:"groups,omitempty"`
	Macros         []*HostMacro   `json:"macros,omitempty"`
	TemplateClear  []*TemplateId  `json:"templates_clear,omitempty"`
	TemplateUnlink []*TemplateId  `json:"templates_link,omitempty"`
}

// MassUpdate is used to massively update or overwrite properties from existing Templates.
func (t *TemplateService) MassUpdate(p *TemplateMassUpdateParameters) (*TemplateResponse, error) {
	req := t.Client.NewRequest("template.massupdate", p)

	res, err := t.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := TemplateResponse{}
	err = t.Client.ConvertResponse(*res, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// TemplateUpdateParameters define the properties needed for the Update method.
// Properties using the 'omitempty' json parameters are optional.
type TemplateUpdateParameters struct {
	Id            string         `json:"templateid"`
	Host          string         `json:"host,omitempty"`
	Name          string         `json:"name,omitempty"`
	Description   string         `json:"description,omitempty"`
	Groups        []*HostGroupId `json:"groups,omitempty"`
	Tags          []*TemplateTag `json:"tags,omitempty"`
	Macros        []*HostMacro   `json:"macros,omitempty"`
	TemplateLink  []*TemplateId  `json:"templates,omitempty"`
	TemplateClear []*TemplateId  `json:"templates_clear,omitempty"`
}

// Update is used to update or overwrite properties from an existing Template.
func (t *TemplateService) Update(p *TemplateUpdateParameters) (*TemplateResponse, error) {
	req := t.Client.NewRequest("template.update", p)

	res, err := t.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := TemplateResponse{}
	err = t.Client.ConvertResponse(*res, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
