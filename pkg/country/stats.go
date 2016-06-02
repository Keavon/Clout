package country

// Countries is a list of all defined countries.
// Note that ID must equal array index
var Countries = []Country{
	Country{
		ID:            0,
		Name:          "United States",
		InitialMoney:  50,
		InitialDemand: 1000000000,
		Coal: ResourceStats{
			Capacity: 2,
			Scalar:   1.0,
		},
		Oil: ResourceStats{
			Capacity: 2,
			Scalar:   1.0,
		},
		Gas: ResourceStats{
			Capacity: 2,
			Scalar:   1.0,
		},
		Nuclear: ResourceStats{
			Capacity: 2,
			Scalar:   1.0,
		},
		Geothermal: ResourceStats{
			Capacity: 2,
			Scalar:   1.0,
		},
		Solar: ResourceStats{
			Capacity: -1,
			Scalar:   1.0,
		},
		Wind: ResourceStats{
			Capacity: -1,
			Scalar:   1.0,
		},
		Hydroelectric: ResourceStats{
			Capacity: 2,
			Scalar:   1.0,
		},
	},
	Country{
		ID:            1,
		Name:          "China",
		InitialMoney:  40,
		InitialDemand: 1000000000,
		Coal: ResourceStats{
			Capacity: 2,
			Scalar:   1.0,
		},
		Oil: ResourceStats{
			Capacity: 2,
			Scalar:   1.0,
		},
		Gas: ResourceStats{
			Capacity: 2,
			Scalar:   1.0,
		},
		Nuclear: ResourceStats{
			Capacity: 2,
			Scalar:   1.0,
		},
		Geothermal: ResourceStats{
			Capacity: 2,
			Scalar:   1.0,
		},
		Solar: ResourceStats{
			Capacity: -1,
			Scalar:   1.0,
		},
		Wind: ResourceStats{
			Capacity: -1,
			Scalar:   1.0,
		},
		Hydroelectric: ResourceStats{
			Capacity: 2,
			Scalar:   1.0,
		},
	},
	Country{
		ID:            2,
		Name:          "Mother Russia",
		InitialDemand: 1000000000,
		InitialMoney:  40,
		Coal: ResourceStats{
			Capacity: 2,
			Scalar:   1.0,
		},
		Oil: ResourceStats{
			Capacity: 2,
			Scalar:   1.0,
		},
		Gas: ResourceStats{
			Capacity: 2,
			Scalar:   1.0,
		},
		Nuclear: ResourceStats{
			Capacity: 2,
			Scalar:   1.0,
		},
		Geothermal: ResourceStats{
			Capacity: 2,
			Scalar:   1.0,
		},
		Solar: ResourceStats{
			Capacity: -1,
			Scalar:   1.0,
		},
		Wind: ResourceStats{
			Capacity: -1,
			Scalar:   1.0,
		},
		Hydroelectric: ResourceStats{
			Capacity: 2,
			Scalar:   1.0,
		},
	},
	Country{
		ID:            3,
		Name:          "India",
		InitialMoney:  40,
		InitialDemand: 1000000000,
		Coal: ResourceStats{
			Capacity: 2,
			Scalar:   1.0,
		},
		Oil: ResourceStats{
			Capacity: 2,
			Scalar:   1.0,
		},
		Gas: ResourceStats{
			Capacity: 2,
			Scalar:   1.0,
		},
		Nuclear: ResourceStats{
			Capacity: 2,
			Scalar:   1.0,
		},
		Geothermal: ResourceStats{
			Capacity: 2,
			Scalar:   1.0,
		},
		Solar: ResourceStats{
			Capacity: -1,
			Scalar:   1.0,
		},
		Wind: ResourceStats{
			Capacity: -1,
			Scalar:   1.0,
		},
		Hydroelectric: ResourceStats{
			Capacity: 2,
			Scalar:   1.0,
		},
	},
}
