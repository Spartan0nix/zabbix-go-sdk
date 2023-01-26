package zabbixgosdk

import (
	"testing"
)

const (
	hostGroupName        = "test-host-group"
	updatedHostGroupName = "test-host-group-update"
)

var hostGroupId string

func TestHostGroupCreate(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	g, err := client.HostGroup.Create(hostGroupName)

	if err != nil {
		t.Fatalf("Error when creating host group '%s'.\nReason : %v", hostGroupName, err)
	}

	if g.GroupIds == nil {
		t.Fatalf("Create method should returned a list with the id of the new host groups.\nAn empty list was returned.")
	}

	hostGroupId = g.GroupIds[0]
}

func TestHostGroupList(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	g, err := client.HostGroup.List()

	if err != nil {
		t.Fatalf("Error when listing host groups.\nReason : %v", err)
	}

	if g == nil {
		t.Fatalf("List method should returned a list of existing host group.\nAn empty list was returned.")
	}
}

func TestHostGroupGet(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	g, err := client.HostGroup.Get(&HostGroupGetParameters{
		CommonGetParameters: CommonGetParameters{
			Filter: map[string]string{
				"name": hostGroupName,
			},
		},
	})

	if err != nil {
		t.Fatalf("Error when getting host group '%s'.\nReason : %v", hostGroupName, err)
	}

	if g == nil {
		t.Fatalf("Get method should returned a list of host groups matching the given criteria.\nAn empty list was returned.")
	}

	if g[0].Id != hostGroupId {
		t.Fatalf("Wrong host group returned.\nExpected Id : %s\nId returned : %s", hostGroupId, g[0].Id)
	}
}

func TestHostGroupMassAdd(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	g, err := client.HostGroup.MassAdd(&HostGroupMassAddParameters{
		Groups: []*HostGroupId{
			{
				GroupId: hostGroupId,
			},
		},
		Hosts: []*HostId{
			{
				HostId: "10084",
			},
		},
	})

	if err != nil {
		t.Fatalf("Error when mass adding host group.\nHost group : %s\nHost : %s\nReason : %v", hostGroupId, "10084", err)
	}

	if g == nil {
		t.Fatal("MassAdd method should returned a list of the updated host groups.\nAn empty list was returned.")
	}

	if g.GroupIds[0] != hostGroupId {
		t.Fatalf("Wrong host group returned.\nExpected Id : %s\nId returned : %s", hostGroupId, g.GroupIds[0])
	}
}

func TestHostGroupMassRemove(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	g, err := client.HostGroup.MassRemove(&HostGroupMassRemoveParameters{
		GroupsIds: []string{
			hostGroupId,
		},
		HostIds: []string{
			"10084",
		},
	})

	if err != nil {
		t.Fatalf("Error when mass removing host group.\nHost group : %s\nHosts : %s \nReason : %v", hostGroupId, "10084", err)
	}

	if g == nil {
		t.Fatal("MassRemove method should returned a list of the updated host groups.\nAn empty list was returned.")
	}

	if g.GroupIds[0] != hostGroupId {
		t.Fatalf("Wrong host group returned.\nExpected Id : %s\nId returned : %s", hostGroupId, g.GroupIds[0])
	}
}

func TestHostGroupMassUpdate(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	if err != nil {
		t.Fatalf("Error when getting host group 'Templates'.\nReason : %v", err)
	}

	template, err := client.Template.Get(&TemplateGetParameters{
		Filter: map[string]string{
			"name": "Zabbix server health",
		},
	})

	if err != nil {
		t.Fatalf("Error when retrieving template 'Zabbix server health'.\nReason : %v", err)
	}

	g, err := client.HostGroup.MassUpdate(&HostGroupMassUpdateParameters{
		Groups: []*HostGroupId{
			{
				GroupId: hostGroupId,
			},
		},
		Hosts: []*HostId{
			{
				HostId: "10084",
			},
		},
		Templates: []*TemplateId{
			{
				Id: template[0].TemplateId,
			},
		},
	})

	if err != nil {
		t.Fatalf("Error when mass updating host group.\nHost group : %s\nHosts : %s\nTemplate : %s\nReason : %v", hostGroupId, "10084", template[0].TemplateId, err)
	}

	if g == nil {
		t.Fatal("MassUpdate method should returned a list of the updated host group(s).\nAn empty list was returned.")
	}

	if g.GroupIds[0] != hostGroupId {
		t.Fatalf("Wrong host group returned.\nExpected Id : %s\nId returned : %s", hostGroupId, g.GroupIds[0])
	}
}

func TestHostGroupUpdate(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	g, err := client.HostGroup.Update(hostGroupId, updatedHostGroupName)
	if err != nil {
		t.Fatalf("Error when updating host group '%s'.\nReason : %v", hostGroupName, err)
	}

	if g == nil {
		t.Fatal("Update method should returned a list of the updated host group.\nAn empty list was returned.")
	}

	if g.GroupIds[0] != hostGroupId {
		t.Fatalf("Wrong host group returned.\nExpected Id : %s\nId returned : %s", hostGroupId, g.GroupIds[0])
	}
}

func TestHostGroupDelete(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	g, err := client.HostGroup.Delete([]string{
		hostGroupId,
	})

	if err != nil {
		t.Fatalf("Error when deleting host group '%s'.\nReason : %v", hostGroupName, err)
	}

	if g.GroupIds == nil {
		t.Fatalf("Delete method should returned a list with the id of the deleted host groups.\nAn empty list was returned.")
	}

	if g.GroupIds[0] != hostGroupId {
		t.Fatalf("Wrong host group returned.\nExpected Id : %s\nId returned : %s", hostGroupId, g.GroupIds[0])
	}
}
