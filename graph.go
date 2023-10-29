package zabbixgosdk

// GraphType define the available types for a graph.
type GraphType string

// GraphShowType define how pie and exploded graphs should be displayed.
type GraphShowType string

// GraphShowLegend define if the legend should be displayed for a graph.
type GraphShowLegend string

// GraphShowWorkPeriod define if the working period should be displayed for a graph.
type GraphShowWorkPeriod string

// GraphShowTriggers define if trigger lines should be displayed for a graph.
type GraphShowTrigger string

// GraphCalculationType defined the available types of calculation method for axis.
type GraphCalculationType string

const (
	GraphNormal   GraphType = "0"
	GraphStacked  GraphType = "1"
	GraphPie      GraphType = "2"
	GraphExploded GraphType = "3"

	Graph2D GraphShowType = "0"
	Graph3D GraphShowType = "1"

	GraphLegendHide GraphShowLegend = "0"
	GraphLegendShow GraphShowLegend = "1"

	GraphWorkPeriodHide GraphShowWorkPeriod = "0"
	GraphWorkPeriodShow GraphShowWorkPeriod = "1"

	GraphTriggerHide GraphShowTrigger = "0"
	GraphTriggerShow GraphShowTrigger = "1"

	GraphCalculated GraphCalculationType = "0"
	GraphFixed      GraphCalculationType = "1"
	GraphItem       GraphCalculationType = "2"
)

// Graph properties.
type Graph struct {
	Id             string               `json:"graphid,omitempty"`
	Height         string               `json:"height,omitempty"`
	Name           string               `json:"name,omitempty"`
	Width          string               `json:"width,omitempty"`
	Flags          OriginType           `json:"flags,omitempty"`
	GraphType      GraphType            `json:"graphtype,omitempty"`
	PercentLeft    GraphType            `json:"percent_left,omitempty"`
	PercentRight   GraphType            `json:"percent_right,omitempty"`
	Show3d         GraphShowType        `json:"show_3d,omitempty"`
	ShowLegend     GraphShowLegend      `json:"show_legend,omitempty"`
	ShowWorkPeriod GraphShowWorkPeriod  `json:"show_work_period,omitempty"`
	ShowTriggers   GraphShowTrigger     `json:"show_triggers,omitempty"`
	TemplateId     string               `json:"templateid,omitempty"`
	YAxisMax       string               `json:"yaxismax,omitempty"`
	YAxisMin       string               `json:"yaxismin,omitempty"`
	YMaxItemId     string               `json:"ymax_itemid,omitempty"`
	YMaxType       GraphCalculationType `json:"ymax_type,omitempty"`
	YMinItemId     string               `json:"ymin_itemid,omitempty"`
	YMinType       GraphCalculationType `json:"ymin_type,omitempty"`
	Uuid           string               `json:"uuid,omitempty"`
}
