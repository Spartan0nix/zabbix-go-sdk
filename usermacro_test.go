package zabbixgosdk

import (
	"testing"
)

const (
	HostMacroName   = "{$TEST}"
	GlobalMacroName = "{$TEST_GLOBAL}"
)

var HostMacroId string
var GlobalMacroId string

// --------------------------------------
// HostMacro
// --------------------------------------
func TestUserMacroValidateHostMacro(t *testing.T) {
	m := &HostMacro{
		Macro: HostMacroName,
		Value: "value",
	}

	if err := m.ValidateMacro(); err != nil {
		t.Fatalf("ValidateMacro for HostMacro should not have had return an error.\nError returned : %v", err)
	}
}

func TestUserMacroValidateHostMacroError(t *testing.T) {
	m := &HostMacro{
		Macro: "{$WRONG-FORMAT}",
		Value: "value",
	}

	if err := m.ValidateMacro(); err == nil {
		t.Fatalf("ValidateMacro for HostMacro should have returned an error.")
	}
}

func TestUserMacroCreate(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	h, err := client.UserMacro.Create(&HostMacro{
		HostId: "10084",
		Macro:  HostMacroName,
		Value:  "test",
		Type:   Text,
	})

	if err != nil {
		t.Fatalf("Error when creating user macro.\nReason : %v", err)
	}

	if h == nil {
		t.Fatalf("Create method should returned a list with the id of the new host macros.\nAn empty list was returned.")
	}

	HostMacroId = h.HostMacroIds[0]
}

func TestUserMacroCreateWrongFormat(t *testing.T) {
	client := NewZabbixService()

	_, err := client.UserMacro.Create(&HostMacro{
		HostId: "10084",
		Macro:  "TEST",
		Value:  "test",
	})

	if err == nil {
		t.Fatal("No error returned when trying to create a user macro without following Zabbix required format")
	}
}

func TestUserMacroCreateWrongMissingId(t *testing.T) {
	client := NewZabbixService()

	_, err := client.UserMacro.Create(&HostMacro{
		Macro: "TEST",
		Value: "test",
	})

	if err == nil {
		t.Fatal("No error returned when trying to create a user macro without the required 'Id' field.")
	}
}

func TestUserMacroGet(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	m, err := client.UserMacro.Get(&UserMacroGetParameters{
		HostMacroIds: []string{
			HostMacroId,
		},
	})

	if err != nil {
		t.Fatalf("Error when retrieving host macros.\nReason : %v", err)
	}

	if m == nil {
		t.Fatalf("Get method should returned a list of host macros matching the given criteria.\nAn empty list was returned.")
	}

	if m[0].HostMacroId != HostMacroId {
		t.Fatalf("Wrong host macro returned.\nExpected Id : %s\nId returned : %s", HostMacroId, m[0].HostMacroId)
	}
}

func TestUserMacroUpdate(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	m, err := client.UserMacro.Update(&HostMacroUpdateParameters{
		HostMacroId: HostMacroId,
		Type:        Secret,
		Value:       "new-secret-value",
	})

	if err != nil {
		t.Fatalf("Error when updating host macro.\nReason : %v", err)
	}

	if m == nil {
		t.Fatal("Update method should returned a list of the updated host macros.\nAn empty list was returned.")
	}

	if m.HostMacroIds[0] != HostMacroId {
		t.Fatalf("Wrong host macro returned.\nExpected Id : %s\nId returned : %s", HostMacroId, m.HostMacroIds[0])
	}
}

func TestUserMacroDelete(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	m, err := client.UserMacro.Delete([]string{
		HostMacroId,
	})

	if err != nil {
		t.Fatalf("Error when deleting user macro.\nReason : %v", err)
	}

	if m == nil {
		t.Fatalf("Delete method should returned a list with the id of the deleted host macros.\nAn empty list was returned.")
	}

	if m.HostMacroIds[0] != HostMacroId {
		t.Fatalf("Wrong host macro returned.\nExpected Id : %s\nId returned : %s", HostMacroId, m.HostMacroIds[0])
	}
}

// --------------------------------------
// HostMacro
// --------------------------------------
func TestUserMacroValidateGlobalMacro(t *testing.T) {
	m := &GlobalMacro{
		Macro: GlobalMacroName,
		Value: "value",
	}

	if err := m.ValidateMacro(); err != nil {
		t.Fatalf("ValidateMacro for GlobalMacro should not have had return an error.\nError returned : %v", err)
	}
}

func TestUserMacroValidateGlobalMacroError(t *testing.T) {
	m := &GlobalMacro{
		Macro: "{$WRONG-FORMAT}",
		Value: "value",
	}

	if err := m.ValidateMacro(); err == nil {
		t.Fatalf("ValidateMacro for GlobalMacro should have returned an error.")
	}
}

func TestUserMacroCreateGlobal(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	m, err := client.UserMacro.CreateGlobal(&GlobalMacro{
		Macro: GlobalMacroName,
		Value: "test",
	})

	if err != nil {
		t.Fatalf("Error when creating global macro.\nReason : %v", err)
	}

	if m == nil {
		t.Fatalf("Create method should returned a list with the id of the new global macros.\nAn empty list was returned.")
	}

	GlobalMacroId = m.GlobalMacroIds[0]
}

func TestUserMacroCreateGlobalWrongFormat(t *testing.T) {
	client := NewZabbixService()

	_, err := client.UserMacro.CreateGlobal(&GlobalMacro{
		Macro: "TEST",
		Value: "test",
	})

	if err == nil {
		t.Fatal("No error returned when trying to create a global macro without following Zabbix required format")
	}
}

func TestUserMacroGetGlobal(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	m, err := client.UserMacro.GetGlobal(&UserMacroGetParameters{
		GlobalMacroIds: []string{
			GlobalMacroId,
		},
	})

	if err != nil {
		t.Fatalf("Error when retrieving global macros.\nReason : %v", err)
	}

	if m == nil {
		t.Fatal("Get method should returned a list of global macros matching the given criteria.\nAn empty list was returned.")
	}

	if m[0].GlobalMacroId != GlobalMacroId {
		t.Fatalf("Wrong global macro returned.\nExpected Id : %s\nId returned : %s", HostMacroId, m[0].GlobalMacroId)
	}
}

func TestUserMacroUpdateGlobal(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	m, err := client.UserMacro.UpdateGlobal(&GlobalMacroUpdateParameters{
		GlobalMacroId: GlobalMacroId,
		Type:          Secret,
		Value:         "new-secret-value",
	})

	if err != nil {
		t.Fatalf("Error when updating global macro.\nReason : %v", err)
	}

	if m == nil {
		t.Fatal("MassUpdate method should returned a list of the updated global macros.\nAn empty list was returned.")
	}

	if m.GlobalMacroIds[0] != GlobalMacroId {
		t.Fatalf("Wrong global macro returned.\nExpected Id : %s\nId returned : %s", HostMacroId, m.GlobalMacroIds[0])
	}
}

func TestUserMacroDeleteGlobal(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	m, err := client.UserMacro.DeleteGlobal([]string{
		GlobalMacroId,
	})

	if err != nil {
		t.Fatalf("Error when deleting global macro.\nReason : %v", err)
	}

	if m.GlobalMacroIds[0] == "" {
		t.Fatalf("Delete method should returned a list with the id of the deleted global macros.\nAn empty list was returned.")
	}

	if m.GlobalMacroIds[0] != GlobalMacroId {
		t.Fatalf("Wrong global macro returned.\nExpected Id : %s\nId returned : %s", HostMacroId, m.GlobalMacroIds[0])
	}
}
