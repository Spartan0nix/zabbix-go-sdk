package zabbixgosdk

// MapService create a new service to access map related methods and functions.
type MapService struct {
	Client *ApiClient
}

// MapMacroExpanded define if macros should be expanded on a Map.
type MapMacroExpanded string

// Expanded define how trigger with problems should be displayed on a Map.
type MapProblemExpanded string

// MapGridAlignment define if grid alignment should be enabled on a Map.
type MapGridAlignment string

// MapGriShow define the available modes of grid display for a Map.
type MapGridShow string

// MapHighlight define the available modes of highlight for a Map.
type MapHighlight string

// MapLabelFormat define if advanced label should be used for a Map.
type MapLabelFormat string

// MapLabelLocation define the available locations for a label on a Map.
type MapLabelLocation string

// MapLabelType define the available types of label for element on a Map.
type MapLabelType string

// MapLabelTypeHost define the available types of label for host elements on a Map.
type MapLabelTypeHost string

// MapLabelTypeHostGroup define the available types of label for hostGroup elements on a Map.
type MapLabelTypeHostGroup string

// MapLabelTypeImage define the available types of label for image element on a Map.
type MapLabelTypeImage string

// MapLabelTypeMap define the available types of label for map elements on a Map.
type MapLabelTypeMap string

// MapLabelTypeTrigger define the available types of label for trigger elements on a Map.
type MapLabelTypeTrigger string

// MapProblemUnAck define the available modes of display for non ack problems on a Map.
type MapProblemUnAck string

// MapSharing define the available modes of sharing for a Map.
type MapSharing string

// MapSuppressedProblem define the available mode of display for suppressed problems on a Map.
type MapProblemSuppressed string

// MapElementType define the available types of map element for a Map.
type MapElementType string

// MapAreaType define how hosts from differents HostGroups should be displayed on a Map.
type MapAreaType string

// MapHostGroupDisplay define how a hostGroup element should be displayed on a Map.
type MapHostGroupDisplay string

// MapPermission define the available type of permissions for a Map.
type MapPermission string

// MapIconMap define if icon mapping should be used for host elements on a Map.
type MapElementIconMap string

// MapElementViewType define hostGroup element placing algorithm for a Mep.
type MapElementViewType string

// MapOperator define the available type of operatior for map operations.
type MapOperator string

// MapDrawType define the available type of drawing style for a Map.
type MapDrawType string

// MapShapeType define the available types of shape for a map.
type MapShapeType string

// MapFont define the available type of text font for a map.
type MapFont string

// MapHorizontalAlignment define the available types of text horizontal alignment for a map.
type MapHorizontalAlignment string

// MapVerticalAlignment define the available types of text vertical alignment for a map.
type MapVerticalAlignment string

// MapBorderType define the available types of border for a map.
type MapBorderType string

