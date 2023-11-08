package services

import(
	"context"

	"lambda-person/internal/core/domain"

)

func (s *PersonService) AddPersonAddress(ctx context.Context,personAddress domain.PersonAddress) (*domain.PersonAddress, error) {
	childLogger.Debug().Msg("AddPersonAddress")

	// Setting the keys
	personAddress.Person.ID = "PERSON-" + personAddress.Person.ID

	_, err := s.personRepository.GetPerson(ctx, personAddress.Person.ID)
	if err != nil {
		return nil, err
	}

	p, err := s.personRepository.AddPersonAddress(ctx, personAddress)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (s *PersonService) ListPersonAddress(ctx context.Context) (*[]domain.PersonAddress, error) {
	childLogger.Debug().Msg("ListPersonAddress")

	p, err := s.personRepository.ListPersonAddress(ctx)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (s *PersonService) QueryPersonAddress(ctx context.Context, id string) (*domain.PersonAddress, error) {
	childLogger.Debug().Msg("QueryPersonAddress")

	p, err := s.personRepository.QueryPersonAddress(ctx, id)
	if err != nil {
		return nil, err
	}
	return p, nil
}
