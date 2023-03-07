package main

import(
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"lambda-person/internal/handlers"
	"lambda-person/internal/repository"
	"lambda-person/internal/services"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"

)

var (
	logLevel = zerolog.DebugLevel // InfoLevel DebugLevel
	tableName 		= "person_tenant"
	version 		= "lambda person (github) version 2.0"
	response 			*events.APIGatewayProxyResponse
	personRepository	*repository.PersonRepository
	personService 		*services.PersonService
	personHandler 		*handlers.PersonHandler
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
	zerolog.SetGlobalLevel(logLevel)
	getEnv()
}

func main(){
	log.Debug().Msg("main lambda-card (go) v 1.0")
	
	personRepository, err := repository.NewPersonRepository(tableName)
	if err != nil {
		return
	}
	personService		= services.NewPersonService(*personRepository)
	personHandler		= handlers.NewPersonHandler(*personService)

	lambda.Start(handler)
}

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	log.Debug().Msg("handler")
	log.Debug().Msg("-----------------------------")
	fmt.Println(req.Resource)
	fmt.Println(req.Path)
	fmt.Println(req.HTTPMethod)
	fmt.Println(req.QueryStringParameters)
	fmt.Println(req.PathParameters)
	fmt.Println("-----------------------------")

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
			response, _ =personHandler.DeletePerson(req)
		case "PUT":
			response, _ = personHandler.UpdatePerson(req)
		default:
			response, _ = personHandler.UnhandledMethod()
	}

	return response, nil
}