package zabbixgosdk

import "testing"

const (
	userGroupName       = "test-user-group"
	updateUserGroupName = "updated-test-user-group"
)

var userGroupId string

func TestUserGroupList(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	g, err := client.UserGroup.List()
	if err != nil {
		t.Fatalf("Error when listing user group.\nReason : %v", err)
	}

	if g == nil {
		t.Fatalf("List method should returned a list of existing user groups.\nAn empty list was returned.")
	}
}

func TestUserGroupCreate(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	g, err := client.UserGroup.Create(&UserGroupCreateParameters{
		Name: userGroupName,
	})

	if err != nil {
		t.Fatalf("Error when creating user group '%s'.\nReason : %v", userGroupName, err)
	}

	if g == nil {
		t.Fatalf("Create method should returned a list with the id of the new user groups.\nAn empty list was returned.")
	}

	userGroupId = g.Usrgrpids[0]
}

func TestUserGroupGet(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	g, err := client.UserGroup.Get(&UserGroupGetParameters{
		Filter: map[string]string{
			"name": userGroupName,
		},
	})

	if err != nil {
		t.Fatalf("Error when getting user group '%s'.\nReason : %v", userGroupName, err)
	}

	if g == nil {
		t.Fatalf("Get method should returned a list of user groups matching the given criteria.\nAn empty list was returned.")
	}

	if g[0].UsrGrpId != userGroupId {
		t.Fatalf("Wrong user group returned.\nExpected Id : %s\nId returned : %s", userGroupId, g[0].UsrGrpId)
	}
}

func TestUserGroupUpdate(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	g, err := client.UserGroup.Update(&UserGroupUpdateParameters{
		UsrGrpId:  userGroupId,
		DebugMode: UserGroupDebugEnabled,
		GuiAccess: UserGroupInternal,
	})

	if err != nil {
		t.Fatalf("Error when updating user group '%s'.\nReason : %v", userGroupName, err)
	}

	if g == nil {
		t.Fatalf("Update method should returned a list of the updated user groups.\nAn empty list was returned.")
	}

	if g.Usrgrpids[0] != userGroupId {
		t.Fatalf("Wrong user group returned.\nExpected Id : %s\nId returned : %s", userGroupId, g.Usrgrpids[0])
	}
}

func TestUserGroupDelete(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	g, err := client.UserGroup.Delete([]string{
		userGroupId,
	})

	if err != nil {
		t.Fatalf("Error when deleting user group '%s'.\nReason : %v", userGroupName, err)
	}

	if g == nil {
		t.Fatalf("Update method should returned a list of the removed user groups.\nAn empty list was returned.")
	}

	if g.Usrgrpids[0] != userGroupId {
		t.Fatalf("Wrong user group returned.\nExpected Id : %s\nId returned : %s", userGroupId, g.Usrgrpids[0])
	}
}
