package zabbixgosdk

import (
	"fmt"
	"testing"
)

func TestHostGroupCreate(t *testing.T) {
	g, err := testingClient.HostGroup.Create(fmt.Sprintf("test-host-group-%d", generateId()))
	if err != nil {
		t.Fatalf("error while executing Create function\nReason: %v", err)
	}

	if g == nil || len(g.Result.Ids) == 0 {
		t.Fatal("an empty list of hostGroup ids was returned")
	}
}

func BenchmarkHostGroupCreate(b *testing.B) {
	var name string
	for i := 0; i < b.N; i++ {
		name = fmt.Sprintf("test-host-group-%d", generateId())

		_, err := testingClient.HostGroup.Create(name)
		if err != nil {
			b.Fatalf("error while executing Create function\nReason: %v", err)
		}
	}
}

func TestHostGroupGet(t *testing.T) {
	g, err := testingClient.HostGroup.Get(&HostGroupGetParameters{
		CommonGetParameters: CommonGetParameters{
			Filter: map[string]string{},
		},
		SelectDiscoveryRule:  "extend",
		SelectGroupDiscovery: "extend",
		SelectHosts:          "extend",
		SelectTemplates:      "extend",
	})

	if err != nil {
		t.Fatalf("error while executing Get function\nReason: %v", err)
	}

	if len(g.Result) == 0 {
		t.Fatal("an empty list of HostGroups was returned")
	}
}

func BenchmarkHostGroupGet(b *testing.B) {
	params := HostGroupGetParameters{
		CommonGetParameters: CommonGetParameters{
			Filter: map[string]string{},
		},
		SelectDiscoveryRule:  "extend",
		SelectGroupDiscovery: "extend",
		SelectHosts:          "extend",
		SelectTemplates:      "extend",
	}

	for i := 0; i < b.N; i++ {
		_, err := testingClient.HostGroup.Get(&params)

		if err != nil {
			b.Fatalf("error while executing Get function\nReason: %v", err)
		}
	}
}

func TestHostGroupUpdate(t *testing.T) {
	g, err := testingClient.HostGroup.Get(&HostGroupGetParameters{
		CommonGetParameters: CommonGetParameters{
			Search: map[string]string{
				"name": "Hypervisors",
			},
			Output: []string{
				"groupid",
			},
		},
	})

	if err != nil {
		t.Fatalf("error while executing Get function\nReason: %v", err)
	}

	updatedG, err := testingClient.HostGroup.Update(g.Result[0].Id, "Hypervisors-updated")
	if err != nil {
		t.Fatalf("error while executing Update function\nReason: %v", err)
	}

	if len(updatedG.Result.Ids) == 0 {
		t.Fatalf("an empty list of HostGroups was returned")
	}
}

func BenchmarkHostGroupUpdate(b *testing.B) {
	g, err := testingClient.HostGroup.Get(&HostGroupGetParameters{
		CommonGetParameters: CommonGetParameters{
			Search: map[string]string{
				"name": "Hypervisors",
			},
			Output: []string{
				"groupid",
			},
		},
	})

	if err != nil {
		b.Fatalf("error while executing Get function\nReason: %v", err)
	}

	for i := 0; i < b.N; i++ {
		_, err := testingClient.HostGroup.Update(g.Result[0].Id, fmt.Sprintf("Hypervisors-%d", generateId()))
		if err != nil {
			b.Fatalf("error while executing Update function\nReason: %v", err)
		}
	}
}

func TestHostGroupDelete(t *testing.T) {
	name := fmt.Sprintf("test-host-group-%d", generateId())
	g, err := testingClient.HostGroup.Create(name)
	if err != nil {
		t.Fatalf("error while creating '%s' host group\nReason: %v", name, err)
	}

	removedG, err := testingClient.HostGroup.Delete([]string{
		g.Result.Ids[0],
	})
	if err != nil {
		t.Fatalf("error while executing Delete function\nReason: %v", err)
	}

	if len(removedG.Result.Ids) == 0 {
		t.Fatalf("an empty list of HostGroups was returned")
	}
}

// No benchmark provided for the HostGroup delete method

// -- TODO : implement after refactoring 'host.get'
// func TestHostGroupMassAdd(t *testing.T) {
// 	g, err := testingClient.HostGroup.Get(&HostGroupGetParameters{
// 		CommonGetParameters: CommonGetParameters{
// 			Search: map[string]string{
// 				"name": "Discovered hosts",
// 			},
// 			Output: []string{
// 				"groupid",
// 			},
// 		},
// 	})

// 	if err != nil {
// 		t.Fatalf("error while executing Get function\nReason: %v", err)
// 	}

// 	fmt.Println(g.Result[0].Id)
// 	massAddedG, err := testingClient.HostGroup.MassAdd(&HostGroupMassAddParameters{
// 		Groups: []*HostGroupId{
// 			{
// 				Id: g.Result[0].Id,
// 			},
// 		},
// 	})