const (
	MapDoNotExpand MapMacroExpanded = "0"
	MapExpand      MapMacroExpanded = "1"

	MapAlwaysShow   MapProblemExpanded = "0"
	MapIfOneProblem MapProblemExpanded = "1"

	MapAlignmentDisable MapGridAlignment = "0"
	MapAlignmentEnable  MapGridAlignment = "1"

	MapGridHide    MapGridShow = "0"
	MapGridDisplay MapGridShow = "1"

	MapHighlightDisable MapHighlight = "0"
	MapHighlightEnable  MapHighlight = "1"

	MapLabelNormal   MapLabelFormat = "0"
	MapLabelAdvanced MapLabelFormat = "1"

	// Only for map element.
	MapLabelLocationDefault MapLabelLocation = "-1"
	MapLabelLocationBottom  MapLabelLocation = "0"
	MapLabelLocationLeft    MapLabelLocation = "1"
	MapLabelLocationRight   MapLabelLocation = "2"
	MapLabelLocationTop     MapLabelLocation = "3"

	MapLabelTypeLabel       MapLabelType = "0"
	MapLabelTypeIP          MapLabelType = "1"
	MapLabelTypeElementName MapLabelType = "2"
	MapLabelTypeStatus      MapLabelType = "3"
	MapLabelTypeNothing     MapLabelType = "4"

	MapLabelTypeHostLabel       MapLabelTypeHost = "0"
	MapLabelTypeHostIP          MapLabelTypeHost = "1"
	MapLabelTypeHostElementName MapLabelTypeHost = "2"
	MapLabelTypeHostStatus      MapLabelTypeHost = "3"
	MapLabelTypeHostNothing     MapLabelTypeHost = "4"
	MapLabelTypeHostCustom      MapLabelTypeHost = "5"

	MapLabelTypeHostGroupLabel       MapLabelTypeHostGroup = "0"
	MapLabelTypeHostGroupElementName MapLabelTypeHostGroup = "2"
	MapLabelTypeHostGroupStatus      MapLabelTypeHostGroup = "3"
	MapLabelTypeHostGroupNothing     MapLabelTypeHostGroup = "4"
	MapLabelTypeHostGroupCustom      MapLabelTypeHostGroup = "5"

	MapLabelTypeImageLabel       MapLabelTypeImage = "0"
	MapLabelTypeImageElementName MapLabelTypeImage = "2"
	MapLabelTypeImageNothing     MapLabelTypeImage = "4"
	MapLabelTypeImageCustom      MapLabelTypeImage = "5"

	MapLabelTypeMapLabel       MapLabelTypeMap = "0"
	MapLabelTypeMapElementName MapLabelTypeMap = "2"
	MapLabelTypeMapStatus      MapLabelTypeMap = "3"
	MapLabelTypeMapNothing     MapLabelTypeMap = "4"
	MapLabelTypeMapCustom      MapLabelTypeMap = "5"

	MapLabelTypeTriggerLabel       MapLabelTypeTrigger = "0"
	MapLabelTypeTriggerElementName MapLabelTypeTrigger = "2"
	MapLabelTypeTriggerStatus      MapLabelTypeTrigger = "3"
	MapLabelTypeTriggerNothing     MapLabelTypeTrigger = "4"
	MapLabelTypeTriggerCustom      MapLabelTypeTrigger = "5"

	MapShowAllProblems MapProblemUnAck = "0"
	MapShowUnAck       MapProblemUnAck = "1"
	MapShowSplit       MapProblemUnAck = "2"

	MapPublic  MapSharing = "0"
	MapPrivate MapSharing = "1"

	MapProblemHide MapProblemSuppressed = "0"
	MapProblemShow MapProblemSuppressed = "1"

	MapHost      MapElementType = "0"
	MapSubMap    MapElementType = "1"
	MapTrigger   MapElementType = "2"
	MapHostGroup MapElementType = "3"
	MapImage     MapElementType = "4"

	MapFullSize  MapAreaType = "0"
	MapFixedSize MapAreaType = "1"

	MapSingleElement    MapHostGroupDisplay = "0"
	MapSeparateEachHost MapHostGroupDisplay = "1"

	// Only for map element.
	MapNone      MapPermission = "-1"
	MapReadOnly  MapPermission = "2"
	MapReadWrite MapPermission = "3"

	MapMappingOff MapElementIconMap = "0"
	MapMappingOn  MapElementIconMap = "1"

	MapViewTypeGrid MapElementViewType = "0"

	MapOperatorContains    MapOperator = "0"
	MapOperatorEquals      MapOperator = "1"
	MapOperatorNotContains MapOperator = "2"
	MapOperatorNotequal    MapOperator = "3"
	MapOperatorExists      MapOperator = "4"
	MapOperatorNotExists   MapOperator = "5"

	MapDrawLine       MapDrawType = "0"
	MapDrawBoldLine   MapDrawType = "2"
	MapDrawDottedLine MapDrawType = "3"
	MapDrawDashedLine MapDrawType = "4"

	MapRectangle MapShapeType = "0"
	MapEllipse   MapShapeType = "1"

	MapFontGeorgia           MapFont = "0"
	MapFontPalatinoLinotype  MapFont = "1"
	MapFontTimesNewRoman     MapFont = "2"
	MapFontArialHelvetica    MapFont = "3"
	MapFontArialBlack        MapFont = "4"
	MapFontComicSansMS       MapFont = "5"
	MapFontImpactCharcoal    MapFont = "6"
	MapFontLucidaSansUnicode MapFont = "7"
	MapFontTahomaGeneva      MapFont = "8"
	MapFontTrebuchetMS       MapFont = "9"
	MapFontVerdanaGeneva     MapFont = "10"
	MapFontCourierNew        MapFont = "11"
	MapFontLucidaConsole     MapFont = "12"

	MapCenter MapHorizontalAlignment = "0"
	MapLeft   MapHorizontalAlignment = "1"
	MapRight  MapHorizontalAlignment = "2"

	MapMiddle MapVerticalAlignment = "0"
	MapTop    MapVerticalAlignment = "1"
	MapBottom MapVerticalAlignment = "2"

	MapBorderNone MapBorderType = "0"
	MapBorderLine MapBorderType = "1"
	MapBorderDot  MapBorderType = "2"
	MapBorderDash MapBorderType = "3"
)

