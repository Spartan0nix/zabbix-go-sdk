package zabbixgosdk

// DiscoveryRuleStatus define the available status of a discovery rule.
type DiscoveryRuleStatus string

const (
	DiscoveryRuleEnabled  DiscoveryRuleStatus = "0"
	DiscoveryRuleDisabled DiscoveryRuleStatus = "1"
)

// DiscoveryRule properties.
type DiscoveryRule struct {
	Id          string              `json:"druleid"`
	IpRange     string              `json:"iprange"`
	Name        string              `json:"name"`
	Delay       string              `json:"delay"`
	NextCheck   string              `json:"nextcheck"`
	ProxyHostId string              `json:"proxy_hostid"`
	Status      DiscoveryRuleStatus `json:"status"`
}
