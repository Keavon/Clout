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
	InitialCoalCost = 1000000000
	// InitialOilCost is the initial cost of the resource
	InitialOilCost = 7500000000
	// InitialGasCost is the initial cost of the resource
	InitialGasCost = 1000000000
	// InitialNuclearCost is the initial cost of the resource
	InitialNuclearCost = 5000000000
	// InitialGeothermalCost is the initial cost of the resource
	InitialGeothermalCost = 4000000000
	// InitialSolarCost is the initial cost of the resource
	InitialSolarCost = 3500000000
	// InitialWindCost is the initial cost of the resource
	InitialWindCost = 2000000000
	// InitialHydroelectricCost is the initial cost of the resource
	InitialHydroelectricCost = 2500000000

	// Monetary return per (real) second of a resource

	// CoalReturn is the monetary return of the resource per second
	CoalReturn = 5000000
	// OilReturn is the monetary return of the resource per second
	OilReturn = 6000000
	// GasReturn is the monetary return of the resource per second
	GasReturn = 7000000
	// NuclearReturn is the monetary return of the resource per second
	NuclearReturn = 15000000
	// GeothermalReturn is the monetary return of the resource per second
	GeothermalReturn = 20000000
	// SolarReturn is the monetary return of the resource per second
	SolarReturn = 30000000
	// WindReturn is the monetary return of the resource per second
	WindReturn = 25000000
	// HydroelectricReturn is the monetary return of the resource per second
	HydroelectricReturn = 22500000

	// Environmental Damage per (real) second of a resource

	// CoalDamage is the environmental damage of the resource per second
	CoalDamage = 1000
	// OilDamage is the environmental damage of the resource per second
	OilDamage = 750
	// GasDamage is the environmental damage of the resource per second
	GasDamage = 500
	// NuclearDamage is the environmental damage of the resource per second
	NuclearDamage = 100
	// GeothermalDamage is the environmental damage of the resource per second
	GeothermalDamage = 5
	// SolarDamage is the environmental damage of the resource per second
	SolarDamage = 0
	// WindDamage is the environmental damage of the resource per second
	WindDamage = 0
	// HydroelectricDamage is the environmental damage of the resource per second
	HydroelectricDamage = 50
)

// ErrInvalidResource is an error returned if a given resource is invalid
var ErrInvalidResource = errors.New("Invalid Resource")
