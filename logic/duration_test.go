package logic

import (
	"math"
	"testing"

	"gitlab.com/enda-mullally/codechallenge/models"
)

func Test_ParseNumberOfDays_1Day_Works(t *testing.T) {
	var days = ParseNumberOfDays("1 daY")

	if days != 1 {
		t.Errorf("ParseNumberOfDays incorrect, got: %d, want: %d.", days, 1)
	}
}

func Test_ParseNumberOfDays_5Days_Work(t *testing.T) {
	var days = ParseNumberOfDays("5 dAys")

	if days != 5 {
		t.Errorf("ParseNumberOfDays incorrect, got: %d, want: %d.", days, 5)
	}
}

func Test_ParseNumberOfDays_1Week_Works(t *testing.T) {
	var days = ParseNumberOfDays("1 WeEK")

	if days != 7 {
		t.Errorf("ParseNumberOfDays incorrect, got: %d, want: %d.", days, 7)
	}
}

func Test_ParseNumberOfDays_2Weeks_Works(t *testing.T) {
	var days = ParseNumberOfDays("2 WeEKs")

	if days != 14 {
		t.Errorf("ParseNumberOfDays incorrect, got: %d, want: %d.", days, 14)
	}
}

func Test_ParseNumberOfDays_1Month_Works(t *testing.T) {
	var days = ParseNumberOfDays("1 Month")

	if days != 30 {
		t.Errorf("ParseNumberOfDays incorrect, got: %d, want: %d.", days, 30)
	}
}

func Test_ParseNumberOfDays_2Months_Work(t *testing.T) {
	var days = ParseNumberOfDays("2 MONTHS")

	if days != 60 {
		t.Errorf("ParseNumberOfDays incorrect, got: %d, want: %d.", days, 60)
	}
}

func Test_ParseNumberOfDays_1Year_Works(t *testing.T) {
	var days = ParseNumberOfDays("1 YeaR")

	if days != 365 {
		t.Errorf("ParseNumberOfDays incorrect, got: %d, want: %d.", days, 365)
	}
}

func Test_ParseNumberOfDays_2Years_Work(t *testing.T) {
	var days = ParseNumberOfDays("2 Years")

	if days != 730 {
		t.Errorf("ParseNumberOfDays incorrect, got: %d, want: %d.", days, 730)
	}
}

func Test_ParseNumberOfDays_UnknownString_Work(t *testing.T) {
	var days = ParseNumberOfDays("UnKnown")

	if days >= 0 {
		t.Errorf("ParseNumberOfDays incorrect, got: %d, want: %d.", days, -1)
	}
}

func Test_ParseNumberOfDays_Part1_Valid_Fails(t *testing.T) {
	var days = ParseNumberOfDays("2 eons")

	if days >= 0 {
		t.Errorf("ParseNumberOfDays incorrect, got: %d, want: %d.", days, -1)
	}
}

func Test_ParseNumberOfDays_Invalid_Fails(t *testing.T) {
	AssertPanic(t, func() { ParseNumberOfDays("Invalid") }, "Test_ParseNumberOfDays did not panic!")
}

func Test_ParseNumberOfDays_InvalidInvalid_Fails(t *testing.T) {
	AssertPanic(t, func() { ParseNumberOfDays("Invalid Invalid") }, "Test_ParseNumberOfDays did not panic!")
}

func Test_ParseNumberOfDays_Empty_Fails(t *testing.T) {
	AssertPanic(t, func() { ParseNumberOfDays("") }, "Test_ParseNumberOfDays did not panic!")
}

func Test_CalculateNumberOfStops_Works(t *testing.T) {

	var sutShipModels []models.ShipModel

	sutShipModel := models.ShipModel{
		Name:                 "Test",
		MgltHoursString:      "75",
		MgltHoursValue:       75,
		ConsumableDaysString: "1 week",
		ConsumableDaysValue:  7,
	}

	sutShipModels = append(sutShipModels, sutShipModel)

	results := CalculateNumberOfStops(sutShipModels, 1000000)

	result := round2places(results[0].NumberOfStops)

	if result != 79.36 {
		t.Errorf("CalculateNumberOfStops incorrect, got: %f, want: %f.", result, 79.36)
	}
}

func Test_CalculateNumberOfStops_DivByZero(t *testing.T) {

	var sutShipModels []models.ShipModel

	sutShipModel := models.ShipModel{
		Name:                 "Test",
		MgltHoursString:      "75",
		MgltHoursValue:       75,
		ConsumableDaysString: "1 week",
		ConsumableDaysValue:  0, // div by zero, go can handle it!
	}

	sutShipModels = append(sutShipModels, sutShipModel)

	results := CalculateNumberOfStops(sutShipModels, 1000000)

	result := round2places(results[0].NumberOfStops)

	if !math.IsInf(result, 0) {
		t.Errorf("CalculateNumberOfStops incorrect, got: %f, want: Inf", result)
	}
}

func AssertPanic(t *testing.T, f func(), errorString string) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf(errorString)
		}
	}()
	f()
}

func round2places(in float64) float64 {
	return math.Round(in*100) / 100
}
