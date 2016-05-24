package player

import (
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/keavon/clout/pkg/country"
)

// expiration is the time until the redis object will expire if unused
const expiration = 2 * time.Minute

var resources = [8]string{
	"coal",
	"oil",
	"gas",
	"nuclear",
	"geothermal",
	"solar",
	"wind",
	"hydroelectric",
}

// Player defines the structure of the player object
type Player struct {
	ID            string                `json:"-"`
	Name          string                `json:"name"`
	Admin         bool                  `json:"admin"`
	Money         int                   `json:"money"`
	Country       country.Country       `json:"country"`
	Coal          ResourceInstallations `json:"coal"`
	Oil           ResourceInstallations `json:"oil"`
	Gas           ResourceInstallations `json:"gas"`
	Nuclear       ResourceInstallations `json:"nuclear"`
	Geothermal    ResourceInstallations `json:"geothermal"`
	Solar         ResourceInstallations `json:"solar"`
	Wind          ResourceInstallations `json:"wind"`
	Hydroelectric ResourceInstallations `json:"hydroelectric"`
}

// ResourceInstallations defines installations of resources.
type ResourceInstallations struct {
	Owned       int `json:"owned"`
	Operational int `json:"operational"`
}

func (p Player) key() string {
	return "player:" + p.ID
}

// TODO: Load() is hardcoded to this order
func (p Player) resourceKeys() []string {
	keys := []string{}

	for _, resource := range resources {
		root := p.key() + ":" + resource
		keys = append(keys, root)
	}

	return keys
}

// Load player from redis
func Load(rc redis.Conn, ID string) (Player, error) {
	p := Player{ID: ID}

	name, err := redis.String(rc.Do("GET", p.key()))
	if err != nil {
		return p, err
	}
	p.Name = name

	admin, err := redis.Bool(rc.Do("GET", p.key()+":admin"))
	if err != nil {
		return p, err
	}
	p.Admin = admin

	money, err := redis.Int(rc.Do("GET", p.key()+":money"))
	if err != nil {
		return p, err
	}
	p.Money = money

	countryID, err := redis.Int(rc.Do("GET", p.key()+":country"))
	if err != nil {
		return p, err
	}

	if countryID > len(country.Countries) {
		return p, err
	}

	p.Country = country.Countries[countryID]

	installations := []ResourceInstallations{}

	for _, key := range p.resourceKeys() {
		owned, err := redis.Int(rc.Do("GET", key))
		if err != nil {
			return p, err
		}

		operational, err := redis.Int(rc.Do("GET", key+":operational"))
		if err != nil {
			return p, err
		}

		installations = append(installations, ResourceInstallations{Owned: owned, Operational: operational})
	}

	p.Coal = installations[0]
	p.Oil = installations[1]
	p.Gas = installations[2]
	p.Nuclear = installations[3]
	p.Geothermal = installations[4]
	p.Solar = installations[5]
	p.Wind = installations[6]
	p.Hydroelectric = installations[7]

	return p, nil
}

// Save saves the player to redis
func (p Player) Save(rc redis.Conn) error {
	rc.Send("MULTI")
	rc.Send("SET", p.key(), p.Name)
	rc.Send("SET", p.key()+":admin", p.Admin)
	rc.Send("SET", p.key()+":money", p.Money)
	rc.Send("SET", p.key()+":country", p.Country.ID)

	saveResource(rc, p.key(), "coal", p.Coal)
	saveResource(rc, p.key(), "oil", p.Oil)
	saveResource(rc, p.key(), "gas", p.Gas)
	saveResource(rc, p.key(), "nuclear", p.Nuclear)
	saveResource(rc, p.key(), "geothermal", p.Geothermal)
	saveResource(rc, p.key(), "solar", p.Solar)
	saveResource(rc, p.key(), "wind", p.Wind)
	saveResource(rc, p.key(), "hydroelectric", p.Hydroelectric)

	if _, err := rc.Do("EXEC"); err != nil {
		return err
	}

	return p.Touch(rc)
}

func saveResource(rc redis.Conn, rkey string, name string, value ResourceInstallations) {
	key := rkey + ":" + name

	rc.Send("SET", key, value.Owned)
	rc.Send("SET", key+":operational", value.Operational)
}

// Touch updates the TTL of the player to keep it from expiring
func (p Player) Touch(rc redis.Conn) error {
	rc.Send("MULTI")
	rc.Send("EXPIRE", p.key(), expiration.Seconds())
	rc.Send("EXPIRE", p.key()+":admin", expiration.Seconds())
	rc.Send("EXPIRE", p.key()+":money", expiration.Seconds())
	rc.Send("EXPIRE", p.key()+":country", expiration.Seconds())

	for _, key := range p.resourceKeys() {
		rc.Send("EXPIRE", key, expiration.Seconds())
		rc.Send("EXPIRE", key+":operational", expiration.Seconds())
	}

	_, err := rc.Do("EXEC")
	return err
}