// Map properties.
type Map struct {
	// ReadOnly
	Id            string             `json:"sysmapid,omitempty"`
	Height        string             `json:"height"`
	Name          string             `json:"name"`
	Width         string             `json:"width"`
	BackgroundId  string             `json:"backgroundid,omitempty"`
	ExpandMacros  MapMacroExpanded   `json:"expand_macros,omitempty"`
	ExpandProblem MapProblemExpanded `json:"expandproblem,omitempty"`
	GridAlign     MapGridAlignment   `json:"grid_align,omitempty"`
	GridShow      MapGridShow        `json:"grid_show,omitempty"`
	// Supported values : 20, 40, 50, 75 and 100
	GridSize             string                `json:"grid_size,omitempty"`
	Highlight            MapHighlight          `json:"highlight,omitempty"`
	IconMapId            string                `json:"iconmapid,omitempty"`
	LabelFormat          MapLabelFormat        `json:"label_format,omitempty"`
	LabelLocation        MapLabelLocation      `json:"label_location,omitempty"`
	LabelStringHost      string                `json:"label_string_host,omitempty"`
	LabelStringHostGroup string                `json:"label_string_hostgroup,omitempty"`
	LabelStringImage     string                `json:"label_string_image,omitempty"`
	LabelStringMap       string                `json:"label_string_map,omitempty"`
	LabelStringTrigger   string                `json:"label_string_trigger,omitempty"`
	LabelType            MapLabelType          `json:"label_type,omitempty"`
	LabelTypeHost        MapLabelTypeHost      `json:"label_type_host,omitempty"`
	LabelTypeHostGroup   MapLabelTypeHostGroup `json:"label_type_hostgroup,omitempty"`
	LabelTypeImage       MapLabelTypeImage     `json:"label_type_image,omitempty"`
	LabelTypeMap         MapLabelTypeMap       `json:"label_type_map,omitempty"`
	LabelTypeTrigger     MapLabelTypeTrigger   `json:"label_type_trigger,omitempty"`
	MarkElements         MapHighlight          `json:"markelements,omitempty"`
	SeverityMin          TriggerSeverity       `json:"severity_min,omitempty"`
	ShowUnAck            MapProblemUnAck       `json:"show_unack,omitempty"`
	UserId               string                `json:"userid,omitempty"`
	Private              MapSharing            `json:"private,omitempty"`
	ShowSuppressed       MapProblemSuppressed  `json:"show_suppressed,omitempty"`
}

// type MapElementSubType struct {
// }

// MapElementHost define the properties of a Host element for a Map.
type MapElementHost struct {
	Id string `json:"hostid"`
}

// MapElementHostGroup define the properties of a HostGroup element for a Map.
type MapElementHostGroup struct {
	Id string `json:"groupid"`
}

// MapElementMap define the properties of a Map element for a Map.
type MapElementMap struct {
	Id string `json:"sysmapid"`
}

// MapElementTrigger define the properties of a Trigger element for a Map.
type MapElementTrigger struct {
	Id string `json:"triggerid"`
}

// MapElementTag define the properties of a Tag element for a Map.
// Some properties are read-only, which means they are only accessible after creation
// and should not be passed as arguments in other methods.
type MapElementTag struct {
	Tag       string      `json:"tag"`
	Operation MapOperator `json:"operator,omitempty"`
	Value     string      `json:"value,omitempty"`
}

// MapElementUrl define the properties of a Url element for a Map.
// Some properties are read-only, which means they are only accessible after creation
// and should not be passed as arguments in other methods.
type MapElementUrl struct {
	// ReadOnly
	Id         string `json:"sysmapurlid,omitempty"`
	Name       string `json:"name"`
	Url        string `json:"url"`
	SelementId string `json:"selementid,omitempty"`
}

