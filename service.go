package zabbixgosdk

// // ZbxService create a new service to access service related methods and functions.
// type ZbxService struct {
// 	Client *ApiClient
// }

// // ServiceAlgorithm defined the available type of status calculation rule
// type ServiceAlgorithm string

// // ServicePropagationRule defined the available type of status propagation rule
// // See : https://www.zabbix.com/documentation/6.0/en/manual/api/reference/service/object#service for more details.
// type ServicePropagationRule string

// // ServicePropagationValue defined the available type of propagation value for a service
// type ServicePropagationValue string

// // ServiceStatusType defined the available type of status for a service
// type ServiceStatusType string

// // ServiceStatusRule defined the available type of condition for the status rule
// // See : https://www.zabbix.com/documentation/6.0/en/manual/api/reference/service/object#status-rule for more details.
// type ServiceStatusRuleType string

// // ServiceOperator defined the available type of condition operator
// type ServiceOperator string

// // ServiceTagOperator defined the available operator when using tag search
// type ServiceTagOperator string

// const (
// 	ServiceAlgorithmOk                      ServiceAlgorithm = "0"
// 	ServiceAlgorithmMostCriticalAllChildren ServiceAlgorithm = "1"
// 	ServiceAlgorithmMostCritical            ServiceAlgorithm = "2"

// 	ServicePropagationRuleAsIs     ServicePropagationRule = "0"
// 	ServicePropagationRuleIncrease ServicePropagationRule = "1"
// 	ServicePropagationRuleDecrease ServicePropagationRule = "2"
// 	ServicePropagationRuleIgnore   ServicePropagationRule = "3"
// 	ServicePropagationRuleFixed    ServicePropagationRule = "4"

// 	ServicePropagationValueOk ServicePropagationValue = "-1"
// 	// Only if PropagationRule = ServicePropagationRuleAsIs || ServicePropagationRuleIgnore
// 	ServicePropagationValueNotClassified ServicePropagationValue = "0"
// 	// Only if PropagationRule = ServicePropagationRuleIncrease || ServicePropagationRuleDecrease
// 	ServicePropagationValueInfo ServicePropagationValue = "1"
// 	// Only if PropagationRule = ServicePropagationRuleIncrease || ServicePropagationRuleDecrease
// 	ServicePropagationValueWarning ServicePropagationValue = "2"
// 	// Only if PropagationRule = ServicePropagationRuleIncrease || ServicePropagationRuleDecrease
// 	ServicePropagationValueAverage ServicePropagationValue = "3"
// 	// Only if PropagationRule = ServicePropagationRuleIncrease || ServicePropagationRuleDecrease
// 	ServicePropagationValueHigh ServicePropagationValue = "4"
// 	// Only if PropagationRule = ServicePropagationRuleIncrease || ServicePropagationRuleDecrease
// 	ServicePropagationValueDisaster ServicePropagationValue = "5"

// 	ServiceStatusOk            ServiceStatusType = "-1"
// 	ServiceStatusNotClassified ServiceStatusType = "0"
// 	ServiceStatusInfo          ServiceStatusType = "1"
// 	ServiceStatusWarning       ServiceStatusType = "2"
// 	ServiceStatusAverage       ServiceStatusType = "3"
// 	ServiceStatusHigh          ServiceStatusType = "4"
// 	ServiceStatusDisaster      ServiceStatusType = "5"

// 	ServiceStatusRuleN_AtLeast            ServiceStatusRuleType = "0"
// 	ServiceStatusRuleN_AtLeastPercentage  ServiceStatusRuleType = "1"
// 	ServiceStatusRuleN_LessThan           ServiceStatusRuleType = "2"
// 	ServiceStatusRuleN_LessThanPercentage ServiceStatusRuleType = "3"
// 	ServiceStatusRuleW_AtLeast            ServiceStatusRuleType = "4"
// 	ServiceStatusRuleW_AtLeastPercentage  ServiceStatusRuleType = "5"
// 	ServiceStatusRuleW_Below              ServiceStatusRuleType = "6"
// 	ServiceStatusRuleW_BelowPercentage    ServiceStatusRuleType = "7"

// 	ServiceOperatorEqual ServiceOperator = "0"
// 	ServiceOperatorLike  ServiceOperator = "2"

