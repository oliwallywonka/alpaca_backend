package shared

type LocationField struct {
	Lat float64
	Lng float64
}
type PriceField []struct {
	MinPersons  int
	MaxPersons  int
	Price       float64
	Currency    string
	IsPerPerson bool
}

type ContactField []struct {
	Type  string
	Value string
}

type PermissionField []string

type ImageField []string

type LangField map[string]string