// 	if err != nil {
// 		t.Fatalf("error while executing MassAdd function\nReason: %v", err)
// 	}

// 	if len(massAddedG.Result.Ids) == 0 {
// 		t.Fatalf("an empty list of HostGroups was returned")
// 	}
// }

// func BenchmarkHostGroupMassAdd(b *testing.B) {
// 	g, err := testingClient.HostGroup.Get(&HostGroupGetParameters{
// 		CommonGetParameters: CommonGetParameters{
// 			Search: map[string]string{
// 				"name": "Discovered hosts",
// 			},
// 			Output: []string{
// 				"groupid",
// 			},
// 		},
// 	})

// 	if err != nil {
// 		b.Fatalf("error while executing Get function\nReason: %v", err)
// 	}

// 	for i := 0; i < b.N; i++ {
// 		_, err = testingClient.HostGroup.MassAdd(&HostGroupMassAddParameters{
// 			Groups: []*HostGroupId{
// 				{
// 					Id: g.Result[0].Id,
// 				},
// 			},
// 		})

// 		if err != nil {
// 			b.Fatalf("error while executing MassAdd function\nReason: %v", err)
// 		}
// 	}
// }

// func TestHostGroupMassRemove(t *testing.T) {
// 	g, err := testingClient.HostGroup.Get(&HostGroupGetParameters{
// 		CommonGetParameters: CommonGetParameters{
// 			Search: map[string]string{
// 				"name": "Discovered hosts",
// 			},
// 			Output: []string{
// 				"groupid",
// 			},
// 		},
// 	})

// 	if err != nil {
// 		t.Fatalf("error while executing Get function\nReason: %v", err)
// 	}

// 	massRemovedG, err := testingClient.HostGroup.MassRemove(&HostGroupMassRemoveParameters{
// 		GroupsIds: []string{
// 			g.Result[0].Id,
// 		},
// 	})

// 	if err != nil {
// 		t.Fatalf("error while executing MassRemove function\nReason: %v", err)
// 	}

// 	if len(massRemovedG.Result.Ids) == 0 {
// 		t.Fatalf("an empty list of HostGroups was returned")
// 	}
// }

// func BenchmarkHostGroupMassRemove(b *testing.B) {
// 	g, err := testingClient.HostGroup.Get(&HostGroupGetParameters{
// 		CommonGetParameters: CommonGetParameters{
// 			Search: map[string]string{
// 				"name": "Discovered hosts",
// 			},
// 			Output: []string{
// 				"groupid",
// 			},
// 		},
// 	})

// 	if err != nil {
// 		b.Fatalf("error while executing Get function\nReason: %v", err)
// 	}

// 	for i := 0; i < b.N; i++ {
// 		_, err = testingClient.HostGroup.MassRemove(&HostGroupMassRemoveParameters{
// 			GroupsIds: []string{
// 				g.Result[0].Id,
// 			},
// 		})

// 		if err != nil {
// 			b.Fatalf("error while executing MassRemove function\nReason: %v", err)
// 		}
// 	}
// }

// func TestHostGroupMassUpdate(t *testing.T) {
// 	g, err := testingClient.HostGroup.Get(&HostGroupGetParameters{
// 		CommonGetParameters: CommonGetParameters{
// 			Search: map[string]string{
// 				"name": "Discovered hosts",
// 			},
// 			Output: []string{
// 				"groupid",
// 			},
// 		},
// 	})

// 	if err != nil {
// 		t.Fatalf("error while executing Get function\nReason: %v", err)
// 	}

// 	massUpdatedG, err := testingClient.HostGroup.MassUpdate(&HostGroupMassUpdateParameters{
// 		Groups: []*HostGroupId{
// 			{
// 				Id: g.Result[0].Id,
// 			},
// 		},
// 	})

// 	if err != nil {
// 		t.Fatalf("error while executing MassUpdate function\nReason: %v", err)
// 	}

// 	if len(massUpdatedG.Result.Ids) == 0 {
// 		t.Fatalf("an empty list of HostGroups was returned")
// 	}
// }

// func BenchmarkHostGroupMassUpdate(b *testing.B) {
// 	g, err := testingClient.HostGroup.Get(&HostGroupGetParameters{
// 		CommonGetParameters: CommonGetParameters{
// 			Search: map[string]string{
// 				"name": "Discovered hosts",
// 			},
// 			Output: []string{
// 				"groupid",
// 			},
// 		},
// 	})

// 	if err != nil {
// 		b.Fatalf("error while executing Get function\nReason: %v", err)
// 	}

// 	for i := 0; i < b.N; i++ {
// 		_, err := testingClient.HostGroup.MassUpdate(&HostGroupMassUpdateParameters{
// 			Groups: []*HostGroupId{
// 				{
// 					Id: g.Result[0].Id,
// 				},
// 			},
// 		})

// 		if err != nil {
// 			b.Fatalf("error while executing MassUpdate function\nReason: %v", err)
// 		}
// 	}
// }
