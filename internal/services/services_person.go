package services

import(
	"lambda-person/internal/core/domain"
	"context"

)

var(
	eventTypeCreated =  "personCreated"
	eventTypeUpdated = 	"personUpdated"
	eventTypeDeleted = 	"personDeleted"
)

func (s *PersonService) AddPerson(ctx context.Context, person domain.Person) (*domain.Person, error) {
	childLogger.Debug().Msg("AddPerson")

	// Setting the keys
	person.ID = "PERSON-" + person.ID
	person.SK = person.ID

	p, err := s.personRepository.AddPerson(ctx, person)
	if err != nil {
		return nil, err
	}

	// Stream new person
	err = s.personNotification.PutEvent(ctx ,*p, eventTypeCreated)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (s *PersonService) DeletePerson(ctx context.Context, id string, sk string) (error) {
	childLogger.Debug().Msg("DeletePerson")

	// Setting the keys
	id = "PERSON-" + id
	sk = "PERSON-" + sk

	// Verify is person exists
	_, err := s.personRepository.GetPerson(ctx, id)
	if err != nil {
		return err
	}

	err = s.personRepository.DeletePerson(ctx, id, sk)
	if err != nil {
		return err
	}

	p := domain.NewPerson(id,sk, "", "")
	// Stream delete person
	err = s.personNotification.PutEvent(ctx, *p, eventTypeDeleted)
	if err != nil {
		return err
	}

	return nil
}

func (s *PersonService) UpdatePerson(ctx context.Context, person domain.Person) (*domain.Person, error) {
	childLogger.Debug().Msg("UpdatePerson")

	// Setting the keys
	person.ID = "PERSON-" + person.ID
	person.SK = person.ID

	p, err := s.personRepository.AddPerson(ctx, person)
	if err != nil {
		return nil, err
	}

	// Stream update person
	err = s.personNotification.PutEvent(ctx, *p, eventTypeUpdated)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (s *PersonService) GetPerson(ctx context.Context, id string) (*domain.Person, error) {
	childLogger.Debug().Msg("GetPerson")

	// Setting the keys
	id = "PERSON-" + id

	p, err := s.personRepository.GetPerson(ctx, id)
	if err != nil {
		return nil, err
	}
	
	//log.Printf("- services.GetPerson - p : %v ", p)
	return p, nil
}

func (s *PersonService) ListPerson(ctx context.Context) (*[]domain.Person, error) {
	childLogger.Debug().Msg("ListPerson")

	p, err := s.personRepository.ListPerson(ctx)
	if err != nil {
		return nil, err
	}

	//log.Printf("- services.ListPerson - p : %v ", p)
	return p, nil
}

func Sum(x int, y int) int {
    return x + y
}