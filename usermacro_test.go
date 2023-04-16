package zabbixgosdk

import (
	"testing"
)

const (
	hostMacroName   = "{$TEST}"
	globalMacroName = "{$TEST_GLOBAL}"
)

var hostMacroId string
var globalMacroId string

// --------------------------------------
// HostMacro
// --------------------------------------
func TestUserMacroValidateHostMacro(t *testing.T) {
	m := &HostMacro{
		Macro: hostMacroName,
		Value: "value",
	}

	if err := m.ValidateMacro(); err != nil {
		t.Fatalf("ValidateMacro for HostMacro should not have had returned an error.\nError returned : %v", err)
	}
}

func TestUserMacroValidateHostMacroError(t *testing.T) {
	m := &HostMacro{
		Macro: "{$WRONG-FORMAT}",
		Value: "value",
	}

	if err := m.ValidateMacro(); err == nil {
		t.Fatalf("ValidateMacro for HostMacro should have had returned an error.")
	}
}

func TestUserMacroCreate(t *testing.T) {
	h, err := testingClient.UserMacro.Create(&HostMacro{
		HostId: "10084",
		Macro:  hostMacroName,
		Value:  "test",
		Type:   Text,
	})

	if err != nil {
		t.Fatalf("Error when creating usermacro '%s'.\nReason : %v", hostMacroName, err)
	}

	if h == nil || len(h.HostMacroIds) == 0 {
		t.Fatalf("Create method should return a list of the created hostmacros.\nAn empty list was returned.")
	}

	hostMacroId = h.HostMacroIds[0]
}

func TestUserMacroCreateWrongFormat(t *testing.T) {
	_, err := testingClient.UserMacro.Create(&HostMacro{
		HostId: "10084",
		Macro:  "TEST",
		Value:  "test",
	})

	if err == nil {
		t.Fatal("No error returned when trying to create a usermacro without following Zabbix required format.")
	}
}

func TestUserMacroCreateWrongMissingId(t *testing.T) {
	_, err := testingClient.UserMacro.Create(&HostMacro{
		Macro: "TEST",
		Value: "test",
	})

	if err == nil {
		t.Fatal("No error returned when trying to create a usermacro without the required 'Id' field.")
	}
}

func TestUserMacroGet(t *testing.T) {
	m, err := testingClient.UserMacro.Get(&UserMacroGetParameters{
		HostMacroIds: []string{
			hostMacroId,
		},
	})

	if err != nil {
		t.Fatalf("Error when getting hostmacro '%s'.\nReason : %v", hostMacroId, err)
	}

	if len(m) == 0 {
		t.Fatalf("Get method should return a list of hostmacros matching the given criteria.\nAn empty list was returned.")
	}

	if m[0].HostMacroId != hostMacroId {
		t.Fatalf("Wrong host macro returned.\nExpected : %s\nReturned : %s", hostMacroId, m[0].HostMacroId)
	}
}

func TestUserMacroUpdate(t *testing.T) {
	m, err := testingClient.UserMacro.Update(&HostMacroUpdateParameters{
		HostMacroId: hostMacroId,
		Type:        Secret,
		Value:       "new-secret-value",
	})

	if err != nil {
		t.Fatalf("Error when updating hostmacro '%s'.\nReason : %v", hostMacroId, err)
	}

	if m == nil || len(m.HostMacroIds) == 0 {
		t.Fatal("Update method should return a list of the updated hostmacros.\nAn empty list was returned.")
	}

	if m.HostMacroIds[0] != hostMacroId {
		t.Fatalf("Wrong host macro returned.\nExpected : %s\nReturned : %s", hostMacroId, m.HostMacroIds[0])
	}
}

func TestUserMacroDelete(t *testing.T) {
	m, err := testingClient.UserMacro.Delete([]string{
		hostMacroId,
	})

	if err != nil {
		t.Fatalf("Error when deleting hostmacro '%s'.\nReason : %v", hostMacroId, err)
	}

	if m == nil || len(m.HostMacroIds) == 0 {
		t.Fatalf("Delete method should return a list of the deleted hostmacros.\nAn empty list was returned.")
	}

	if m.HostMacroIds[0] != hostMacroId {
		t.Fatalf("Wrong host macro returned.\nExpected : %s\nReturned : %s", hostMacroId, m.HostMacroIds[0])
	}
}