// 	ServiceTagOperatorContain        ServiceTagOperator = "0"
// 	ServiceTagOperatorEqual          ServiceTagOperator = "1"
// 	ServiceTagOperatorDoesNotContain ServiceTagOperator = "2"
// 	ServiceTagOperatorDoesNotEqual   ServiceTagOperator = "3"
// 	ServiceTagOperatorExist          ServiceTagOperator = "4"
// 	ServiceTagOperatorDoesNotExist   ServiceTagOperator = "5"
// )

// // Service properties.
// // Some properties are read-only, which means they are only accessible after creation
// // and should not be passed as arguments in other methods.
// type Service struct {
// 	// ReadOnly
// 	ServiceId string           `json:"serviceid"`
// 	Algorithm ServiceAlgorithm `json:"algorithm"`
// 	Name      string           `json:"name"`
// 	SortOrder string           `json:"sortorder"`
// 	Weight    string           `json:"weight,omitempty"`
// 	// Must be set with along PropagationValue
// 	PropagationRule ServicePropagationRule `json:"propagation_rule,omitempty"`
// 	// Must be set with along PropagationRule
// 	//
// 	// ServicePropagationValueNotClassified only if PropagationRule = ServicePropagationRuleAsIs || ServicePropagationRuleIgnore
// 	//
// 	// ServicePropagationValueInfo only if PropagationRule = ServicePropagationRuleIncrease || ServicePropagationRuleDecrease
// 	//
// 	// ServicePropagationValueWarning only if PropagationRule = ServicePropagationRuleIncrease || ServicePropagationRuleDecrease
// 	//
// 	// ServicePropagationValueAverage only if PropagationRule = ServicePropagationRuleIncrease || ServicePropagationRuleDecrease
// 	//
// 	// ServicePropagationValueHigh only if PropagationRule = ServicePropagationRuleIncrease || ServicePropagationRuleDecrease
// 	//
// 	// ServicePropagationValueDisaster only if PropagationRule = ServicePropagationRuleIncrease || ServicePropagationRuleDecrease
// 	PropagationValue ServicePropagationValue `json:"propagation_value,omitempty"`
// 	// ReadOnly
// 	Status      string `json:"status,omitempty"`
// 	Description string `json:"description,omitempty"`
// 	// ReadOnly
// 	Uuid string `json:"uuid,omitempty"`
// 	// ReadOnly
// 	CreatedAt string `json:"created_at,omitempty"`
// 	// ReadOnly
// 	ReadOnly bool `json:"readonly,omitempty"`
// }

// // Service Status Rule properties.
// // Some properties are read-only, which means they are only accessible after creation
// // and should not be passed as arguments in other methods.
// type ServiceStatusRule struct {
// 	Type ServiceStatusRuleType `json:"type"`
// 	// 	Possible values: for N and W: 1-100000; for N%: 1-100.
// 	LimitValue  string            `json:"limit_value"`
// 	LimitStatus ServiceStatusType `json:"limit_status"`
// 	// ServiceStatusOk is not supported for this field
// 	NewStatus ServiceStatusType `json:"new_status"`
// }

// // Service Tag properties.
// // Some properties are read-only, which means they are only accessible after creation
// // and should not be passed as arguments in other methods.
// type ServiceTag struct {
// 	Tag   string `json:"tag"`
// 	Value string `json:"value,omitempty"`
// }

// // Service Alarm properties.
// // Some properties are read-only, which means they are only accessible after creation
// // and should not be passed as arguments in other methods.
// type ServiceAlarm struct {
// 	Clock string `json:"clock,omitempty"`
// 	Value string `json:"value,omitempty"`
// }

// // Service Problem Tag properties.
// // Some properties are read-only, which means they are only accessible after creation
// // and should not be passed as arguments in other methods.
// type ServiceProblemTag struct {
// 	Tag      string          `json:"tag"`
// 	Operator ServiceOperator `json:"operator,omitempty"`
// 	Value    string          `json:"value,omitempty"`
// }

// // ServiceId define a representation for certain methods that only requires the 'serviceid' property.
// type ServiceId struct {
// 	ServiceId string `json:"serviceid"`
// }

// // ServiceReponse define the server response format for Service methods.
// type ServiceReponse struct {
// 	ServiceIds []string `json:"serviceids"`
// }