// MapElement properties.
type MapElement struct {
	// ReadOnly
	Id                string              `json:"selementid,omitempty"`
	Elements          interface{}         `json:"elements"`
	ElementType       MapElementType      `json:"elementtype"`
	IconIdOff         string              `json:"iconid_off"`
	AreaType          MapAreaType         `json:"areatype,omitempty"`
	ElementSubType    MapHostGroupDisplay `json:"elementsubtype,omitempty"`
	EvalType          EvalType            `json:"evaltype,omitempty"`
	Height            string              `json:"height,omitempty"`
	IconIdDisabled    string              `json:"iconid_disabled,omitempty"`
	IconIdMaintenance string              `json:"iconid_maintenance,omitempty"`
	IconIdOn          string              `json:"iconid_on,omitempty"`
	Label             string              `json:"label,omitempty"`
	LabelLocation     MapLabelLocation    `json:"label_location,omitempty"`
	Permission        MapPermission       `json:"permission,omitempty"`
	// ReadOnly
	SysMapId   string             `json:"sysmapid,omitempty"`
	Urls       []*MapElementUrl   `json:"urls,omitempty"`
	UseIconMap MapElementIconMap  `json:"use_iconmap,omitempty"`
	ViewType   MapElementViewType `json:"viewtype,omitempty"`
	Width      string             `json:"width,omitempty"`
	X          string             `json:"x,omitempty"`
	Y          string             `json:"y,omitempty"`
	Tags       []*MapElementTag   `json:"tags,omitempty"`
}

// Map link trigger properties.
// Some properties are read-only, which means they are only accessible after creation
// and should not be passed as arguments in other methods.
type MapLinkTrigger struct {
	// ReadOnly
	Id        string      `json:"linktriggerid,omitempty"`
	TriggerId string      `json:"triggerid"`
	Color     string      `json:"color,omitempty"`
	DrawType  MapDrawType `json:"drawtype,omitempty"`
	LinkId    string      `json:"linkid,omitempty"`
}

// Map link properties.
// Some properties are read-only, which means they are only accessible after creation
// and should not be passed as arguments in other methods.
type MapLink struct {
	// ReadOnly
	Id           string            `json:"linkid,omitempty"`
	SelementId1  string            `json:"selementid1"`
	SelementId2  string            `json:"selementid2"`
	Color        string            `json:"color,omitempty"`
	DrawType     MapDrawType       `json:"drawtype,omitempty"`
	Label        string            `json:"label,omitempty"`
	LinkTriggers []*MapLinkTrigger `json:"linktriggers,omitempty"`
	Permission   MapPermission     `json:"permission,omitempty"`
	SysMapId     string            `json:"sysmapid,omitempty"`
}

// Map user properties
// Some properties are read-only, which means they are only accessible after creation
// and should not be passed as arguments in other methods.
type MapUser struct {
	// ReadOnly
	Id         string        `json:"sysmapuserid,omitempty"`
	UserId     string        `json:"userid"`
	Permission MapPermission `json:"permission"`
}

// Map user group properties
// Some properties are read-only, which means they are only accessible after creation
// and should not be passed as arguments in other methods.
type MapUserGroup struct {
	// ReadOnly
	Id         string        `json:"sysmapusrgrpid,omitempty"`
	UsrGrpId   string        `json:"usrgrpid"`
	Permission MapPermission `json:"permission"`
}

// Map shape properties
// Some properties are read-only, which means they are only accessible after creation
// and should not be passed as arguments in other methods.
type MapShape struct {
	// ReadOnly
	Id              string                 `json:"sysmap_shapeid,omitempty"`
	Type            MapShapeType           `json:"type"`
	X               string                 `json:"x,omitempty"`
	Y               string                 `json:"y,omitempty"`
	Width           string                 `json:"width,omitempty"`
	Height          string                 `json:"height,omitempty"`
	Text            string                 `json:"text,omitempty"`
	Font            MapFont                `json:"font,omitempty"`
	FontSize        string                 `json:"font_size,omitempty"`
	FontColor       string                 `json:"font_color,omitempty"`
	TextHAlign      MapHorizontalAlignment `json:"text_halign,omitempty"`
	TextVAlign      MapVerticalAlignment   `json:"text_valign,omitempty"`
	BorderType      MapBorderType          `json:"border_type,omitempty"`
	BorderWidth     string                 `json:"border_width,omitempty"`
	BorderColor     string                 `json:"border_color,omitempty"`
	BackgroundColor string                 `json:"background_color,omitempty"`
	ZIndex          string                 `json:"zindex,omitempty"`
}

