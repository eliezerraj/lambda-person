package handler

import(
	"github.com/rs/zerolog/log"
	"net/http"
	"encoding/json"
	"context"

	"lambda-person/internal/services"
	"lambda-person/internal/erro"
	"lambda-person/internal/core/domain"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-lambda-go/events"

)

var childLogger = log.With().Str("handler", "PersonHandler").Logger()

var transactionSuccess	= "Transação com sucesso"

type PersonHandler struct {
	personService services.PersonService
}

type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

type MessageBody struct {
	Msg *string `json:"message,omitempty"`
}

func (h *PersonHandler) UnhandledMethod() (*events.APIGatewayProxyResponse, error){
	return ApiHandlerResponse(http.StatusMethodNotAllowed, ErrorBody{aws.String(erro.ErrMethodNotAllowed.Error())})
}

func NewPersonHandler(personService services.PersonService) *PersonHandler{
	childLogger.Debug().Msg("NewPersonHandler")

	return &PersonHandler{
		personService: personService,
	}
}

func (h *PersonHandler) AddPerson(ctx context.Context, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	childLogger.Debug().Msg("AddPerson")

    var person domain.Person
    if err := json.Unmarshal([]byte(req.Body), &person); err != nil {
        return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
    }

	response, err := h.personService.AddPerson(ctx, person)
	if err != nil {
		return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}

	handlerResponse, err := ApiHandlerResponse(http.StatusOK, response)
	if err != nil {
		return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}
	return handlerResponse, nil
}

func (h *PersonHandler) UpdatePerson(ctx context.Context,req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	childLogger.Debug().Msg("UpdatePerson")

	var person domain.Person
    if err := json.Unmarshal([]byte(req.Body), &person); err != nil {
        return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
    }

	response, err := h.personService.UpdatePerson(ctx, person)
	if err != nil {
		return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}

	handlerResponse, err := ApiHandlerResponse(http.StatusOK, response)
	if err != nil {
		return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}
	return handlerResponse, nil
}

func (h *PersonHandler) GetPerson(ctx context.Context,req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	childLogger.Debug().Msg("GetPerson")

	id := req.PathParameters["id"]
	if len(id) == 0 {
		return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(erro.ErrQueryEmpty.Error())})
	}

	response, err := h.personService.GetPerson(ctx, id)
	if err != nil {
		return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}

	handlerResponse, err := ApiHandlerResponse(http.StatusOK, response)
	if err != nil {
		return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}
	//log.Printf("- handlers.GetPerson - handlerResponse :", handlerResponse)

	return handlerResponse, nil
}

func (h *PersonHandler) DeletePerson(ctx context.Context, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	childLogger.Debug().Msg("DeletePerson")

	id := req.PathParameters["id"]
	sk := req.PathParameters["sk"]
	if len(id) == 0 || len(sk) == 0 {
		return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(erro.ErrQueryEmpty.Error())})
	}
	
	err := h.personService.DeletePerson(ctx, id, sk)
	if err != nil {
		return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}
	response := MessageBody { Msg: &transactionSuccess }
	handlerResponse, err := ApiHandlerResponse(http.StatusOK, response)
	if err != nil {
		return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}
 	log.Printf("- handlers.DeletePerson - handlerResponse :", handlerResponse)
	return handlerResponse, nil
}

func (h *PersonHandler) ListPerson(ctx context.Context) (*events.APIGatewayProxyResponse, error) {
	childLogger.Debug().Msg("ListPerson")

	response, err := h.personService.ListPerson(ctx)
	if err != nil {
		return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}

	handlerResponse, err := ApiHandlerResponse(http.StatusOK, response)
	if err != nil {
		return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}
	//log.Printf("- handlers.ListPerson - handlerResponse :", handlerResponse)

	return handlerResponse, nil
}

func (h *PersonHandler) GetVersion(version string) (*events.APIGatewayProxyResponse, error) {
	childLogger.Debug().Msg("GetVersion")

	response := MessageBody { Msg: &version }
	handlerResponse, err := ApiHandlerResponse(http.StatusOK, response)
	if err != nil {
		return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}
	//log.Printf("- handlers.GetVersion - handlerResponse :", handlerResponse)

	return handlerResponse, nil
}

//------------------------------------------

func (h *PersonHandler) ListPersonAddress(ctx context.Context) (*events.APIGatewayProxyResponse, error) {
	childLogger.Debug().Msg("ListPersonAddress")

	response, err := h.personService.ListPersonAddress(ctx)
	if err != nil {
		return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}

	handlerResponse, err := ApiHandlerResponse(http.StatusOK, response)
	if err != nil {
		return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}
	//log.Printf("- handlers.ListPerson - handlerResponse :", handlerResponse)

	return handlerResponse, nil
}

func (h *PersonHandler) GetPersonAddress(ctx context.Context, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	childLogger.Debug().Msg("GetPersonAddress")

	id := req.PathParameters["id"]
	if len(id) == 0 {
		return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(erro.ErrQueryEmpty.Error())})
	}

	response, err := h.personService.QueryPersonAddress(ctx, id)
	if err != nil {
		return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}

	handlerResponse, err := ApiHandlerResponse(http.StatusOK, response)
	if err != nil {
		return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}
	//log.Printf("- handlers.GetPersonAddress - handlerResponse :", handlerResponse)

	return handlerResponse, nil
}

func (h *PersonHandler) AddPersonAddress(ctx context.Context, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	childLogger.Debug().Msg("AddPersonAddress")

    var personAddress domain.PersonAddress
    if err := json.Unmarshal([]byte(req.Body), &personAddress); err != nil {
        return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
    }

	response, err := h.personService.AddPersonAddress(ctx, personAddress)
	if err != nil {
		return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}

	handlerResponse, err := ApiHandlerResponse(http.StatusOK, response)
	if err != nil {
		return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}
	return handlerResponse, nil
}