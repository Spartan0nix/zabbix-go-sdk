package main

import "fmt"

type TemplateService struct {
	Client *ApiClient
}

type Template struct {
	Id          string `json:"templateid"`
	Host        string `json:"host"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Uuid        string `json:"uuid"`
}

type TemplateTag struct {
	Name  string `json:"tag"`
	Value string `json:"value,omitempty"`
}

type TemplateId struct {
	Id string `json:"templateid"`
}

type TemplateGetParameters struct {
	Templateids            []string    `json:"templateids,omitempty"`
	Groupids               []string    `json:"groupids,omitempty"`
	ParentTemplateids      []string    `json:"parentTemplateids,omitempty"`
	Hostids                []string    `json:"hostids,omitempty"`
	Graphids               []string    `json:"graphids,omitempty"`
	Itemids                []string    `json:"itemids,omitempty"`
	Triggerids             []string    `json:"triggerids,omitempty"`
	With_items             bool        `json:"with_items,omitempty"`
	With_triggers          bool        `json:"with_triggers,omitempty"`
	With_graphs            bool        `json:"with_graphs,omitempty"`
	With_httptests         bool        `json:"with_httptests,omitempty"`
	Evaltype               string      `json:"evaltype,omitempty"`
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

type TemplateResponse struct {
	Templateids []string `json:"templateids"`
}

type TemplateGetResponse struct {
	Id                 string `json:"templateid,omitempty"`
	Host               string `json:"host"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	Proxy_hostid       string `json:"proxy_hostid"`
	Status             string `json:"status"`
	Available          string `json:"available"`
	Disable_until      string `json:"disable_until"`
	Error              string `json:"error"`
	Errors_from        string `json:"errors_from"`
	Lastaccess         string `json:"lastaccess"`
	Flags              string `json:"flags"`
	Snmp_disable_until string `json:"snmp_disable_until"`
	Snmp_available     string `json:"snmp_available"`
	Snmp_errors_from   string `json:"snmp_errors_from"`
	Snmp_error         string `json:"snmp_error"`
	Ipmi_authtype      string `json:"ipmi_authtype"`
	Ipmi_privilege     string `json:"ipmi_privilege"`
	Ipmi_username      string `json:"ipmi_username"`
	Ipmi_password      string `json:"ipmi_password"`
	Ipmi_available     string `json:"ipmi_available"`
	Ipmi_disable_until string `json:"ipmi_disable_until"`
	Ipmi_error         string `json:"Ipmi_error"`
	Ipmi_errors_from   string `json:"ipmi_errors_from"`
	Jmx_available      string `json:"jmx_available"`
	Jmx_disable_until  string `json:"jmx_disable_until"`
	Jmx_error          string `json:"jmx_error"`
	Jmx_errors_from    string `json:"jmx_errors_from"`
	Maintenanceid      string `json:"maintenanceid"`
	Maintenance_status string `json:"maintenance_status"`
	Maintenance_type   string `json:"maintenance_type"`
	Maintenance_from   string `json:"maintenance_from"`
	Tls_connect        string `json:"tls_connect"`
	Tls_accept         string `json:"tls_accept"`
	Tls_issuer         string `json:"tls_issuer"`
	Tls_subject        string `json:"tls_subject"`
	Tls_psk_identity   string `json:"tls_psk_identity"`
	Tls_psk            string `json:"tls_psk"`
	Uuid               string `json:"uuid"`
}

type TemplateCreateParameters struct {
	Host        string         `json:"host"`
	Groups      []*HostGroupId `json:"groups"`
	Name        string         `json:"name,omitempty"`
	Description string         `json:"description,omitempty"`
	Tags        []*TemplateTag `json:"tags,omitempty"`
	Templates   []*TemplateId  `json:"templates,omitempty"`
	Macros      []*HostMacro   `json:"macros,omitempty"`
}

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

type TemplateMassAddParameters struct {
	Templates     []*TemplateId  `json:"templates"`
	Groups        []*HostGroupId `json:"groups,omitempty"`
	Macros        []*HostMacro   `json:"macros,omitempty"`
	TemplatesLink []*TemplateId  `json:"templates_link,omitempty"`
}

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

type TemplateMassRemoveParameters struct {
	TemplateIds       []string `json:"templateids"`
	GroupIds          []string `json:"groupids,omitempty"`
	Macros            []string `json:"macros,omitempty"`
	TemplateIdsClear  []string `json:"templateids_clear,omitempty"`
	TemplateIdsUnlink []string `json:"templateids_link,omitempty"`
}

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

type TemplateMassUpdateParameters struct {
	Templates      []*TemplateId  `json:"templates"`
	Groups         []*HostGroupId `json:"groups,omitempty"`
	Macros         []*HostMacro   `json:"macros,omitempty"`
	TemplateClear  []*TemplateId  `json:"templates_clear,omitempty"`
	TemplateUnlink []*TemplateId  `json:"templates_link,omitempty"`
}

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