// Map line properties
// Some properties are read-only, which means they are only accessible after creation
// and should not be passed as arguments in other methods.
type MapLine struct {
	// ReadOnly
	Id        string        `json:"sysmap_shapeid,omitempty"`
	X1        string        `json:"x1,omitempty"`
	Y1        string        `json:"y1,omitempty"`
	X2        string        `json:"x2,omitempty"`
	Y2        string        `json:"y2,omitempty"`
	LineType  MapBorderType `json:"line_type,omitempty"`
	LineWidth string        `json:"line_width,omitempty"`
	LineColor string        `json:"line_color,omitempty"`
	ZIndex    string        `json:"zindex,omitempty"`
}

// MapResponse define the server response format for Map methods.
type MapResponse struct {
	MapIds []string `json:"sysmapids"`
}

// MapCreateParameters define the properties needed to create a new Map.
// Properties using the 'omitempty' json parameters are optional
type MapCreateParameters struct {
	Map
	Links      []*MapLink       `json:"links,omitempty"`
	Elements   []*MapElement    `json:"selements,omitempty"`
	Urls       []*MapElementUrl `json:"urls,omitempty"`
	Users      []*MapUser       `json:"users,omitempty"`
	UserGroups []*MapUserGroup  `json:"userGroups,omitempty"`
	Shapes     []*MapShape      `json:"shapes,omitempty"`
	Lines      []*MapLine       `json:"lines,omitempty"`
}

