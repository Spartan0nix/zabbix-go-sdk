package zabbixgosdk

import (
	"testing"
)

const (
	interfaceName        = "test-interface"
	updatedInterfaceName = "test-interface-update"
)

var InterfaceId string
var TemplateHostId = "10084"

func TestHostInterfaceCreate(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	i, err := client.HostInterface.Create(&HostInterface{
		Hostid: TemplateHostId,
		Ip:     "127.0.0.2",
		Dns:    "localhost",
		Main:   NotDefault,
		Port:   "1234",
		Type:   Agent,
		Useip:  "0",
	})

	if err != nil {
		t.Fatalf("Error when creating hostInterface '%s' for host '%s'.\nReason : %v", interfaceName, TemplateHostId, err)
	}

	if i == nil {
		t.Fatalf("Create method should returned a list with the Id of the new host interfaces.\nAn empty list was returned.")
	}

	InterfaceId = i.InterfaceIds[0].String()
}

func TestHostInterfaceGet(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	i, err := client.HostInterface.Get(&HostInterfaceGetParameters{
		Hostids: []string{
			TemplateHostId,
		},
		Filter: map[string]string{
			"name": interfaceName,
		},
	})

	if err != nil {
		t.Fatalf("Error when getting hostInterface '%s' for host '%s'.\nReason : %v", interfaceName, TemplateHostId, err)
	}

	if i == nil {
		t.Fatalf("Get method should returned a list of interfaces matching the given criteria.\nAn empty list was returned.")
	}

	var id string
	for _, y := range i {
		if y.Ip == "127.0.0.2" {
			id = y.Interfaceid
		}
	}

	if id != InterfaceId {
		t.Fatalf("Wrong interface returned.\nExpected Id : %s\nId returned : %s", InterfaceId, i[0].Interfaceid)
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
				Hostid: TemplateHostId,
			},
		},
		Interfaces: []*HostInterfaceMassProperties{
			{
				Ip:   "127.0.0.3",
				Dns:  "zabbix-server",
				Main: Default,
				Port: "2345",
				Type: SNMP,
				Details: &HostInterfaceDetail{
					Version:   SNMPv2c,
					Community: "test",
				},
				Useip: "1",
			},
		},
	})

	if err != nil {
		t.Fatalf("Error when mass interface to host '%s'.\nReason : %v", TemplateHostId, err)
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

// ---------------------------------
// TODO
// ---------------------------------
// Wait for host.create support
// ---------------------------------
// func TestHostInterfaceReplace(t *testing.T) {
// 	client, err := NewTestingService()
// 	if err != nil {
// 		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
// 	}

// 	i, err := client.HostInterface.ReplaceHostInterfaces(&HostInterfaceReplaceParameters{
// 		Host: TemplateHostId,
// 		Interfaces: []*HostInterfaceMassProperties{
// 			{
// 				Ip:    "127.0.0.1",
// 				Dns:   "localhost",
// 				Main:  Default,
// 				Port:  "10050",
// 				Type:  Agent,
// 				Useip: "0",
// 			},
// 		},
// 	})

// 	if err != nil {
// 		t.Fatalf("Error when replacing interfaces on host '%s'.\nReason : %v", TemplateHostId, err)
// 	}

// 	if i == nil {
// 		t.Fatal("MassUpdate method should returned a list of the updated interfaces.\nAn empty list was returned.")
// 	}
// }

func TestHostInterfaceUpdate(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	i, err := client.HostInterface.Update(&HostInterfaceUpdateParameters{
		Interfaceid: InterfaceId,
		Ip:          "127.0.0.1",
	},
	)

	if err != nil {
		t.Fatalf("Error when updating interface '%s' for host '%s'.\nReason : %v", InterfaceId, TemplateHostId, err)
	}

	if i == nil {
		t.Fatal("Update method should returned a list of with the updated interface.\nAn empty list was returned.")
	}
}

func TestHostInterfaceDelete(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	i, err := client.HostInterface.Delete([]string{
		InterfaceId,
	})

	if err != nil {
		t.Fatalf("Error when deleting template '%s'.\nReason : %v", templateId, err)
	}

	if i == nil {
		t.Fatalf("Delete method should returned a list of with the updated interface.\nAn empty list was returned.")
	}
}
