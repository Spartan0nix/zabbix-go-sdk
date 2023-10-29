package zabbixgosdk

type LowLevelDiscovery struct {
	Id              string               `json:"itemid,omitempty"`
	Delay           string               `json:"delay,omitempty"`
	HostId          string               `json:"hostid,omitempty"`
	InterfaceId     string               `json:"interfaceid,omitempty"`
	Key             string               `json:"key_,omitempty"`
	Name            string               `json:"name,omitempty"`
	Type            ItemType             `json:"type,omitempty"`
	Url             string               `json:"url,omitempty"`
	AllowTraps      ItemTraps            `json:"allow_traps,omitempty"`
	AuthType        ItemAuthType         `json:"authtype,omitempty"`
	Description     string               `json:"description,omitempty"`
	Error           string               `json:"error,omitempty"`
	FollowRedirects ItemHttpRedirect     `json:"follow_redirects,omitempty"`
	Headers         []string             `json:"headers,omitempty"`
	HttpProxy       string               `json:"http_proxy,omitempty"`
	IpmiSensor      string               `json:"ipmi_sensor,omitempty"`
	JmxEndpoint     string               `json:"jmx_endpoint,omitempty"`
	Lifetime        string               `json:"lifetime,omitempty"`
	MasterItemId    string               `json:"master_itemid,omitempty"`
	OutputFormat    ItemHttpOutputFormat `json:"output_format,omitempty"`
	Params          string               `json:"params,omitempty"`
	Parameters      []string             `json:"parameters,omitempty"`
	Password        string               `json:"password,omitempty"`
	PostType        ItemHttpPostType     `json:"post_type,omitempty"`
	Posts           string               `json:"posts,omitempty"`
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
	Username        string               `json:"username,omitempty"`
	Uuid            string               `json:"uuid,omitempty"`
	VerifyHost      HostVerification     `json:"verify_host,omitempty"`
	VerifyPeer      PeerVerification     `json:"verify_peer,omitempty"`
}
