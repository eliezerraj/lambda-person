package repository

import(
	"context"
	"lambda-person/internal/core/domain"
	"lambda-person/internal/erro"
	"fmt"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

type PersonAddressRecord struct {
    ID     	string `json:"id,omitempty"`
	SK     	string `json:"sk,omitempty"`
    Address domain.Address  `json:"address,omitempty"`
}

func (r *PersonRepository) AddPersonAddress(ctx context.Context,personAddress domain.PersonAddress) (*domain.PersonAddress, error){
	childLogger.Debug().Msg("AddPersonAddress")
	//log.Printf("- repository.AddPersonAddress - adresses : ", adresses)

	for _, item_address := range personAddress.Addresses {

		concat_sk := fmt.Sprintf("ADDRESS:%s", item_address.ID)

		item_to_marshall := PersonAddressRecord{
			ID: personAddress.Person.ID,
			SK: concat_sk,

			Address: item_address,
		}

		item, err := dynamodbattribute.MarshalMap(item_to_marshall)
		if err != nil {
			childLogger.Error().Err(err).Msg("error message")
			return nil, erro.ErrUnmarshal
		}

		transactItems := []*dynamodb.TransactWriteItem{}
		transactItems = append(transactItems, &dynamodb.TransactWriteItem{Put: &dynamodb.Put{
			TableName: r.tableName,
			Item:      item,
		}})
	
		transaction := &dynamodb.TransactWriteItemsInput{TransactItems: transactItems}
		if err := transaction.Validate(); err != nil {
			childLogger.Error().Err(err).Msg("error message")
			return nil, erro.ErrInsert
		}
	
		_, err = r.client.TransactWriteItemsWithContext(ctx, transaction)
		if err != nil {
			childLogger.Error().Err(err).Msg("error message")
			return nil, erro.ErrInsert
		}
	}

	return &personAddress , nil
}

func (r *PersonRepository) ListPersonAddress(ctx context.Context) (*[]domain.PersonAddress, error){
	childLogger.Debug().Msg("ListPersonAddress")

	key := &dynamodb.ScanInput{
		TableName:	r.tableName,
	}

	result, err := r.client.ScanWithContext(ctx, key)
	if err != nil {
		childLogger.Error().Err(err).Msg("error message")
		return nil, erro.ErrList
	}
	//log.Printf("result => ", result)

	personAddress := []domain.PersonAddress{}
	personAddressRecord := PersonAddressRecord{}
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &personAddressRecord)
    if err != nil {
		childLogger.Error().Err(err).Msg("error message")
		return nil, erro.ErrUnmarshal
    }

	//log.Printf("personAddressRecord => ", personAddressRecord)
	if len(personAddress) == 0 {
		return nil, erro.ErrNotFound
	} else {
		return &personAddress, nil
	}
}

func (r *PersonRepository) QueryPersonAddress(ctx context.Context,id string) (*domain.PersonAddress, error){
	childLogger.Debug().Msg("QueryPersonAddress")

	var keyCond expression.KeyConditionBuilder
	id = fmt.Sprintf("PERSON-%s", id)
	sk := "ADDRESS"
	keyCond = expression.KeyAnd(
		expression.Key("id").Equal(expression.Value(id)),
		expression.Key("sk").BeginsWith(sk),
	)

	expr, err := expression.NewBuilder().
							WithKeyCondition(keyCond).
							Build()
	if err != nil {
		return nil, err
	}

	key := &dynamodb.QueryInput{
		TableName:                 r.tableName,
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
	}

	result, err := r.client.QueryWithContext(ctx, key)
	if err != nil {
		childLogger.Error().Err(err).Msg("error message")
		return nil, erro.ErrList
	}
	//log.Printf("result => ", result)

	personAddressRecord := []PersonAddressRecord{}
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &personAddressRecord)
    if err != nil {
		childLogger.Error().Err(err).Msg("error message")
		return nil, erro.ErrUnmarshal
    }
	//log.Printf("personAddressRecord => ", personAddressRecord)

	listAdresses := []domain.Address{}
	person := domain.Person{}
	for _, result_personAddressRecord := range personAddressRecord{
		person.ID = result_personAddressRecord.ID
		address := domain.NewAddress(result_personAddressRecord.Address.ID,
									result_personAddressRecord.Address.SK,
									result_personAddressRecord.Address.Street,
									result_personAddressRecord.Address.StreetNumber,
									result_personAddressRecord.Address.ZipCode)
		
		listAdresses = append(listAdresses, *address)	
	}

	personAddress := domain.NewPersonAddress(person, listAdresses)
	return personAddress, nil
}
