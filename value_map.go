package zabbixgosdk

// ValueMapType define the available types of value for a value map entry.
type ValueMapType string

const (
	ValueExactMatch             ValueMapType = "0"
	ValueGreaterOrEqual         ValueMapType = "1"
	ValueLessOrEqual            ValueMapType = "2"
	ValueinRange                ValueMapType = "3"
	ValueMatchRegularExpression ValueMapType = "4"
	ValueIfNoMatch              ValueMapType = "5"
)

// ValueMap properties.
type ValueMap struct {
	Id       string          `json:"valuemapid,omitempty"`
	HostId   string          `json:"hostid,omitempty"`
	Name     string          `json:"name,omitempty"`
	Mappings []*ValueMapping `json:"mappings,omitempty"`
	Uuid     string          `json:"uuid,omitempty"`
}

// ValueMapping properties.
type ValueMapping struct {
	Type     ValueMapType `json:"type,omitempty"`
	Value    string       `json:"value,omitempty"`
	NewValue string       `newvalue:"uuid,omitempty"`
}
