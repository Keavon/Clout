package player

import (
	"errors"
	"time"

	"github.com/keavon/clout/pkg/country"
	"github.com/keavon/clout/pkg/resource"
)

// ErrCost is an error returned if a given cost is greater than availible money
var ErrCost = errors.New("Cost Exceeds Money")

//ErrCapacity is an error returned when trying to purchase more installations than capacity
var ErrCapacity = errors.New("Capacity met")

// Purchase an resource
func (p *Player) Purchase(id int) error {
	ri := &ResourceInstallations{}
	stats := country.ResourceStats{}
	cost := 0

	switch id {
	case resource.Coal:
		ri = &p.Coal
		cost, _ = p.Country.ResourceCost(resource.Coal, time.Second)
		stats = p.Country.Coal
	case resource.Oil:
		ri = &p.Oil
		cost, _ = p.Country.ResourceCost(resource.Oil, time.Second)
		stats = p.Country.Oil
	case resource.Gas:
		ri = &p.Gas
		cost, _ = p.Country.ResourceCost(resource.Gas, time.Second)
		stats = p.Country.Gas
	case resource.Nuclear:
		ri = &p.Nuclear
		cost, _ = p.Country.ResourceCost(resource.Nuclear, time.Second)
		stats = p.Country.Nuclear
	case resource.Geothermal:
		ri = &p.Geothermal
		cost, _ = p.Country.ResourceCost(resource.Geothermal, time.Second)
		stats = p.Country.Geothermal
	case resource.Solar:
		ri = &p.Solar
		cost, _ = p.Country.ResourceCost(resource.Solar, time.Second)
		stats = p.Country.Solar
	case resource.Wind:
		ri = &p.Wind
		cost, _ = p.Country.ResourceCost(resource.Wind, time.Second)
		stats = p.Country.Wind
	case resource.Hydroelectric:
		ri = &p.Hydroelectric
		cost, _ = p.Country.ResourceCost(resource.Hydroelectric, time.Second)
		stats = p.Country.Hydroelectric
	default:
		return resource.ErrInvalidResource
	}

	if cost > p.Money {
		return ErrCost
	}

	if ri.Owned >= stats.Capacity {
		return ErrCapacity
	}

	p.Money = p.Money - cost
	ri.Owned = ri.Owned + 1
	// Make the new installation active
	ri.Operational = ri.Operational + 1
	return nil
}
