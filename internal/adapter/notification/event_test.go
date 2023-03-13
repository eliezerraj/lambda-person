package notification

import (
	"testing"

	"github.com/lambda-card/internal/core/domain"

)

var (
	tableName = "person"
	eventSource	= "lambda-person"
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
	err = notification.PutEvent(*person, eventType)
	if err != nil {
		t.Errorf("Error -TestPutEvent Access EventBridge %v ", err)
	}
}
