package main

import(
//	"fmt"
	"os"
	"context"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"lambda-person/internal/adapter/handler"
	"lambda-person/internal/repository"
	"lambda-person/internal/services"
	"lambda-person/internal/adapter/notification"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"

)

var (
	logLevel = zerolog.DebugLevel // InfoLevel DebugLevel
	tableName 		= "person_tenant"
	version 		= "lambda person (github) version 2.3"
	eventSource		=	"lambda.person"
	eventBusName	=	"event-bus-person"
	response 			*events.APIGatewayProxyResponse
	personRepository	*repository.PersonRepository
	personService 		*services.PersonService
	personHandler 		*handler.PersonHandler
	personNotification *notification.PersonNotification
)

func getEnv() {
	if os.Getenv("TABLE_NAME") !=  "" {
		tableName = os.Getenv("TABLE_NAME")
	}
	if os.Getenv("LOG_LEVEL") !=  "" {
		if (os.Getenv("LOG_LEVEL") == "DEBUG"){
			logLevel = zerolog.DebugLevel
		}else if (os.Getenv("LOG_LEVEL") == "INFO"){
			logLevel = zerolog.InfoLevel
		}else if (os.Getenv("LOG_LEVEL") == "ERROR"){
				logLevel = zerolog.ErrorLevel
		}else {
			logLevel = zerolog.InfoLevel
		}
	}
	if os.Getenv("VERSION") !=  "" {
		version = os.Getenv("VERSION")
	}
}

func init(){
	log.Debug().Msg("*** init")
	zerolog.SetGlobalLevel(logLevel)
	getEnv()
}

func main(){
	log.Debug().Msg("*** main lambda-card (go) v 2.3")
	log.Debug().Msg("-------------------")
	log.Debug().Str("version", version).
				Str("tableName", tableName).
				Msg("Enviroment Variables")
	log.Debug().Msg("--------------------")

	personRepository, err := repository.NewPersonRepository(tableName)
	if err != nil {
		panic("configuration error NewPersonRepository(), " + err.Error())
	}
	personNotification, err = notification.NewPersonNotification(eventSource,eventBusName)
	if err != nil{
		panic("configuration error NewPersonNotification(), " + err.Error())
	}
	personService		= services.NewPersonService(*personRepository, *personNotification)
	personHandler		= handler.NewPersonHandler(*personService)

	lambda.Start(lambdaHandler)
}

func lambdaHandler(ctx context.Context, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	log.Debug().Msg("lambdaHandler")
	log.Debug().Msg("-------------------")
	log.Debug().Str("req.Body", req.Body).
				Msg("APIGateway Request.Body")
	log.Debug().Msg("--------------------")

	switch req.HTTPMethod {
		case "GET":
			if (req.Resource == "/person/list"){
				response, _ = personHandler.ListPerson()
			}else if (req.Resource == "/personaddress/list"){
				response, _ = personHandler.ListPersonAddress()
			}else if (req.Resource == "/person/{id}"){
				response, _ = personHandler.GetPerson(req)
			}else if (req.Resource == "/personaddress/{id}"){
				response, _ = personHandler.GetPersonAddress(req)
			}else if (req.Resource == "/version"){
				response, _ = personHandler.GetVersion(version)
			}else {
				response, _ = personHandler.UnhandledMethod()
			}
		case "POST":
			if (req.Resource == "/person"){
				response, _ = personHandler.AddPerson(req)
			} else if (req.Resource == "/personaddress") {
				response, _ = personHandler.AddPersonAddress(req)
			}else {
				response, _ = personHandler.UnhandledMethod()
			}
		case "DELETE":
			response, _ = personHandler.DeletePerson(req)
		case "PUT":
			response, _ = personHandler.UpdatePerson(req)
		default:
			response, _ = personHandler.UnhandledMethod()
	}

	return response, nil
}