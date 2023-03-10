package services

import(
	"github.com/rs/zerolog/log"

	"lambda-person/internal/repository"

)

var childLogger = log.With().Str("service", "PersonService").Logger()

type PersonService struct {
	personRepository repository.PersonRepository
}

func NewPersonService(personRepository repository.PersonRepository) *PersonService{
	childLogger.Debug().Msg("NewCardsService")

	return &PersonService{
		personRepository: personRepository,
	}
}