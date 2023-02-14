package handlers

import(
	"log"
	"net/http"
	"encoding/json"

	"lambda-person/internal/services"
	"lambda-person/internal/erro"
	//"lambda-person/internal/ports"
	"lambda-person/internal/core/domain"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-lambda-go/events"

)

type PersonHandler struct {
	personService services.PersonService
}

type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

func (h *PersonHandler) UnhandledMethod() (*events.APIGatewayProxyResponse, error){
	return ApiHandlerResponse(http.StatusMethodNotAllowed, ErrorBody{aws.String(erro.ErrMethodNotAllowed.Error())})
}

func NewPersonHandler(personService services.PersonService) *PersonHandler{
	log.Printf("----------------------------")
	log.Print("- handler.NewPersonHandler") 

	return &PersonHandler{
		personService: personService,
	}
}

func (h *PersonHandler) AddPerson(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	log.Printf("+++++++++++++++++++++++++++++++++")
	log.Printf("- handlers.AddPerson -")

	//person := domain.NewPerson("333","mahone","M")
    var person domain.Person
    if err := json.Unmarshal([]byte(req.Body), &person); err != nil {
        return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
    }

	response, err := h.personService.AddPerson(person)
	if err != nil {
		return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}

	handlerResponse, err := ApiHandlerResponse(http.StatusOK, response)
	if err != nil {
		return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}
	return handlerResponse, nil
}

func (h *PersonHandler) UpdatePerson(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	log.Printf("+++++++++++++++++++++++++++++++++")
	log.Printf("- handlers.UpdatePerson -")

	//person := domain.NewPerson("333","mahone champion","M")
	
	var person domain.Person
    if err := json.Unmarshal([]byte(req.Body), &person); err != nil {
        return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
    }

	response, err := h.personService.UpdatePerson(person)
	if err != nil {
		return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}

	handlerResponse, err := ApiHandlerResponse(http.StatusOK, response)
	if err != nil {
		return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}
	return handlerResponse, nil
}

func (h *PersonHandler) GetPerson(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	log.Printf("+++++++++++++++++++++++++++++++++")
	log.Printf("- handlers.GetPerson -")

	id := req.PathParameters["id"]
	if len(id) == 0 {
		return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(erro.ErrQueryEmpty.Error())})
	}

	response, err := h.personService.GetPerson(id)
	if err != nil {
		return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}

	handlerResponse, err := ApiHandlerResponse(http.StatusOK, response)
	if err != nil {
		return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}
	log.Printf("- handlers.GetPerson - handlerResponse :", handlerResponse)

	return handlerResponse, nil
}

func (h *PersonHandler) DeletePerson(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	log.Printf("+++++++++++++++++++++++++++++++++")
	log.Printf("- handlers.DeletePerson -")

	id := req.PathParameters["id"]
	if len(id) == 0 {
		return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(erro.ErrQueryEmpty.Error())})
	}
	
	err := h.personService.DeletePerson(id)
	if err != nil {
		return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}

	return ApiHandlerResponse(http.StatusOK, nil)
}

func (h *PersonHandler) ListPerson() (*events.APIGatewayProxyResponse, error) {
	log.Printf("+++++++++++++++++++++++++++++++++")
	log.Printf("- handlers.ListPerson -")

	response, err := h.personService.ListPerson()
	if err != nil {
		return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}

	handlerResponse, err := ApiHandlerResponse(http.StatusOK, response)
	if err != nil {
		return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}
	log.Printf("- handlers.ListPerson - handlerResponse :", handlerResponse)

	return handlerResponse, nil
}