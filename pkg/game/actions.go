package game

import (
	"math/rand"
	"time"

	"github.com/keavon/clout/pkg/country"
	"github.com/keavon/clout/pkg/player"
)

// NewPlayer adds a new player to a game
func (g *Game) NewPlayer(ID string, name string, admin bool) (player.Player, error) {
	player := player.Player{ID: ID, Name: name, Admin: admin}

	// Select a country that are not already in use
	countries := []int{}

	for i := range country.Countries {
		countries = append(countries, i)
	}

	// Loop through players and remove countries that have been chosen
	for _, player := range g.Players {
		for i, country := range countries {

			if player.Country.ID == country {
				if i == 0 {
					countries = countries[1:]
				} else if i == len(countries)-1 {
					countries = countries[:len(countries)-1]
				} else {
					countries = append(countries[:i], countries[i+1:]...)
				}

				break
			}
		}
	}

	if len(countries) == 0 {
		return player, ErrFull
	}

	// Seed the RNG
	rand.Seed(time.Now().UnixNano())
	// Select a country from the array of remaining countries.
	player.Country = country.Countries[countries[rand.Intn(len(countries))]]
	player.Money = player.Country.InitialMoney

	// Add player to game
	g.Players = append(g.Players, player)

	return player, nil
}

// Start starts a game
func (g *Game) Start() {
	g.Status = Started
	g.StartTime = time.Now()
}

// Update game fields that change over time
func (g *Game) Update() {
	if g.StartTime.IsZero() {
		return
	}

	gl := time.Since(g.StartTime)
	if gl.Nanoseconds() > duration.Nanoseconds() {
		g.Status = Stopped
	}
}
