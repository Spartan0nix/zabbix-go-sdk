package zabbixgosdk

// HostGroupService create a new service to access host related methods and functions.
type HostGroupService struct {
	client *apiClient
}

// HostGroupInternal define if the HostGroup is internal or not.
type HostGroupInternal string

const (
	HostGroupNotInternal HostGroupInternal = "0"
	HostGroupIsInternal  HostGroupInternal = "1"
)

// HostGroup properties.
type HostGroup struct {
	Id       string            `json:"groupid,omitempty"`
	Name     string            `json:"name"`
	Flags    OriginType        `json:"flags,omitempty"`
	Internal HostGroupInternal `json:"internal,omitempty"`
	Uuid     string            `json:"uuid,omitempty"`
}

// HostGroupDiscovery define the properties when "selectGroupDiscovery" is used with the hostgroup.get method.
type HostGroupDiscovery struct {
	Id                     string `json:"groupid"`
	LastCheck              string `json:"lastcheck"`
	Name                   string `json:"name"`
	ParentGroupPrototypeId string `json:"parent_group_prototypeid"`
	TsDelete               string `json:"ts_delete "`
}

// HostGroupId define the format for certain methods that required a reference to an HostGroup using it's groupId.
type HostGroupId struct {
	Id string `json:"groupid"`
}

// HostGroupIds define the response format for most hostgroup methods.
// Ids of the affected hostGroups are returned as a list of string mapped under the 'groupids' property.
type HostGroupIds struct {
	Ids []string `json:"groupids"`
}

// Create is used to create a new HostGroup.
func (h *HostGroupService) Create(name string) (*Response[*HostGroupIds], error) {
	// Zabbix only need the name of the Hostgroup.
	params := map[string]string{
		"name": name,
	}

	r := Response[*HostGroupIds]{}
	err := h.client.Post("hostgroup.create", params, &r)
	if err != nil {
		return nil, err
	}

	err = r.Validate()
	return &r, err
}

// HostGroupGetParameters define the properties used to search hostGroups.
// Properties using the 'omitempty' json parameters are optional.
type HostGroupGetParameters struct {
	GraphIds                      []string    `json:"graphids,omitempty"`
	GroupIds                      []string    `json:"groupids,omitempty"`
	HostIds                       []string    `json:"hostids,omitempty"`
	MaintenanceIds                []string    `json:"maintenanceids,omitempty"`
	MonitoredHosts                bool        `json:"monitored_hosts,omitempty"`
	RealHosts                     bool        `json:"real_hosts,omitempty"`
	TemplatedHosts                bool        `json:"templated_hosts,omitempty"`
	TemplateIds                   []string    `json:"templateids,omitempty"`
	TriggerIds                    []string    `json:"triggerids,omitempty"`
	WithGraphs                    bool        `json:"with_graphs,omitempty"`
	WithGraphPrototypes           bool        `json:"with_graph_prototypes,omitempty"`
	WithHostsAndTemplates         bool        `json:"with_hosts_and_templates,omitempty"`
	WithHttpTests                 bool        `json:"with_httptests,omitempty"`
	WithItemPrototypes            bool        `json:"with_item_prototypes,omitempty"`
	WithSimpleGraphItemPrototypes bool        `json:"with_simple_graph_item_prototypes,omitempty"`
	WithMonitoredHttpTests        bool        `json:"with_monitored_httptests,omitempty"`
	WithMonitoredItems            bool        `json:"with_monitored_items,omitempty"`
	WithMonitoredTriggers         bool        `json:"with_monitored_triggers,omitempty"`
	WithSimpleGraphItems          bool        `json:"with_simple_graph_items,omitempty"`
	WithTriggers                  bool        `json:"with_triggers,omitempty"`
	SelectDiscoveryRule           interface{} `json:"selectDiscoveryRule,omitempty"`
	SelectGroupDiscovery          interface{} `json:"selectGroupDiscovery,omitempty"`
	SelectHosts                   interface{} `json:"selectHosts,omitempty"`
	SelectTemplates               interface{} `json:"selectTemplates,omitempty"`
	LimitSelects                  string      `json:"limitSelects,omitempty"`
	CommonGetParameters
}

