package domain

// Person class
type PersonAddress struct {
	Person		Person		`json:"person,omitempty"`
	Addresses	[]Address	`json:"addresses,omitempty"`
}

//Person Constructor
func NewPersonAddress(person Person, adresses []Address) *PersonAddress{
	return &PersonAddress{
		Person: person,
		Addresses: adresses,
	}
}
