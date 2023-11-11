package zabbixgosdk

import (
	"fmt"
	"testing"
)

func TestHostCreate(t *testing.T) {
	g, err := testingClient.HostGroup.Get(&HostGroupGetParameters{
		CommonGetParameters: CommonGetParameters{
			Search: map[string]string{
				"name": "Virtual machines",
			},
			Output: []string{
				"groupid",
			},
		},
	})

	if err != nil {
		t.Fatalf("error while retrieving hostGroup 'Virtual machines'\nReason: %v", err)
	}

	if len(g.Result) == 0 {
		t.Fatal("error while retrieving hostGroup 'Virtual machines'\nReason: an empty list of hostGroups was returned")
	}

	template, err := testingClient.Template.Get(&TemplateGetParameters{
		CommonGetParameters: CommonGetParameters{
			Search: map[string]string{
				"name": "Linux by Zabbix agent",
			},
			Output: []string{
				"templateid",
			},
		},
	})

	if err != nil {
		t.Fatalf("error while retrieving template 'Linux by Zabbix agent'\nReason: %v", err)
	}

	if len(template.Result) == 0 {
		t.Fatal("error while retrieving template 'Linux by Zabbix agent'\nReason: an empty list of templates was returned")
	}

	name := fmt.Sprintf("test-host-%d", generateId())
	h, err := testingClient.Host.Create(&HostCreateParameters{
		Host: name,
		Name: name,
		Groups: []*HostGroupId{
			{
				Id: g.Result[0].Id,
			},
		},
		Description:   "This host was created during tests",
		InventoryMode: HostManual,
		IpmiAuthType:  IpmiAuthNone,
		Status:        HostUnmonitored,
		TlsConnect:    HostNoEncryption,
		Interfaces: []*HostInterfaceCreateParameters{
			{
				Dns:   "localhost",
				Ip:    "127.0.0.1",
				Main:  HostInterfaceDefault,
				Port:  "10051",
				Type:  HostInterfaceAgent,
				UseIp: HostInterfaceUseIp,
			},
			{
				Dns:   "localhost",
				Ip:    "127.0.0.1",
				Main:  HostInterfaceDefault,
				Port:  "161",
				Type:  HostInterfaceSNMP,
				UseIp: HostInterfaceUseIp,
				Details: &HostInterfaceDetail{
					Version:        HostInterfaceSnmpV3,
					Bulk:           HostInterfaceUseSnmpBulk,
					SecurityName:   "public",
					ContextName:    "public",
					SecurityLevel:  HostInterfaceAuthPriv,
					AuthPassPhrase: "public",
					PrivPassPhrase: "public",
					AuthProtocol:   HostInterfaceSHA256,
					PrivProtocol:   HostInterfaceAES256,
				},
			},
		},
		Tags: []*HostTag{
			{
				Tag:   "source",
				Value: "tests",
			},
		},
		Templates: []*TemplateId{
			{
				Id: template.Result[0].Id,
			},
		},
		Macros: []*HostMacroCreateParamaters{
			{
				Macro:       "{$MY.MACRO}",
				Value:       "my-macro-value",
				Type:        UserMacroSecret,
				Description: "Macro created during tests",
			},
		},
		Inventory: map[HostInventory]string{
			Notes:            "Created during tests",
			DeploymentStatus: "can-be-removed",
		},
	})

	if err != nil {
		t.Fatalf("error while executing Create function\nReason: %v", err)
	}

	if h == nil || len(h.Result.Ids) == 0 {
		t.Fatal("an empty list of host ids was returned")
	}
}