// hostGroupGetResponse define the format of the Result field for the Response struct.
type hostGroupGetResponse struct {
	HostGroup
	DiscoveryRules   []*DiscoveryRule      `json:"discoveryRule,omitempty"`
	GroupDiscoveries []*HostGroupDiscovery `json:"groupDiscovery,omitempty"`
	Hosts            []*Host               `json:"hosts,omitempty"`
	Templates        []*Template           `json:"templates,omitempty"`
}

// Get is used to retrieve one or multiple HostGroups matching the given criterias.
func (h *HostGroupService) Get(p *HostGroupGetParameters) (*Response[[]*hostGroupGetResponse], error) {
	r := Response[[]*hostGroupGetResponse]{}
	err := h.client.Post("hostgroup.get", p, &r)
	if err != nil {
		return nil, err
	}

	err = r.Validate()
	return &r, err
}

// Update is used to update the name of an HostGroup.
func (h *HostGroupService) Update(id string, name string) (*Response[*HostGroupIds], error) {
	params := map[string]string{
		"groupid": id,
		"name":    name,
	}

	r := Response[*HostGroupIds]{}
	err := h.client.Post("hostgroup.update", params, &r)
	if err != nil {
		return nil, err
	}

	err = r.Validate()
	return &r, err
}

// Delete is used to delete one or multiple HostGroups.
func (h *HostGroupService) Delete(ids []string) (*Response[*HostGroupIds], error) {
	r := Response[*HostGroupIds]{}
	err := h.client.Post("hostgroup.delete", ids, &r)
	if err != nil {
		return nil, err
	}

	err = r.Validate()
	return &r, err
}

// HostGroupMassAddParameters define the properties used for the MassAdd method.
// Properties using the 'omitempty' json parameters are optional.
type HostGroupMassAddParameters struct {
	Groups    []*HostGroupId `json:"groups"`
	Hosts     []*HostId      `json:"hosts,omitempty"`
	Templates []*TemplateId  `json:"templates,omitempty"`
}

// MassAdd is used to massively add properties to a given list of HostGroups.
func (h *HostGroupService) MassAdd(p *HostGroupMassAddParameters) (*Response[*HostGroupIds], error) {
	r := Response[*HostGroupIds]{}
	err := h.client.Post("hostgroup.massadd", p, &r)
	if err != nil {
		return nil, err
	}

	err = r.Validate()
	return &r, err
}

// HostGroupMassRemoveParameters define the properties used for the MassRemove method.
// Properties using the 'omitempty' json parameters are optional.
type HostGroupMassRemoveParameters struct {
	GroupsIds   []string `json:"groupids"`
	HostIds     []string `json:"hostids,omitempty"`
	TemplateIds []string `json:"templateids,omitempty"`
}

// MassRemove is used to massively remove properties from a given list of HostGroups.
func (h *HostGroupService) MassRemove(p *HostGroupMassRemoveParameters) (*Response[*HostGroupIds], error) {
	r := Response[*HostGroupIds]{}
	err := h.client.Post("hostgroup.remove", p, &r)
	if err != nil {
		return nil, err
	}

	err = r.Validate()
	return &r, err
}

// HostGroupMassUpdateParameters define the properties used for the MassUpdate method.
type HostGroupMassUpdateParameters struct {
	Groups    []*HostGroupId `json:"groups"`
	Hosts     []*HostId      `json:"hosts"`
	Templates []*TemplateId  `json:"templates"`
}

// MassUpdate is used to massively replace properties from a list of HostGroups.
func (h *HostGroupService) MassUpdate(p *HostGroupMassUpdateParameters) (*Response[*HostGroupIds], error) {
	r := Response[*HostGroupIds]{}
	err := h.client.Post("hostgroup.update", p, &r)
	if err != nil {
		return nil, err
	}

	err = r.Validate()
	return &r, err
}
