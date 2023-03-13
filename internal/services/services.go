package services

import(
	"github.com/rs/zerolog/log"

	"lambda-person/internal/repository"
	"lambda-person/internal/adapter/notification"

)

var childLogger = log.With().Str("service", "PersonService").Logger()

type PersonService struct {
	personRepository 	repository.PersonRepository
	personNotification 	notification.PersonNotification
}

func NewPersonService(	personRepository repository.PersonRepository, 
						personNotification notification.PersonNotification) *PersonService{
	childLogger.Debug().Msg("NewCardsService")

	return &PersonService{
		personRepository: personRepository,
		personNotification: personNotification,
	}
}