// type ServiceCreateParameters struct {
// 	Algorithm ServiceAlgorithm `json:"algorithm"`
// 	Name      string           `json:"name"`
// 	SortOrder string           `json:"sortorder"`
// 	Weight    string           `json:"weight,omitempty"`
// 	// Must be set with along PropagationValue
// 	PropagationRule ServicePropagationRule `json:"propagation_rule,omitempty"`
// 	// Must be set with along PropagationRule
// 	//
// 	// ServicePropagationValueNotClassified only if PropagationRule = ServicePropagationRuleAsIs || ServicePropagationRuleIgnore
// 	//
// 	// ServicePropagationValueInfo only if PropagationRule = ServicePropagationRuleIncrease || ServicePropagationRuleDecrease
// 	//
// 	// ServicePropagationValueWarning only if PropagationRule = ServicePropagationRuleIncrease || ServicePropagationRuleDecrease
// 	//
// 	// ServicePropagationValueAverage only if PropagationRule = ServicePropagationRuleIncrease || ServicePropagationRuleDecrease
// 	//
// 	// ServicePropagationValueHigh only if PropagationRule = ServicePropagationRuleIncrease || ServicePropagationRuleDecrease
// 	//
// 	// ServicePropagationValueDisaster only if PropagationRule = ServicePropagationRuleIncrease || ServicePropagationRuleDecrease
// 	PropagationValue ServicePropagationValue `json:"propagation_value,omitempty"`
// 	Description      string                  `json:"description,omitempty"`
// 	Children         []*ServiceId            `json:"children,omitempty"`
// 	Parents          []*ServiceId            `json:"parents,omitempty"`
// 	Tags             []*ServiceTag           `json:"tags,omitempty"`
// 	ProblemTags      []*ServiceProblemTag    `json:"problem_tags,omitempty"`
// 	StatusRules      []*ServiceStatusRule    `json:"status_rules,omitempty"`
// }

// // Create is used to create a new Service.
// func (s *ZbxService) Create(p *ServiceCreateParameters) (*ServiceReponse, error) {
// 	req := s.Client.NewRequest("service.create", p)

// 	res, err := s.Client.Post(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	r := ServiceReponse{}
// 	err = s.Client.ConvertResponse(*res, &r)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &r, nil
// }

// // Delete is used to delete one or multiples Services.
// func (s *ZbxService) Delete(ids []string) (*ServiceReponse, error) {
// 	req := s.Client.NewRequest("service.delete", ids)

// 	res, err := s.Client.Post(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	r := ServiceReponse{}
// 	err = s.Client.ConvertResponse(*res, &r)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &r, nil
// }

// // ServiceGetTags defined the format used for tag search using Get method
// type ServiceGetTags struct {
// 	Tag      string             `json:"tag"`
// 	Value    string             `json:"value"`
// 	Operator ServiceTagOperator `json:"operator"`
// }

// // ServiceGetStatusTimeline defined the format used when searcing service with specified periods
// type ServiceGetStatusTimeline struct {
// 	PeriodFrom string `json:"period_from"`
// 	PeriodTo   string `json:"period_to"`
// }

