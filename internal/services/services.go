package services

import(
	"log"

	"lambda-person/internal/repository"
//	"lambda-person/internal/ports"

)

type PersonService struct {
	personRepository repository.PersonRepository
}

func NewPersonService(personRepository repository.PersonRepository) *PersonService{
	log.Printf("----------------------------")
	log.Print("- services.NewPersonService") 

	return &PersonService{
		personRepository: personRepository,
	}
}