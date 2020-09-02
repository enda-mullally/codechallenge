package logic

import (
	"strconv"
	"strings"

	"github.com/leejarvis/swapi"
	models "gitlab.com/enda-mullally/codechallenge/models"
)

// ParseNumberOfDays - Given a string like "2 months" return 60 (days)
func ParseNumberOfDays(durationStr string) int {

	if durationStr == "" || len(durationStr) == 0 {
		panic("ParseHours() invalid duration string")
	}

	// special case, some fields are "unknown", instead of panic'ing return negative
	if strings.ToLower(durationStr) == "unknown" {
		return -1
	}

	durationParts := strings.Split(
		strings.ToLower(durationStr), " ")

	if len(durationParts) != 2 {
		panic("ParseHours() invalid duration string")
	}

	integral, err := strconv.Atoi(durationParts[0])

	if err != nil {
		panic(err)
	}

	var days = 0

	switch durationParts[1] {
	case "day", "days":
		days = 1
	case "week", "weeks":
		days = 7
	case "month", "months":
		days = 30
	case "year", "years":
		days = 365
	default:
		days = -1
	}

	return integral * days
}

// CalculateNumberOfStops - Iterate the shipModels and do the calculations
func CalculateNumberOfStops(shipModels []models.ShipModel, totalMqltHours int) (calculatedShipModels []models.ShipModel) {

	for shipIndex := 0; shipIndex < len(shipModels); shipIndex++ {

		var mgltHoursOfTravel = (float64)(totalMqltHours / shipModels[shipIndex].MgltHoursValue)
		var mgltDaysOfTravel = (float64)(mgltHoursOfTravel / 24)

		nextCalculatedShipModel := models.ShipModel{
			// Clone existing model data
			Name:                 shipModels[shipIndex].Name,
			MgltHoursString:      shipModels[shipIndex].MgltHoursString,
			MgltHoursValue:       shipModels[shipIndex].MgltHoursValue,
			ConsumableDaysString: shipModels[shipIndex].ConsumableDaysString,
			ConsumableDaysValue:  shipModels[shipIndex].ConsumableDaysValue,
			// Calculate these values
			MgltHoursOfTravel: mgltHoursOfTravel,
			MgltDaysOfTravel:  mgltDaysOfTravel,
			NumberOfStops:     (float64)(mgltDaysOfTravel / (float64)(shipModels[shipIndex].ConsumableDaysValue)),
		}

		calculatedShipModels = append(calculatedShipModels, nextCalculatedShipModel)
	}

	return calculatedShipModels
}

// UniqueStarShipURLS Iterate the array of Films and foreach extact the list of unique starShipUrls
func UniqueStarShipURLS(allFilms []swapi.Film) (uniqueStarShipURLS []string) {

	uniqueShips := make(map[string]string)

	for filmIndex := 0; filmIndex < len(allFilms); filmIndex++ {
		for shipIndex := 0; shipIndex < len(allFilms[filmIndex].Starships); shipIndex++ {

			nextStarShipURL := allFilms[filmIndex].Starships[shipIndex]

			_, shipAlreadyMapped := uniqueShips[nextStarShipURL]

			if !shipAlreadyMapped {
				uniqueShips[nextStarShipURL] = nextStarShipURL

				uniqueStarShipURLS = append(uniqueStarShipURLS, nextStarShipURL)
			}
		}
	}

	return uniqueStarShipURLS
}
