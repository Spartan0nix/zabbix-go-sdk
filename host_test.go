package main

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
		Host:           hostName,
		Name:           hostName,
		Description:    "Testing description",
		Inventory_mode: Host_Manual,
		Status:         Unmonitored,
		Groups: []*HostGroupId{
			{
				Groupid: "1",
			},
		},
		Interfaces: []*HostInterface{
			{
				Ip:    "127.0.0.1",
				Dns:   "localhost",
				Main:  Default,
				Port:  "10050",
				Type:  Agent,
				Useip: "0",
			},
		},
		Tags: []*HostTag{
			{
				Tag:   "{$TEST_TAG}",
				Value: "test-host-tag-value",
			},
		},
		Inventory: map[HostInventory]string{
			name:  hostName,
			alias: "testing-host-alias",
		},
	})

	if err != nil {
		t.Fatalf("Error when creating host '%s'.\nReason : %v", hostName, err)
	}

	if h == nil {
		t.Fatalf("Create method should returned a list with the Id of the new hosts.\nAn empty list was returned.")
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

	if h[0].Host != hostName {
		t.Fatalf("Wrong host returned.\nExpected name : %s\nName returned : %s", hostName, h[0].Host)
	}

	if h[0].Hostid != hostId {
		t.Fatalf("Wrong host returned.\nExpected Id : %s\nId returned : %s", hostId, h[0].Hostid)
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
				Hostid: hostId,
			},
		},
		Interfaces: []*HostInterface{
			{
				Ip:    "127.0.0.2",
				Dns:   "",
				Main:  NotDefault,
				Port:  "10893",
				Type:  Agent,
				Useip: "1",
			},
		},
		Macros: []*HostMacro{
			{
				Macro: "{$TEST_MACRO}",
				Value: "test-host-macro-value",
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
		Hostids: []string{
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
			m[0].Id,
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
				Hostid: hostId,
			},
		},
		Status: Unmonitored,
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
		Hostid:         hostId,
		Status:         Monitored,
		Host:           updatedHostName,
		Name:           updatedHostName,
		Inventory_mode: Host_Manual,
		Inventory: map[HostInventory]string{
			alias:   updatedHostName,
			contact: "random@localhost",
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
		t.Fatalf("Delete method should returned a list with the Ids of the deleted host.\nAn empty list was returned.")
	}

	if h.HostIds[0] != hostId {
		t.Fatalf("Wrong host returned.\nExpected Id : %s\nId returned : %s", hostId, h.HostIds[0])
	}
}