// type ServiceGetParameters struct {
// 	ServiceIds             []string                    `json:"serviceids,omitempty"`
// 	ParentIds              []string                    `json:"parentids,omitempty"`
// 	DeepParentIds          bool                        `json:"deep_parentids,omitempty"`
// 	ChildIds               []string                    `json:"childids,omitempty"`
// 	EvalType               EvalType                    `json:"evaltype,omitempty"`
// 	Tags                   []*ServiceGetTags           `json:"tags,omitempty"`
// 	ProblemTags            []*ServiceGetTags           `json:"problem_tags,omitempty"`
// 	WithoutProblemTags     bool                        `json:"without_problem_tags,omitempty"`
// 	SlaIds                 []string                    `json:"slaids,omitempty"`
// 	SelectChildren         interface{}                 `json:"selectChildren,omitempty"`
// 	SelectParents          interface{}                 `json:"selectParents,omitempty"`
// 	SelectTags             interface{}                 `json:"selectTags,omitempty"`
// 	SelectProblemEvents    interface{}                 `json:"selectProblemEvents,omitempty"`
// 	SelectProblemTags      interface{}                 `json:"selectProblemTags,omitempty"`
// 	SelectStatusRules      interface{}                 `json:"selectStatusRules,omitempty"`
// 	SelectStatusTimeline   []*ServiceGetStatusTimeline `json:"selectStatusTimeline,omitempty"`
// 	CountOutput            bool                        `json:"countOutput,omitempty"`
// 	Editable               bool                        `json:"editable,omitempty"`
// 	ExcludeSearch          bool                        `json:"excludeSearch,omitempty"`
// 	Filter                 interface{}                 `json:"filter,omitempty"`
// 	Limit                  string                      `json:"limit,omitempty"`
// 	Output                 interface{}                 `json:"output,omitempty"`
// 	PreserveKeys           bool                        `json:"preservekeys,omitempty"`
// 	Search                 interface{}                 `json:"search,omitempty"`
// 	SearchByAny            bool                        `json:"searchByAny,omitempty"`
// 	SearchWildcardsEnabled bool                        `json:"searchWildcardsEnabled,omitempty"`
// 	SortField              []string                    `json:"sortfield,omitempty"`
// 	SortOrder              []string                    `json:"sortorder,omitempty"`
// 	StartSearch            bool                        `json:"startSearch,omitempty"`
// }

// type ServiceGetResponse struct {
// 	Service
// 	Parents  []*Service `json:"parents,omitempty"`
// 	Children []*Service `json:"children,omitempty"`
// }

// // Get is used to retrieve one or multiples Services matching the given criteria.
// func (s *ZbxService) Get(p *ServiceGetParameters) ([]*ServiceGetResponse, error) {
// 	req := s.Client.NewRequest("service.get", p)

// 	res, err := s.Client.Post(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	r := make([]*ServiceGetResponse, 0)
// 	err = s.Client.ConvertResponse(*res, &r)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return r, nil
// }

// type ServiceUpdateParameters struct {
// 	ServiceId string           `json:"serviceid"`
// 	Algorithm ServiceAlgorithm `json:"algorithm,omitempty"`
// 	Name      string           `json:"name,omitempty"`
// 	SortOrder string           `json:"sortorder,omitempty"`
// 	Weight    string           `json:"weight,omitempty"`
// 	// Must be set with along PropagationValue
// 	PropagationRule ServicePropagationRule `json:"propagation_rule,omitempty"`
// 	// Must be set with along PropagationRule
// 	//
// 	// ServicePropagationValueNotClassified only if PropagationRule = ServicePropagationRuleAsIs || ServicePropagationRuleIgnore
// 	//
// 	// ServicePropagationValueInfo only if PropagationRule = ServicePropagationRuleIncrease || ServicePropagationRuleDecrease
// 	//
// 	// ServicePropagationValueWarning only if PropagationRule = ServicePropagationRuleIncrease || ServicePropagationRuleDecrease
// 	//
// 	// ServicePropagationValueAverage only if PropagationRule = ServicePropagationRuleIncrease || ServicePropagationRuleDecrease
// 	//
// 	// ServicePropagationValueHigh only if PropagationRule = ServicePropagationRuleIncrease || ServicePropagationRuleDecrease
// 	//
// 	// ServicePropagationValueDisaster only if PropagationRule = ServicePropagationRuleIncrease || ServicePropagationRuleDecrease
// 	PropagationValue ServicePropagationValue `json:"propagation_value,omitempty"`
// 	Description      string                  `json:"description,omitempty"`
// 	Children         []*ServiceId            `json:"children,omitempty"`
// 	Parents          []*ServiceId            `json:"parents,omitempty"`
// 	Tags             []*ServiceTag           `json:"tags,omitempty"`
// 	ProblemTags      []*ServiceProblemTag    `json:"problem_tags,omitempty"`
// 	StatusRules      []*ServiceStatusRule    `json:"status_rules,omitempty"`
// }

// // Update is used to update or overwrite properties from an existing Host.
// func (s *ZbxService) Update(p *ServiceUpdateParameters) (*ServiceReponse, error) {
// 	req := s.Client.NewRequest("service.update", p)

// 	res, err := s.Client.Post(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	r := ServiceReponse{}
// 	err = s.Client.ConvertResponse(*res, &r)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &r, nil
// }
