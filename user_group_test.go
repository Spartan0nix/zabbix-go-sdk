package main

import "testing"

const (
	name = "test-user-group"
)

func TestUserGroupList(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	groups, err := client.UserGroup.List()
	if err != nil {
		t.Fatalf("Error when listing user group.\nReason : %v", err)
	}

	if groups[0].Id == "" {
		t.Fail()
	}
}

func TestUserGroupCreate(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	group, err := client.UserGroup.Create(&UserGroupExtendedParameters{
		Name: name,
	})

	if err != nil {
		t.Fatalf("Error when creating user group '%s'.\nReason : %v", name, err)
	}

	if group.Usrgrpids[0] == "" {
		t.Fail()
	}
}

func TestUserGroupGet(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	group, err := client.UserGroup.Get(&UserGroupGetParameters{
		Filter: map[string]string{
			"name": name,
		},
	})

	if err != nil {
		t.Fatalf("Error when getting user group '%s'.\nReason : %v", name, err)
	}

	if group[0].Name != name {
		t.Fail()
	}
}

func TestUserGroupUpdate(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	group, err := client.UserGroup.Get(&UserGroupGetParameters{
		Filter: map[string]string{
			"name": name,
		},
		Output: "extend",
	})

	if err != nil {
		t.Fatalf("Error when getting user group '%s'.\nReason : %v", name, err)
	}

	if group[0].Name != name {
		t.Fatalf("Failed to get user group '%s'.", name)
	}

	updated_group, err := client.UserGroup.Update(&UserGroupExtendedParameters{
		Id:         group[0].Id,
		Debug_mode: 1,
		Gui_access: 0,
	})

	if err != nil {
		t.Fatalf("Error when updating user group '%s'.\nReason : %v", name, err)
	}

	if updated_group.Usrgrpids[0] == "" {
		t.Fail()
	}
}

func TestUserGroupDelete(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	group, err := client.UserGroup.Get(&UserGroupGetParameters{
		Filter: map[string]string{
			"name": name,
		},
		Output: "extend",
	})

	if err != nil {
		t.Fatalf("Error when getting user group '%s'.\nReason : %v", name, err)
	}

	if group[0].Name != name {
		t.Fatalf("Failed to get user group '%s'.", name)
	}

	deleted_group, err := client.UserGroup.Delete([]string{
		group[0].Id,
	})

	if err != nil {
		t.Fatalf("Error when deleting user group '%s'.\nReason : %v", name, err)
	}

	if deleted_group.Usrgrpids[0] == "" {
		t.Fail()
	}
}
