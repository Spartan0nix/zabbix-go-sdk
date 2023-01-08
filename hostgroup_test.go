package main

import "testing"

const (
	hostGroupName        = "test-host-group"
	updatedHostGroupName = "test-host-group-update"
)

func TestHostGroupCreate(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	g, err := client.HostGroup.Create(hostGroupName)

	if err != nil {
		t.Fatalf("Error when creating host group '%s'.\nReason : %v", hostGroupName, err)
	}

	if g.Groupids == nil {
		t.Fatalf("Create method should returned a list with the Id of the new host group.\nAn empty list was returned.")
	}
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
		Filter: map[string]string{
			"name": hostGroupName,
		},
	})

	if err != nil {
		t.Fatalf("Error when getting host group '%s'.\nReason : %v", hostGroupName, err)
	}

	if g[0].Name != hostGroupName {
		t.Fatalf("Wrong host returned.\nRequested name : %s\nName returned : %s", hostGroupName, g[0].Name)
	}
}

func TestHostGroupMassAdd(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	g, err := client.HostGroup.Get(&HostGroupGetParameters{
		Filter: map[string]string{
			"name": "Templates",
		},
	})

	if err != nil {
		t.Fatalf("Error when getting host group 'Templates'.\nReason : %v", err)
	}

	mass_add_g, err := client.HostGroup.MassAdd(&HostGroupMassAddParameters{
		Groups: []*HostGroupId{
			{
				Groupid: g[0].Id,
			},
		},
		Hosts: []*HostId{
			{
				Hostid: "10084",
			},
		},
	})

	if err != nil {
		t.Fatalf("Error when mass adding host group.\nHost group : %s\nHost : %s\nReason : %v", g[0].Name, "10084", err)
	}

	if mass_add_g == nil {
		t.Fatal("MassAdd method should returned a list of the updated host group(s).\nAn empty list was returned.")
	}
}

func TestHostGroupMassRemove(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	g, err := client.HostGroup.Get(&HostGroupGetParameters{
		Filter: map[string]string{
			"name": "Templates",
		},
	})

	if err != nil {
		t.Fatalf("Error when getting host group 'Templates'.\nReason : %v", err)
	}

	mass_rm_g, err := client.HostGroup.MassRemove(&HostGroupMassRemoveParameters{
		GroupsIds: []string{
			g[0].Id,
		},
		HostIds: []string{
			"10084",
		},
	})

	if err != nil {
		t.Fatalf("Error when mass removing host group.\nHost group : %s\nHosts : %s \nReason : %v", g[0].Name, "10084", err)
	}

	if mass_rm_g == nil {
		t.Fatal("MassRemove method should returned a list of the removed host group(s).\nAn empty list was returned.")
	}
}

func TestHostGroupMassUpdate(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	g, err := client.HostGroup.Get(&HostGroupGetParameters{
		Filter: map[string]string{
			"name": "Templates",
		},
	})

	if err != nil {
		t.Fatalf("Error when getting host group 'Templates'.\nReason : %v", err)
	}

	mass_rm_g, err := client.HostGroup.MassUpdate(&HostGroupMassUpdateParameters{
		Groups: []*HostGroupId{
			{
				Groupid: g[0].Id,
			},
		},
		Hosts: []*HostId{
			{
				Hostid: "10084",
			},
		},
		Templates: []*TemplateId{},
	})

	if err != nil {
		t.Fatalf("Error when mass updating host group.\nHost group : %s\nHosts : %s \nReason : %v", g[0].Name, "10084", err)
	}

	if mass_rm_g == nil {
		t.Fatal("MassUpdate method should returned a list of the updated host group(s).\nAn empty list was returned.")
	}
}

func TestHostGroupUpdate(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	g, err := client.HostGroup.Get(&HostGroupGetParameters{
		Filter: map[string]string{
			"name": hostGroupName,
		},
	})

	if err != nil {
		t.Fatalf("Error when getting host group '%s'.\nReason : %v", hostGroupName, err)
	}

	if g[0].Name != hostGroupName {
		t.Fatalf("Wrong host returned.\nRequested name : %s\nName returned : %s", hostGroupName, g[0].Name)
	}

	updated_g, err := client.HostGroup.Update(g[0].Id, updatedHostGroupName)
	if err != nil {
		t.Fatalf("Error when updating host group '%s'.\nReason : %v", hostGroupName, err)
	}

	if updated_g.Groupids == nil {
		t.Fatal("Update method should returned a list of with the updated host group id.\nAn empty list was returned.")
	}

}

func TestHostGroupDelete(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	g, err := client.HostGroup.Get(&HostGroupGetParameters{
		Filter: map[string]string{
			"name": updatedHostGroupName,
		},
	})

	if err != nil {
		t.Fatalf("Error when getting host group '%s'.\nReason : %v", hostGroupName, err)
	}

	if g[0].Name != updatedHostGroupName {
		t.Fatalf("Wrong host returned.\nRequested name : %s\nName returned : %s", updatedHostGroupName, g[0].Name)
	}

	deleted_g, err := client.HostGroup.Delete([]string{
		g[0].Id,
	})

	if err != nil {
		t.Fatalf("Error when deleting host group '%s'.\nReason : %v", hostGroupName, err)
	}

	if deleted_g.Groupids == nil {
		t.Fatalf("Delete method should returned a list with the Ids of the deleted host groups.\nAn empty list was returned.")
	}
}
