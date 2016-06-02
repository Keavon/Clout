package player

import (
	"errors"
	"fmt"
	"time"

	"github.com/keavon/clout/pkg/country"
	"github.com/keavon/clout/pkg/resource"
)

// ErrCost is an error returned if a given cost is greater than availible money
var ErrCost = errors.New("Cost Exceeds Money")

//ErrCapacity is an error returned when trying to purchase more installations than capacity
var ErrCapacity = errors.New("Capacity met")

//ErrInvalidOperational is an error returned when a player sets an invalid number of installations operational
var ErrInvalidOperational = errors.New("Invalid Number Operational")

// Purchase an resource
func (p *Player) Purchase(id int) error {
	ri := &ResourceInstallations{}
	stats := country.ResourceStats{}

	switch id {
	case resource.Coal:
		ri = &p.Coal
		stats = p.Country.Coal
	case resource.Oil:
		ri = &p.Oil
		stats = p.Country.Oil
	case resource.Gas:
		ri = &p.Gas
		stats = p.Country.Gas
	case resource.Nuclear:
		ri = &p.Nuclear
		stats = p.Country.Nuclear
	case resource.Geothermal:
		ri = &p.Geothermal
		stats = p.Country.Geothermal
	case resource.Solar:
		ri = &p.Solar
		stats = p.Country.Solar
	case resource.Wind:
		ri = &p.Wind
		stats = p.Country.Wind
	case resource.Hydroelectric:
		ri = &p.Hydroelectric
		stats = p.Country.Hydroelectric
	default:
		return resource.ErrInvalidResource
	}

	if ri.Cost > p.Money {
		return ErrCost
	}

	if stats.Capacity != -1 && ri.Owned >= stats.Capacity {
		return ErrCapacity
	}

	p.Money = p.Money - ri.Cost
	ri.Owned = ri.Owned + 1
	// Make the new installation active
	ri.Operational = ri.Operational + 1
	return nil
}

// SetOperational sets the number of installations operational
func (p *Player) SetOperational(id int, num int) error {
	ri := &ResourceInstallations{}

	switch id {
	case resource.Coal:
		ri = &p.Coal
	case resource.Oil:
		ri = &p.Oil
	case resource.Gas:
		ri = &p.Gas
	case resource.Nuclear:
		ri = &p.Nuclear
	case resource.Geothermal:
		ri = &p.Geothermal
	case resource.Solar:
		ri = &p.Solar
	case resource.Wind:
		ri = &p.Wind
	case resource.Hydroelectric:
		ri = &p.Hydroelectric
	default:
		return resource.ErrInvalidResource
	}

	if num < 0 || num > ri.Owned {
		return ErrInvalidOperational
	}

	fmt.Println(num)

	ri.Operational = num
	return nil
}

// Update environmental damage and money
func (p *Player) Update(start time.Time) {
	p.updateMoney()
	p.updateDamage(start)

	p.LastUpdated = time.Now()
}

type resCalc struct {
	res *ResourceInstallations
	ret int
}

func (p *Player) updateMoney() {
	resources := []resCalc{
		resCalc{
			res: &p.Coal,
			ret: resource.CoalReturn,
		},
		resCalc{
			res: &p.Oil,
			ret: resource.OilReturn,
		},
		resCalc{
			res: &p.Gas,
			ret: resource.GasReturn,
		},
		resCalc{
			res: &p.Nuclear,
			ret: resource.NuclearReturn,
		},
		resCalc{
			res: &p.Geothermal,
			ret: resource.GeothermalReturn,
		},
		resCalc{
			res: &p.Solar,
			ret: resource.SolarReturn,
		},
		resCalc{
			res: &p.Wind,
			ret: resource.WindReturn,
		},
		resCalc{
			res: &p.Hydroelectric,
			ret: resource.HydroelectricReturn,
		},
	}

	for _, r := range resources {
		income := float64(r.res.Operational) * float64(r.ret) * time.Since(p.LastUpdated).Seconds()
		// Round income to nearest integer and add to money
		p.Money += int(income + 0.5)
	}
}

func (p *Player) updateDamage(start time.Time) {
	resources := []resCalc{
		resCalc{
			res: &p.Coal,
			ret: resource.CoalDamage,
		},
		resCalc{
			res: &p.Oil,
			ret: resource.OilDamage,
		},
		resCalc{
			res: &p.Gas,
			ret: resource.GasDamage,
		},
		resCalc{
			res: &p.Nuclear,
			ret: resource.NuclearDamage,
		},
		resCalc{
			res: &p.Geothermal,
			ret: resource.GeothermalDamage,
		},
		resCalc{
			res: &p.Solar,
			ret: resource.SolarDamage,
		},
		resCalc{
			res: &p.Wind,
			ret: resource.WindDamage,
		},
		resCalc{
			res: &p.Hydroelectric,
			ret: resource.HydroelectricDamage,
		},
	}

	capacity := 0

	for _, r := range resources {
		current := float64(r.res.Operational) * float64(r.ret) * time.Since(p.LastUpdated).Seconds()
		// Round current damage and add to current damage
		p.Damage += int(current + 0.5)

		// Add operational power plants to capacity
		capacity += r.res.Operational * 1000000000
	}

	p.Demand = p.Country.AvgDemand(start, p.LastUpdated)

	if p.Demand > capacity {
		p.Damage += p.Demand - capacity
	}
}
