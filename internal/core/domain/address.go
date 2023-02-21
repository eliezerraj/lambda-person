package domain

// Address class
type Address struct {
	ID				string	`json:"id,omitempty"`
	SK				string	`json:"sk,omitempty"`
	Street			string	`json:"street,omitempty"`
	StreetNumber	int		`json:"street_number,omitempty"`
	ZipCode			string	`json:"zip_code,omitempty"`
}

//Person Constructor
func NewAddress(id string, sk string, street string, streetnumber int, zipcode string) *Address{
	return &Address{
		ID:	id,
		SK:	sk,
		Street: street,
		StreetNumber: streetnumber,
		ZipCode: zipcode,
	}
}
