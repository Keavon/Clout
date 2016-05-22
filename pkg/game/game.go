package game

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/keavon/clout/pkg/country"
	"github.com/keavon/clout/pkg/player"
)

// expiration is the time until the redis object will expire if unused
const expiration = 5 * time.Minute

// ErrFull is an error returned when there is no room left in game
var ErrFull = errors.New("No room left in game")

// Game defines the structure of the game object
type Game struct {
	ID      string
	Status  int
	Players []player.Player
}

func (g Game) key() string {
	return "game:" + g.ID
}

// Load game from redis
func Load(rc redis.Conn, ID string) (Game, error) {
	g := Game{ID: ID}
	//rc.Send("GET", g.key())
	//rc.Send("SMEMBERS", g.key()+":players")
	//rc.Flush()

	fmt.Println("_________> In game")

	status, err := redis.Int(rc.Do("GET", g.key()))
	if err != nil {
		fmt.Println("///////////> Status error")
		fmt.Println()
		return g, err
	}
	g.Status = status

	fmt.Println(">>>>>>>>>>> After Status")

	players, err := redis.Values(rc.Do("SMEMBERS", g.key()+":players"))
	if err != nil {
		return g, err
	}

	fmt.Println("~~~~~~~~> Started player load")

	playerIDs := []string{}
	if err := redis.ScanSlice(players, &playerIDs); err != nil {
		return g, err
	}

	for _, playerID := range playerIDs {
		player, err := player.Load(rc, playerID)
		if err != nil {
			fmt.Println(err)
			continue
		}

		g.Players = append(g.Players, player)
	}

	return g, nil
}

// Save saves the game to redis
func (g Game) Save(rc redis.Conn) error {
	rc.Send("MULTI")
	rc.Send("SET", g.key(), g.Status)

	for _, player := range g.Players {
		rc.Send("SADD", g.key()+":players", player.ID)
	}

	if _, err := rc.Do("EXEC"); err != nil {
		return err
	}

	return g.Touch(rc)
}

// Touch updates the TTL of the game to keep it from expiring
func (g Game) Touch(rc redis.Conn) error {
	rc.Send("MULTI")
	rc.Send("EXPIRE", g.key(), expiration.Seconds())
	rc.Send("EXPIRE", g.key()+":players", expiration.Seconds())
	_, err := rc.Do("EXEC")
	return err
}

// New creates a new game
func New(ID string) Game {
	return Game{ID: ID, Status: 0}
}

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
