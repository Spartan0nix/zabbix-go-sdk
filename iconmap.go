package zabbixgosdk

// // IconMapService create a new service to access icon map related methods and functions.
// type IconMapService struct {
// 	Client *ApiClient
// }

// // IconMap properties.
// type IconMap struct {
// 	// ReadOnly
// 	Id            string `json:"iconmapid"`
// 	Name          string `json:"name"`
// 	DefaultIconId string `json:"default_iconid"`
// }

// // IconMapping define a mapping between an icon and hosts using specific inventory field value.
// type IconMapping struct {
// 	// ReadOnly
// 	IconMappingId string          `json:"iconmappingid"`
// 	IconId        string          `json:"iconid"`
// 	Expression    string          `json:"expression"`
// 	InventoryLink HostInventoryId `json:"inventory_link"`
// 	// ReadOnly
// 	IconMapId string `json:"iconmapid"`
// 	// ReadOnly
// 	SortOrder string `json:"sortorder"`
// }

// // IconMapResponse define the server response format for IconMap methods.
// type IconMapResponse struct {
// 	IconMapIds []string `json:"iconmapids"`
// }

// // IconMapCreateParameters define the properties needed to create a new Icon Mapping
// type IconMappingParameters struct {
// 	IconId        string          `json:"iconid"`
// 	Expression    string          `json:"expression"`
// 	InventoryLink HostInventoryId `json:"inventory_link"`
// }

// // IconMapCreateParameters define the properties needed to create a new Icon Map
// type IconMapCreateParameters struct {
// 	Name          string                  `json:"name"`
// 	DefaultIconId string                  `json:"default_iconid"`
// 	Mappings      []IconMappingParameters `json:"mappings"`
// }

// // Create is used to create a new IconMap.
// func (i *IconMapService) Create(p *IconMapCreateParameters) (*IconMapResponse, error) {
// 	req := i.Client.NewRequest("iconmap.create", p)

// 	res, err := i.Client.Post(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	r := IconMapResponse{}
// 	err = i.Client.ConvertResponse(*res, &r)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &r, nil
// }

// // Delete is used to delete one or multiples IconMaps.
// func (i *IconMapService) Delete(ids []string) (*IconMapResponse, error) {
// 	req := i.Client.NewRequest("iconmap.delete", ids)

// 	res, err := i.Client.Post(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	r := IconMapResponse{}
// 	err = i.Client.ConvertResponse(*res, &r)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &r, nil
// }

// // IconMapGetParameters define the properties used to search IconMaps.
// // Properties using the 'omitempty' json parameters are optional
// type IconMapGetParameters struct {
// 	Iconmapids             []string    `json:"iconmapids,omitempty"`
// 	Sysmapids              []string    `json:"sysmapids,omitempty"`
// 	SelectMappings         interface{} `json:"selectMappings,omitempty"`
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
// 	SortField              []string    `json:"sortfield,omitempty"`
// 	SortOrder              []string    `json:"sortorder,omitempty"`
// 	StartSearch            bool        `json:"startSearch,omitempty"`
// }

// // Get is used to retrieve one or multiples IconMappings matching the given criteria(s).
// func (i *IconMapService) Get(p *IconMapGetParameters) ([]*IconMapping, error) {
// 	req := i.Client.NewRequest("iconmap.get", p)

// 	res, err := i.Client.Post(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	r := make([]*IconMapping, 0)
// 	err = i.Client.ConvertResponse(*res, &r)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return r, nil
// }

// // IconMapUpdateParameters define the properties used to update IconMaps.
// // Properties using the 'omitempty' json parameters are optional
// type IconMapUpdateParameters struct {
// 	Id            string                  `json:"iconmapid"`
// 	Name          string                  `json:"name,omitempty"`
// 	DefaultIconId string                  `json:"default_iconid,omitempty"`
// 	Mappings      []IconMappingParameters `json:"mappings,omitempty"`
// }

// // Update is used to update one or multiples IconMap.
// func (i *IconMapService) Update(p *IconMapUpdateParameters) (*IconMapResponse, error) {
// 	req := i.Client.NewRequest("iconmap.update", p)

// 	res, err := i.Client.Post(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	r := IconMapResponse{}
// 	err = i.Client.ConvertResponse(*res, &r)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &r, nil
// }
