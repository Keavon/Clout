package game

import (
	"errors"
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/keavon/clout/pkg/player"
)

const (
	// expiration is the time until the redis object will expire if unused
	expiration = 5 * time.Minute
	// duration is the length of an entire game
	duration = 10 * time.Minute
)

const (
	// Waiting is the game state when waiting for players
	Waiting = 0
	// Started is the game state when game is in progress
	Started = 1
	// Stopped is the game state when the game is over
	Stopped = 2
)

// ErrFull is an error returned when there is no room left in game
var ErrFull = errors.New("No room left in game")

// Game defines the structure of the game object
type Game struct {
	ID        string          `json:"id"`
	Status    int             `json:"status"`
	StartTime time.Time       `json:"startTime"`
	Players   []player.Player `json:"-"`
}

func (g Game) key() string {
	return "game:" + g.ID
}

// Load game from redis
func Load(rc redis.Conn, ID string) (Game, error) {
	g := Game{ID: ID}

	status, err := redis.Int(rc.Do("GET", g.key()))
	if err != nil {
		return g, err
	}
	g.Status = status

	start, err := redis.Int64(rc.Do("GET", g.key()+":start"))
	if err != nil {
		return g, err
	}
	if start != -1 {
		g.StartTime = time.Unix(0, start)
	}

	players, err := redis.Values(rc.Do("SMEMBERS", g.key()+":players"))
	if err != nil {
		return g, err
	}

	playerIDs := []string{}
	if err := redis.ScanSlice(players, &playerIDs); err != nil {
		return g, err
	}

	for _, playerID := range playerIDs {
		// Since resource cost doesn't matter in this context, pass a dummy value
		player, err := player.Load(rc, playerID, 0*time.Second)
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

	if g.StartTime.IsZero() {
		rc.Send("SET", g.key()+":start", -1)
	} else {
		rc.Send("SET", g.key()+":start", g.StartTime.UnixNano())
	}

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
	rc.Send("EXPIRE", g.key()+":start", expiration.Seconds())
	rc.Send("EXPIRE", g.key()+":players", expiration.Seconds())
	_, err := rc.Do("EXEC")
	return err
}

// New creates a new game
func New(ID string) Game {
	return Game{ID: ID, Status: Waiting}
}
