package country

import (
	"time"

	"github.com/keavon/clout/pkg/resource"
)

// ResourceCost gives the cost of a resource at a given time
// TODO: Implement price drop over time
func (c Country) ResourceCost(id int, length time.Duration) (int, error) {
	// Time discount is the discount given to resources over time
	td := length.Seconds() / 4
	// Resource cost is the initial cost * the countries scalar for the resource
	rc := 0.0

	switch id {
	case resource.Coal:
		rc = resource.InitialCoalCost * c.Coal.Scalar
	case resource.Oil:
		rc = resource.InitialOilCost * c.Oil.Scalar
	case resource.Gas:
		rc = resource.InitialGasCost * c.Gas.Scalar
	case resource.Nuclear:
		rc = resource.InitialNuclearCost * c.Nuclear.Scalar
	case resource.Geothermal:
		rc = resource.InitialGeothermalCost * c.Geothermal.Scalar
	case resource.Solar:
		rc = resource.InitialSolarCost * c.Solar.Scalar
	case resource.Wind:
		rc = resource.InitialWindCost * c.Wind.Scalar
	case resource.Hydroelectric:
		rc = resource.InitialHydroelectricCost * c.Hydroelectric.Scalar
	default:
		return -1, resource.ErrInvalidResource
	}

	cost := int(rc - td + 0.5)

	// Cost can't be lower than 0
	if cost < 0 {
		cost = 0
	}

	return cost, nil
}

// AvgDemand calculates average power demand between two times.
// TODO: Add a growth constant between times
func (c Country) AvgDemand(start time.Time, previous time.Time) int {
	growth := 6000000.0
	cd := time.Since(start).Seconds() * growth
	pd := previous.Sub(start).Seconds() * growth

	avg := (cd + pd) / 2

	// Round growth demand to int and add it to inital demand
	return c.InitialDemand + int(avg+0.5)
}
