package zabbixgosdk

// import "testing"

// const (
// 	iconMapName        = "test-iconmap"
// 	updatedIconMapName = "test-iconmap-update"
// )

// var iconMapId string

// func TestIconMapCreate(t *testing.T) {
// 	i, err := testingClient.IconMap.Create(&IconMapCreateParameters{
// 		Name:          iconMapName,
// 		DefaultIconId: "2",
// 		Mappings: []IconMappingParameters{
// 			{
// 				IconId:        "1",
// 				Expression:    "main-site",
// 				InventoryLink: GetHostInventoryId(Location),
// 			},
// 		},
// 	})

// 	if err != nil {
// 		t.Fatalf("Error when creating iconmap '%s'.\nReason: %v", iconMapName, err)
// 	}

// 	if i.IconMapIds == nil || len(i.IconMapIds) == 0 {
// 		t.Fatalf("Create method should return a list of the created iconmaps.\nAn empty list was returned.")
// 	}

// 	iconMapId = i.IconMapIds[0]
// }

// func TestIconMapGet(t *testing.T) {
// 	i, err := testingClient.IconMap.Get(&IconMapGetParameters{
// 		Iconmapids: []string{
// 			iconMapId,
// 		},
// 	})

// 	if err != nil {
// 		t.Fatalf("Error when getting iconmap '%s'.\nReason: %v", iconMapId, err)
// 	}

// 	if len(i) == 0 {
// 		t.Fatalf("Get method should return a list of iconmaps matching the given criteria.\nAn empty list was returned.")
// 	}

// 	if i[0].IconMapId != iconMapId {
// 		t.Fatalf("Wrong iconmap returned.\nExpected : %s\nReturned : %s", iconMapId, i[0].IconMapId)
// 	}
// }

// func TestIconMapUpdate(t *testing.T) {
// 	i, err := testingClient.IconMap.Update(&IconMapUpdateParameters{
// 		Id:   iconMapId,
// 		Name: updatedIconMapName,
// 		Mappings: []IconMappingParameters{
// 			{
// 				IconId:        "2",
// 				Expression:    "updated-field",
// 				InventoryLink: GetHostInventoryId(Alias),
// 			},
// 		},
// 	})

// 	if err != nil {
// 		t.Fatalf("Error when updating iconmap '%s'.\nReason: %v", iconMapId, err)
// 	}

// 	if i == nil || len(i.IconMapIds) == 0 {
// 		t.Fatalf("Update method should return a list of  the updated iconmaps.\nAn empty list was returned.")
// 	}

// 	if i.IconMapIds[0] != iconMapId {
// 		t.Fatalf("Wrong iconmap returned.\nExpected : %s\nReturned : %s", iconMapId, i.IconMapIds[0])
// 	}
// }

// func TestIconMapDelete(t *testing.T) {
// 	i, err := testingClient.IconMap.Delete([]string{
// 		iconMapId,
// 	})

// 	if err != nil {
// 		t.Fatalf("Error when deleting iconmap '%s'.\nReason: %v", iconMapId, err)
// 	}

// 	if i.IconMapIds == nil || len(i.IconMapIds) == 0 {
// 		t.Fatalf("Delete method should return a list with the id of the deleted iconmaps.\nAn empty list was returned.")
// 	}

// 	if i.IconMapIds[0] != iconMapId {
// 		t.Fatalf("Wrong id returned.\nExpected : %s\nReturned : %s", iconMapId, i.IconMapIds[0])
// 	}
// }