func TestHostParseDetails(t *testing.T) {
	res := &hostGetResponseInterface{
		InterfaceId: "1",
		HostId:      "1",
		Ip:          "127.0.0.1",
		Dns:         "localhost",
		Main:        HostInterfaceDefault,
		Port:        "161",
		Type:        HostInterfaceSNMP,
		UseIp:       HostInterfaceUseIp,
		Details: map[string]interface{}{
			"version":        "2",
			"bulk":           "1",
			"community":      "",
			"securityname":   "sec",
			"contextname":    "cont",
			"securitylevel":  "2",
			"authpassphrase": "pass",
			"privpassphrase": "priv",
			"authprotocol":   "5",
			"privprotocol":   "3",
		},
		Available:    HostInterfaceAvailable,
		DisableUntil: "",
		Error:        "",
		ErrorsFrom:   "",
	}

	detail := res.ParseDetails()
	fmt.Println(detail)

	if detail.Version != HostInterfaceSnmpV2c {
		t.Fatalf("wrong value returned for 'Version' field\nExpected: %s\nReturned: %s", "2", detail.Version)
	}
	if detail.Bulk != HostInterfaceUseSnmpBulk {
		t.Fatalf("wrong value returned for 'Bulk' field\nExpected: %s\nReturned: %s", "1", detail.Bulk)
	}
	if detail.Community != "" {
		t.Fatalf("wrong value returned for 'Community' field\nExpected: %s\nReturned: %s", "", detail.Community)
	}
	if detail.SecurityName != "sec" {
		t.Fatalf("wrong value returned for 'SecurityName' field\nExpected: %s\nReturned: %s", "sec", detail.SecurityName)
	}
	if detail.ContextName != "cont" {
		t.Fatalf("wrong value returned for 'ContextName' field\nExpected: %s\nReturned: %s", "cont", detail.ContextName)
	}
	if detail.SecurityLevel != HostInterfaceAuthPriv {
		t.Fatalf("wrong value returned for 'SecurityLevel' field\nExpected: %s\nReturned: %s", "2", detail.SecurityLevel)
	}
	if detail.AuthPassPhrase != "pass" {
		t.Fatalf("wrong value returned for 'AuthPassPhrase' field\nExpected: %s\nReturned: %s", "pass", detail.AuthPassPhrase)
	}
	if detail.PrivPassPhrase != "priv" {
		t.Fatalf("wrong value returned for 'PrivPassPhrase' field\nExpected: %s\nReturned: %s", "priv", detail.PrivPassPhrase)
	}
	if detail.AuthProtocol != HostInterfaceSHA512 {
		t.Fatalf("wrong value returned for 'AuthProtocol' field\nExpected: %s\nReturned: %s", "5", detail.AuthProtocol)
	}
	if detail.PrivProtocol != HostInterfaceAES256 {
		t.Fatalf("wrong value returned for 'PrivProtocol' field\nExpected: %s\nReturned: %s", "3", detail.PrivProtocol)
	}
}

func TestHostParseInventory(t *testing.T) {
	res := hostGetResponseInventory{
		Alias: "alias-value",
		Type:  "type-value",
	}

	inventory := res.ParseInventory()

	if inventory[Alias] != "alias-value" {
		t.Fatalf("wrong value returned for 'Alias' key\nExpected: %s\nReturned: %s", "alias-value", inventory[Alias])
	}
	if inventory[Type] != "type-value" {
		t.Fatalf("wrong value returned for 'Type' key\nExpected: %s\nReturned: %s", "type-value", inventory[Type])
	}

}

func TestHostGet(t *testing.T) {
	h, err := testingClient.Host.Get(&HostGetParameters{
		CommonGetParameters: CommonGetParameters{
			Filter: map[string]string{},
			Output: "extend",
		},
		SelectDiscoveryRule:   "extend",
		SelectDiscoveries:     "extend",
		SelectGraphs:          "extend",
		SelectGroups:          "extend",
		SelectHostDiscovery:   "extend",
		SelectHttpTests:       "extend",
		SelectInterfaces:      "extend",
		SelectInventory:       "extend",
		SelectItems:           "extend",
		SelectMacros:          "extend",
		SelectParentTemplates: "extend",
		SelectDashboards:      "extend",
		SelectTags:            "extend",
		SelectInheritedTags:   "extend",
		SelectTriggers:        "extend",
		SelectValueMaps:       "extend",
	})

	if err != nil {
		t.Fatalf("error while executing Get function\nReason: %v", err)
	}

	if len(h.Result) == 0 {
		t.Fatal("an empty list of hosts was returned")
	}
}

