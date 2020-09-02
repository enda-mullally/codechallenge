## StarWars code challenge ##

**The Challenge**

Using the API available here: [https://swapi.co/](https://swapi.co/)  
We want to know for all SW star ships, to cover a given distance, how many stops for resupply are required.
The application will take as input a distance in mega lights (MGLT).
The output should be a collection of all the star ships and the total amount of stops required to make the distance between the planets.

Sample output for 1000000 input:

Y-wing: 74

Millennium Falcon: 9

Rebel Transport: 11


## The Result  ##

**It works! (spoiler alert, results below)**

![Screenshot](/screenshots/Screen_01.png)

The calculations, solution can be found in [StarWars-Distance.xlsx](/workbook/StarWars-Distance.xlsx)

The calculations in the code version is in: [duration.go](/logic/duration.go)

    func CalculateNumberOfStops(shipModels []models.ShipModel, totalMqltHours int) (calculatedShipModels []models.ShipModel)

## Running The Tests  ##

    go test -v github.com/enda-mullally/codechallenge/logic github.com/enda-mullally/codechallenge/models 

**Tests Results**

![Screenshot](/screenshots/Test_Screen_01.png)

## Test Coverage  ##

Running the test coverage report

    go test -timeout 30s -coverprofile=test_coverage.txt github.com/enda-mullally/codechallenge/logic github.com/enda-mullally/codechallenge/models
    go tool cover -html=test_coverage.txt -o test_coverage.html
    
**Coverage Results**

Package: logic, 100%

Coverage report: [test_coverage.html](/test_coverage.html)

![Screenshot](/screenshots/Test_Screen_02.png)

## Compiling and running the code ##

    go get https://github.com/enda-mullally/codechallenge
    go get https://github.com/leejarvis/swapi
    
    cd %GOPATH%\src\github.com\enda-mullally\codechallenge
    
	go build starwars.go

    go run starwars.go

## Installer  ##

I've included a small InnoSetup installer here:
[codechallenge.iss](/installer/codechallenge.iss)

![Screenshot](/screenshots/Setup_01.png)

After installation search for CodeChallenge

## Installer source  ##

Built using [InnoSetup - https://jrsoftware.org/isinfo.php](https://jrsoftware.org/isinfo.php)

Install script here: [codechallenge.iss](/installer/codechallenge.iss) 

![Screenshot](/screenshots/Setup_Source.png)

## Special Thanks  ##

The swapi go package @ [https://github.com/leejarvis/swapi](https://github.com/leejarvis/swapi)

 


