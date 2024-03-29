package notification

import (
	"testing"

	"lambda-person/internal/core/domain"

)

var (
	tableName = "person"
	eventSource	= "lambda.person"
	eventTypeCreated =  "personCreated"
	eventTypeUpdated = 	"personUpdated"
	eventTypeDeleted = 	"personDeleted"
	eventBusName	= "event-bus-person"
	person = domain.NewPerson("PERSON-001","PERSON-001","Mr Luigi","F")
)

func TestPutEvent(t *testing.T) {
	t.Setenv("AWS_REGION", "us-east-2")

	notification, err := NewPersonNotification(eventSource,eventBusName)
	if err != nil {
		t.Errorf("Error -TestPutEvent Access EventBridge %v ", err)
	}

	eventType := "add-new-person"
	err = notification.PutEvent(*person, eventTypeCreated)
	if err != nil {
		t.Errorf("Error -TestPutEvent Access EventBridge %v ", err)
	}
}
