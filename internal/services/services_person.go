package services

import(

	"lambda-person/internal/core/domain"

)

func (s *PersonService) AddPerson(person domain.Person) (*domain.Person, error) {
	childLogger.Debug().Msg("AddPerson")

	p, err := s.personRepository.AddPerson(person)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (s *PersonService) DeletePerson(id string, sk string) (error) {
	childLogger.Debug().Msg("DeletePerson")

	err := s.personRepository.DeletePerson(id, sk)
	if err != nil {
		return err
	}
	return nil
}

func (s *PersonService) UpdatePerson(person domain.Person) (*domain.Person, error) {
	childLogger.Debug().Msg("UpdatePerson")

	p, err := s.personRepository.AddPerson(person)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (s *PersonService) GetPerson(id string) (*domain.Person, error) {
	childLogger.Debug().Msg("GetPerson")

	p, err := s.personRepository.GetPerson(id)
	if err != nil {
		return nil, err
	}
	
	//log.Printf("- services.GetPerson - p : %v ", p)
	return p, nil
}

func (s *PersonService) ListPerson() (*[]domain.Person, error) {
	childLogger.Debug().Msg("ListPerson")

	p, err := s.personRepository.ListPerson()
	if err != nil {
		return nil, err
	}

	//log.Printf("- services.ListPerson - p : %v ", p)
	return p, nil
}

func Sum(x int, y int) int {
    return x + y
}