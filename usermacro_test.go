package main

import "testing"

func TestUserMacroCreate(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	h, err := client.UserMacro.Create(HostMacro{
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

	_, err := client.UserMacro.Create(HostMacro{
		Hostid: "10084",
		Macro:  "TEST",
		Value:  "test",
	})

	if err == nil {
		t.Fatal("No error returned when trying to create a user macro without following Zabbix required format")
	}
}

func TestGlobalMacroCreate(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	h, err := client.UserMacro.CreateGlobal(GlobalMacro{
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

func TestGlobalMacroCreateWrongFormat(t *testing.T) {
	client := NewZabbixService()

	_, err := client.UserMacro.CreateGlobal(GlobalMacro{
		Macro: "TEST",
		Value: "test",
	})

	if err == nil {
		t.Fatal("No error returned when trying to create a global macro without following Zabbix required format")
	}
}

func TestUserMacroGetHostMacro(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	m, err := client.UserMacro.GetHostMacro(&UserMacroGetParameters{
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

func TestUserMacroGetGlobalMacro(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	m, err := client.UserMacro.GetGlobalMacro(&UserMacroGetParameters{})

	if err != nil {
		t.Fatalf("Error when retrieving global macros.\nReason : %v", err)
	}

	if m == nil {
		t.Fatal("An empty response was returned when retrieving server global macros.")
	}
}

func TestUserMacroDelete(t *testing.T) {
	// client, err := NewTestingService()
	// if err != nil {
	// 	t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	// }

	/*
	* Waiting for usermacro.get implementation
	 */

	// h, err = client.UserMacro.Delete([]string{
	// 	h.Hostmacroids[0],
	// })

	// if err != nil {
	// 	t.Fatalf("Error when deleting user macro.\nReason : %v", err)
	// }

	// if h.Hostmacroids[0] == "" {
	// 	t.Fatalf("Delete request returned an empty response.")
	// }
}

func TestUserMacroDeleteGlobal(t *testing.T) {
	// client, err := NewTestingService()
	// if err != nil {
	// 	t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	// }

	/*
	* Waiting for usermacro.get implementation
	 */

	// h, err = client.UserMacro.DeleteGlobal([]string{
	// 	h.Hostmacroids[0],
	// })

	// if err != nil {
	// 	t.Fatalf("Error when deleting global macro.\nReason : %v", err)
	// }

	// if h.Globalmacroids[0] == "" {
	// 	t.Fatalf("Delete request returned an empty response.")
	// }
}
