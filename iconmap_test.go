package zabbixgosdk

import "testing"

const (
	iconMapName        = "test-iconmap"
	updatedIconMapName = "test-iconmap-update"
)

var iconMapId string

func TestIconMapCreate(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	i, err := client.IconMap.Create(&IconMapCreateParameters{
		Name:          iconMapName,
		DefaultIconId: "2",
		Mappings: []IconMappingParameters{
			{
				IconId:        "1",
				Expression:    "main-site",
				InventoryLink: GetHostInventoryId(Location),
			},
		},
	})

	if err != nil {
		t.Fatalf("Error when creating icon map '%s'.\nReason : %v", iconMapName, err)
	}

	if i.IconMapIds == nil {
		t.Fatalf("Create method should returned a list with the id of the new icon mappings.\nAn empty list was returned.")
	}

	iconMapId = i.IconMapIds[0]
}

func TestIconMapGet(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	i, err := client.IconMap.Get(&IconMapGetParameters{
		Iconmapids: []string{
			iconMapId,
		},
	})

	if err != nil {
		t.Fatalf("Error when getting icon map '%s'.\nReason : %v", iconMapId, err)
	}

	if len(i) == 0 {
		t.Fatalf("Get method should returned a list of icon map matching the given criteria.\nAn empty list was returned.")
	}

	if i[0].IconMapId != iconMapId {
		t.Fatalf("Wrong icon map returned.\nExpected : %s\nReturned : %s", iconMapId, i[0].IconMapId)
	}
}

func TestIconMapUpdate(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	i, err := client.IconMap.Update(&IconMapUpdateParameters{
		Id:   iconMapId,
		Name: updatedIconMapName,
		Mappings: []IconMappingParameters{
			{
				IconId:        "2",
				Expression:    "updated-field",
				InventoryLink: GetHostInventoryId(Alias),
			},
		},
	})

	if err != nil {
		t.Fatalf("Error when updating icon map '%s'.\nReason : %v", iconMapId, err)
	}

	if i == nil {
		t.Fatalf("Update method should returned a list of update icon map.\nAn empty list was returned.")
	}

	if i.IconMapIds[0] != iconMapId {
		t.Fatalf("Wrong icon map returned.\nExpected : %s\nReturned : %s", iconMapId, i.IconMapIds[0])
	}
}

func TestIconMapDelete(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	i, err := client.IconMap.Delete([]string{
		iconMapId,
	})

	if err != nil {
		t.Fatalf("Error when deleting icon map '%s'.\nReason : %v", iconMapId, err)
	}

	if i.IconMapIds == nil {
		t.Fatalf("Delete method should returned a list with the id of the deleted icon mappings.\nAn empty list was returned.")
	}

	if i.IconMapIds[0] != iconMapId {
		t.Fatalf("Wrong id returned.\nExpected : %s\nReturned : %s", iconMapId, i.IconMapIds[0])
	}
}
