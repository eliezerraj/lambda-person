package domain

// Person class
type Person struct {
	ID		string	`json:"id,omitempty"`
	SK		string	`json:"sk,omitempty"`
	Name	string	`json:"name,omitempty"`
	Gender	string	`json:"gender,omitempty"`
}

//Person Constructor
func NewPerson(id string, sk string ,name string, gender string) *Person{
	return &Person{
		ID:	id,
		SK: sk,
		Name: name,
		Gender: gender,
	}
}