func TestHostUpdate(t *testing.T) {
	templates, err := testingClient.Template.Get(&TemplateGetParameters{
		CommonGetParameters: CommonGetParameters{
			Filter: map[string][]string{
				"name": {
					"Linux by SNMP",
					"Linux by Zabbix agent",
					"Zabbix server health",
				},
			},
			Output: []string{
				"templateid",
				"name",
			},
		},
	})

	if err != nil {
		t.Fatalf("error while retrieving templates 'Linux by SNMP' / 'Linux by Zabbix agent' / 'Zabbix server health'\nReason: %v", err)
	}

	if len(templates.Result) == 0 {
		t.Fatal("error while retrieving templates 'Linux by SNMP' / 'Linux by Zabbix agent' / 'Zabbix server health'\nReason: an empty list of templates was returned")
	}

	h, err := testingClient.Host.Get(&HostGetParameters{
		CommonGetParameters: CommonGetParameters{
			Filter: map[string]string{
				"host": "Zabbix server",
			},
			Output: []string{
				"hostid",
			},
		},
	})

	if err != nil {
		t.Fatalf("error while retrieving host 'Zabbix server'\nReason: %v", err)
	}

	if len(h.Result) == 0 {
		t.Fatal("error while retrieving host 'Zabbix server'\nReason: an empty list of hosts was returned")
	}

	var snmpTemplate, agentTemplate, serverTemplate string
	for _, template := range templates.Result {
		switch template.Name {
		case "Linux by SNMP":
			snmpTemplate = template.Id
		case "Linux by Zabbix agent":
			agentTemplate = template.Id
		case "Zabbix server health":
			serverTemplate = template.Id
		default:
			t.Fatalf("unsupported template '%s' returned", template.Name)
		}
	}

	updatedHost, err := testingClient.Host.Update(&HostUpdateParameters{
		HostId:        h.Result[0].Id,
		Host:          h.Result[0].Host.Host,
		Name:          h.Result[0].Name,
		Description:   h.Result[0].Description,
		InventoryMode: HostDisabled,
		IpmiAuthType:  IpmiAuthNone,
		Status:        HostUnmonitored,
		Groups: []*HostGroupId{
			{
				Id: h.Result[0].Groups[0].Id,
			},
		},
		TlsConnect: HostNoEncryption,
		Tags: []*HostTag{
			{
				Tag:   "source",
				Value: "tests",
			},
		},
		Interfaces: []*HostInterfaceCreateParameters{
			{
				Dns:  "localhost",
				Ip:   "127.0.0.1",
				Main: HostInterfaceDefault,
				Port: "161",
				Type: HostInterfaceSNMP,
			},
		},
		Templates: []*TemplateId{
			{
				Id: snmpTemplate,
			},
		},
		TemplatesClear: []*TemplateId{
			{
				Id: agentTemplate,
			},
			{
				Id: serverTemplate,
			},
		},
		Macros: []*HostMacroCreateParamaters{
			{
				Macro:       "{$MY.MACRO}",
				Value:       "my-macro-value",
				Type:        UserMacroSecret,
				Description: "Macro created during tests",
			},
		},
		Inventory: map[HostInventory]string{
			Notes:            "Updated during tests",
			DeploymentStatus: "can-be-removed",
		},
	})

	if err != nil {
		t.Fatalf("error while executing Update function\nReason: %v", err)
	}

	if len(updatedHost.Result.Ids) == 0 {
		t.Fatalf("an empty list of hosts was returned")
	}
}

func TestHostDelete(t *testing.T) {
	g, err := testingClient.HostGroup.Get(&HostGroupGetParameters{
		CommonGetParameters: CommonGetParameters{
			Search: map[string]string{
				"name": "Virtual machines",
			},
			Output: []string{
				"groupid",
			},
		},
	})

	if err != nil {
		t.Fatalf("error while retrieving hostGroup 'Virtual machines'\nReason: %v", err)
	}

	if len(g.Result) == 0 {
		t.Fatal("error while retrieving hostGroup 'Virtual machines'\nReason: an empty list of hostGroups was returned")
	}

	name := fmt.Sprintf("test-host-%d", generateId())
	h, err := testingClient.Host.Create(&HostCreateParameters{
		Host: name,
		Groups: []*HostGroupId{
			{
				Id: g.Result[0].Id,
			},
		},
	})

	if err != nil {
		t.Fatalf("error while creating '%s' host\nReason: %v", name, err)
	}

	removedHost, err := testingClient.Host.Delete([]string{
		h.Result.Ids[0],
	})

	if err != nil {
		t.Fatalf("error while executing Delete function\nReason: %v", err)
	}

	if len(removedHost.Result.Ids) == 0 {
		t.Fatalf("an empty list of hosts was returned")
	}
}

