package zabbixgosdk

import (
	"testing"
)

const (
	TemplateHostId = "10084"
)

var interfaceId string

func TestHostInterfaceCreate(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	i, err := client.HostInterface.Create(&HostInterfaceCreateParameters{
		HostId: TemplateHostId,
		Ip:     "127.0.0.2",
		Dns:    "localhost",
		Main:   HostInterfaceNotDefault,
		Port:   "1234",
		Type:   HostInterfaceAgent,
		UseIp:  HostInterfaceDns,
	})

	if err != nil {
		t.Fatalf("Error when creating hostInterface for host '%s'.\nReason : %v", TemplateHostId, err)
	}

	if i == nil {
		t.Fatalf("Create method should returned a list with the id of the new host interfaces.\nAn empty list was returned.")
	}

	interfaceId = i.InterfaceIds[0].String()
}

func TestHostInterfaceGet(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	i, err := client.HostInterface.Get(&HostInterfaceGetParameters{
		HostIds: []string{
			TemplateHostId,
		},
	})

	if err != nil {
		t.Fatalf("Error when getting hostInterface for host '%s'.\nReason : %v", TemplateHostId, err)
	}

	if i == nil {
		t.Fatalf("Get method should returned a list of interfaces matching the given criteria.\nAn empty list was returned.")
	}

	var id string
	for _, y := range i {
		if y.Ip == "127.0.0.2" {
			id = y.InterfaceId
		}
	}

	if id != interfaceId {
		t.Fatalf("Wrong interface returned.\nExpected Id : %s\nId returned : %s", interfaceId, i[0].InterfaceId)
	}
}

func TestHostInterfaceMassAdd(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	i, err := client.HostInterface.MassAdd(&HostInterfaceMassAddParameters{
		Hosts: []*HostId{
			{
				HostId: TemplateHostId,
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
		t.Fatalf("Error when mass adding interface to host '%s'.\nReason : %v", TemplateHostId, err)
	}

	if i == nil {
		t.Fatal("MassAdd method should returned a list of the updated interfaces.\nAn empty list was returned.")
	}
}

func TestHostInterfaceMassRemove(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	i, err := client.HostInterface.MassRemove(&HostInterfaceMassRemoveParameters{
		HostIds: []string{
			TemplateHostId,
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
		t.Fatalf("Error when mass removing interface from host '%s'.\nReason : %v", TemplateHostId, err)
	}

	if i == nil {
		t.Fatal("MassRemove method should returned a list of the updated interfaces.\nAn empty list was returned.")
	}
}

func TestHostInterfaceReplace(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	h, err := client.Host.Create(&HostCreateParameters{
		Host:   "test-host-interface",
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
		t.Fatalf("Error when ading host 'test-host-interface'.\nReason : %v", err)
	}

	if h == nil {
		t.Fatalf("'host.create' method returned an empty list of host.")
	}

	i, err := client.HostInterface.ReplaceHostInterfaces(&HostInterfaceReplaceParameters{
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
		t.Fatalf("Error when replacing interfaces on host '%s'.\nReason : %v", TemplateHostId, err)
	}

	if i == nil {
		t.Fatal("MassUpdate method should returned a list of the updated interfaces.\nAn empty list was returned.")
	}
}

func TestHostInterfaceUpdate(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	i, err := client.HostInterface.Update(&HostInterfaceUpdateParameters{
		InterfaceId: interfaceId,
		Ip:          "127.0.0.4",
	},
	)

	if err != nil {
		t.Fatalf("Error when updating interface '%s' for host '%s'.\nReason : %v", interfaceId, TemplateHostId, err)
	}

	if i == nil {
		t.Fatal("Update method should returned a list of the updated interfaces.\nAn empty list was returned.")
	}

	if string(i.InterfaceIds[0]) != interfaceId {
		t.Fatalf("Wrong interface returned.\nExpected Id : %s\nId returned : %s", interfaceId, string(i.InterfaceIds[0]))
	}
}

func TestHostInterfaceDelete(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	i, err := client.HostInterface.Delete([]string{
		interfaceId,
	})

	if err != nil {
		t.Fatalf("Error when deleting interface '%s'.\nReason : %v", templateId, err)
	}

	if i == nil {
		t.Fatalf("Delete method should returned a list with the id of the deleted host interfaces.\nAn empty list was returned.")
	}

	if string(i.InterfaceIds[0]) != interfaceId {
		t.Fatalf("Wrong interface returned.\nExpected Id : %s\nId returned : %s", interfaceId, string(i.InterfaceIds[0]))
	}
}
