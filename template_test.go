package main

import "testing"

const (
	templateName        = "test-template"
	updatedTemplateName = "test-template-update"
)

var templateId string
var apacheTemplateId string

func TestTemplateCreate(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	template, err := client.Template.Create(&TemplateCreateParameters{
		Host: templateName,
		Groups: []*TemplateGroup{
			{
				Id: "1",
			},
		},
	})

	if err != nil {
		t.Fatalf("Error when creating template '%s'.\nReason : %v", templateName, err)
	}

	if template == nil {
		t.Fatalf("Create method should returned a list with the Id of the new template.\nAn empty list was returned.")
	}

	templateId = template.Templateids[0]
}

func TestTemplateList(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	template, err := client.Template.List()

	if err != nil {
		t.Fatalf("Error when listing templates.\nReason : %v", err)
	}

	if template == nil {
		t.Fatalf("List method should returned a list of existing templates.\nAn empty list was returned.")
	}
}

func TestTemplateGet(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	template, err := client.Template.Get(&TemplateGetParameters{
		Filter: map[string]string{
			"name": templateName,
		},
	})

	if err != nil {
		t.Fatalf("Error when getting template '%s'.\nReason : %v", templateName, err)
	}

	if template == nil {
		t.Fatalf("Get method should returned a list of templates matching the given criteria.\nAn empty list was returned.")
	}

	if template[0].Host != templateName {
		t.Fatalf("Wrong host returned.\nExpected name : %s\nName returned : %s", templateName, template[0].Host)
	}

	if template[0].Id != templateId {
		t.Fatalf("Wrong host returned.\nExpected Id : %s\nId returned : %s", templateId, template[0].Id)
	}
}

func TestTemplateMassAdd(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	template, err := client.Template.Get(&TemplateGetParameters{
		Filter: map[string]string{
			"name": "Apache by HTTP",
		},
	})

	if template == nil {
		t.Fatalf("Error when getting template 'Apache by HTTP'.\nAn empty list was returned.")
	}

	apacheTemplateId = template[0].Id

	if err != nil {
		t.Fatalf("Error when getting template 'Apache by HTTP'.\nReason : %v", err)
	}

	mass_add_t, err := client.Template.MassAdd(&TemplateMassAddParameters{
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
		t.Fatalf("Error when mass adding template.\nReason : %v", err)
	}

	if mass_add_t == nil {
		t.Fatal("MassAdd method should returned a list of the updated templates.\nAn empty list was returned.")
	}
}

func TestTemplateMassRemove(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	mass_rm_t, err := client.Template.MassRemove(&TemplateMassRemoveParameters{
		TemplateIds: []string{
			templateId,
		},
		TemplateIdsClear: []string{
			apacheTemplateId,
		},
	})

	if err != nil {
		t.Fatalf("Error when mass removing template.\nTemplate : %s\nTemplate to unlink : %s \nReason : %v", templateId, apacheTemplateId, err)
	}

	if mass_rm_t == nil {
		t.Fatal("MassRemove method should returned a list of the updated templates.\nAn empty list was returned.")
	}
}

func TestTemplateMassUpdate(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	mass_rm_t, err := client.Template.MassUpdate(&TemplateMassUpdateParameters{
		Templates: []*TemplateId{
			{
				Id: templateId,
			},
		},
		Macros: []*HostMacro{
			{
				Macro: "{$MASS_UPDATE_TEMPLATE_MACRO}",
				Value: "massUpdate-test-template-value",
			},
		},
	})

	if err != nil {
		t.Fatalf("Error when mass updating template.\nReason : %v", err)
	}

	if mass_rm_t == nil {
		t.Fatal("MassUpdate method should returned a list of the updated templates.\nAn empty list was returned.")
	}
}

func TestTemplateUpdate(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	updated_t, err := client.Template.Update(&TemplateUpdateParameters{
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
			},
		},
	})

	if err != nil {
		t.Fatalf("Error when updating template '%s'.\nReason : %v", templateId, err)
	}

	if updated_t == nil {
		t.Fatal("Update method should returned a list of with the updated templates id.\nAn empty list was returned.")
	}
}

func TestTemplateDelete(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	deleted_t, err := client.Template.Delete([]string{
		templateId,
	})

	if err != nil {
		t.Fatalf("Error when deleting template '%s'.\nReason : %v", templateId, err)
	}

	if deleted_t == nil {
		t.Fatalf("Delete method should returned a list with the Ids of the deleted templates.\nAn empty list was returned.")
	}
}