func TestHostMassAdd(t *testing.T) {
	groups, err := testingClient.HostGroup.Get(&HostGroupGetParameters{
		CommonGetParameters: CommonGetParameters{
			Filter: map[string][]string{
				"name": {
					"Linux servers",
					"Zabbix servers",
				},
			},
			Output: []string{
				"groupid",
			},
		},
	})

	if err != nil {
		t.Fatalf("error while executing Get (HostGroup) function\nReason: %v", err)
	}

	templates, err := testingClient.Template.Get(&TemplateGetParameters{
		CommonGetParameters: CommonGetParameters{
			Filter: map[string][]string{
				"name": {
					"Linux by Zabbix agent",
					"Zabbix server health",
				},
			},
			Output: []string{
				"templateid",
			},
		},
	})

	if err != nil {
		t.Fatalf("error while executing Get (Template) function\nReason: %v", err)
	}

	hosts, err := testingClient.Host.Get(&HostGetParameters{
		CommonGetParameters: CommonGetParameters{
			Search: map[string]string{
				"name": "Zabbix server",
			},
			Output: []string{
				"hostid",
			},
		},
	})

	if err != nil {
		t.Fatalf("error while executing Get (Host) function\nReason: %v", err)
	}

	massAddedHosts, err := testingClient.Host.MassAdd(&HostMassAddParameters{
		Hosts: []*HostId{
			{
				Id: hosts.Result[0].Id,
			},
		},
		Groups: []*HostGroupId{
			{
				Id: groups.Result[0].Id,
			},
			{
				Id: groups.Result[1].Id,
			},
		},
		Interfaces: []*HostInterfaceCreateParameters{
			{
				Dns:   "",
				Ip:    "127.0.0.1",
				Main:  HostInterfaceNotDefault,
				Port:  "10050",
				Type:  HostInterfaceAgent,
				UseIp: HostInterfaceUseIp,
			},
		},
		Macros: []*HostMacroCreateParamaters{
			{
				Macro:       "{$MY.MACRO}",
				Value:       "testing-value",
				Type:        UserMacroText,
				Description: "macro created during tests",
			},
		},
		Templates: []*TemplateId{
			{
				Id: templates.Result[0].Id,
			},
			{
				Id: templates.Result[1].Id,
			},
		},
	})

	if err != nil {
		t.Fatalf("error while executing MassAdd function\nReason: %v", err)
	}

	if len(massAddedHosts.Result.Ids) == 0 {
		t.Fatalf("an empty list of HostGroups was returned")
	}
}

func TestHostMassUpdate(t *testing.T) {
	groups, err := testingClient.HostGroup.Get(&HostGroupGetParameters{
		CommonGetParameters: CommonGetParameters{
			Filter: map[string][]string{
				"name": {
					"Linux servers",
					"Zabbix servers",
				},
			},
			Output: []string{
				"groupid",
			},
		},
	})

	if err != nil {
		t.Fatalf("error while executing Get (HostGroup) function\nReason: %v", err)
	}

	templates, err := testingClient.Template.Get(&TemplateGetParameters{
		CommonGetParameters: CommonGetParameters{
			Filter: map[string][]string{
				"name": {
					"Linux by Zabbix agent",
					"Zabbix server health",
				},
			},
			Output: []string{
				"templateid",
			},
		},
	})

	if err != nil {
		t.Fatalf("error while executing Get (Template) function\nReason: %v", err)
	}

	hostName := fmt.Sprintf("testing-host-%d", generateId())
	hosts, err := testingClient.Host.Create(&HostCreateParameters{
		Host: hostName,
		Groups: []*HostGroupId{
			{
				groups.Result[0].Id,
			},
		},
		Templates: []*TemplateId{
			{
				Id: templates.Result[0].Id,
			},
		},
		Interfaces: []*HostInterfaceCreateParameters{
			{
				Dns:   "localhost",
				Ip:    "",
				Main:  HostInterfaceDefault,
				Port:  "10050",
				Type:  HostInterfaceAgent,
				UseIp: HostInterfaceUseDns,
			},
		},
	})

	if err != nil {
		t.Fatalf("error while executing Create (Host) function\nReason: %v", err)
	}

	massUpdatedHosts, err := testingClient.Host.MassUpdate(&HostMassUpdateParameters{
		Hosts: []*HostId{
			{
				Id: hosts.Result.Ids[0],
			},
		},
		Host:          fmt.Sprintf("%s-updated", hostName),
		Description:   "updated description",
		InventoryMode: HostManual,
		Status:        HostUnmonitored,
		Inventory: map[HostInventory]string{
			Type:             "docker",
			DeploymentStatus: "can-be-removed",
		},
		Groups: []*HostGroupId{
			{
				Id: groups.Result[1].Id,
			},
		},
		Interfaces: []*HostInterfaceCreateParameters{
			{
				Dns:   "",
				Ip:    "127.0.0.1",
				Main:  HostInterfaceDefault,
				Port:  "161",
				Type:  HostInterfaceSNMP,
				UseIp: HostInterfaceUseIp,
				Details: &HostInterfaceDetail{
					Version:        HostInterfaceSnmpV3,
					Bulk:           HostInterfaceUseSnmpBulk,
					Community:      "",
					SecurityName:   "public",
					ContextName:    "",
					SecurityLevel:  HostInterfaceAuthPriv,
					AuthPassPhrase: "public",
					PrivPassPhrase: "public",
					AuthProtocol:   HostInterfaceSHA512,
					PrivProtocol:   HostInterfaceAES256,
				},
			},
		},
		Macros: []*HostMacroCreateParamaters{
			{
				Macro:       "{$MY.MACRO}",
				Value:       "testing-value",
				Type:        UserMacroText,
				Description: "macro created during tests",
			},
		},
		Templates: []*TemplateId{
			{
				Id: templates.Result[1].Id,
			},
		},
		TemplatesClear: []*TemplateId{
			{
				Id: templates.Result[0].Id,
			},
		},
	})

	if err != nil {
		t.Fatalf("error while executing MassUpdate function\nReason: %v", err)
	}

	if len(massUpdatedHosts.Result.Ids) == 0 {
		t.Fatalf("an empty list of HostGroups was returned")
	}
}

