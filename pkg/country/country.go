package country

// Country defines the structure of the country object
type Country struct {
	ID            int           `json:"id"`
	Name          string        `json:"name"`
	InitialMoney  int           `json:"initialMoney"`
	InitialDemand int           `json:"initalDemand"`
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
	Capacity     int     `json:"capacity"`
	Scalar       float64 `json:"scalar"`
	YearlyDamage int     `json:"yearlyDamage"`
	YearlyIncome int     `json:"yearlyIncome"`
}
