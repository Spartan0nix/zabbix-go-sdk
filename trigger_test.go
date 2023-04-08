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
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	trigger, err := client.Trigger.Create(&TriggerCreateParameters{
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
		t.Fatalf("Create method should returned a list with the id of the new trigger.\nAn empty list was returned.")
	}

	triggerId = trigger.TriggerIds[0]

	trigger, err = client.Trigger.Create(&TriggerCreateParameters{
		Description: dependentTriggerDescription,
		Expression:  dependentTriggerExpression,
	})

	if err != nil {
		t.Fatalf("Error when creating trigger '%s'.\nReason : %v", dependentTriggerDescription, err)
	}

	if trigger == nil || len(trigger.TriggerIds) == 0 {
		t.Fatalf("Create method should returned a list with the id of the new trigger.\nAn empty list was returned.")
	}

	dependentTriggerId = trigger.TriggerIds[0]
}

func TestTriggerGet(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	trigger, err := client.Trigger.Get(&TriggerGetParameters{
		TriggerIds: []string{
			triggerId,
		},
		MinSeverity: TriggerHigh,
	})

	if err != nil {
		t.Fatalf("Error when getting trigger '%s'.\nReason : %v", triggerId, err)
	}

	if len(trigger) == 0 {
		t.Fatalf("Get method should returned a list of trigger matching the given criteria.\nAn empty list was returned.")
	}

	if trigger[0].Id != triggerId {
		t.Fatalf("Wrong trigger returned.\nExpected : %s\nReturned : %s", triggerId, trigger[0].Id)
	}
}

func TestTriggerUpdate(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	trigger, err := client.Trigger.Update(&TriggerUpdateParameters{
		Id:          triggerId,
		Description: updatedTriggerDescription,
		Priority:    TriggerAverage,
	})

	if err != nil {
		t.Fatalf("Error when updating trigger '%s'.\nReason : %v", triggerId, err)
	}

	if trigger == nil || len(trigger.TriggerIds) == 0 {
		t.Fatalf("Update method should returned a list with the id of the updated trigger.\nAn empty list was returned.")
	}

	if trigger.TriggerIds[0] != triggerId {
		t.Fatalf("Wrong trigger returned.\nExpected : %s\nReturned : %s", triggerId, trigger.TriggerIds[0])
	}
}

func TestTriggerDelete(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	trigger, err := client.Trigger.Delete([]string{
		triggerId,
	})

	if err != nil {
		t.Fatalf("Error when deleting trigger '%s'.\nReason : %v", triggerId, err)
	}

	if trigger == nil || len(trigger.TriggerIds) == 0 {
		t.Fatalf("Delete method should returned a list with the id of the deleted trigger.\nAn empty list was returned.")
	}

	if trigger.TriggerIds[0] != triggerId {
		t.Fatalf("Wrong trigger returned.\nExpected : %s\nReturned : %s", triggerId, trigger.TriggerIds[0])
	}
}
