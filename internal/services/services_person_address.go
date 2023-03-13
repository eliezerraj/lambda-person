package services

import(
	//"log"

	"lambda-person/internal/core/domain"

)

func (s *PersonService) AddPersonAddress(personAddress domain.PersonAddress) (*domain.PersonAddress, error) {
	childLogger.Debug().Msg("AddPersonAddress")

	// Setting the keys
	personAddress.Person.ID = "PERSON-" + personAddress.Person.ID

	_, err := s.personRepository.GetPerson(personAddress.Person.ID)
	if err != nil {
		return nil, err
	}

	p, err := s.personRepository.AddPersonAddress(personAddress)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (s *PersonService) ListPersonAddress() (*[]domain.PersonAddress, error) {
	childLogger.Debug().Msg("ListPersonAddress")

	p, err := s.personRepository.ListPersonAddress()
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (s *PersonService) QueryPersonAddress(id string) (*domain.PersonAddress, error) {
	childLogger.Debug().Msg("QueryPersonAddress")

	p, err := s.personRepository.QueryPersonAddress(id)
	if err != nil {
		return nil, err
	}
	return p, nil
}
