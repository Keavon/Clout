package authorization

import (
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/keavon/clout/pkg/game"
	"github.com/keavon/clout/pkg/player"
)

// expiration is the time until the redis object will expire if unused
const expration = 2 * time.Minute

// Authorization defines the structure of the auth object
type Authorization struct {
	ID     string        `json:"-"`
	Game   game.Game     `json:"game"`
	Player player.Player `json:"player"`
}

func (a Authorization) key() string {
	return "auth:" + a.ID
}

// Load authorization from redis
func Load(rc redis.Conn, ID string) (Authorization, error) {
	a := Authorization{ID: ID}

	gameID, err := redis.String(rc.Do("GET", a.key()+":game"))
	if err != nil {
		return a, err
	}

	g, err := game.Load(rc, gameID)
	if err != nil {
		return a, err
	}
	a.Game = g

	playerID, err := redis.String(rc.Do("GET", a.key()+":player"))
	if err != nil {
		return a, err
	}

	length := time.Since(g.StartTime)
	if g.StartTime.IsZero() {
		length = 0 * time.Second
	}

	p, err := player.Load(rc, playerID, length)
	if err != nil {
		return a, err
	}
	a.Player = p

	return a, nil
}

// Save saves the authorization to redis
func (a Authorization) Save(rc redis.Conn) error {
	rc.Send("MULTI")
	rc.Send("SET", a.key()+":game", a.Game.ID)
	rc.Send("SET", a.key()+":player", a.Player.ID)

	if _, err := rc.Do("EXEC"); err != nil {
		return err
	}

	return a.Touch(rc)
}

// Touch updates the TTL of the authorization to keep it from expiring
func (a Authorization) Touch(rc redis.Conn) error {
	rc.Send("MULTI")
	rc.Send("EXPIRE", a.key()+":game", expration.Seconds())
	rc.Send("EXPIRE", a.key()+":player", expration.Seconds())
	_, err := rc.Do("EXEC")
	return err
}

// New creates a new authorization
func New(ID string, g game.Game, p player.Player) Authorization {
	return Authorization{ID: ID, Game: g, Player: p}
}
