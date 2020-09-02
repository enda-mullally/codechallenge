package models

// ShipModel - represents a ship model with only the required fields for this challenge + calculated fields
type ShipModel struct {
	Name                 string
	MgltHoursString      string
	MgltHoursValue       int
	ConsumableDaysString string
	ConsumableDaysValue  int
	// Calculated fields
	MgltHoursOfTravel float64
	MgltDaysOfTravel  float64
	NumberOfStops     float64
}
