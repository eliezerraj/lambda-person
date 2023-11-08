package repository

import(
	"context"
	"lambda-person/internal/core/domain"
	"lambda-person/internal/erro"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"

)

func (r *PersonRepository) DeletePerson(ctx context.Context, id string, sk string) (error) {
	childLogger.Debug().Msg("DeletePerson")

	key := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
			"sk": {
				S: aws.String(sk),
			},
		},
		TableName: r.tableName,
	}

	_, err := r.client.DeleteItemWithContext(ctx, key)
	if err != nil {
		childLogger.Error().Err(err).Msg("error message") 
		return erro.ErrDelete
	}

	return nil
}

func (r *PersonRepository) AddPerson(ctx context.Context, person domain.Person) (*domain.Person, error){
	childLogger.Debug().Msg("AddPerson")

	item, err := dynamodbattribute.MarshalMap(person)
	if err != nil {
		childLogger.Error().Err(err).Msg("error message")
		return nil, erro.ErrUnmarshal
	}

	//log.Printf("- repository.AddPerson - item : ", item)

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

	return &person , nil
}

func (r *PersonRepository) ListPerson(ctx context.Context) (*[]domain.Person, error){
	childLogger.Debug().Msg("ListPerson")

	expr, err := expression.NewBuilder().
							WithFilter(       expression.And(
								expression.Contains(expression.Name("sk"), "PERSON"),
								expression.Contains(expression.Name("id"), "PERSON"),
								),).
							Build()
	if err != nil {
		childLogger.Error().Err(err).Msg("error message")
		return nil, erro.ErrPreparedQuery
	}

	key := &dynamodb.ScanInput{
		TableName:                 	r.tableName,
		ExpressionAttributeNames:  	expr.Names(),
		ExpressionAttributeValues: 	expr.Values(),
		FilterExpression:    		expr.Filter(),
	}

	result, err := r.client.ScanWithContext(ctx, key)
	if err != nil {
		childLogger.Error().Err(err).Msg("error message")
		return nil, erro.ErrList
	}
	//log.Printf("result => ", result)

	persons := []domain.Person{}
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &persons)
    if err != nil {
		childLogger.Error().Err(err).Msg("error message")
		return nil, erro.ErrUnmarshal
    }

	if len(persons) == 0 {
		return nil, erro.ErrNotFound
	} else {
		return &persons, nil
	}
}

func (r *PersonRepository) GetPerson(ctx context.Context, id string) (*domain.Person, error){
	childLogger.Debug().Msg("GetPerson")

	var keyCond expression.KeyConditionBuilder

	keyCond = expression.KeyAnd(
		expression.Key("id").Equal(expression.Value(id)),
		expression.Key("sk").BeginsWith(id),
	)

	expr, err := expression.NewBuilder().
							WithKeyCondition(keyCond).
							Build()
	if err != nil {
		childLogger.Error().Err(err).Msg("error message")
		return nil, erro.ErrPreparedQuery
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
		return nil, erro.ErrQuery
	}

	person := []domain.Person{}
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &person)
    if err != nil {
		childLogger.Error().Err(err).Msg("error message")
		return nil, erro.ErrUnmarshal
    }

	if len(person) == 0 {
		return nil, erro.ErrNotFound
	} else {
		return &person[0], nil
	}
}