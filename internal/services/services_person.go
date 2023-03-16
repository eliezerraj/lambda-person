package services

import(

	"lambda-person/internal/core/domain"

)

var(
	eventTypeCreated =  "personCreated"
	eventTypeUpdated = 	"personUpdated"
	eventTypeDeleted = 	"personDeleted"
)

func (s *PersonService) AddPerson(person domain.Person) (*domain.Person, error) {
	childLogger.Debug().Msg("AddPerson")

	// Setting the keys
	person.ID = "PERSON-" + person.ID
	person.SK = person.ID

	p, err := s.personRepository.AddPerson(person)
	if err != nil {
		return nil, err
	}

	// Stream new person
	err = s.personNotification.PutEvent(*p, eventTypeCreated)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (s *PersonService) DeletePerson(id string, sk string) (error) {
	childLogger.Debug().Msg("DeletePerson")

	// Setting the keys
	id = "PERSON-" + id
	sk = "PERSON-" + sk
	err := s.personRepository.DeletePerson(id, sk)
	if err != nil {
		return err
	}

	p := domain.NewPerson(id,sk, "", "")
	// Stream delete person
	err = s.personNotification.PutEvent(*p, eventTypeDeleted)
	if err != nil {
		return err
	}

	return nil
}

func (s *PersonService) UpdatePerson(person domain.Person) (*domain.Person, error) {
	childLogger.Debug().Msg("UpdatePerson")

	// Setting the keys
	person.ID = "PERSON-" + person.ID
	person.SK = person.ID

	p, err := s.personRepository.AddPerson(person)
	if err != nil {
		return nil, err
	}

	// Stream update person
	err = s.personNotification.PutEvent(*p, eventTypeUpdated)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (s *PersonService) GetPerson(id string) (*domain.Person, error) {
	childLogger.Debug().Msg("GetPerson")

	// Setting the keys
	id = "PERSON-" + id

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