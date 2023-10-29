package zabbixgosdk

// WebAuthentication define the available types of authentication for a Web Scenario.
type WebAuthentication string

// WebScenarioStatus define the available status for a Web Scenario.
type WebScenarioStatus string

const (
	WebAuthenticationNone  WebAuthentication = "0"
	WebAuthenticationBasic WebAuthentication = "1"
	WebAuthenticationNTLM  WebAuthentication = "2"

	WebScenarioEnable  WebScenarioStatus = "0"
	WebScenarioDisable WebScenarioStatus = "1"
)

// WebHTTPField define a variable, HTTP header, POST form field data of query field data.
type WebHTTPField struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// WebScenario properties.
type WebScenario struct {
	Id             string            `json:"httptestid,omitempty"`
	HostId         string            `json:"hostid,omitempty"`
	Name           string            `json:"name,omitempty"`
	UserAgent      string            `json:"agent,omitempty"`
	Authentication WebAuthentication `json:"authentication,omitempty"`
	Delay          string            `json:"delay,omitempty"`
	Headers        []*WebHTTPField   `json:"headers,omitempty"`
	HttpPassword   string            `json:"http_password,omitempty"`
	HttpProxy      string            `json:"http_proxy,omitempty"`
	HttpUser       string            `json:"http_user,omitempty"`
	NextCheck      string            `json:"nextcheck,omitempty"`
	Retries        string            `json:"retries,omitempty"`
	SslCertFile    string            `json:"ssl_cert_file,omitempty"`
	SslKeyFile     string            `json:"ssl_key_file,omitempty"`
	SslKeyPassword string            `json:"ssl_key_password,omitempty"`
	Status         WebScenarioStatus `json:"status,omitempty"`
	TemplateId     string            `json:"templateid,omitempty"`
	Variables      []*WebHTTPField   `json:"variables,omitempty"`
	VerifyHost     HostVerification  `json:"verify_host,omitempty"`
	VerifyPeer     PeerVerification  `json:"verify_peer,omitempty"`
	Uuid           string            `json:"uuid,omitempty"`
}
