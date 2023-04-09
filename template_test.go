package zabbixgosdk

import "testing"

const (
	templateName        = "test-template"
	updatedTemplateName = "test-template-update"
)

var templateId string
var apacheTemplateId string

func TestTemplateCreate(t *testing.T) {
	template, err := testingClient.Template.Create(&TemplateCreateParameters{
		Host: templateName,
		Groups: []*HostGroupId{
			{
				GroupId: "1",
			},
		},
	})

	if err != nil {
		t.Fatalf("Error when creating template '%s'.\nReason : %v", templateName, err)
	}

	if template == nil || len(template.TemplateIds) == 0 {
		t.Fatalf("Create method should return a list of the created templates.\nAn empty list was returned.")
	}

	templateId = template.TemplateIds[0]
}

func TestTemplateList(t *testing.T) {
	template, err := testingClient.Template.List()
	if err != nil {
		t.Fatalf("Error when listing templates.\nReason : %v", err)
	}

	if len(template) == 0 {
		t.Fatalf("List method should return a list with all the existing templates on the server.\nAn empty list was returned.")
	}
}

func TestTemplateGet(t *testing.T) {
	template, err := testingClient.Template.Get(&TemplateGetParameters{
		TemplateIds: []string{
			templateId,
		},
	})

	if err != nil {
		t.Fatalf("Error when getting template '%s'.\nReason : %v", templateName, err)
	}

	if len(template) == 0 {
		t.Fatalf("Get method should return a list of templates matching the given criteria.\nAn empty list was returned.")
	}

	if template[0].TemplateId != templateId {
		t.Fatalf("Wrong template returned.\nExpected : %s\nReturned : %s", templateId, template[0].TemplateId)
	}
}

func TestTemplateMassAdd(t *testing.T) {
	apacheTemplate, err := testingClient.Template.Get(&TemplateGetParameters{
		Filter: map[string]string{
			"name": "Apache by HTTP",
		},
	})

	if err != nil {
		t.Fatalf("Error when getting template 'Apache by HTTP'.\nReason : %v", err)
	}

	if len(apacheTemplate) == 0 {
		t.Fatalf("Error when getting template 'Apache by HTTP'.\nAn empty list was returned.")
	}

	apacheTemplateId = apacheTemplate[0].TemplateId

	template, err := testingClient.Template.MassAdd(&TemplateMassAddParameters{
		Templates: []*TemplateId{
			{
				Id: templateId,
			},
		},
		TemplatesLink: []*TemplateId{
			{
				Id: apacheTemplateId,
			},
		},
		Macros: []*HostMacro{
			{
				Macro: "{$TEST_TEMPLATE_MACRO}",
				Value: "test-template-value",
			},
		},
	})

	if err != nil {
		t.Fatalf("Error when massadding properties to template '%s'.\nReason : %v", templateId, err)
	}

	if template == nil || len(template.TemplateIds) == 0 {
		t.Fatal("MassAdd method should return a list of the updated templates.\nAn empty list was returned.")
	}

	if template.TemplateIds[0] != templateId {
		t.Fatalf("Wrong template returned.\nExpected : %s\nReturned : %s", templateId, template.TemplateIds[0])
	}
}

func TestTemplateMassRemove(t *testing.T) {
	template, err := testingClient.Template.MassRemove(&TemplateMassRemoveParameters{
		TemplateIds: []string{
			templateId,
		},
		TemplateIdsClear: []string{
			apacheTemplateId,
		},
	})

	if err != nil {
		t.Fatalf("Error when massremoving properties from template '%s'.\nReason : %v", templateId, err)
	}

	if template == nil || len(template.TemplateIds) == 0 {
		t.Fatal("MassRemove method should return a list of the updated templates.\nAn empty list was returned.")
	}

	if template.TemplateIds[0] != templateId {
		t.Fatalf("Wrong template returned.\nExpected : %s\nReturned : %s", templateId, template.TemplateIds[0])
	}
}

func TestTemplateMassUpdate(t *testing.T) {
	template, err := testingClient.Template.MassUpdate(&TemplateMassUpdateParameters{
		Templates: []*TemplateId{
			{
				Id: templateId,
			},
		},
		Macros: []*HostMacro{
			{
				Macro: "{$MASS_UPDATE_TEMPLATE_MACRO}",
				Value: "massUpdate-test-template-value",
				Type:  Text,
			},
		},
	})

	if err != nil {
		t.Fatalf("Error when massupdating properties from template '%s'.\nReason : %v", templateId, err)
	}

	if template == nil || len(template.TemplateIds) == 0 {
		t.Fatal("MassUpdate method should return a list of the updated templates.\nAn empty list was returned.")
	}

	if template.TemplateIds[0] != templateId {
		t.Fatalf("Wrong template returned.\nExpected : %s\nReturned : %s", templateId, template.TemplateIds[0])
	}
}

func TestTemplateUpdate(t *testing.T) {
	template, err := testingClient.Template.Update(&TemplateUpdateParameters{
		Id: templateId,
		TemplateLink: []*TemplateId{
			{
				Id: apacheTemplateId,
			},
		},
		Macros: []*HostMacro{
			{
				Macro: "{$UPDATED_TEMPLATE_MACRO}",
				Value: "updated-test-template-value",
				Type:  Secret,
			},
		},
	})

	if err != nil {
		t.Fatalf("Error when updating template '%s'.\nReason : %v", templateId, err)
	}

	if template == nil || len(template.TemplateIds) == 0 {
		t.Fatal("Update method should return a list of the updated templates.\nAn empty list was returned.")
	}

	if template.TemplateIds[0] != templateId {
		t.Fatalf("Wrong template returned.\nExpected : %s\nReturned : %s", templateId, template.TemplateIds[0])
	}
}

func TestTemplateDelete(t *testing.T) {
	template, err := testingClient.Template.Delete([]string{
		templateId,
	})

	if err != nil {
		t.Fatalf("Error when deleting template '%s'.\nReason : %v", templateId, err)
	}

	if template == nil || len(template.TemplateIds) == 0 {
		t.Fatalf("Delete method should return a list of the deleted templates.\nAn empty list was returned.")
	}

	if template.TemplateIds[0] != templateId {
		t.Fatalf("Wrong template returned.\nExpected Id : %s\nId returned : %s", templateId, template.TemplateIds[0])
	}
}
