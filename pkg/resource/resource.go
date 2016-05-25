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

	// Monetary return per (real) second of a resource

	// CoalReturn is the monetary return of the resource per second
	CoalReturn = 1000
	// OilReturn is the monetary return of the resource per second
	OilReturn = 1000
	// GasReturn is the monetary return of the resource per second
	GasReturn = 1000
	// NuclearReturn is the monetary return of the resource per second
	NuclearReturn = 1000
	// GeothermalReturn is the monetary return of the resource per second
	GeothermalReturn = 2000
	// SolarReturn is the monetary return of the resource per second
	SolarReturn = 2000
	// WindReturn is the monetary return of the resource per second
	WindReturn = 2000
	// HydroelectricReturn is the monetary return of the resource per second
	HydroelectricReturn = 2000

	// Environmental Damage per (real) second of a resource

	// CoalDamage is the environmental damage of the resource per second
	CoalDamage = 1000
	// OilDamage is the environmental damage of the resource per second
	OilDamage = 500
	// GasDamage is the environmental damage of the resource per second
	GasDamage = 400
	// NuclearDamage is the environmental damage of the resource per second
	NuclearDamage = 200
	// GeothermalDamage is the environmental damage of the resource per second
	GeothermalDamage = 10
	// SolarDamage is the environmental damage of the resource per second
	SolarDamage = 10
	// WindDamage is the environmental damage of the resource per second
	WindDamage = 10
	// HydroelectricDamage is the environmental damage of the resource per second
	HydroelectricDamage = 50
)

// ErrInvalidResource is an error returned if a given resource is invalid
var ErrInvalidResource = errors.New("Invalid Resource")
