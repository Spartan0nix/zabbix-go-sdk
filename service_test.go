package zabbixgosdk

import "testing"

const (
	serviceName       = "test-service"
	updateServiceName = "updated-test-service"
)

var serviceId string

func TestServiceCreate(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	s, err := client.Service.Create(&ServiceCreateParameters{
		Algorithm:        ServiceAlgorithmMostCriticalAllChildren,
		Name:             serviceName,
		SortOrder:        "0",
		Weight:           "0",
		PropagationRule:  ServicePropagationRuleAsIs,
		PropagationValue: ServicePropagationValueNotClassified,
		Tags: []*ServiceTag{
			{
				Tag:   "service",
				Value: "root",
			},
		},
		ProblemTags: []*ServiceProblemTag{
			{
				Tag:      "scope",
				Operator: ServiceOperatorEqual,
				Value:    "performance",
			},
		},
	})

	if err != nil {
		t.Fatalf("Error when creating service.\nReason : %v", err)
	}

	if s == nil {
		t.Fatalf("Create method should returned a list with the id of the services.\nAn empty list was returned.")
	}

	serviceId = s.ServiceIds[0]
}

func TestServiceGet(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	s, err := client.Service.Get(&ServiceGetParameters{
		ServiceIds: []string{
			serviceId,
		},
	})

	if err != nil {
		t.Fatalf("Error when retrieving service '%s'.\nReason : %v", serviceId, err)
	}

	if s == nil {
		t.Fatalf("Get method should returned a list of services matching the given criteria.\nAn empty list was returned.")
	}

	if s[0].ServiceId != serviceId {
		t.Fatalf("Wrong service returned.\nExpected Id : %s\nId returned : %s", serviceId, s[0].ServiceId)
	}
}

func TestServiceUpdate(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	s, err := client.Service.Update(&ServiceUpdateParameters{
		ServiceId: serviceId,
		Name:      updateServiceName,
		Tags: []*ServiceTag{
			{
				Tag:   "service",
				Value: "updated-root",
			},
		},
	})

	if err != nil {
		t.Fatalf("Error when updating service '%s'.\nReason : %v", serviceId, err)
	}

	if s == nil {
		t.Fatal("Update method should returned a list of the updated service.\nAn empty list was returned.")
	}

	if s.ServiceIds[0] != serviceId {
		t.Fatalf("Wrong service returned.\nExpected Id : %s\nId returned : %s", serviceId, s.ServiceIds[0])
	}
}

func TestServiceDelete(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	s, err := client.Service.Delete([]string{
		serviceId,
	})

	if err != nil {
		t.Fatalf("Error when deleting service '%s'.\nReason : %v", serviceId, err)
	}

	if s == nil {
		t.Fatalf("Wrong service returned.\nExpected Id : %s\nId returned : %s", serviceId, s.ServiceIds[0])
	}

	if s.ServiceIds[0] != serviceId {
		t.Fatalf("Wrong service returned.\nExpected Id : %s\nId returned : %s", serviceId, s.ServiceIds[0])
	}
}
