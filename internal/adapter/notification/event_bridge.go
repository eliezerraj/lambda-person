package notification

import (
	"os"
	"encoding/json"

	"github.com/rs/zerolog/log"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/eventbridge"

	"lambda-person/internal/erro"
	"lambda-person/internal/core/domain"

)

var childLogger = log.With().Str("notification", "eventBridge").Logger()

type PersonNotification struct {
	client			*eventbridge.EventBridge
	eventSource		string
	eventBusName 	string
}

func NewPersonNotification(eventSource string, eventBusName string ) (*PersonNotification, error){
	childLogger.Debug().Msg("NewPersonNotification")

	region := os.Getenv("AWS_REGION")
    awsSession, err := session.NewSession(&aws.Config{
        Region: aws.String(region)},
    )
	if err != nil {
		childLogger.Error().Err(err).Msg("error message") 
		return nil, erro.ErrCreateSession
	}
	return &PersonNotification{
		client: eventbridge.New(awsSession),
		eventSource: eventSource,
		eventBusName: eventBusName,
	}, nil
}

func (n *PersonNotification) PutEvent(person domain.Person, eventType string) error {
	childLogger.Debug().Msg("PutEvent")

	payload, err := json.Marshal(person)
	if err != nil {
		childLogger.Error().Err(err).Msg("error message") 
		return erro.ErrUnmarshal
	}

	entries := []*eventbridge.PutEventsRequestEntry{{
		EventBusName: aws.String(n.eventBusName),
		Source:       aws.String(n.eventSource),
		DetailType:   aws.String(eventType),
		Detail:       aws.String(string(payload)),
	}}

	_, err = n.client.PutEvents(&eventbridge.PutEventsInput{Entries: entries})
	if err != nil {
		childLogger.Error().Err(err).Msg("error message") 
		return erro.ErrPutEvent
	}

	return nil
}