package domain

// Person class
type Person struct {
	ID		string	`json:"id,omitempty"`
	Name	string	`json:"name,omitempty"`
	Gender	string	`json:"gender,omitempty"`
}

//Person Constructor
func NewPerson(id string, name string, gender string) *Person{
	return &Person{
		ID:	id,
		Name: name,
		Gender: gender,
	}
}
