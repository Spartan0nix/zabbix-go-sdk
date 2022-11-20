package main

type HostGroupService struct {
	Client *ApiClient
}

type HostGroup struct {
	Id       string `json:"groupid"`
	Name     string `json:"name"`
	Flags    string `json:"flags"`
	Internal string `json:"internal"`
	Uuid     string `json:"uuid"`
}

type HostGroupResponse struct {
	Groupids []string `json:"groupids"`
}

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

func (h *HostGroupService) Create(name string) (*HostGroupResponse, error) {
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

type HostGroupMassGroup struct {
	Groupid string `json:"groupid"`
}

type HostGroupMassHost struct {
	Hostid string `json:"hostid"`
}

type HostGroupMassTemplate struct {
	Templateid string `json:"templateid "`
}

type HostGroupMassAddParameters struct {
	Groups    []*HostGroupMassGroup    `json:"groups"`
	Hosts     []*HostGroupMassHost     `json:"hosts,omitempty"`
	Templates []*HostGroupMassTemplate `json:"templates,omitempty"`
}

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

type HostGroupMassRemoveParameters struct {
	GroupsIds   []string `json:"groupids"`
	HostIds     []string `json:"hostids,omitempty"`
	TemplateIds []string `json:"templateids,omitempty"`
}

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

type HostGroupMassUpdateParameters struct {
	Groups    []*HostGroupMassGroup    `json:"groups"`
	Hosts     []*HostGroupMassHost     `json:"hosts"`
	Templates []*HostGroupMassTemplate `json:"templates"`
}

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
