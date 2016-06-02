package country

import (
	"github.com/keavon/clout/pkg/resource"
)

// mutiply by 10 to get year rather than per second
const (
	coalIncomeYearly          = resource.CoalReturn * 10
	oilIncomeYearly           = resource.OilReturn * 10
	gasIncomeYearly           = resource.GasReturn * 10
	nuclearIncomeYearly       = resource.NuclearReturn * 10
	geothermalIncomeYearly    = resource.GeothermalReturn * 10
	solarIncomeYearly         = resource.SolarReturn * 10
	windIncomeYearly          = resource.WindReturn * 10
	hydroelectricIncomeYearly = resource.HydroelectricReturn * 10

	coalDamageYearly          = resource.CoalDamage * 10
	oilDamageYearly           = resource.OilDamage * 10
	gasDamageYearly           = resource.GasDamage * 10
	nuclearDamageYearly       = resource.NuclearDamage * 10
	geothermalDamageYearly    = resource.GeothermalDamage * 10
	solarDamageYearly         = resource.SolarDamage * 10
	windDamageYearly          = resource.WindDamage * 10
	hydroelectricDamageYearly = resource.HydroelectricDamage * 10
)

// Countries is a list of all defined countries.
// Note that ID must equal array index
var Countries = []Country{
	Country{
		ID:            0,
		Name:          "United States",
		InitialMoney:  50,
		InitialDemand: 1000000000,
		Coal: ResourceStats{
			Capacity:     2,
			Scalar:       1.0,
			YearlyIncome: coalIncomeYearly,
			YearlyDamage: coalDamageYearly,
		},
		Oil: ResourceStats{
			Capacity:     2,
			Scalar:       1.0,
			YearlyIncome: oilIncomeYearly,
			YearlyDamage: oilDamageYearly,
		},
		Gas: ResourceStats{
			Capacity:     2,
			Scalar:       1.0,
			YearlyIncome: gasIncomeYearly,
			YearlyDamage: gasDamageYearly,
		},
		Nuclear: ResourceStats{
			Capacity:     2,
			Scalar:       1.0,
			YearlyIncome: nuclearIncomeYearly,
			YearlyDamage: nuclearDamageYearly,
		},
		Geothermal: ResourceStats{
			Capacity:     2,
			Scalar:       1.0,
			YearlyIncome: geothermalIncomeYearly,
			YearlyDamage: geothermalDamageYearly,
		},
		Solar: ResourceStats{
			Capacity:     -1,
			Scalar:       1.0,
			YearlyIncome: solarIncomeYearly,
			YearlyDamage: solarDamageYearly,
		},
		Wind: ResourceStats{
			Capacity:     -1,
			Scalar:       1.0,
			YearlyIncome: windIncomeYearly,
			YearlyDamage: windDamageYearly,
		},
		Hydroelectric: ResourceStats{
			Capacity:     2,
			Scalar:       1.0,
			YearlyIncome: hydroelectricIncomeYearly,
			YearlyDamage: hydroelectricDamageYearly,
		},
	},
	Country{
		ID:            1,
		Name:          "China",
		InitialMoney:  40,
		InitialDemand: 1000000000,
		Coal: ResourceStats{
			Capacity:     2,
			Scalar:       1.0,
			YearlyIncome: coalIncomeYearly,
			YearlyDamage: coalDamageYearly,
		},
		Oil: ResourceStats{
			Capacity:     2,
			Scalar:       1.0,
			YearlyIncome: oilIncomeYearly,
			YearlyDamage: oilDamageYearly,
		},
		Gas: ResourceStats{
			Capacity:     2,
			Scalar:       1.0,
			YearlyIncome: gasIncomeYearly,
			YearlyDamage: gasDamageYearly,
		},
		Nuclear: ResourceStats{
			Capacity:     2,
			Scalar:       1.0,
			YearlyIncome: nuclearIncomeYearly,
			YearlyDamage: nuclearDamageYearly,
		},
		Geothermal: ResourceStats{
			Capacity:     2,
			Scalar:       1.0,
			YearlyIncome: geothermalIncomeYearly,
			YearlyDamage: geothermalDamageYearly,
		},
		Solar: ResourceStats{
			Capacity:     -1,
			Scalar:       1.0,
			YearlyIncome: solarIncomeYearly,
			YearlyDamage: solarDamageYearly,
		},
		Wind: ResourceStats{
			Capacity:     -1,
			Scalar:       1.0,
			YearlyIncome: windIncomeYearly,
			YearlyDamage: windDamageYearly,
		},
		Hydroelectric: ResourceStats{
			Capacity:     2,
			Scalar:       1.0,
			YearlyIncome: hydroelectricIncomeYearly,
			YearlyDamage: hydroelectricDamageYearly,
		},
	},
	Country{
		ID:            2,
		Name:          "Mother Russia",
		InitialDemand: 1000000000,
		InitialMoney:  40,
		Coal: ResourceStats{
			Capacity:     2,
			Scalar:       1.0,
			YearlyIncome: coalIncomeYearly,
			YearlyDamage: coalDamageYearly,
		},
		Oil: ResourceStats{
			Capacity:     2,
			Scalar:       1.0,
			YearlyIncome: oilIncomeYearly,
			YearlyDamage: oilDamageYearly,
		},
		Gas: ResourceStats{
			Capacity:     2,
			Scalar:       1.0,
			YearlyIncome: gasIncomeYearly,
			YearlyDamage: gasDamageYearly,
		},
		Nuclear: ResourceStats{
			Capacity:     2,
			Scalar:       1.0,
			YearlyIncome: nuclearIncomeYearly,
			YearlyDamage: nuclearDamageYearly,
		},
		Geothermal: ResourceStats{
			Capacity:     2,
			Scalar:       1.0,
			YearlyIncome: geothermalIncomeYearly,
			YearlyDamage: geothermalDamageYearly,
		},
		Solar: ResourceStats{
			Capacity:     -1,
			Scalar:       1.0,
			YearlyIncome: solarIncomeYearly,
			YearlyDamage: solarDamageYearly,
		},
		Wind: ResourceStats{
			Capacity:     -1,
			Scalar:       1.0,
			YearlyIncome: windIncomeYearly,
			YearlyDamage: windDamageYearly,
		},
		Hydroelectric: ResourceStats{
			Capacity:     2,
			Scalar:       1.0,
			YearlyIncome: hydroelectricIncomeYearly,
			YearlyDamage: hydroelectricDamageYearly,
		},
	},
	Country{
		ID:            3,
		Name:          "India",
		InitialMoney:  40,
		InitialDemand: 1000000000,
		Coal: ResourceStats{
			Capacity:     2,
			Scalar:       1.0,
			YearlyIncome: coalIncomeYearly,
			YearlyDamage: coalDamageYearly,
		},
		Oil: ResourceStats{
			Capacity:     2,
			Scalar:       1.0,
			YearlyIncome: oilIncomeYearly,
			YearlyDamage: oilDamageYearly,
		},
		Gas: ResourceStats{
			Capacity:     2,
			Scalar:       1.0,
			YearlyIncome: gasIncomeYearly,
			YearlyDamage: gasDamageYearly,
		},
		Nuclear: ResourceStats{
			Capacity:     2,
			Scalar:       1.0,
			YearlyIncome: nuclearIncomeYearly,
			YearlyDamage: nuclearDamageYearly,
		},
		Geothermal: ResourceStats{
			Capacity:     2,
			Scalar:       1.0,
			YearlyIncome: geothermalIncomeYearly,
			YearlyDamage: geothermalDamageYearly,
		},
		Solar: ResourceStats{
			Capacity:     -1,
			Scalar:       1.0,
			YearlyIncome: solarIncomeYearly,
			YearlyDamage: solarDamageYearly,
		},
		Wind: ResourceStats{
			Capacity:     -1,
			Scalar:       1.0,
			YearlyIncome: windIncomeYearly,
			YearlyDamage: windDamageYearly,
		},
		Hydroelectric: ResourceStats{
			Capacity:     2,
			Scalar:       1.0,
			YearlyIncome: hydroelectricIncomeYearly,
			YearlyDamage: hydroelectricDamageYearly,
		},
	},
}
