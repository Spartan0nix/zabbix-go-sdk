package zabbixgosdk

// EvalType define the available evaluation operators.
type EvalType string

// IpmiAuthType define the available modes of Ipmi authentication.
type IpmiAuthType string

// IpmiPrivilege define the available level of Ipmi privilege.
type IpmiPrivilege string

// MaintenanceStatus define the available effective maintenance status.
type MaintenanceStatus string

// MaintenanceType define the available effective type of maintenance.
type MaintenanceType string

// SearchTagOperator define the available operators while searching with tags.
type SearchTagOperator string

// HostVerification define the available types of host verification.
type HostVerification string

// PeerVerification define the available types of peer verification.
type PeerVerification string

// Origin define the available types of origin for an element.
type OriginType string

// SearchTag define an object used to search elements using tags.
type SearchTag struct {
	Tag      string            `json:"tag"`
	Value    string            `json:"value"`
	Operator SearchTagOperator `json:"operator"`
}

const (
	EvalAndOr EvalType = "0"
	EvalOr    EvalType = "2"

	IpmiAuthDefault  IpmiAuthType = "-1"
	IpmiAuthNone     IpmiAuthType = "0"
	IpmiAuthMD2      IpmiAuthType = "1"
	IpmiAuthMD5      IpmiAuthType = "2"
	IpmiAuthStraight IpmiAuthType = "4"
	IpmiAuthOEM      IpmiAuthType = "5"
	IpmiAuthRMCP     IpmiAuthType = "6"

	IpmiPrivCallback IpmiPrivilege = "1"
	IpmiPrivUser     IpmiPrivilege = "2"
	IpmiPrivOperator IpmiPrivilege = "3"
	IpmiPrivAdmin    IpmiPrivilege = "4"
	IpmiPrivOEM      IpmiPrivilege = "5"

	MaintenanceNotInEffect MaintenanceStatus = "0"
	MaintenanceInEffect    MaintenanceStatus = "1"

	MaintenanceWithDataCollection    MaintenanceType = "0"
	MaintenanceWithoutDataCollection MaintenanceType = "1"

	SearchTagContains  SearchTagOperator = "0"
	SearchTagEquals    SearchTagOperator = "1"
	SearchTagNotLike   SearchTagOperator = "2"
	SearchTagNotEqual  SearchTagOperator = "3"
	SearchTagExists    SearchTagOperator = "4"
	SearchTagNotExists SearchTagOperator = "5"

	SkipHostVerification HostVerification = "0"
	VerifyHost           HostVerification = "1"

	SkipPeerVerification PeerVerification = "0"
	VerifyPeer           PeerVerification = "1"

	PlainElement      OriginType = "0"
	DiscoveredElement OriginType = "4"
)