// Create is used to create a new Map.
func (i *MapService) Create(p *MapCreateParameters) (*MapResponse, error) {
	req := i.Client.NewRequest("map.create", p)

	res, err := i.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := MapResponse{}
	err = i.Client.ConvertResponse(*res, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// Delete is used to delete one or multiples Maps.
func (i *MapService) Delete(ids []string) (*MapResponse, error) {
	req := i.Client.NewRequest("map.delete", ids)

	res, err := i.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := MapResponse{}
	err = i.Client.ConvertResponse(*res, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// MapGetParameters define the properties used to search Maps.
// Properties using the 'omitempty' json parameters are optional
type MapGetParameters struct {
	Sysmapids              []string    `json:"sysmapids,omitempty"`
	UserIds                []string    `json:"userids,omitempty"`
	ExpandUrls             bool        `json:"expandUrls,omitempty"`
	SelectIconMap          interface{} `json:"selectIconMap,omitempty"`
	SelectLinks            interface{} `json:"selectLinks,omitempty"`
	SelectSelements        interface{} `json:"selectSelements,omitempty"`
	SelectUrls             interface{} `json:"selectUrls,omitempty"`
	SelectUsers            interface{} `json:"selectUsers,omitempty"`
	SelectUserGroups       interface{} `json:"selectUserGroups,omitempty"`
	SelectShapes           interface{} `json:"selectShapes,omitempty"`
	SelectLines            interface{} `json:"selectLines,omitempty"`
	CountOutput            bool        `json:"countOutput,omitempty"`
	Editable               bool        `json:"editable,omitempty"`
	ExcludeSearch          bool        `json:"excludeSearch,omitempty"`
	Filter                 interface{} `json:"filter,omitempty"`
	Limit                  string      `json:"limit,omitempty"`
	Output                 interface{} `json:"output,omitempty"`
	PreserveKeys           bool        `json:"preservekeys,omitempty"`
	Search                 interface{} `json:"search,omitempty"`
	SearchByAny            bool        `json:"searchByAny,omitempty"`
	SearchWildcardsEnabled bool        `json:"searchWildcardsEnabled,omitempty"`
	Sortfield              []string    `json:"sortfield,omitempty"`
	SortOrder              []string    `json:"sortorder,omitempty"`
	StartSearch            bool        `json:"startSearch,omitempty"`
}

// MapGetResponse define the server response format for the Map Get method.
type MapGetResponse struct {
	Map
	Links      []*MapLink       `json:"links,omitempty"`
	Elements   []*MapElement    `json:"selements,omitempty"`
	Urls       []*MapElementUrl `json:"urls,omitempty"`
	Users      []*MapUser       `json:"users,omitempty"`
	UserGroups []*MapUserGroup  `json:"userGroups,omitempty"`
	Shapes     []*MapShape      `json:"shapes,omitempty"`
	Lines      []*MapLine       `json:"lines,omitempty"`
}

// Get is used to retrieve one or multiples Maps matching the given criteria(s).
func (i *MapService) Get(p *MapGetParameters) ([]*MapGetResponse, error) {
	req := i.Client.NewRequest("map.get", p)

	res, err := i.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := make([]*MapGetResponse, 0)
	err = i.Client.ConvertResponse(*res, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// IconMapUpdateParameters define the properties used to update IconMaps.
// Properties using the 'omitempty' json parameters are optional
type MapUpdateParameters struct {
	Id            string             `json:"sysmapid"`
	Height        string             `json:"height,omitempty"`
	Name          string             `json:"name,omitempty"`
	Width         string             `json:"width,omitempty"`
	BackgroundId  string             `json:"backgroundid,omitempty"`
	ExpandMacros  MapMacroExpanded   `json:"expand_macros,omitempty"`
	ExpandProblem MapProblemExpanded `json:"expandproblem,omitempty"`
	GridAlign     MapGridAlignment   `json:"grid_align,omitempty"`
	GridShow      MapGridShow        `json:"grid_show,omitempty"`
	// Supported values : 20, 40, 50, 75 and 100
	GridSize             string                `json:"grid_size,omitempty"`
	Highlight            MapHighlight          `json:"highlight,omitempty"`
	IconMapId            string                `json:"iconmapid,omitempty"`
	LabelFormat          MapLabelFormat        `json:"label_format,omitempty"`
	LabelLocation        MapLabelLocation      `json:"label_location,omitempty"`
	LabelStringHost      string                `json:"label_string_host,omitempty"`
	LabelStringHostGroup string                `json:"label_string_hostgroup,omitempty"`
	LabelStringImage     string                `json:"label_string_image,omitempty"`
	LabelStringMap       string                `json:"label_string_map,omitempty"`
	LabelStringTrigger   string                `json:"label_string_trigger,omitempty"`
	LabelType            MapLabelType          `json:"label_type,omitempty"`
	LabelTypeHost        MapLabelTypeHost      `json:"label_type_host,omitempty"`
	LabelTypeHostGroup   MapLabelTypeHostGroup `json:"label_type_hostgroup,omitempty"`
	LabelTypeImage       MapLabelTypeImage     `json:"label_type_image,omitempty"`
	LabelTypeMap         MapLabelTypeMap       `json:"label_type_map,omitempty"`
	LabelTypeTrigger     MapLabelTypeTrigger   `json:"label_type_trigger,omitempty"`
	MarkElements         MapHighlight          `json:"markelements,omitempty"`
	SeverityMin          TriggerSeverity       `json:"severity_min,omitempty"`
	ShowUnAck            MapProblemUnAck       `json:"show_unack,omitempty"`
	UserId               string                `json:"userid,omitempty"`
	Private              MapSharing            `json:"private,omitempty"`
	ShowSuppressed       MapProblemSuppressed  `json:"show_suppressed,omitempty"`
	Links                []*MapLink            `json:"links,omitempty"`
	Elements             []*MapElement         `json:"selements,omitempty"`
	Urls                 []*MapElementUrl      `json:"urls,omitempty"`
	Users                []*MapUser            `json:"users,omitempty"`
	UserGroups           []*MapUserGroup       `json:"userGroups,omitempty"`
	Shapes               []*MapShape           `json:"shapes,omitempty"`
	Lines                []*MapLine            `json:"lines,omitempty"`
}

// Update is used to update a Map.
func (i *MapService) Update(p *MapUpdateParameters) (*MapResponse, error) {
	req := i.Client.NewRequest("map.update", p)

	res, err := i.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := MapResponse{}
	err = i.Client.ConvertResponse(*res, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
