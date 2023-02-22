package services

import(
	"log"

	"lambda-person/internal/core/domain"

)

func (s *PersonService) AddPersonAddress(personAddress domain.PersonAddress) (*domain.PersonAddress, error) {
	log.Printf("+++++++++++++++++++++++++++++++++")
	log.Print("- services.AddPersonAddress - personAddress   : ", personAddress)

	p, err := s.personRepository.AddPersonAddress(personAddress)
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

func (s *PersonService) QueryPersonAddress(id string) (*domain.PersonAddress, error) {
	log.Printf("+++++++++++++++++++++++++++++++++")
	log.Print("- services.QueryPersonAddress")

	p, err := s.personRepository.QueryPersonAddress(id)
	if err != nil {
		return nil, err
	}
	return p, nil
}
