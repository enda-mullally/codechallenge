package logic

import (
	"errors"
	"fmt"
	"sync"

	"github.com/leejarvis/swapi"
)

func getStarshipParallel(url string, c chan swapi.Starship, wg *sync.WaitGroup) {
	defer (*wg).Done()

	var e error
	var s swapi.Starship
	if e = swapi.Get(url, &s); e != nil {
		return
	}

	if e == nil {
		fmt.Printf("*")
		c <- s // push this starhip to the channel
	}
}

// GetStarshipsParallel - GetStarships wrapped in a goroutine
func GetStarshipsParallel(urls []string) (starships []swapi.Starship, err error) {
	starshipChannel := make(chan swapi.Starship)
	var wg sync.WaitGroup

	for _, link := range urls {
		wg.Add(1)
		go getStarshipParallel(link, starshipChannel, &wg)
	}

	go func() {
		wg.Wait()
		close(starshipChannel)
	}()

	var starShipsArray []swapi.Starship
	for nextStarship := range starshipChannel {
		starShipsArray = append(starShipsArray, nextStarship)
	}

	if len(starShipsArray) <= 0 {
		return nil, errors.New("Failed, expected results")
	}

	return starShipsArray, nil
}

func getFilmParallel(url string, c chan swapi.Film, wg *sync.WaitGroup) {
	defer (*wg).Done()

	var e error
	var f swapi.Film
	if e = swapi.Get(url, &f); e != nil {
		return
	}

	if e == nil {
		fmt.Printf("+")
		c <- f // push this film to the channel
	}
}

// GetFilmsParallel - GetFilms wrapped in a goroutine
func GetFilmsParallel(urls []string) (films []swapi.Film, err error) {
	filmsChannel := make(chan swapi.Film)
	var wg sync.WaitGroup

	for _, link := range urls {
		wg.Add(1)
		go getFilmParallel(link, filmsChannel, &wg)
	}

	go func() {
		wg.Wait()
		close(filmsChannel)
	}()

	var filmsArray []swapi.Film
	for nextFilm := range filmsChannel {
		filmsArray = append(filmsArray, nextFilm)
	}

	if len(filmsArray) <= 0 {
		return nil, errors.New("Failed, expected results")
	}
	return filmsArray, nil
}
