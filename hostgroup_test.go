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
	g, err := testingClient.HostGroup.Create(hostGroupName)
	if err != nil {
		t.Fatalf("Error when creating hostgroup '%s'.\nReason : %v", hostGroupName, err)
	}

	if g == nil || len(g.GroupIds) == 0 {
		t.Fatalf("Create method should return a list of the created hostgroups.\nAn empty list was returned.")
	}

	hostGroupId = g.GroupIds[0]
}

func TestHostGroupList(t *testing.T) {
	g, err := testingClient.HostGroup.List()
	if err != nil {
		t.Fatalf("Error when listing hostgroups.\nReason : %v", err)
	}

	if len(g) == 0 {
		t.Fatalf("List method should return a list with all the existing hostgroups on the server.\nAn empty list was returned.")
	}
}

func TestHostGroupGet(t *testing.T) {
	g, err := testingClient.HostGroup.Get(&HostGroupGetParameters{
		GroupIds: []string{
			hostGroupId,
		},
	})

	if err != nil {
		t.Fatalf("Error when getting host group '%s'.\nReason : %v", hostGroupName, err)
	}

	if len(g) == 0 {
		t.Fatalf("Get method should return a list of hostgroups matching the given criteria.\nAn empty list was returned.")
	}

	if g[0].Id != hostGroupId {
		t.Fatalf("Wrong host group returned.\nExpected : %s\nReturned : %s", hostGroupId, g[0].Id)
	}
}

func TestHostGroupMassAdd(t *testing.T) {
	g, err := testingClient.HostGroup.MassAdd(&HostGroupMassAddParameters{
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
		t.Fatalf("Error when massadding hosts to the hostgroup '%s'.\nReason : %v", hostGroupId, err)
	}

	if g == nil || len(g.GroupIds) == 0 {
		t.Fatal("MassAdd method should return a list of the updated hostgroups.\nAn empty list was returned.")
	}

	if g.GroupIds[0] != hostGroupId {
		t.Fatalf("Wrong hostgroup returned.\nExpected : %s\nRreturned : %s", hostGroupId, g.GroupIds[0])
	}
}

func TestHostGroupMassRemove(t *testing.T) {
	g, err := testingClient.HostGroup.MassRemove(&HostGroupMassRemoveParameters{
		GroupsIds: []string{
			hostGroupId,
		},
		HostIds: []string{
			"10084",
		},
	})

	if err != nil {
		t.Fatalf("Error when massremoving hosts from the hostgroup '%s'.\nReason : %v", hostGroupId, err)
	}

	if g == nil || len(g.GroupIds) == 0 {
		t.Fatal("MassRemove method should return a list of the updated hostgroups.\nAn empty list was returned.")
	}

	if g.GroupIds[0] != hostGroupId {
		t.Fatalf("Wrong host group returned.\nExpected : %s\nReturned : %s", hostGroupId, g.GroupIds[0])
	}
}

func TestHostGroupMassUpdate(t *testing.T) {
	template, err := testingClient.Template.Get(&TemplateGetParameters{
		Filter: map[string]string{
			"name": "Zabbix server health",
		},
	})

	if err != nil {
		t.Fatalf("Error when getting template 'Zabbix server health'.\nReason : %v", err)
	}

	if len(template) == 0 {
		t.Fatalf("An empty list was returned when getting the template 'Zabbix server health'")
	}

	g, err := testingClient.HostGroup.MassUpdate(&HostGroupMassUpdateParameters{
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
		t.Fatalf("Error when mass updating hostgroup '%s'.\nReason : %v", hostGroupId, err)
	}

	if g == nil || len(g.GroupIds) == 0 {
		t.Fatal("MassUpdate method should return a list of the updated hostgroups.\nAn empty list was returned.")
	}

	if g.GroupIds[0] != hostGroupId {
		t.Fatalf("Wrong host group returned.\nExpected : %s\nReturned : %s", hostGroupId, g.GroupIds[0])
	}
}

func TestHostGroupUpdate(t *testing.T) {
	g, err := testingClient.HostGroup.Update(hostGroupId, updatedHostGroupName)
	if err != nil {
		t.Fatalf("Error when updating name of the hostgroup '%s'.\nReason : %v", hostGroupName, err)
	}

	if g == nil || len(g.GroupIds) == 0 {
		t.Fatal("Update method should return a list of the updated hostgroups.\nAn empty list was returned.")
	}

	if g.GroupIds[0] != hostGroupId {
		t.Fatalf("Wrong host group returned.\nExpected : %s\nReturned : %s", hostGroupId, g.GroupIds[0])
	}
}

func TestHostGroupDelete(t *testing.T) {
	g, err := testingClient.HostGroup.Delete([]string{
		hostGroupId,
	})

	if err != nil {
		t.Fatalf("Error when deleting hostgroup '%s'.\nReason : %v", hostGroupName, err)
	}

	if g.GroupIds == nil || len(g.GroupIds) == 0 {
		t.Fatalf("Delete method should return a list with the id of the deleted hostgroups.\nAn empty list was returned.")
	}

	if g.GroupIds[0] != hostGroupId {
		t.Fatalf("Wrong host group returned.\nExpected : %s\nReturned : %s", hostGroupId, g.GroupIds[0])
	}
}
