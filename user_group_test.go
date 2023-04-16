package zabbixgosdk

import "testing"

const (
	userGroupName       = "test-user-group"
	updateUserGroupName = "updated-test-user-group"
)

var userGroupId string

func TestUserGroupCreate(t *testing.T) {
	g, err := testingClient.UserGroup.Create(&UserGroupCreateParameters{
		Name: userGroupName,
	})

	if err != nil {
		t.Fatalf("Error when creating usergroup '%s'.\nReason : %v", userGroupName, err)
	}

	if g == nil || len(g.Usrgrpids) == 0 {
		t.Fatalf("Create method should return a list of the created usergroups.\nAn empty list was returned.")
	}

	userGroupId = g.Usrgrpids[0]
}

func TestUserGroupList(t *testing.T) {
	g, err := testingClient.UserGroup.List()
	if err != nil {
		t.Fatalf("Error when listing usergroup.\nReason : %v", err)
	}

	if len(g) == 0 {
		t.Fatalf("List method should return a list with all the existing usergroups on the server.\nAn empty list was returned.")
	}
}

func TestUserGroupGet(t *testing.T) {
	g, err := testingClient.UserGroup.Get(&UserGroupGetParameters{
		UsrGrpIds: []string{
			userGroupId,
		},
	})

	if err != nil {
		t.Fatalf("Error when getting usergroup '%s'.\nReason : %v", userGroupName, err)
	}

	if len(g) == 0 {
		t.Fatalf("Get method should return a list of usergroups matching the given criteria.\nAn empty list was returned.")
	}

	if g[0].UsrGrpId != userGroupId {
		t.Fatalf("Wrong user group returned.\nExpected : %s\nReturned : %s", userGroupId, g[0].UsrGrpId)
	}
}

func TestUserGroupUpdate(t *testing.T) {
	g, err := testingClient.UserGroup.Update(&UserGroupUpdateParameters{
		UsrGrpId:  userGroupId,
		DebugMode: UserGroupDebugEnabled,
		GuiAccess: UserGroupInternal,
	})

	if err != nil {
		t.Fatalf("Error when updating usergroup '%s'.\nReason : %v", userGroupName, err)
	}

	if g == nil || len(g.Usrgrpids) == 0 {
		t.Fatalf("Update method should return a list of the updated usergroups.\nAn empty list was returned.")
	}

	if g.Usrgrpids[0] != userGroupId {
		t.Fatalf("Wrong user group returned.\nExpected : %s\nReturned : %s", userGroupId, g.Usrgrpids[0])
	}
}

func TestUserGroupDelete(t *testing.T) {
	g, err := testingClient.UserGroup.Delete([]string{
		userGroupId,
	})

	if err != nil {
		t.Fatalf("Error when deleting usergroup '%s'.\nReason : %v", userGroupName, err)
	}

	if g == nil || len(g.Usrgrpids) == 0 {
		t.Fatalf("Update method should return a list of the deleted usergroups.\nAn empty list was returned.")
	}

	if g.Usrgrpids[0] != userGroupId {
		t.Fatalf("Wrong user group returned.\nExpected : %s\nReturned : %s", userGroupId, g.Usrgrpids[0])
	}
}
