// Package graphics provides console printing methods and simple console input
package graphics

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	consts "gitlab.com/enda-mullally/codechallenge/consts"
)

// PrintBanner - Prints the application banner - ascii text from: http://www.patorjk.com/software/taag/#p=display&h=1&f=Graffiti&t=StarWars
func PrintBanner() {
	fmt.Println(`  _________  __                   __      __                       `)
	fmt.Println(` /   _____/_/  |_ _____  _______ /  \    /  \_____  _______  ______`)
	fmt.Println(` \_____  \ \   __\\__  \ \_  __ \\   \/\/   /\__  \ \_  __ \/  ___/`)
	fmt.Println(` /        \ |  |   / __ \_|  | \/ \        /  / __ \_|  | \/\___ \ `)
	fmt.Println(`/_______  / |__|  (____  /|__|     \__/\  /  (____  /|__|  /____  >`)
	fmt.Println(`        \/             \/               \/        \/            \/ `)
	fmt.Println()

	fmt.Println("StarWars - code challenge - enda.mullally@outlook.com - Version: " + consts.AppVersion)
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("Displays a list of all current StarWars ships and calculates the number of necessary stops between planets for each based on the number of light years input (MGLT). Calculations are based on data retrieved from 'https://swapi.co/'")
	fmt.Println()
}

// ReadInputMglt - ASk the user to enter the number of MGLT years
func ReadInputMglt() int {
	fmt.Printf("Please enter the total MGLT years (0 to Cancel): ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	result, err := strconv.Atoi(text)

	if err != nil {
		panic(err)
	}

	if result > 0 {
		fmt.Println("Ok, got it:", result, "MGLT it is!")
	}

	return result
}

// Complete - Display cancelled/complete message and wait for input
func Complete(successfullyCompleted bool) {
	fmt.Printf("----------------------------------------\n")

	if successfullyCompleted {
		fmt.Printf("Complete :-)\n\n")
	} else {
		fmt.Printf("Cancelled :-(\n\n")
	}

	fmt.Printf("<Return> to close.\n")
	fmt.Scanln()
}
