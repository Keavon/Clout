package country

import (
	"time"

	"github.com/keavon/clout/pkg/resource"
)

// ResourceCost gives the cost of a resource at a given time
// TODO: Implement price drop over time
func (c Country) ResourceCost(id int, time time.Duration) (int, error) {
	switch id {
	case resource.Coal:
		return resource.InitialCoalCost, nil
	case resource.Oil:
		return resource.InitialOilCost, nil
	case resource.Gas:
		return resource.InitialGasCost, nil
	case resource.Nuclear:
		return resource.InitialNuclearCost, nil
	case resource.Geothermal:
		return resource.InitialGeothermalCost, nil
	case resource.Solar:
		return resource.InitialSolarCost, nil
	case resource.Wind:
		return resource.InitialWindCost, nil
	case resource.Hydroelectric:
		return resource.InitialHydroelectricCost, nil
	default:
		return -1, resource.ErrInvalidResource
	}
}

// Demand calculates average power demand between two times.
// TODO: Add a growth constant between times
func (c Country) AvgDemand(start time.Time, previous time.Time) int {
	growth := 1.1
	cd := time.Since(start).Seconds() * growth
	pd := previous.Sub(start).Seconds() * growth

	avg := (cd + pd) / 2

	// Round growth demand to int and add it to inital demand
	return c.InitialDemand + int(avg+0.5)
}
