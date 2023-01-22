package port

type Port struct {
	Code
	Details
}

type Code string

type Details struct {
	Name        string
	City        string
	Country     string
	Alias       []interface{} // TODO: in example doesn't exist any value for determine type
	Regions     []interface{} // TODO: in example doesn't exist any value for determine type
	Coordinates []float64
	Province    string
	Timezone    string
	Unlocs      []string
	Code        string
}
