package models

import (
	"testing"
)

func Test_ShipModel_Works(t *testing.T) {

	sut := ShipModel{
		Name:                 "Test",
		MgltHoursString:      "25",
		MgltHoursValue:       25,
		ConsumableDaysString: "2 weeks",
		ConsumableDaysValue:  7,
	}

	if sut.Name != "Test" {
		t.Errorf("Model Name incorrect, got: %s, want: %s.", sut.Name, "Test")
	}

	if sut.MgltHoursString != "25" {
		t.Errorf("Model MgltHoursString incorrect, got: %s, want: %s.", sut.MgltHoursString, "25")
	}

	if sut.MgltHoursValue != 25 {
		t.Errorf("Model MgltHoursValue incorrect, got: %d, want: %d.", sut.MgltHoursValue, 25)
	}

	if sut.ConsumableDaysString != "2 weeks" {
		t.Errorf("Model ConsumableDaysString incorrect, got: %s, want: %s.", sut.ConsumableDaysString, "2 weeks")
	}

	if sut.ConsumableDaysValue != 7 {
		t.Errorf("Model ConsumableDaysValue incorrect, got: %d, want: %d.", sut.ConsumableDaysValue, 7)
	}
}
