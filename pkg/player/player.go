package player

import (
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/keavon/clout/pkg/country"
)

// expiration is the time until the redis object will expire if unused
const expration = 2 * time.Minute

// Player defines the structure of the player object
type Player struct {
	ID      string          `json:"-"`
	Name    string          `json:"name"`
	Admin   bool            `json:"admin"`
	Money   int             `json:"money"`
	Country country.Country `json:"country"`
}

func (p Player) key() string {
	return "player:" + p.ID
}

// Load player from redis
func Load(rc redis.Conn, ID string) (Player, error) {
	p := Player{ID: ID}
	rc.Send("GET", p.key())
	rc.Send("GET", p.key()+":admin")
	rc.Send("GET", p.key()+":money")
	rc.Send("GET", p.key()+":country")
	rc.Flush()

	name, err := redis.String(rc.Receive())
	if err != nil {
		return p, err
	}
	p.Name = name

	admin, err := redis.Bool(rc.Receive())
	if err != nil {
		return p, err
	}
	p.Admin = admin

	money, err := redis.Int(rc.Receive())
	if err != nil {
		return p, err
	}
	p.Money = money

	countryID, err := redis.Int(rc.Receive())
	if err != nil {
		return p, err
	}

	if countryID > len(country.Countries) {
		return p, err
	}

	p.Country = country.Countries[countryID]

	return p, nil
}

// Save saves the player to redis
func (p Player) Save(rc redis.Conn) error {
	rc.Send("MULTI")
	rc.Send("SET", p.key(), p.Name)
	rc.Send("SET", p.key()+":admin", p.Admin)
	rc.Send("SET", p.key()+":money", p.Money)
	rc.Send("SET", p.key()+":country", p.Country.ID)

	if _, err := rc.Do("EXEC"); err != nil {
		return err
	}

	return p.Touch(rc)
}

// Touch updates the TTL of the player to keep it from expiring
func (p Player) Touch(rc redis.Conn) error {
	rc.Send("MULTI")
	rc.Send("EXPIRE", p.key(), expration.Seconds())
	rc.Send("EXPIRE", p.key()+":admin", expration.Seconds())
	rc.Send("EXPIRE", p.key()+":money", expration.Seconds())
	rc.Send("EXPIRE", p.key()+":country", expration.Seconds())
	_, err := rc.Do("EXEC")
	return err
}
