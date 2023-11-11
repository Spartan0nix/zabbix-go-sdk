package zabbixgosdk

import (
	"fmt"
	"testing"
)

func TestTemplateCreate(t *testing.T) {
	sysHostGroup, err := testingClient.HostGroup.Get(&HostGroupGetParameters{
		CommonGetParameters: CommonGetParameters{
			Search: map[string]string{
				"name": "Templates",
			},
			Output: "groupid",
		},
	})

	if err != nil {
		t.Fatalf("error while retrieving the 'Templates' host group\nReason: %v", err)
	}

	template, err := testingClient.Template.Create(&TemplateCreateParameters{
		Host: fmt.Sprintf("test-template-%d", generateId()),
		Name: "testing-template",
		Tags: []*TemplateTag{
			{
				Tag:   "source",
				Value: "tests",
			},
		},
		Description: "my-description",
		Groups: []*HostGroupId{
			{
				Id: sysHostGroup.Result[0].Id,
			},
		},
		Macros: []*HostMacroCreateParamaters{
			{
				Macro:       "{$TESTS.MACRO}",
				Value:       "testing-macro",
				Type:        UserMacroText,
				Description: "macro added during tests",
			},
		},
	})

	if err != nil {
		t.Fatalf("error while executing Create function\nReason: %v", err)
	}

	if template == nil || len(template.Result.Ids) == 0 {
		t.Fatal("an empty list of template ids was returned")
	}
}

func TestTemplateGet(t *testing.T) {
	template, err := testingClient.Template.Get(&TemplateGetParameters{
		CommonGetParameters: CommonGetParameters{
			Search: map[string]string{
				"name": "Linux by SNMP",
			},
		},
		SelectGroups:          "extend",
		SelectTags:            "extend",
		SelectHosts:           "extend",
		SelectTemplates:       "extend",
		SelectParentTemplates: "extend",
		SelectHttpTests:       "extend",
		SelectItems:           "extend",
		SelectDiscoveries:     "extend",
		SelectTriggers:        "extend",
		SelectGraphs:          "extend",
		SelectMacros:          "extend",
		SelectDashboards:      "extend",
		SelectValueMaps:       "extend",
	})

	if err != nil {
		t.Fatalf("error while executing Get function\nReason: %v", err)
	}

	if len(template.Result) == 0 {
		t.Fatalf("an empty list of Templates was returned")
	}
}

func TestTemplateUpdate(t *testing.T) {
	sysTemplate, err := testingClient.Template.Get(&TemplateGetParameters{
		CommonGetParameters: CommonGetParameters{
			Search: map[string]string{
				"name": "Linux by Zabbix agent active*",
			},
			SearchWildcardsEnabled: true,
		},
	})

	if err != nil {
		t.Fatalf("error while retrieving the 'Linux by Zabbix agent active*' template\nReason: %v", err)
	}

	id := generateId()
	groups := make([]*HostGroupId, len(sysTemplate.Result[0].HostGroups))
	for _, g := range sysTemplate.Result[0].HostGroups {
		groups = append(groups, &HostGroupId{
			Id: g.Id,
		})
	}

	macros := make([]*HostMacroUpdateParamaters, len(sysTemplate.Result[0].Macros))
	for _, m := range sysTemplate.Result[0].Macros {
		macros = append(macros, &HostMacroUpdateParamaters{
			Macro:       m.Macro,
			Value:       m.Value,
			Type:        m.Type,
			Description: m.Description,
		})
	}

	template, err := testingClient.Template.Update(&TemplateUpdateParameters{
		Id:          sysTemplate.Result[0].Id,
		Host:        fmt.Sprintf("Linux by Zabbix agent active - %d", id),
		Description: fmt.Sprintf("updated description - %d", id),
		Groups:      groups,
		Tags:        sysTemplate.Result[0].Tags,
		Macros:      macros,
	})

	if err != nil {
		t.Fatalf("error while executing Update function\nReason: %v", err)
	}

	if len(template.Result.Ids) == 0 {
		t.Fatalf("an empty list of Templates was returned")
	}
}

func TestTemplateDelete(t *testing.T) {
	sysHostGroup, err := testingClient.HostGroup.Get(&HostGroupGetParameters{
		CommonGetParameters: CommonGetParameters{
			Search: map[string]string{
				"name": "Templates",
			},
			Output: "groupid",
		},
	})

	if err != nil {
		t.Fatalf("error while retrieving the 'Templates' host group\nReason: %v", err)
	}

	name := fmt.Sprintf("test-template-%d", generateId())
	template, err := testingClient.Template.Create(&TemplateCreateParameters{
		Host: name,
		Groups: []*HostGroupId{
			{
				Id: sysHostGroup.Result[0].Id,
			},
		},
	})

	if err != nil {
		t.Fatalf("error while creating '%s' template\nReason: %v", name, err)
	}

	template, err = testingClient.Template.Delete([]string{
		template.Result.Ids[0],
	})

	if err != nil {
		t.Fatalf("error while executing Delete function\nReason: %v", err)
	}

	if len(template.Result.Ids) == 0 {
		t.Fatalf("an empty list of Templates was returned")
	}
}

