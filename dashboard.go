package zabbixgosdk

// DashboardDisplayPeriod define the available types of display periods for a page.
type DashboardDisplayPeriod string

// DashboardStartType define the available types of auto start slideshow for a dashboard.
type DashboardStartType string

const (
	DashboardDisplay10s   DashboardDisplayPeriod = "10"
	DashboardDisplay30s   DashboardDisplayPeriod = "30"
	DashboardDisplay60s   DashboardDisplayPeriod = "60"
	DashboardDisplay120s  DashboardDisplayPeriod = "120"
	DashboardDisplay600s  DashboardDisplayPeriod = "600"
	DashboardDisplay1800s DashboardDisplayPeriod = "1800"
	DashboardDisplay3600s DashboardDisplayPeriod = "3600"

	DashboardNoAutoStart DashboardStartType = "0"
	DashboardAutoStart   DashboardStartType = "1"
)

// Dashboard properties.
type Dashboard struct {
	Id            string                 `json:"dashboardid,omitempty"`
	Name          string                 `json:"name,omitempty"`
	TemplateId    string                 `json:"templateid,omitempty"`
	DisplayPeriod DashboardDisplayPeriod `json:"display_period,omitempty"`
	AutoStart     DashboardStartType     `json:"auto_start,omitempty"`
	Uuid          string                 `json:"uuid,omitempty"`
}
