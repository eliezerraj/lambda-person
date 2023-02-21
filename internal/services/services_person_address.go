package services

import(
	"log"

	"lambda-person/internal/core/domain"

)

func (s *PersonService) AddPersonAddress(person domain.Person, adresses []domain.Address ) (*domain.PersonAddress, error) {
	log.Printf("+++++++++++++++++++++++++++++++++")
	log.Print("- services.AddPersonAddress - person   : ", person)
	log.Print("- services.AddPersonAddress - adresses: ", adresses)

	p, err := s.personRepository.AddPersonAddress(person, adresses)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (s *PersonService) ListPersonAddress() (*[]domain.PersonAddress, error) {
	log.Printf("+++++++++++++++++++++++++++++++++")
	log.Print("- services.ListPersonAddress")

	p, err := s.personRepository.ListPersonAddress()
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (s *PersonService) QueryPersonAddress(person domain.Person) (*[]domain.PersonAddress, error) {
	log.Printf("+++++++++++++++++++++++++++++++++")
	log.Print("- services.QueryPersonAddress")

	p, err := s.personRepository.QueryPersonAddress(person)
	if err != nil {
		return nil, err
	}
	return p, nil
}
