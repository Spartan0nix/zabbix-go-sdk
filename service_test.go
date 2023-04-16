package zabbixgosdk

import "testing"

const (
	serviceName       = "test-service"
	updateServiceName = "updated-test-service"
)

var serviceId string

func TestServiceCreate(t *testing.T) {
	s, err := testingClient.Service.Create(&ServiceCreateParameters{
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
		t.Fatalf("Error when creating service '%s'.\nReason : %v", serviceName, err)
	}

	if s == nil || len(s.ServiceIds) == 0 {
		t.Fatalf("Create method should return a list of the created services.\nAn empty list was returned.")
	}

	serviceId = s.ServiceIds[0]
}

func TestServiceGet(t *testing.T) {
	s, err := testingClient.Service.Get(&ServiceGetParameters{
		ServiceIds: []string{
			serviceId,
		},
	})

	if err != nil {
		t.Fatalf("Error when getting service '%s'.\nReason : %v", serviceId, err)
	}

	if len(s) == 0 {
		t.Fatalf("Get method should return a list of services matching the given criteria.\nAn empty list was returned.")
	}

	if s[0].ServiceId != serviceId {
		t.Fatalf("Wrong service returned.\nExpected : %s\nReturned : %s", serviceId, s[0].ServiceId)
	}
}

func TestServiceUpdate(t *testing.T) {
	s, err := testingClient.Service.Update(&ServiceUpdateParameters{
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

	if s == nil || len(s.ServiceIds) == 0 {
		t.Fatal("Update method should return a list of the updated services.\nAn empty list was returned.")
	}

	if s.ServiceIds[0] != serviceId {
		t.Fatalf("Wrong service returned.\nExpected Id : %s\nId returned : %s", serviceId, s.ServiceIds[0])
	}
}

func TestServiceDelete(t *testing.T) {
	s, err := testingClient.Service.Delete([]string{
		serviceId,
	})

	if err != nil {
		t.Fatalf("Error when deleting service '%s'.\nReason : %v", serviceId, err)
	}

	if s == nil || len(s.ServiceIds) == 0 {
		t.Fatal("Delete method should return a list of the deleted services.\nAn empty list was returned.")
	}

	if s.ServiceIds[0] != serviceId {
		t.Fatalf("Wrong service returned.\nExpected : %s\nReturned : %s", serviceId, s.ServiceIds[0])
	}
}
