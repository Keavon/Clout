package country

// Country defines the structure of the country object
type Country struct {
	ID            int           `json:"id"`
	Name          string        `json:"name"`
	InitialMoney  int           `json:"initialMoney"`
	Coal          ResourceStats `json:"coal"`
	Oil           ResourceStats `json:"oil"`
	Gas           ResourceStats `json:"gas"`
	Nuclear       ResourceStats `json:"nuclear"`
	Geothermal    ResourceStats `json:"geothermal"`
	Solar         ResourceStats `json:"solar"`
	Wind          ResourceStats `json:"wind"`
	Hydroelectric ResourceStats `json:"hydroelectric"`
}

// ResourceStats defines the structure of country resources.
type ResourceStats struct {
	Capacity int     `json:"capacity"`
	Scalar   float64 `json:"scalar"`
}

// Countries is a list of all defined countries.
// Note that ID must equal array index
var Countries = []Country{
	Country{
		ID:           0,
		Name:         "United States",
		InitialMoney: 50,
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
		ID:           1,
		Name:         "China",
		InitialMoney: 40,
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
