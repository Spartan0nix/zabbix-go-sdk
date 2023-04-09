package zabbixgosdk

import "testing"

const (
	hostName        = "test-host"
	updatedHostName = "test-host-update"
)

var hostId string

func TestHostCreate(t *testing.T) {
	h, err := testingClient.Host.Create(&HostCreateParameters{
		Host:          hostName,
		Name:          hostName,
		Description:   "Testing description",
		InventoryMode: HostManual,
		Status:        HostUnmonitored,
		Groups: []*HostGroupId{
			{
				GroupId: "1",
			},
		},
		Interfaces: []*HostInterface{
			{
				Ip:    "127.0.0.1",
				Dns:   "localhost",
				Main:  HostInterfaceDefault,
				Port:  "10050",
				Type:  HostInterfaceAgent,
				UseIp: HostInterfaceDns,
			},
		},
		Tags: []*HostTag{
			{
				Tag:   "{$TEST_TAG}",
				Value: "test-host-tag-value",
			},
		},
		Inventory: map[HostInventory]string{
			Name:  hostName,
			Alias: "testing-host-alias",
		},
	})

	if err != nil {
		t.Fatalf("Error when creating host '%s'.\nReason : %v", hostName, err)
	}

	if h == nil || len(h.HostIds) == 0 {
		t.Fatalf("Create method should return a list of the created hosts.\nAn empty list was returned.")
	}

	hostId = h.HostIds[0]
}

func TestHostGet(t *testing.T) {
	h, err := testingClient.Host.Get(&HostGetParameters{
		Filter: map[string]string{
			"name": hostName,
		},
	})

	if err != nil {
		t.Fatalf("Error when getting host '%s'.\nReason : %v", hostName, err)
	}

	if len(h) == 0 {
		t.Fatalf("Get method should return a list of hosts matching the given criteria.\nAn empty list was returned.")
	}

	if h[0].HostId != hostId {
		t.Fatalf("Wrong host returned.\nExpected : %s\nReturned : %s", hostId, h[0].HostId)
	}
}

func TestHostMassAdd(t *testing.T) {
	h, err := testingClient.Host.MassAdd(&HostMassAddParameters{
		Hosts: []*HostId{
			{
				HostId: hostId,
			},
		},
		Interfaces: []*HostInterface{
			{
				Ip:    "127.0.0.2",
				Dns:   "",
				Main:  HostInterfaceNotDefault,
				Port:  "10893",
				Type:  HostInterfaceAgent,
				UseIp: HostInterfaceIp,
			},
		},
		Macros: []*HostMacro{
			{
				Macro: "{$TEST_MACRO}",
				Value: "test-host-macro-value",
				Type:  Text,
			},
		},
	})

	if err != nil {
		t.Fatalf("Error when mass adding host '%s'.\nReason : %v", hostName, err)
	}

	if h == nil || len(h.HostIds) == 0 {
		t.Fatalf("MassAdd method should return a list of the updated hosts.\nAn empty list was returned.")
	}

	if h.HostIds[0] != hostId {
		t.Fatalf("Wrong host returned.\nExpected : %s\nReturned : %s", hostId, h.HostIds[0])
	}
}

func TestHostMassRemove(t *testing.T) {
	m, err := testingClient.UserMacro.Get(&UserMacroGetParameters{
		HostIds: []string{
			hostId,
		},
	})

	if err != nil {
		t.Fatalf("Error when getting macro for host '%s'.\nReason : %v", hostName, err)
	}

	h, err := testingClient.Host.MassRemove(&HostMassRemoveParameters{
		HostIds: []string{
			hostId,
		},
		Interfaces: []*HostInterface{
			{
				Ip:   "127.0.0.2",
				Dns:  "",
				Port: "10893",
			},
		},
		Macros: []string{
			m[0].HostId,
		},
	})

	if err != nil {
		t.Fatalf("Error when mass removing interface  and macro from host '%s'.\nReason : %v", hostName, err)
	}

	if h == nil || len(h.HostIds) == 0 {
		t.Fatalf("MassRemove method should return a list of the updated hosts.\nAn empty list was returned.")
	}

	if h.HostIds[0] != hostId {
		t.Fatalf("Wrong host returned.\nExpected : %s\nReturned : %s", hostId, h.HostIds[0])
	}
}

