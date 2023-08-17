package address

import (
	"testing"
)

func TestTypeOfAddressValid(t *testing.T) {
	address := "Avenue Random"

	iSTypeOfAddressExpected := true
	typeOfAddressReceived := TypeOfAddress(address)

	if iSTypeOfAddressExpected != typeOfAddressReceived {
		t.Error("Type is Different from expected")
	}

}

func TestTypeOfAddressInvalid(t *testing.T) {
	address := "Wrond random address"

	iSTypeOfAddressExpected := false
	typeOfAddressReceived := TypeOfAddress(address)

	if iSTypeOfAddressExpected != typeOfAddressReceived {
		t.Error("Type is Different from expected")
	}

}

type testScenario struct {
	addressReceived string
	addressExpected bool
}

func TestTypeOfAddressStruct(t *testing.T) {

	testScenarios := []testScenario{
		{"Street ABC", true},
		{"Avenue ABC", true},
		{"Rod ABC", true},
		{"Wrong ABC", false},
		{"STREET ABC", true},
		{"", false},
		{"AVENUE ABC", true},
	}

	for _, scenario := range testScenarios {
		typeOfAddressReceived := TypeOfAddress(scenario.addressReceived)

		if typeOfAddressReceived != scenario.addressExpected {
			t.Errorf("Type of %t is different then expected.",
				typeOfAddressReceived)
		}
	}

}
