package zabbixgosdk

// TriggerService create a new service to access trigger related methods and functions.
type TriggerService struct {
	Client *ApiClient
}

// TriggerFlag define the available type of trigger
type TriggerFlag string

// TriggerSeverity define the available levels of trigger severity
type TriggerSeverity string

// TriggerState define the available type of trigger state
type TriggerState string

// TriggerStatus define the available type of trigger status
type TriggerStatus string

// TriggerType define the available type of trigger
type TriggerType string

// TriggerValue define the available type of value for a trigger
type TriggerValue string

// TriggerRecoveryMode define the available type of recovery mode for a trigger
type TriggerRecoveryMode string

// TriggerCorrelationMode define the available type of correlation mode for a trigger
type TriggerCorrelationMode string

// TriggerManualClose define the available type if a trigger can be manually closed
type TriggerManualClose string

const (
	TriggerPlan       TriggerFlag = "0"
	TriggerDiscovered TriggerFlag = "4"

	TriggerNotClassified TriggerSeverity = "0"
	TriggerInfo          TriggerSeverity = "1"
	TriggerWarning       TriggerSeverity = "2"
	TriggerAverage       TriggerSeverity = "3"
	TriggerHigh          TriggerSeverity = "4"
	TriggerDisaster      TriggerSeverity = "5"

	TriggerUpToDate TriggerState = "0"
	TriggerUnknown  TriggerState = "1"

	TriggerEnable  TriggerStatus = "0"
	TriggerDisable TriggerStatus = "1"

	TriggerSingleEvent    TriggerType = "0"
	TriggerMultipleEvents TriggerType = "1"

	TriggerOk      TriggerValue = "0"
	TriggerProblem TriggerValue = "1"

	TriggerExpression         TriggerRecoveryMode = "0"
	TriggerRecoveryExpression TriggerRecoveryMode = "1"
	TriggerRecoveryNone       TriggerRecoveryMode = "2"

	TriggerAllProblems     TriggerCorrelationMode = "0"
	TriggerIfTagValueMatch TriggerCorrelationMode = "1"

	TriggerCloseNo  TriggerManualClose = "0"
	TriggerCloseYes TriggerManualClose = "1"
)

// Trigger properties.
// Some properties are read-only, which means they are only accessible after creation
// and should not be passed as arguments in other methods.
type Trigger struct {
	// ReadOnly
	Id          string `json:"triggerid"`
	Description string `json:"description"`
	Expression  string `json:"expression"`
	EventName   string `json:"event_name"`
	OpData      string `json:"opdata"`
	Comments    string `json:"comments"`
	// ReadOnly
	Error string `json:"error"`
	// ReadOnly
	Flags TriggerFlag `json:"flags"`
	// ReadOnly
	LastChange string          `json:"lastchange"`
	Priority   TriggerSeverity `json:"priority"`
	// ReadOnly
	State      TriggerState  `json:"state"`
	Status     TriggerStatus `json:"status"`
	TemplateId string        `json:"templateid"`
	Type       TriggerType   `json:"type"`
	Url        string        `json:"url"`
	// ReadOnly
	Value              TriggerValue           `json:"value"`
	RecoveryMode       TriggerRecoveryMode    `json:"recovery_mode"`
	RecoveryExpression string                 `json:"recovery_expression"`
	CorrelationMode    TriggerCorrelationMode `json:"correlation_mode"`
	CorrelationTag     string                 `json:"correlation_tag"`
	ManualClose        TriggerManualClose     `json:"manual_close"`
	Uuid               string                 `json:"uuid"`
}

