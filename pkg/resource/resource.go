package resource

import "errors"

const (
	// Coal resource
	Coal = 0
	// Oil resource
	Oil = 1
	// Gas resource
	Gas = 2
	// Nuclear resource
	Nuclear = 3
	// Geothermal resource
	Geothermal = 4
	// Solar resource
	Solar = 5
	// Wind resource
	Wind = 6
	// Hydroelectric resource
	Hydroelectric = 7

	// InitialCoalCost is the initial cost of the resource
	InitialCoalCost = 10
	// InitialOilCost is the initial cost of the resource
	InitialOilCost = 10
	// InitialGasCost is the initial cost of the resource
	InitialGasCost = 10
	// InitialNuclearCost is the initial cost of the resource
	InitialNuclearCost = 10
	// InitialGeothermalCost is the initial cost of the resource
	InitialGeothermalCost = 10
	// InitialSolarCost is the initial cost of the resource
	InitialSolarCost = 10
	// InitialWindCost is the initial cost of the resource
	InitialWindCost = 10
	// InitialHydroelectricCost is the initial cost of the resource
	InitialHydroelectricCost = 10
)

// ErrInvalidResource is an error returned if a given resource is invalid
var ErrInvalidResource = errors.New("Invalid Resource")
