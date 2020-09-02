package logic

import (
	"strings"
	"testing"
)

func Test_GetFilmsParallel_Works(t *testing.T) {

	var results, err = GetFilmsParallel([]string{"https://swapi.co/api/films/1/"})

	if err != nil {
		t.Fail()
	}

	if len(results) <= 0 {
		t.Error("GetFilms failed!")
	}

	if strings.ToLower(results[0].Title) != "a new hope" {
		t.Errorf("GetFilms[0] Film Title incorrect, got: %s, want: %s.", results[0].Title, "A New Hope")
	}
}

func Test_GetFilmsParallel_Fails(t *testing.T) {
	_, err := GetFilmsParallel([]string{"https://bad_url_097908987/invalid"})

	if err == nil {
		t.Fail()
	}
}

func Test_GetStarShip_Works(t *testing.T) {
	var result, err = GetStarshipsParallel([]string{"https://swapi.co/api/starships/2/"})

	if err != nil {
		t.Fail()
	}

	if strings.ToLower(result[0].Name) != "cr90 corvette" {
		t.Errorf("GetStarship Name incorrect, got: %s, want: %s.", result[0].Name, "CR90 corvette")
	}
}

func Test_GetStarShip_Fails(t *testing.T) {
	_, err := GetStarshipsParallel([]string{"https://bad_url_87876576/invalid"})

	if err == nil {
		t.Errorf("Expected api error!")
	}
}