func TestHostMassUpdate(t *testing.T) {
	h, err := testingClient.Host.MassUpdate(&HostMassUpdateParameters{
		Hosts: []*HostId{
			{
				HostId: hostId,
			},
		},
		Status: HostUnmonitored,
	})

	if err != nil {
		t.Fatalf("Error when mass updating status from host '%s'.\nReason : %v", hostName, err)
	}

	if h == nil || len(h.HostIds) == 0 {
		t.Fatalf("MassUpdate method should return a list of the updated hosts.\nAn empty list was returned.")
	}

	if h.HostIds[0] != hostId {
		t.Fatalf("Wrong host returned.\nExpected : %s\nReturned : %s", hostId, h.HostIds[0])
	}
}

func TestHostUpdate(t *testing.T) {
	h, err := testingClient.Host.Update(&HostUpdateParameters{
		HostId:        hostId,
		Status:        HostMonitored,
		Host:          updatedHostName,
		Name:          updatedHostName,
		InventoryMode: HostManual,
		Inventory: map[HostInventory]string{
			Alias:   updatedHostName,
			Contact: "random@localhost",
		},
	})

	if err != nil {
		t.Fatalf("Error when updating host '%s'.\nReason : %v", hostName, err)
	}

	if h == nil || len(h.HostIds) == 0 {
		t.Fatalf("Update method should return a list of the updated hosts.\nAn empty list was returned.")
	}

	if h.HostIds[0] != hostId {
		t.Fatalf("Wrong host returned.\nExpected Id : %s\nId returned : %s", hostId, h.HostIds[0])
	}
}

func TestHostDelete(t *testing.T) {
	h, err := testingClient.Host.Delete([]string{
		hostId,
	})

	if err != nil {
		t.Fatalf("Error when deleting host '%s'.\nReason : %v", hostName, err)
	}

	if h == nil || len(h.HostIds) == 0 {
		t.Fatalf("Delete method should return a list with the id of the deleted hosts.\nAn empty list was returned.")
	}

	if h.HostIds[0] != hostId {
		t.Fatalf("Wrong host returned.\nExpected : %s\nReturned : %s", hostId, h.HostIds[0])
	}
}

