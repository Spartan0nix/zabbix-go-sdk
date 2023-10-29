package zabbixgosdk

// import "testing"

// const (
// 	mapName        = "test-map"
// 	updatedMapName = "test-map-update"
// )

// var mapId string

// func TestMapCreate(t *testing.T) {
// 	m, err := testingClient.Map.Create(&MapCreateParameters{
// 		Map: Map{
// 			Name:                 mapName,
// 			Height:               "800",
// 			Width:                "800",
// 			ExpandMacros:         MapExpand,
// 			ExpandProblem:        MapIfOneProblem,
// 			GridShow:             MapGridHide,
// 			GridAlign:            MapAlignmentDisable,
// 			Highlight:            MapHighlightEnable,
// 			LabelFormat:          MapLabelAdvanced,
// 			LabelLocation:        MapLabelLocationRight,
// 			LabelType:            MapLabelTypeElementName,
// 			LabelTypeHost:        MapLabelTypeHostCustom,
// 			LabelStringHost:      "test-label-host",
// 			LabelTypeHostGroup:   MapLabelTypeHostGroupCustom,
// 			LabelStringHostGroup: "test-label-hostgroup",
// 			LabelTypeImage:       MapLabelTypeImageCustom,
// 			LabelStringImage:     "test-label-image",
// 			LabelTypeMap:         MapLabelTypeMapCustom,
// 			LabelStringMap:       "test-label-map",
// 			LabelTypeTrigger:     MapLabelTypeTriggerCustom,
// 			LabelStringTrigger:   "test-label-trigger",
// 			MarkElements:         MapHighlightEnable,
// 			SeverityMin:          TriggerInfo,
// 			ShowUnAck:            MapShowAllProblems,
// 			Private:              MapPublic,
// 			ShowSuppressed:       MapProblemHide,
// 		},
// 		Elements: []*MapElement{
// 			{
// 				Id: "1",
// 				Elements: []MapElementHostGroup{
// 					{
// 						Id: "1",
// 					},
// 				},
// 				ElementType: MapHostGroup,
// 				IconIdOff:   "2",
// 			},
// 			{
// 				Id: "2",
// 				Elements: []MapElementHostGroup{
// 					{
// 						Id: "2",
// 					},
// 				},
// 				ElementType: MapHostGroup,
// 				IconIdOff:   "2",
// 			},
// 		},
// 		Links: []*MapLink{
// 			{
// 				SelementId1: "1",
// 				SelementId2: "2",
// 			},
// 		},
// 		Shapes: []*MapShape{
// 			{
// 				Type:   MapRectangle,
// 				X:      "100",
// 				Y:      "100",
// 				Width:  "100",
// 				Height: "10",
// 				Text:   "test-line",
// 			},
// 		},
// 	})

// 	if err != nil {
// 		t.Fatalf("Error when creating map '%s'.\nReason: %v", mapName, err)
// 	}

// 	if m == nil || len(m.MapIds) == 0 {
// 		t.Fatalf("Create method should return a list of the created maps.\nAn empty list was returned.")
// 	}

// 	mapId = m.MapIds[0]
// }

// func TestMapGet(t *testing.T) {
// 	m, err := testingClient.Map.Get(&MapGetParameters{
// 		Sysmapids: []string{
// 			mapId,
// 		},
// 	})

// 	if err != nil {
// 		t.Fatalf("Error when getting map '%s'.\nReason: %v", mapId, err)
// 	}

// 	if len(m) == 0 {
// 		t.Fatalf("Get method should return a list of maps matching the given criteria.\nAn empty list was returned.")
// 	}

// 	if m[0].Id != mapId {
// 		t.Fatalf("Wrong map returned.\nExpected : %s\nReturned : %s", mapId, m[0].Id)
// 	}
// }

// func TestMapUpdate(t *testing.T) {
// 	m, err := testingClient.Map.Update(&MapUpdateParameters{
// 		Id:                 mapId,
// 		Name:               updatedMapName,
// 		Height:             "600",
// 		Width:              "600",
// 		ExpandMacros:       MapDoNotExpand,
// 		ExpandProblem:      MapAlwaysShow,
// 		GridShow:           MapGridDisplay,
// 		GridAlign:          MapAlignmentEnable,
// 		Highlight:          MapHighlightDisable,
// 		LabelFormat:        MapLabelNormal,
// 		LabelLocation:      MapLabelLocationTop,
// 		LabelType:          MapLabelTypeStatus,
// 		LabelTypeHost:      MapLabelTypeHostIP,
// 		LabelTypeHostGroup: MapLabelTypeHostGroupElementName,
// 		LabelTypeImage:     MapLabelTypeImageNothing,
// 		LabelTypeMap:       MapLabelTypeMapStatus,
// 		LabelTypeTrigger:   MapLabelTypeTriggerStatus,
// 		MarkElements:       MapHighlightDisable,
// 		SeverityMin:        TriggerAverage,
// 		ShowUnAck:          MapShowUnAck,
// 		ShowSuppressed:     MapProblemShow,
// 	})

// 	if err != nil {
// 		t.Fatalf("Error when updating map '%s'.\nReason: %v", mapId, err)
// 	}

// 	if m == nil || len(m.MapIds) == 0 {
// 		t.Fatalf("Update method should return a list of the updated maps.\nAn empty list was returned.")
// 	}

// 	if m.MapIds[0] != mapId {
// 		t.Fatalf("Wrong map returned.\nExpected : %s\nReturned : %s", mapId, m.MapIds[0])
// 	}
// }

// func TestMapDelete(t *testing.T) {
// 	m, err := testingClient.Map.Delete([]string{
// 		mapId,
// 	})

// 	if err != nil {
// 		t.Fatalf("Error when deleting map '%s'.\nReason: %v", mapId, err)
// 	}

// 	if m == nil || len(m.MapIds) == 0 {
// 		t.Fatalf("Delete method should return a list with the id of the deleted maps.\nAn empty list was returned.")
// 	}

// 	if m.MapIds[0] != mapId {
// 		t.Fatalf("Wrong id returned.\nExpected : %s\nReturned : %s", mapId, m.MapIds[0])
// 	}
// }