// --------------------------------------
// GlobalMacro
// --------------------------------------
func TestUserMacroValidateGlobalMacro(t *testing.T) {
	m := &GlobalMacro{
		Macro: globalMacroName,
		Value: "value",
	}

	if err := m.ValidateMacro(); err != nil {
		t.Fatalf("ValidateMacro for GlobalMacro should not have had returned an error.\nError returned : %v", err)
	}
}

func TestUserMacroValidateGlobalMacroError(t *testing.T) {
	m := &GlobalMacro{
		Macro: "{$WRONG-FORMAT}",
		Value: "value",
	}

	if err := m.ValidateMacro(); err == nil {
		t.Fatalf("ValidateMacro for GlobalMacro should have had returned an error.")
	}
}

func TestUserMacroCreateGlobal(t *testing.T) {
	m, err := testingClient.UserMacro.CreateGlobal(&GlobalMacro{
		Macro: globalMacroName,
		Value: "test",
	})

	if err != nil {
		t.Fatalf("Error when creating globalmacro '%s'.\nReason : %v", globalMacroName, err)
	}

	if m == nil || len(m.GlobalMacroIds) == 0 {
		t.Fatalf("Create method should return a list of the created globalmacros.\nAn empty list was returned.")
	}

	globalMacroId = m.GlobalMacroIds[0]
}

func TestUserMacroCreateGlobalWrongFormat(t *testing.T) {
	_, err := testingClient.UserMacro.CreateGlobal(&GlobalMacro{
		Macro: "TEST",
		Value: "test",
	})

	if err == nil {
		t.Fatal("No error returned when trying to create a global macro without following Zabbix required format.")
	}
}

func TestUserMacroGetGlobal(t *testing.T) {
	m, err := testingClient.UserMacro.GetGlobal(&UserMacroGetParameters{
		GlobalMacroIds: []string{
			globalMacroId,
		},
	})

	if err != nil {
		t.Fatalf("Error when getting globalmacro '%s'.\nReason : %v", globalMacroId, err)
	}

	if len(m) == 0 {
		t.Fatal("Get method should return a list of globalmacros matching the given criteria.\nAn empty list was returned.")
	}

	if m[0].GlobalMacroId != globalMacroId {
		t.Fatalf("Wrong globalmacro returned.\nExpected : %s\nReturned : %s", globalMacroId, m[0].GlobalMacroId)
	}
}

func TestUserMacroUpdateGlobal(t *testing.T) {
	m, err := testingClient.UserMacro.UpdateGlobal(&GlobalMacroUpdateParameters{
		GlobalMacroId: globalMacroId,
		Type:          Secret,
		Value:         "new-secret-value",
	})

	if err != nil {
		t.Fatalf("Error when updating globalmacro '%s'.\nReason : %v", globalMacroId, err)
	}

	if m == nil || len(m.GlobalMacroIds) == 0 {
		t.Fatal("MassUpdate method should return a list of the updated globalmacros.\nAn empty list was returned.")
	}

	if m.GlobalMacroIds[0] != globalMacroId {
		t.Fatalf("Wrong globalmacro returned.\nExpected : %s\nReturned : %s", globalMacroId, m.GlobalMacroIds[0])
	}
}

func TestUserMacroDeleteGlobal(t *testing.T) {
	m, err := testingClient.UserMacro.DeleteGlobal([]string{
		globalMacroId,
	})

	if err != nil {
		t.Fatalf("Error when deleting globalmacro '%s'.\nReason : %v", globalMacroId, err)
	}

	if m == nil || len(m.GlobalMacroIds) == 0 {
		t.Fatalf("Delete method should return a list of the deleted globalmacros.\nAn empty list was returned.")
	}

	if m.GlobalMacroIds[0] != globalMacroId {
		t.Fatalf("Wrong globalmacro returned.\nExpected : %s\nReturned : %s", globalMacroId, m.GlobalMacroIds[0])
	}
}
