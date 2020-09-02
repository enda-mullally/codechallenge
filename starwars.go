package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	console "github.com/enda-mullally/codechallenge/console"
	consts "github.com/enda-mullally/codechallenge/consts"
	logic "github.com/enda-mullally/codechallenge/logic"
	models "github.com/enda-mullally/codechallenge/models"
)

func main() {

	defer func() { // catch or finally
		if err := recover(); err != nil { // catch
			fmt.Fprintf(os.Stderr, "StarWars exception: %v\n", err)
			os.Exit(-1)
		}
	}()

	console.PrintBanner()

	totalMglt := console.ReadInputMglt()

	if totalMglt <= 0 {
		console.Complete(false)
		return
	}

	displayShipDistances(totalMglt)

	console.Complete(true)
}

func getAllShips() (shipModels []models.ShipModel) {

	fmt.Printf("\nPlease wait reading api data (Films/StarShips)...\n")

	allFilms, err := logic.GetFilmsParallel(consts.FilmUrls)

	if err != nil {
		panic(err)
	}

	var uniqueStarShipURLS = logic.UniqueStarShipURLS(allFilms)

	starShips, err := logic.GetStarshipsParallel(uniqueStarShipURLS)

	for shipIndex := 0; shipIndex < len(starShips); shipIndex++ {

		var mgltHoursValue = -1
		if strings.ToLower(starShips[shipIndex].MGLT) != "unknown" { // some of the ships data doesn't have a valid MQLT value, "unknown", default to -1 in this case, negative will indicate unknown later
			mgltHoursValue, _ = strconv.Atoi(starShips[shipIndex].MGLT)
		}

		nextShipModel := models.ShipModel{
			Name:                 starShips[shipIndex].Name,
			MgltHoursString:      starShips[shipIndex].MGLT,
			MgltHoursValue:       mgltHoursValue,
			ConsumableDaysString: starShips[shipIndex].Consumables,
			ConsumableDaysValue:  logic.ParseNumberOfDays(starShips[shipIndex].Consumables),
		}

		shipModels = append(shipModels, nextShipModel)
	}

	fmt.Println()

	return shipModels
}

func displayShipDistances(totalMqlt int) {

	allShips := getAllShips()

	allCalculatedShips := logic.CalculateNumberOfStops(allShips, totalMqlt)

	fmt.Println()
	fmt.Println("Results...")

	for _, starShip := range allCalculatedShips {

		var numberOfStops = starShip.NumberOfStops
		var numberOfStopsStr = ""
		if numberOfStops <= 0 {
			numberOfStopsStr = "unknown, missing data"
		} else {
			numberOfStopsStr = strconv.FormatFloat(numberOfStops, 'f', 2, 64)
		}

		fmt.Printf("Ship name: '%s', Number of stops between planets: '%s'\n", starShip.Name, numberOfStopsStr)
	}

	fmt.Println()
}
