package zabbixgosdk

import (
	"testing"
)

const (
	interfacehostId   = "10084"
	interfaceHostName = "test-host-interface"
)

var interfaceId string

func TestHostInterfaceCreate(t *testing.T) {
	i, err := testingClient.HostInterface.Create(&HostInterfaceCreateParameters{
		HostId: interfacehostId,
		Ip:     "127.0.0.2",
		Dns:    "localhost",
		Main:   HostInterfaceNotDefault,
		Port:   "1234",
		Type:   HostInterfaceAgent,
		UseIp:  HostInterfaceDns,
	})

	if err != nil {
		t.Fatalf("Error when creating hostInterface for host '%s'.\nReason : %v", interfacehostId, err)
	}

	if i == nil || len(i.InterfaceIds) == 0 {
		t.Fatalf("Create method should return a list of the created hostinterfaces.\nAn empty list was returned.")
	}

	interfaceId = i.InterfaceIds[0].String()
}

func TestHostInterfaceGet(t *testing.T) {
	i, err := testingClient.HostInterface.Get(&HostInterfaceGetParameters{
		HostIds: []string{
			interfacehostId,
		},
	})

	if err != nil {
		t.Fatalf("Error when getting hostInterface for host '%s'.\nReason : %v", interfacehostId, err)
	}

	if len(i) == 0 {
		t.Fatalf("Get method should return a list of hostinterfaces matching the given criteria.\nAn empty list was returned.")
	}

	var id string
	for _, y := range i {
		if y.Ip == "127.0.0.2" {
			id = y.InterfaceId
		}
	}

	if id != interfaceId {
		t.Fatalf("Wrong interface returned.\nExpected : %s\nReturned : %s", interfaceId, i[0].InterfaceId)
	}
}

func TestHostInterfaceMassAdd(t *testing.T) {
	i, err := testingClient.HostInterface.MassAdd(&HostInterfaceMassAddParameters{
		Hosts: []*HostId{
			{
				HostId: interfacehostId,
			},
		},
		Interfaces: []*HostInterfaceMassProperties{
			{
				Ip:   "127.0.0.3",
				Dns:  "zabbix-server",
				Main: HostInterfaceDefault,
				Port: "2345",
				Type: HostInterfaceSNMP,
				Details: &HostInterfaceDetail{
					Version:   HostInterfaceSNMPv2c,
					Community: "test",
				},
				UseIp: HostInterfaceIp,
			},
		},
	})

	if err != nil {
		t.Fatalf("Error when massadding interfaces to host '%s'.\nReason : %v", interfacehostId, err)
	}

	if i == nil || len(i.InterfaceIds) == 0 {
		t.Fatal("MassAdd method should return a list of the updated interfaces.\nAn empty list was returned.")
	}
}

func TestHostInterfaceMassRemove(t *testing.T) {
	i, err := testingClient.HostInterface.MassRemove(&HostInterfaceMassRemoveParameters{
		HostIds: []string{
			interfacehostId,
		},
		Interfaces: []*HostInterfaceMassProperties{
			{
				Ip:   "127.0.0.3",
				Dns:  "zabbix-server",
				Port: "2345",
			},
		},
	})

	if err != nil {
		t.Fatalf("Error when massremoving interfaces from host '%s'.\nReason : %v", interfacehostId, err)
	}

	if i == nil || len(i.InterfaceIds) == 0 {
		t.Fatal("MassRemove method should return a list of the updated interfaces.\nAn empty list was returned.")
	}
}

func TestHostInterfaceReplace(t *testing.T) {
	h, err := testingClient.Host.Create(&HostCreateParameters{
		Host:   interfaceHostName,
		Status: HostUnmonitored,
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
	})

	if err != nil {
		t.Fatalf("Error when creating host '%s'.\nReason : %v", interfaceHostName, err)
	}

	if h == nil || len(h.HostIds) == 0 {
		t.Fatalf("Error when creating host '%s'.\nAn empty list was of hosts returned.", interfaceHostName)
	}

	i, err := testingClient.HostInterface.ReplaceHostInterfaces(&HostInterfaceReplaceParameters{
		Host: h.HostIds[0],
		Interfaces: []*HostInterfaceMassProperties{
			{
				Ip:    "127.0.0.1",
				Dns:   "localhost",
				Main:  HostInterfaceDefault,
				Port:  "161",
				Type:  HostInterfaceSNMP,
				UseIp: HostInterfaceIp,
				Details: &HostInterfaceDetail{
					Version:   HostInterfaceSNMPv2c,
					Community: "public",
					Bulk:      HostInterfaceBulkTrue,
				},
			},
		},
	})

	if err != nil {
		t.Fatalf("Error when replacing interfaces on host '%s'.\nReason : %v", interfacehostId, err)
	}

	if i == nil || len(i.InterfaceIds) == 0 {
		t.Fatal("MassUpdate method should return a list of the updated interfaces.\nAn empty list was returned.")
	}
}

func TestHostInterfaceUpdate(t *testing.T) {
	i, err := testingClient.HostInterface.Update(&HostInterfaceUpdateParameters{
		InterfaceId: interfaceId,
		Ip:          "127.0.0.4",
	})

	if err != nil {
		t.Fatalf("Error when updating interface '%s'.\nReason : %v", interfaceId, err)
	}

	if i == nil || len(i.InterfaceIds) == 0 {
		t.Fatal("Update method should return a list of the updated interfaces.\nAn empty list was returned.")
	}

	if string(i.InterfaceIds[0]) != interfaceId {
		t.Fatalf("Wrong interface returned.\nExpected : %s\nReturned : %s", interfaceId, string(i.InterfaceIds[0]))
	}
}

func TestHostInterfaceDelete(t *testing.T) {
	i, err := testingClient.HostInterface.Delete([]string{
		interfaceId,
	})

	if err != nil {
		t.Fatalf("Error when deleting interface '%s'.\nReason : %v", interfaceId, err)
	}

	if i == nil || len(i.InterfaceIds) == 0 {
		t.Fatalf("Delete method should return a list of the deleted interfaces.\nAn empty list was returned.")
	}

	if string(i.InterfaceIds[0]) != interfaceId {
		t.Fatalf("Wrong interface returned.\nExpected : %s\nReturned : %s", interfaceId, string(i.InterfaceIds[0]))
	}
}