func TestGetHostInventoryId(t *testing.T) {
	if id := GetHostInventoryId(Alias); id != "4" {
		t.Fatalf("Wrong id returned for 'Alias' field.\nExpected '4'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(AssetTag); id != "11" {
		t.Fatalf("Wrong id returned for 'AssetTag' field.\nExpected '11'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(Chassis); id != "28" {
		t.Fatalf("Wrong id returned for 'Chassis' field.\nExpected '28'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(Contact); id != "23" {
		t.Fatalf("Wrong id returned for 'Contact' field.\nExpected '23'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(ContractNumber); id != "32" {
		t.Fatalf("Wrong id returned for 'ContractNumber' field.\nExpected '32'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(DateHwDecomm); id != "47" {
		t.Fatalf("Wrong id returned for 'DateHwDecomm' field.\nExpected '47'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(DateHwExpiry); id != "46" {
		t.Fatalf("Wrong id returned for 'DateHwExpiry' field.\nExpected '46'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(DateHwInstall); id != "45" {
		t.Fatalf("Wrong id returned for 'DateHwInstall' field.\nExpected '45'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(DateHwPurchase); id != "44" {
		t.Fatalf("Wrong id returned for 'DateHwPurchase' field.\nExpected '44'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(DeploymentStatus); id != "34" {
		t.Fatalf("Wrong id returned for 'DeploymentStatus' field.\nExpected '34'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(Hardware); id != "14" {
		t.Fatalf("Wrong id returned for 'Hardware' field.\nExpected '14'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(HardwareFull); id != "15" {
		t.Fatalf("Wrong id returned for 'HardwareFull' field.\nExpected '15'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(HostNetmask); id != "39" {
		t.Fatalf("Wrong id returned for 'HostNetmask' field.\nExpected '39'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(HostNetworks); id != "38" {
		t.Fatalf("Wrong id returned for 'HostNetworks' field.\nExpected '38'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(HostRouter); id != "40" {
		t.Fatalf("Wrong id returned for 'HostRouter' field.\nExpected '40'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(HwArch); id != "30" {
		t.Fatalf("Wrong id returned for 'HwArch' field.\nExpected '30'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(InstallerName); id != "33" {
		t.Fatalf("Wrong id returned for 'InstallerName' field.\nExpected '33'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(Location); id != "24" {
		t.Fatalf("Wrong id returned for 'Location' field.\nExpected '24'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(LocationLat); id != "25" {
		t.Fatalf("Wrong id returned for 'LocationLat' field.\nExpected '25'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(LocationLon); id != "26" {
		t.Fatalf("Wrong id returned for 'LocationLon' field.\nExpected '26'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(MacAddressA); id != "12" {
		t.Fatalf("Wrong id returned for 'MacAddressA' field.\nExpected '12'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(MacAddressB); id != "13" {
		t.Fatalf("Wrong id returned for 'MacAddressB' field.\nExpected '13'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(Model); id != "29" {
		t.Fatalf("Wrong id returned for 'Model' field.\nExpected '29'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(Name); id != "3" {
		t.Fatalf("Wrong id returned for 'Name' field.\nExpected '3'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(Notes); id != "27" {
		t.Fatalf("Wrong id returned for 'Notes' field.\nExpected '27'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(OobIp); id != "41" {
		t.Fatalf("Wrong id returned for 'OobIp' field.\nExpected '41'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(OobNetmask); id != "42" {
		t.Fatalf("Wrong id returned for 'OobNetmask' field.\nExpected '42'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(OobRouter); id != "43" {
		t.Fatalf("Wrong id returned for 'OobRouter' field.\nExpected '43'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(Os); id != "5" {
		t.Fatalf("Wrong id returned for 'Os' field.\nExpected '5'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(OsFull); id != "6" {
		t.Fatalf("Wrong id returned for 'OsFull' field.\nExpected '6'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(OsShort); id != "7" {
		t.Fatalf("Wrong id returned for 'OsShort' field.\nExpected '7'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(Poc1Cell); id != "61" {
		t.Fatalf("Wrong id returned for 'Poc1Cell' field.\nExpected '61'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(Poc1Email); id != "58" {
		t.Fatalf("Wrong id returned for 'Poc1Email' field.\nExpected '58'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(Poc1Name); id != "57" {
		t.Fatalf("Wrong id returned for 'Poc1Name' field.\nExpected '57'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(Poc1Notes); id != "63" {
		t.Fatalf("Wrong id returned for 'Poc1Notes' field.\nExpected '63'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(Poc1PhoneA); id != "59" {
		t.Fatalf("Wrong id returned for 'Poc1PhoneA' field.\nExpected '59'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(Poc1PhoneB); id != "60" {
		t.Fatalf("Wrong id returned for 'Poc1PhoneB' field.\nExpected '60'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(Poc1Screen); id != "62" {
		t.Fatalf("Wrong id returned for 'Poc1Screen' field.\nExpected '62'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(Poc2Cell); id != "68" {
		t.Fatalf("Wrong id returned for 'Poc2Cell' field.\nExpected '68'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(Poc2Email); id != "65" {
		t.Fatalf("Wrong id returned for 'Poc2Email' field.\nExpected '65'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(Poc2Name); id != "64" {
		t.Fatalf("Wrong id returned for 'Poc2Name' field.\nExpected '64'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(Poc2Notes); id != "70" {
		t.Fatalf("Wrong id returned for 'Poc2Notes' field.\nExpected '70'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(Poc2PhoneA); id != "66" {
		t.Fatalf("Wrong id returned for 'Poc2PhoneA' field.\nExpected '66'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(Poc2PhoneB); id != "67" {
		t.Fatalf("Wrong id returned for 'Poc2PhoneB' field.\nExpected '67'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(Poc2Screen); id != "69" {
		t.Fatalf("Wrong id returned for 'Poc2Screen' field.\nExpected '69'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(SerialNoA); id != "8" {
		t.Fatalf("Wrong id returned for 'SerialNoA' field.\nExpected '8'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(SerialNoB); id != "9" {
		t.Fatalf("Wrong id returned for 'SerialNoB' field.\nExpected '9'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(SiteAddressA); id != "48" {
		t.Fatalf("Wrong id returned for 'SiteAddressA' field.\nExpected '48'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(SiteAddressB); id != "49" {
		t.Fatalf("Wrong id returned for 'SiteAddressB' field.\nExpected '49'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(SiteAddressC); id != "50" {
		t.Fatalf("Wrong id returned for 'SiteAddressC' field.\nExpected '50'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(SiteCity); id != "51" {
		t.Fatalf("Wrong id returned for 'SiteCity' field.\nExpected '51'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(SiteCountry); id != "53" {
		t.Fatalf("Wrong id returned for 'SiteCountry' field.\nExpected '53'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(SitNotes); id != "56" {
		t.Fatalf("Wrong id returned for 'SitNotes' field.\nExpected '56'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(SiteRack); id != "55" {
		t.Fatalf("Wrong id returned for 'SiteRack' field.\nExpected '55'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(SiteState); id != "52" {
		t.Fatalf("Wrong id returned for 'SiteState' field.\nExpected '52'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(SiteZip); id != "54" {
		t.Fatalf("Wrong id returned for 'SiteZip' field.\nExpected '54'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(Software); id != "16" {
		t.Fatalf("Wrong id returned for 'Software' field.\nExpected '16'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(SoftwareAppA); id != "18" {
		t.Fatalf("Wrong id returned for 'SoftwareAppA' field.\nExpected '18'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(SoftwareAppB); id != "19" {
		t.Fatalf("Wrong id returned for 'SoftwareAppB' field.\nExpected '19'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(SoftwareAppC); id != "20" {
		t.Fatalf("Wrong id returned for 'SoftwareAppC' field.\nExpected '20'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(SoftwareAppD); id != "21" {
		t.Fatalf("Wrong id returned for 'SoftwareAppD' field.\nExpected '21'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(SoftwareAppE); id != "22" {
		t.Fatalf("Wrong id returned for 'SoftwareAppE' field.\nExpected '22'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(SoftwareFull); id != "17" {
		t.Fatalf("Wrong id returned for 'SoftwareFull' field.\nExpected '17'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(Tag); id != "10" {
		t.Fatalf("Wrong id returned for 'Tag' field.\nExpected '10'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(Type); id != "1" {
		t.Fatalf("Wrong id returned for 'Type' field.\nExpected '1'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(TypeFull); id != "2" {
		t.Fatalf("Wrong id returned for 'TypeFull' field.\nExpected '2'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(UrlA); id != "35" {
		t.Fatalf("Wrong id returned for 'UrlA' field.\nExpected '35'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(UrlB); id != "36" {
		t.Fatalf("Wrong id returned for 'UrlB' field.\nExpected '36'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(UrlC); id != "37" {
		t.Fatalf("Wrong id returned for 'UrlC' field.\nExpected '37'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId(Vendor); id != "31" {
		t.Fatalf("Wrong id returned for 'Vendor' field.\nExpected '31'.\nReturned : '%s' ", id)
	}

	if id := GetHostInventoryId("Unsupported"); id != "0" {
		t.Fatalf("Wrong id returned for unsupported field.\nExpected '0'.\nReturned : '%s' ", id)
	}
}
