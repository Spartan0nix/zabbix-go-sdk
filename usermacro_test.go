package main

import (
	"testing"
)

func TestUserMacroValidateHostMacro(t *testing.T) {
	m := &HostMacro{
		Macro: "{$GOOD_FORMAT}",
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
		Hostid: "10084",
		Macro:  "{$TEST}",
		Value:  "test",
	})

	if err != nil {
		t.Fatalf("Error when creating user macro.\nReason : %v", err)
	}

	if h.Hostmacroids[0] == "" {
		t.Fatalf("Create request returned an empty response.")
	}
}

func TestUserMacroCreateWrongFormat(t *testing.T) {
	client := NewZabbixService()

	_, err := client.UserMacro.Create(&HostMacro{
		Hostid: "10084",
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
		Hostids: []string{
			"10084",
		},
	})

	if err != nil {
		t.Fatalf("Error when retrieving host macros.\nReason : %v", err)
	}

	if m == nil {
		t.Fatalf("An empty response was returned when retrieving macros for host '%s'.", "10084")
	}
}

func TestUserMacroUpdate(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	m, err := client.UserMacro.Get(&UserMacroGetParameters{
		Hostids: []string{
			"10084",
		},
	})

	if err != nil {
		t.Fatalf("Error when retrieving host macros.\nReason : %v", err)
	}

	if m == nil {
		t.Fatalf("An empty response was returned when retrieving macros for host '%s'.", "10084")
	}

	m[0].Type = Secret
	m[0].Value = "new-secret-value"

	updated_m, err := client.UserMacro.Update(m[0])
	if err != nil {
		t.Fatalf("Error when updating host macro.\nReason : %v", err)
	}

	if updated_m == nil {
		t.Fatalf("An empty response was returned when updating host macro.\nDesired state : %v", m[0])
	}
}

func TestUserMacroDelete(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	m, err := client.UserMacro.Get(&UserMacroGetParameters{
		Hostids: []string{
			"10084",
		},
	})

	if err != nil {
		t.Fatalf("Error when retrieving host macros.\nReason : %v", err)
	}

	if m == nil {
		t.Fatalf("An empty response was returned when retrieving macros for host '%s'.", "10084")
	}

	d, err := client.UserMacro.Delete([]string{
		m[0].Id,
	})

	if err != nil {
		t.Fatalf("Error when deleting user macro.\nReason : %v", err)
	}

	if d.Hostmacroids[0] == "" {
		t.Fatalf("Delete request returned an empty response.")
	}

	if d.Hostmacroids[0] != m[0].Id {
		t.Fatalf("Error during delete request.\nId of the deleted host macro is '%s'.\nHost macro with id '%s' was requested for deletion.", d.Hostmacroids[0], m[0].Id)
	}
}

// -----------------------------------
func TestUserMacroValidateGlobalMacro(t *testing.T) {
	m := &GlobalMacro{
		Macro: "{$GOOD_FORMAT}",
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

	h, err := client.UserMacro.CreateGlobal(&GlobalMacro{
		Macro: "{$TEST}",
		Value: "test",
	})

	if err != nil {
		t.Fatalf("Error when creating global macro.\nReason : %v", err)
	}

	if h.Globalmacroids[0] == "" {
		t.Fatalf("Create request returned an empty response.")
	}
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

	m, err := client.UserMacro.GetGlobal(&UserMacroGetParameters{})

	if err != nil {
		t.Fatalf("Error when retrieving global macros.\nReason : %v", err)
	}

	if m == nil {
		t.Fatal("An empty response was returned when retrieving server global macros.")
	}
}

func TestUserMacroUpdateGlobal(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	m, err := client.UserMacro.GetGlobal(&UserMacroGetParameters{})

	if err != nil {
		t.Fatalf("Error when retrieving global macros.\nReason : %v", err)
	}

	if m == nil {
		t.Fatal("An empty response was returned when retrieving server global macros.")
	}

	m[0].Type = Secret
	m[0].Value = "new-secret-value"

	updated_m, err := client.UserMacro.UpdateGlobal(m[0])
	if err != nil {
		t.Fatalf("Error when updating global macro.\nReason : %v", err)
	}

	if updated_m == nil {
		t.Fatalf("An empty response was returned when updating global macro.\nDesired state : %v", m[0])
	}
}

func TestUserMacroDeleteGlobal(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	m, err := client.UserMacro.GetGlobal(&UserMacroGetParameters{})

	if err != nil {
		t.Fatalf("Error when retrieving global macros.\nReason : %v", err)
	}

	if m == nil {
		t.Fatal("An empty response was returned when retrieving server global macros.")
	}

	d, err := client.UserMacro.DeleteGlobal([]string{
		m[0].Id,
	})

	if err != nil {
		t.Fatalf("Error when deleting global macro.\nReason : %v", err)
	}

	if d.Globalmacroids[0] == "" {
		t.Fatalf("Delete request returned an empty response.")
	}

	if d.Globalmacroids[0] != m[0].Id {
		t.Fatalf("Error during delete request.\nId of the deleted global macro is '%s'.\nGlobal macro with id '%s' was requested for deletion.", d.Globalmacroids[0], m[0].Id)
	}
}
