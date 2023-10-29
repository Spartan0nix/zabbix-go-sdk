package zabbixgosdk

// ItemType define the available types of item.
type ItemType string

// ItemValueType define the available types of information for an item.
type ItemValueType string

// ItemTraps define if an HTTP item is allowed to populate value as in the SNMP trapper item.
type ItemTraps string

// ItemAuthType define the available types of authentication method for an item.
type ItemAuthType string

// ItemOriginType define the available types of origin for an item.
type ItemOriginType string

// ItemHttpRedirect define the available configurations when it comes to HTTP redirect.
type ItemHttpRedirect string

// ItemOutputFormat define the available types an HTTP agent item can be converted to.
type ItemHttpOutputFormat string

// ItemHttpPostType define the available types of Post body format for an Http item.
type ItemHttpPostType string

// ItemHttpMethod define the available types of Http method for an Http item.
type ItemHttpMethod string

// ItemHttpResponse define what part of the Http response should be stored.
type ItemHttpRetrieveMode string

// ItemState define the available types of state for an item.
type ItemState string

// ItemState define the available type of status for an item.
type ItemStatus string

const (
	ItemZabbixAgent       ItemType = "0"
	ItemZabbixTrapper     ItemType = "2"
	ItemSimpleCheck       ItemType = "3"
	ItemZabbixInternal    ItemType = "5"
	ItemZabbixAgentActive ItemType = "7"
	ItemWeb               ItemType = "9"
	ItemExternalCheck     ItemType = "10"
	ItemDatabaseMonitor   ItemType = "11"
	ItemIpmiAgent         ItemType = "12"
	ItemSshAgent          ItemType = "13"
	ItemTelnetAgent       ItemType = "14"
	ItemCalculated        ItemType = "15"
	ItemJmxAgent          ItemType = "16"
	ItemSnmpTrap          ItemType = "17"
	ItemDependent         ItemType = "18"
	ItemHttpAgent         ItemType = "19"
	ItemSnmpAgent         ItemType = "20"
	ItemScript            ItemType = "21"

	ItemNumericFloat    ItemValueType = "0"
	ItemCharacter       ItemValueType = "1"
	ItemLog             ItemValueType = "2"
	ItemNumericUnsigned ItemValueType = "3"
	ItemText            ItemValueType = "4"

	ItemDontAllowTraps ItemTraps = "0"
	ItemAllowTraps     ItemTraps = "1"

	ItemSshPassword  ItemAuthType = "0"
	ItemSshPublicKey ItemAuthType = "1"
	ItemHttpNone     ItemAuthType = "0"
	ItemHttpBasic    ItemAuthType = "1"
	ItemHttpNtml     ItemAuthType = "2"
	ItemHttpKerberos ItemAuthType = "3"

	ItemPlain      ItemOriginType = "0"
	ItemDiscovered ItemOriginType = "1"

	ItemDontFollowRedirect ItemHttpRedirect = "0"
	ItemFollowRedirect     ItemHttpRedirect = "1"

	ItemStoreRaw    ItemHttpOutputFormat = "0"
	ItemConvertJson ItemHttpOutputFormat = "1"

	ItemPostRaw  ItemHttpPostType = "0"
	ItemPostJson ItemHttpPostType = "2"
	ItemPostXml  ItemHttpPostType = "3"

	ItemHttpGet  ItemHttpMethod = "0"
	ItemHttpPost ItemHttpMethod = "1"
	ItemHttPut   ItemHttpMethod = "2"
	ItemHttpHead ItemHttpMethod = "3"

	ItemStoreBody        ItemHttpRetrieveMode = "0"
	ItemStoreHeaders     ItemHttpRetrieveMode = "1"
	ItemStoreBodyHeaders ItemHttpRetrieveMode = "2"

	ItemNormal       ItemState = "0"
	ItemNotSupported ItemState = "1"

	ItemEnable  ItemStatus = "0"
	ItemDisable ItemStatus = "1"
)

// Item properties.
type Item struct {
	Id              string               `json:"itemid,omitempty"`
	Delay           string               `json:"delay,omitempty"`
	HostId          string               `json:"hostid,omitempty"`
	InterfaceId     string               `json:"interfaceid,omitempty"`
	Key             string               `json:"key_,omitempty"`
	Name            string               `json:"name,omitempty"`
	Type            ItemType             `json:"type,omitempty"`
	Url             string               `json:"url,omitempty"`
	ValueType       ItemValueType        `json:"value_type,omitempty"`
	AllowTraps      ItemTraps            `json:"allow_traps,omitempty"`
	AuthType        ItemAuthType         `json:"authtype,omitempty"`
	Description     string               `json:"description,omitempty"`
	Error           string               `json:"error,omitempty"`
	Flags           ItemOriginType       `json:"flags,omitempty"`
	FollowRedirects ItemHttpRedirect     `json:"follow_redirects,omitempty"`
	Headers         []string             `json:"headers,omitempty"`
	History         string               `json:"history,omitempty"`
	HttpProxy       string               `json:"http_proxy,omitempty"`
	InventoryLink   HostInventoryId      `json:"inventory_link,omitempty"`
	IpmiSensor      string               `json:"ipmi_sensor,omitempty"`
	JmxEndpoint     string               `json:"jmx_endpoint,omitempty"`
	LastClock       string               `json:"lastclock,omitempty"`
	LastNs          string               `json:"lastns,omitempty"`
	LastValue       string               `json:"lastvalue,omitempty"`
	LogTimeFmt      string               `json:"logtimefmt,omitempty"`
	MasterItemId    string               `json:"master_itemid,omitempty"`
	OutputFormat    ItemHttpOutputFormat `json:"output_format,omitempty"`
	Params          string               `json:"params,omitempty"`
	Parameters      []string             `json:"parameters,omitempty"`
	Password        string               `json:"password,omitempty"`
	PostType        ItemHttpPostType     `json:"post_type,omitempty"`
	Posts           string               `json:"posts,omitempty"`
	PrevValue       string               `json:"prevvalue,omitempty"`
	PrivateKey      string               `json:"privatekey,omitempty"`
	PublicKey       string               `json:"publickey,omitempty"`
	QueryFields     []string             `json:"query_fields,omitempty"`
	RequestMethod   ItemHttpMethod       `json:"request_method,omitempty"`
	RetrieveMode    ItemHttpRetrieveMode `json:"retrieve_mode,omitempty"`
	SnmpOid         string               `json:"snmp_oid,omitempty"`
	SslCertFile     string               `json:"ssl_cert_file,omitempty"`
	SslKeyFile      string               `json:"ssl_key_file,omitempty"`
	SslKeyPassword  string               `json:"ssl_key_password,omitempty"`
	State           ItemState            `json:"state,omitempty"`
	Status          ItemStatus           `json:"status,omitempty"`
	StatusCodes     string               `json:"status_codes,omitempty"`
	TemplateId      string               `json:"templateid,omitempty"`
	Timeout         string               `json:"timeout,omitempty"`
	TrapperHosts    string               `json:"trapper_hosts,omitempty"`
	Trends          string               `json:"trends,omitempty"`
	Units           string               `json:"units,omitempty"`
	Username        string               `json:"username,omitempty"`
	Uuid            string               `json:"uuid,omitempty"`
	ValueMapId      string               `json:"valuemapid,omitempty"`
	VerifyHost      HostVerification     `json:"verify_host,omitempty"`
	VerifyPeer      PeerVerification     `json:"verify_peer,omitempty"`
}
