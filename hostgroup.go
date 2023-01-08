package main

// HostGroupService create a new service to access host related methods and functions.
type HostGroupService struct {
	Client *ApiClient
}

// HostGroup properties.
type HostGroup struct {
	Id   string `json:"groupid"`
	Name string `json:"name"`
	// ReadOnly
	Flags string `json:"flags"`
	// ReadOnly
	Internal string `json:"internal"`
	// ReadOnly
	Uuid string `json:"uuid"`
}

// HostGroupId define a representation for certain methods that only requires the 'groupid' property.
type HostGroupId struct {
	Groupid string `json:"groupid"`
}

// HostGroupResponse define the server response format for HostGroup methods.
type HostGroupResponse struct {
	Groupids []string `json:"groupids"`
}

// Create is used to create a new HostGroup.
func (h *HostGroupService) Create(name string) (*HostGroupResponse, error) {
	// Zabbix only need the name of the Hostgroup.
	params := map[string]string{
		"name": name,
	}

	req := h.Client.NewRequest("hostgroup.create", params)

	res, err := h.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := HostGroupResponse{}
	err = h.Client.ConvertResponse(*res, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// Delete is used to delete one or multiples HostGroups.
func (h *HostGroupService) Delete(ids []string) (*HostGroupResponse, error) {
	req := h.Client.NewRequest("hostgroup.delete", ids)

	res, err := h.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := HostGroupResponse{}
	err = h.Client.ConvertResponse(*res, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// HostGroupGetParameters define the properties used to search HostGroup(s).
// Properties using the 'omitempty' json parameters are optional
type HostGroupGetParameters struct {
	Graphids                          []string    `json:"graphids,omitempty"`
	Groupids                          []string    `json:"groupids,omitempty"`
	Hostids                           []string    `json:"hostids,omitempty"`
	Maintenanceids                    []string    `json:"maintenanceids,omitempty"`
	Monitored_hosts                   bool        `json:"monitored_hosts,omitempty"`
	Real_hosts                        bool        `json:"real_hosts,omitempty"`
	Templated_hosts                   bool        `json:"templated_hosts,omitempty"`
	Templateids                       []string    `json:"templateids,omitempty"`
	Triggerids                        []string    `json:"triggerids,omitempty"`
	With_graphs                       bool        `json:"with_graphs,omitempty"`
	With_graph_prototypes             bool        `json:"with_graph_prototypes,omitempty"`
	With_hosts_and_templates          bool        `json:"with_hosts_and_templates,omitempty"`
	With_httptests                    bool        `json:"with_httptests,omitempty"`
	With_item_prototypes              bool        `json:"with_item_prototypes,omitempty"`
	With_simple_graph_item_prototypes bool        `json:"with_simple_graph_item_prototypes,omitempty"`
	With_monitored_httptests          bool        `json:"with_monitored_httptests,omitempty"`
	With_monitored_items              bool        `json:"with_monitored_items,omitempty"`
	With_monitored_triggers           bool        `json:"with_monitored_triggers,omitempty"`
	With_simple_graph_items           bool        `json:"with_simple_graph_items,omitempty"`
	With_triggers                     bool        `json:"with_triggers,omitempty"`
	SelectDiscoveryRule               interface{} `json:"selectDiscoveryRule,omitempty"`
	SelectGroupDiscovery              interface{} `json:"selectGroupDiscovery,omitempty"`
	SelectHosts                       interface{} `json:"selectHosts,omitempty"`
	SelectTemplates                   interface{} `json:"selectTemplates,omitempty"`
	LimitSelects                      int         `json:"limitSelects,omitempty"`
	Sortfield                         []string    `json:"sortfield,omitempty"`
	CountOutput                       bool        `json:"countOutput,omitempty"`
	Editable                          bool        `json:"editable,omitempty"`
	ExcludeSearch                     bool        `json:"excludeSearch,omitempty"`
	Filter                            interface{} `json:"filter,omitempty"`
	Limit                             int         `json:"limit,omitempty"`
	Output                            interface{} `json:"output,omitempty"`
	Preservekeys                      bool        `json:"preservekeys,omitempty"`
	Search                            interface{} `json:"search,omitempty"`
	SearchByAny                       bool        `json:"searchByAny,omitempty"`
	SearchWildcardsEnabled            bool        `json:"searchWildcardsEnabled,omitempty"`
	Sortorder                         []string    `json:"sortorder,omitempty"`
	StartSearch                       bool        `json:"startSearch,omitempty"`
}

// List is used to retrieve all HostGroups for the server.
func (h *HostGroupService) List() ([]*HostGroup, error) {
	p := &HostGroupGetParameters{
		Filter: map[string]string{},
	}

	req := h.Client.NewRequest("hostgroup.get", p)

	res, err := h.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := make([]*HostGroup, 0)
	err = h.Client.ConvertResponse(*res, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Get is used to retrieve one or multiples HostGroups matching the given criteria(s).
func (h *HostGroupService) Get(p *HostGroupGetParameters) ([]*HostGroup, error) {
	req := h.Client.NewRequest("hostgroup.get", p)

	res, err := h.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := make([]*HostGroup, 0)
	err = h.Client.ConvertResponse(*res, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// HostGroupMassAddParameters define the properties used for the MassAdd method.
// Properties using the 'omitempty' json parameters are optional.
type HostGroupMassAddParameters struct {
	Groups    []*HostGroupId `json:"groups"`
	Hosts     []*HostId      `json:"hosts,omitempty"`
	Templates []*TemplateId  `json:"templates,omitempty"`
}

// MassAdd is used to massively add properties to existing HostGroups.
func (h *HostGroupService) MassAdd(p *HostGroupMassAddParameters) (*HostGroupResponse, error) {
	req := h.Client.NewRequest("hostgroup.massadd", p)

	res, err := h.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := HostGroupResponse{}
	err = h.Client.ConvertResponse(*res, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// HostGroupMassRemoveParameters define the properties used for the MassRemove method.
// Properties using the 'omitempty' json parameters are optional.
type HostGroupMassRemoveParameters struct {
	GroupsIds   []string `json:"groupids"`
	HostIds     []string `json:"hostids,omitempty"`
	TemplateIds []string `json:"templateids,omitempty"`
}

// MassRemove is used to massively remove properties from existing HostGroups.
func (h *HostGroupService) MassRemove(p *HostGroupMassRemoveParameters) (*HostGroupResponse, error) {
	req := h.Client.NewRequest("hostgroup.massremove", p)

	res, err := h.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := HostGroupResponse{}
	err = h.Client.ConvertResponse(*res, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// HostGroupMassUpdateParameters define the properties used for the MassUpdate method.
// Properties using the 'omitempty' json parameters are optional.
type HostGroupMassUpdateParameters struct {
	Groups    []*HostGroupId `json:"groups"`
	Hosts     []*HostId      `json:"hosts"`
	Templates []*TemplateId  `json:"templates"`
}

// MassUpdate is used to massively update or overwrite properties from existing HostGroups.
func (h *HostGroupService) MassUpdate(p *HostGroupMassUpdateParameters) (*HostGroupResponse, error) {
	req := h.Client.NewRequest("hostgroup.massupdate", p)

	res, err := h.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := HostGroupResponse{}
	err = h.Client.ConvertResponse(*res, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// Update is used to update the name of an existing HostGroup.
func (h *HostGroupService) Update(id string, name string) (*HostGroupResponse, error) {
	params := map[string]string{
		"groupid": id,
		"name":    name,
	}

	req := h.Client.NewRequest("hostgroup.update", params)

	res, err := h.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := HostGroupResponse{}
	err = h.Client.ConvertResponse(*res, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