func TestHostGroupMassRemove(t *testing.T) {
	groups, err := testingClient.HostGroup.Get(&HostGroupGetParameters{
		CommonGetParameters: CommonGetParameters{
			Filter: map[string][]string{
				"name": {
					"Linux servers",
					"Zabbix servers",
				},
			},
			Output: []string{
				"groupid",
			},
		},
	})

	if err != nil {
		t.Fatalf("error while executing Get (HostGroup) function\nReason: %v", err)
	}

	templates, err := testingClient.Template.Get(&TemplateGetParameters{
		CommonGetParameters: CommonGetParameters{
			Filter: map[string][]string{
				"name": {
					"Linux by Zabbix agent",
					"Zabbix server health",
				},
			},
			Output: []string{
				"templateid",
			},
		},
	})

	if err != nil {
		t.Fatalf("error while executing Get (Template) function\nReason: %v", err)
	}

	hosts, err := testingClient.Host.Create(&HostCreateParameters{
		Host: fmt.Sprintf("testing-host-%d", generateId()),
		Groups: []*HostGroupId{
			{
				groups.Result[0].Id,
			},
			{
				groups.Result[1].Id,
			},
		},
		Templates: []*TemplateId{
			{
				Id: templates.Result[0].Id,
			},
			{
				Id: templates.Result[1].Id,
			},
		},
		Interfaces: []*HostInterfaceCreateParameters{
			{
				Dns:   "localhost",
				Ip:    "",
				Main:  HostInterfaceDefault,
				Port:  "10050",
				Type:  HostInterfaceAgent,
				UseIp: HostInterfaceUseDns,
			},
			{
				Dns:   "",
				Ip:    "127.0.0.1",
				Main:  HostInterfaceNotDefault,
				Port:  "10051",
				Type:  HostInterfaceAgent,
				UseIp: HostInterfaceUseIp,
			},
		},
		Macros: []*HostMacroCreateParamaters{
			{
				Macro:       "{$MY.MACRO}",
				Value:       "testing-value",
				Type:        UserMacroText,
				Description: "macro created during tests",
			},
		},
	})

	if err != nil {
		t.Fatalf("error while executing Create (Host) function\nReason: %v", err)
	}

	massRemovedHosts, err := testingClient.Host.MassRemove(&HostMassRemoveParameters{
		HostIds: []string{
			hosts.Result.Ids[0],
		},
		GroupIds: []string{
			groups.Result[1].Id,
		},
		Interfaces: []*HostMassRemoveInterfaces{
			{
				Ip:   "127.0.0.1",
				Dns:  "",
				Port: "10051",
			},
		},
		Macros: []string{
			"{$MY.MACRO}",
		},
		TemplateIds: []string{
			templates.Result[0].Id,
		},
		TemplateIdsClear: []string{
			templates.Result[1].Id,
		},
	})

	if err != nil {
		t.Fatalf("error while executing MassRemove function\nReason: %v", err)
	}

	if len(massRemovedHosts.Result.Ids) == 0 {
		t.Fatalf("an empty list of HostGroups was returned")
	}
}