func TestTemplateMassAdd(t *testing.T) {
	sysTemplate, err := testingClient.Template.Get(&TemplateGetParameters{
		CommonGetParameters: CommonGetParameters{
			Search: map[string]string{
				"name": "Linux by Zabbix agent active*",
			},
			SearchWildcardsEnabled: true,
		},
	})

	if err != nil {
		t.Fatalf("error while retrieving the 'Linux by Zabbix agent active*' template\nReason: %v", err)
	}

	sysHostGroup, err := testingClient.HostGroup.Get(&HostGroupGetParameters{
		CommonGetParameters: CommonGetParameters{
			Search: map[string]string{
				"name": "Templates",
			},
			Output: "groupid",
		},
	})

	if err != nil {
		t.Fatalf("error while retrieving the 'Templates' host group\nReason: %v", err)
	}

	template, err := testingClient.Template.MassAdd(&TemplateMassAddParameters{
		Templates: []*TemplateId{
			{
				Id: sysTemplate.Result[0].Id,
			},
		},
		Groups: []*HostGroupId{
			{
				Id: sysHostGroup.Result[0].Id,
			},
		},
		Macros: []*HostMacroCreateParamaters{
			{
				Macro:       "{$TESTS.MACRO}",
				Value:       "testing-macro",
				Type:        UserMacroText,
				Description: "macro added during tests",
			},
		},
	})

	if err != nil {
		t.Fatalf("error while executing MassAdd function\nReason: %v", err)
	}

	if len(template.Result.Ids) == 0 {
		t.Fatal("an empty list of Templates was returned")
	}
}

func TestTemplateMassUpdate(t *testing.T) {
	sysTemplate, err := testingClient.Template.Get(&TemplateGetParameters{
		CommonGetParameters: CommonGetParameters{
			Search: map[string]string{
				"name": "Linux by Zabbix agent active*",
			},
			SearchWildcardsEnabled: true,
		},
	})

	if err != nil {
		t.Fatalf("error while retrieving the 'Linux by Zabbix agent active*' template\nReason: %v", err)
	}

	groups := make([]*HostGroupId, len(sysTemplate.Result[0].HostGroups))
	for _, g := range sysTemplate.Result[0].HostGroups {
		groups = append(groups, &HostGroupId{
			g.Id,
		})
	}

	macros := make([]*HostMacroUpdateParamaters, len(sysTemplate.Result[0].Macros))
	for _, m := range sysTemplate.Result[0].Macros {
		macros = append(macros, &HostMacroUpdateParamaters{
			Macro:       m.Macro,
			Value:       m.Value,
			Type:        m.Type,
			Description: m.Description,
		})
	}

	template, err := testingClient.Template.MassUpdate(&TemplateMassUpdateParameters{
		Templates: []*TemplateId{
			{
				Id: sysTemplate.Result[0].Id,
			},
		},
		Groups: groups,
		Macros: macros,
	})

	if err != nil {
		t.Fatalf("error while executing MassUpdate function\nReason: %v", err)
	}

	if len(template.Result.Ids) == 0 {
		t.Fatal("an empty list of Templates was returned")
	}
}

func TestTemplateMassRemove(t *testing.T) {
	sysTemplate, err := testingClient.Template.Get(&TemplateGetParameters{
		CommonGetParameters: CommonGetParameters{
			Search: map[string]string{
				"name": "Linux by Zabbix agent active*",
			},
			SearchWildcardsEnabled: true,
		},
		SelectMacros: []string{
			"macro",
		},
	})

	if err != nil {
		t.Fatalf("error while retrieving the 'Linux by Zabbix agent active*' template\nReason: %v", err)
	}

	template, err := testingClient.Template.MassRemove(&TemplateMassRemoveParameters{
		TemplateIds: []string{
			sysTemplate.Result[0].Id,
		},
		Macros: []string{
			sysTemplate.Result[0].Macros[0].Macro,
		},
	})

	if err != nil {
		t.Fatalf("error while executing MassRemove function\nReason: %v", err)
	}

	if len(template.Result.Ids) == 0 {
		t.Fatal("an empty list of Templates was returned")
	}
}
