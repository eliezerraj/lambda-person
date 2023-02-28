package main

import(
	"fmt"

	"lambda-person/internal/handlers"
	"lambda-person/internal/repository"
	"lambda-person/internal/services"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"

)

var (
	tableName = "person_tenant"
	version 	= "lambda person version 1.0"
	response 			*events.APIGatewayProxyResponse
	personRepository	*repository.PersonRepository
	personService 		*services.PersonService
	personHandler 		*handlers.PersonHandler
  )

func main(){
	fmt.Println("Main Person (GO) v 2.0")
	
	personRepository, err := repository.NewPersonRepository(tableName)
	if err != nil {
		return
	}
	personService		= services.NewPersonService(*personRepository)
	personHandler		= handlers.NewPersonHandler(*personService)

	lambda.Start(handler)
}

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	fmt.Println("main.handler")
	fmt.Println("-----------------------------")
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