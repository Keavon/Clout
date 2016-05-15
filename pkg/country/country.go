package country

// Country defines the structure of the country object
type Country struct {
	ID   int    `json:"-"`
	Name string `json:"name"`
}

// Countries is a list of all defined countries.
// Note that ID must equal array index
var Countries = []Country{
	Country{
		ID:   0,
		Name: "United States",
	},
	Country{
		ID:   1,
		Name: "Chine",
	},
}
