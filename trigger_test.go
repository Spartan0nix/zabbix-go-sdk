package zabbixgosdk

import (
	"testing"
)

const (
	triggerDescription          = "test-trigger"
	triggerExpression           = "last(/Zabbix server/vm.memory.size[pavailable])<20"
	updatedTriggerDescription   = "test-trigger-update"
	dependentTriggerDescription = "test-trigger-dependent"
	dependentTriggerExpression  = "last(/Zabbix server/system.swap.size[,pfree])<80"
)

var triggerId string
var dependentTriggerId string

func TestCreateTrigger(t *testing.T) {
	trigger, err := testingClient.Trigger.Create(&TriggerCreateParameters{
		Description: triggerDescription,
		Expression:  triggerExpression,
		EventName:   triggerDescription,
		Comments:    "Testing Trigger Create() function",
		Priority:    TriggerHigh,
		Status:      TriggerDisable,
		Type:        TriggerSingleEvent,
	})

	if err != nil {
		t.Fatalf("Error when creating trigger '%s'.\nReason : %v", triggerDescription, err)
	}

	if trigger == nil || len(trigger.TriggerIds) == 0 {
		t.Fatalf("Create method should return a list of the created triggers.\nAn empty list was returned.")
	}

	triggerId = trigger.TriggerIds[0]

	trigger, err = testingClient.Trigger.Create(&TriggerCreateParameters{
		Description: dependentTriggerDescription,
		Expression:  dependentTriggerExpression,
	})

	if err != nil {
		t.Fatalf("Error when creating trigger '%s'.\nReason : %v", dependentTriggerDescription, err)
	}

	if trigger == nil || len(trigger.TriggerIds) == 0 {
		t.Fatalf("Create method should return a list of the created triggers.\nAn empty list was returned.")
	}

	dependentTriggerId = trigger.TriggerIds[0]
}

func TestTriggerGet(t *testing.T) {
	trigger, err := testingClient.Trigger.Get(&TriggerGetParameters{
		TriggerIds: []string{
			triggerId,
		},
		MinSeverity: TriggerHigh,
	})

	if err != nil {
		t.Fatalf("Error when getting trigger '%s'.\nReason : %v", triggerId, err)
	}

	if len(trigger) == 0 {
		t.Fatalf("Get method should return a list of triggers matching the given criteria.\nAn empty list was returned.")
	}

	if trigger[0].Id != triggerId {
		t.Fatalf("Wrong trigger returned.\nExpected : %s\nReturned : %s", triggerId, trigger[0].Id)
	}
}

func TestTriggerUpdate(t *testing.T) {
	trigger, err := testingClient.Trigger.Update(&TriggerUpdateParameters{
		Id:          triggerId,
		Description: updatedTriggerDescription,
		Priority:    TriggerAverage,
	})

	if err != nil {
		t.Fatalf("Error when updating trigger '%s'.\nReason : %v", triggerId, err)
	}

	if trigger == nil || len(trigger.TriggerIds) == 0 {
		t.Fatalf("Update method should return a list of the updated triggers.\nAn empty list was returned.")
	}

	if trigger.TriggerIds[0] != triggerId {
		t.Fatalf("Wrong trigger returned.\nExpected : %s\nReturned : %s", triggerId, trigger.TriggerIds[0])
	}
}

func TestTriggerDelete(t *testing.T) {
	trigger, err := testingClient.Trigger.Delete([]string{
		triggerId,
	})

	if err != nil {
		t.Fatalf("Error when deleting trigger '%s'.\nReason : %v", triggerId, err)
	}

	if trigger == nil || len(trigger.TriggerIds) == 0 {
		t.Fatalf("Delete method should return a list of the deleted triggers.\nAn empty list was returned.")
	}

	if trigger.TriggerIds[0] != triggerId {
		t.Fatalf("Wrong trigger returned.\nExpected : %s\nReturned : %s", triggerId, trigger.TriggerIds[0])
	}
}