// TriggerTag define a tag assignable to a Trigger
type TriggerTag struct {
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

// TriggerId define a representation for certain methods that only requires the 'triggerid' property.
type TriggerId struct {
	Id string `json:"triggerid "`
}

// TriggerResponse define the server response format for Trigger methods.
type TriggerResponse struct {
	TriggerIds []string `json:"triggerids"`
}

// TriggerCreateParameters define the properties needed to create a new Trigger
type TriggerCreateParameters struct {
	Description        string                 `json:"description"`
	Expression         string                 `json:"expression"`
	EventName          string                 `json:"event_name,omitempty"`
	OpData             string                 `json:"opdata,omitempty"`
	Comments           string                 `json:"comments,omitempty"`
	Priority           TriggerSeverity        `json:"priority,omitempty"`
	Status             TriggerStatus          `json:"status,omitempty"`
	TemplateId         string                 `json:"templateid,omitempty"`
	Type               TriggerType            `json:"type,omitempty"`
	Url                string                 `json:"url,omitempty"`
	RecoveryMode       TriggerRecoveryMode    `json:"recovery_mode,omitempty"`
	RecoveryExpression string                 `json:"recovery_expression,omitempty"`
	CorrelationMode    TriggerCorrelationMode `json:"correlation_mode,omitempty"`
	CorrelationTag     string                 `json:"correlation_tag,omitempty"`
	ManualClose        TriggerManualClose     `json:"manual_close,omitempty"`
	Uuid               string                 `json:"uuid,omitempty"`
	Dependencies       []*TriggerId           `json:"dependencies,omitempty"`
	Tags               []*TriggerTag          `json:"tags,omitempty"`
}

// Create is used to create a new Trigger.
func (t *TriggerService) Create(p *TriggerCreateParameters) (*TriggerResponse, error) {
	req := t.Client.NewRequest("trigger.create", p)

	res, err := t.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := TriggerResponse{}
	err = t.Client.ConvertResponse(*res, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// Delete is used to delete one or multiples Trigger.
func (t *TriggerService) Delete(ids []string) (*TriggerResponse, error) {
	req := t.Client.NewRequest("trigger.delete", ids)

	res, err := t.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := TriggerResponse{}
	err = t.Client.ConvertResponse(*res, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// TriggerGetParameters define the properties used to search Triggers.
// Properties using the 'omitempty' json parameters are optional
type TriggerGetParameters struct {
	TriggerIds                  []string        `json:"triggerids,omitempty"`
	GroupIds                    []string        `json:"groupids,omitempty"`
	TemplateIds                 []string        `json:"templateids,omitempty"`
	HostIds                     []string        `json:"hostids,omitempty"`
	ItemIds                     []string        `json:"itemids,omitempty"`
	Functions                   []string        `json:"functions,omitempty"`
	Group                       string          `json:"group,omitempty"`
	Host                        string          `json:"host,omitempty"`
	Inherited                   bool            `json:"inherited,omitempty"`
	Templated                   bool            `json:"templated,omitempty"`
	Dependent                   bool            `json:"dependent,omitempty"`
	Monitored                   bool            `json:"monitored,omitempty"`
	Active                      bool            `json:"active,omitempty"`
	Maintenance                 bool            `json:"maintenance,omitempty"`
	WithUnacknowledgedEvents    bool            `json:"withUnacknowledgedEvents,omitempty"`
	WithAcknowledgedEvents      bool            `json:"withAcknowledgedEvents,omitempty"`
	WithLastEventUnacknowledged bool            `json:"withLastEventUnacknowledged,omitempty"`
	SkipDependent               bool            `json:"skipDependent,omitempty"`
	LastChangeSince             string          `json:"lastChangeSince,omitempty"`
	LastChangeTill              string          `json:"lastChangeTill,omitempty"`
	OnlyTrue                    bool            `json:"only_true,omitempty"`
	MinSeverity                 TriggerSeverity `json:"min_severity,omitempty"`
	EvalType                    EvalType        `json:"evaltype,omitempty"`
	Tags                        interface{}     `json:"tags,omitempty"`
	ExpandComment               bool            `json:"expandComment,omitempty"`
	ExpandDescription           bool            `json:"expandDescription,omitempty"`
	ExpandExpression            bool            `json:"expandExpression,omitempty"`
	SelectGroups                interface{}     `json:"selectGroups,omitempty"`
	SelectHosts                 interface{}     `json:"selectHosts,omitempty"`
	SelectItems                 interface{}     `json:"selectItems,omitempty"`
	SelectFunctions             interface{}     `json:"selectFunctions,omitempty"`
	SelectDependencies          interface{}     `json:"selectDependencies,omitempty"`
	SelectDiscoveryRule         interface{}     `json:"selectDiscoveryRule,omitempty"`
	SelectLastEvent             interface{}     `json:"selectLastEvent,omitempty"`
	SelectTags                  interface{}     `json:"selectTags,omitempty"`
	SelectTriggerDiscovery      interface{}     `json:"selectTriggerDiscovery,omitempty"`
	LimitSelects                string          `json:"limitSelects,omitempty"`
	CountOutput                 bool            `json:"countOutput,omitempty"`
	Editable                    bool            `json:"editable,omitempty"`
	ExcludeSearch               bool            `json:"excludeSearch,omitempty"`
	Filter                      interface{}     `json:"filter,omitempty"`
	Limit                       string          `json:"limit,omitempty"`
	Output                      interface{}     `json:"output,omitempty"`
	PreserveKeys                bool            `json:"preservekeys,omitempty"`
	Search                      interface{}     `json:"search,omitempty"`
	SearchByAny                 bool            `json:"searchByAny,omitempty"`
	SearchWildcardsEnabled      bool            `json:"searchWildcardsEnabled,omitempty"`
	SortField                   []string        `json:"sortfield,omitempty"`
	SortOrder                   []string        `json:"sortorder,omitempty"`
	StartSearch                 bool            `json:"startSearch,omitempty"`
}

type TriggerFunctions struct {
	FunctionId string `json:"functionid"`
	ItemId     string `json:"itemid"`
	TriggerId  string `json:"triggerid"`
	Parameter  string `json:"parameter"`
	Last       string `json:"last"`
}

type TriggerGetResponse struct {
	Trigger
	Functions []*TriggerFunctions `json:"functions,omitempty"`
	Tags      []*TriggerTag       `json:"tags,omitempty"`
}

// Get is used to retrieve one or multiples Trigger matching the given criteria(s).
func (t *TriggerService) Get(p *TriggerGetParameters) ([]*TriggerGetResponse, error) {
	req := t.Client.NewRequest("trigger.get", p)

	res, err := t.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := make([]*TriggerGetResponse, 0)
	err = t.Client.ConvertResponse(*res, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// TriggerUpdateParameters define the properties used to update Triggers.
// Properties using the 'omitempty' json parameters are optional
type TriggerUpdateParameters struct {
	Id                 string                 `json:"triggerid"`
	Description        string                 `json:"description,omitempty"`
	Expression         string                 `json:"expression,omitempty"`
	EventName          string                 `json:"event_name,omitempty"`
	OpData             string                 `json:"opdata,omitempty"`
	Comments           string                 `json:"comments,omitempty"`
	Priority           TriggerSeverity        `json:"priority,omitempty"`
	Status             TriggerStatus          `json:"status,omitempty"`
	TemplateId         string                 `json:"templateid,omitempty"`
	Type               TriggerType            `json:"type,omitempty"`
	Url                string                 `json:"url,omitempty"`
	RecoveryMode       TriggerRecoveryMode    `json:"recovery_mode,omitempty"`
	RecoveryExpression string                 `json:"recovery_expression,omitempty"`
	CorrelationMode    TriggerCorrelationMode `json:"correlation_mode,omitempty"`
	CorrelationTag     string                 `json:"correlation_tag,omitempty"`
	ManualClose        TriggerManualClose     `json:"manual_close,omitempty"`
	Uuid               string                 `json:"uuid,omitempty"`
	Dependencies       []*TriggerId           `json:"dependencies,omitempty"`
	Tags               []*TriggerTag          `json:"tags,omitempty"`
}

// Update is used to update one or multiples Trigger.
func (t *TriggerService) Update(p *TriggerUpdateParameters) (*TriggerResponse, error) {
	req := t.Client.NewRequest("trigger.update", p)

	res, err := t.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := TriggerResponse{}
	err = t.Client.ConvertResponse(*res, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
