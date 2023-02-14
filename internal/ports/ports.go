package ports

import(

	"lambda-person/internal/core/domain"
)

type PersonService interface {
	GetPerson(id int) (*domain.Person, error)
}

type PersonRepository interface {
	GetPerson(id int) (*domain.Person, error)
}

type PersonHandler interface {
	GetPerson(id int) (*domain.Person, error)
}