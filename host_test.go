package zabbixgosdk

import "testing"

const (
	hostName        = "test-host"
	updatedHostName = "test-host-update"
)

var hostId string

func TestHostCreate(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	h, err := client.Host.Create(&HostCreateParameters{
		Host:          hostName,
		Name:          hostName,
		Description:   "Testing description",
		InventoryMode: HostManual,
		Status:        HostUnmonitored,
		Groups: []*HostGroupId{
			{
				GroupId: "1",
			},
		},
		Interfaces: []*HostInterface{
			{
				Ip:    "127.0.0.1",
				Dns:   "localhost",
				Main:  HostInterfaceDefault,
				Port:  "10050",
				Type:  HostInterfaceAgent,
				UseIp: HostInterfaceDns,
			},
		},
		Tags: []*HostTag{
			{
				Tag:   "{$TEST_TAG}",
				Value: "test-host-tag-value",
			},
		},
		Inventory: map[HostInventory]string{
			Name:  hostName,
			Alias: "testing-host-alias",
		},
	})

	if err != nil {
		t.Fatalf("Error when creating host '%s'.\nReason : %v", hostName, err)
	}

	if h == nil {
		t.Fatalf("Create method should returned a list with the id of the new hosts.\nAn empty list was returned.")
	}

	hostId = h.HostIds[0]
}

func TestHostGet(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	h, err := client.Host.Get(&HostGetParameters{
		Filter: map[string]string{
			"name": hostName,
		},
	})

	if err != nil {
		t.Fatalf("Error when getting host '%s'.\nReason : %v", hostName, err)
	}

	if h == nil {
		t.Fatalf("Get method should returned a list of hosts matching the given criteria.\nAn empty list was returned.")
	}

	if h[0].HostId != hostId {
		t.Fatalf("Wrong host returned.\nExpected Id : %s\nId returned : %s", hostId, h[0].HostId)
	}
}

func TestHostMassAdd(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	h, err := client.Host.MassAdd(&HostMassAddParameters{
		Hosts: []*HostId{
			{
				HostId: hostId,
			},
		},
		Interfaces: []*HostInterface{
			{
				Ip:    "127.0.0.2",
				Dns:   "",
				Main:  HostInterfaceNotDefault,
				Port:  "10893",
				Type:  HostInterfaceAgent,
				UseIp: HostInterfaceIp,
			},
		},
		Macros: []*HostMacro{
			{
				Macro: "{$TEST_MACRO}",
				Value: "test-host-macro-value",
				Type:  Text,
			},
		},
	})

	if err != nil {
		t.Fatalf("Error when mass adding host '%s'.\nReason : %v", hostName, err)
	}

	if h == nil {
		t.Fatalf("MassAdd method should returned a list of the updated hosts.\nAn empty list was returned.")
	}

	if h.HostIds[0] != hostId {
		t.Fatalf("Wrong host returned.\nExpected Id : %s\nId returned : %s", hostId, h.HostIds[0])
	}
}

func TestHostMassRemove(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	m, err := client.UserMacro.Get(&UserMacroGetParameters{
		HostIds: []string{
			hostId,
		},
	})

	if err != nil {
		t.Fatalf("Error when retrieving macro for host '%s'.\nReason : %v", hostName, err)
	}

	h, err := client.Host.MassRemove(&HostMassRemoveParameters{
		HostIds: []string{
			hostId,
		},
		Interfaces: []*HostInterface{
			{
				Ip:   "127.0.0.2",
				Dns:  "",
				Port: "10893",
			},
		},
		Macros: []string{
			m[0].HostId,
		},
	})

	if err != nil {
		t.Fatalf("Error when mass removing host '%s'.\nReason : %v", hostName, err)
	}

	if h == nil {
		t.Fatalf("MassRemove method should returned a list of the updated hosts.\nAn empty list was returned.")
	}

	if h.HostIds[0] != hostId {
		t.Fatalf("Wrong host returned.\nExpected Id : %s\nId returned : %s", hostId, h.HostIds[0])
	}
}

func TestHostMassUpdate(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	h, err := client.Host.MassUpdate(&HostMassUpdateParameters{
		Hosts: []*HostId{
			{
				HostId: hostId,
			},
		},
		Status: HostUnmonitored,
	})

	if err != nil {
		t.Fatalf("Error when mass updating host '%s'.\nReason : %v", hostName, err)
	}

	if h == nil {
		t.Fatalf("MassUpdate method should returned a list of the updated hosts.\nAn empty list was returned.")
	}

	if h.HostIds[0] != hostId {
		t.Fatalf("Wrong host returned.\nExpected Id : %s\nId returned : %s", hostId, h.HostIds[0])
	}
}

func TestHostUpdate(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	h, err := client.Host.Update(&HostUpdateParameters{
		HostId:        hostId,
		Status:        HostMonitored,
		Host:          updatedHostName,
		Name:          updatedHostName,
		InventoryMode: HostManual,
		Inventory: map[HostInventory]string{
			Alias:   updatedHostName,
			Contact: "random@localhost",
		},
	})

	if err != nil {
		t.Fatalf("Error when updating host '%s'.\nReason : %v", hostName, err)
	}

	if h == nil {
		t.Fatalf("Update method should returned a list of the updated hosts.\nAn empty list was returned.")
	}

	if h.HostIds[0] != hostId {
		t.Fatalf("Wrong host returned.\nExpected Id : %s\nId returned : %s", hostId, h.HostIds[0])
	}
}

func TestHostDelete(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	h, err := client.Host.Delete([]string{
		hostId,
	})

	if err != nil {
		t.Fatalf("Error when deleting host '%s'.\nReason : %v", hostName, err)
	}

	if h == nil {
		t.Fatalf("Delete method should returned a list with the id of the deleted host.\nAn empty list was returned.")
	}

	if h.HostIds[0] != hostId {
		t.Fatalf("Wrong host returned.\nExpected Id : %s\nId returned : %s", hostId, h.HostIds[0])
	}
}
