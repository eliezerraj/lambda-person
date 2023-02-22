package repository

import(
	"log"
	"lambda-person/internal/core/domain"
	"lambda-person/internal/erro"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"

)

func (r *PersonRepository) DeletePerson(id string, sk string) (error) {
	log.Printf("+++++++++++++++++++++++++++++++++")
	log.Printf("- repository.DeletePerson - id : ", id)

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

	_, err := r.client.DeleteItem(key)
	if err != nil {
		log.Print("erro :", err) 
		return erro.ErrDelete
	}

	return nil
}

func (r *PersonRepository) AddPerson(person domain.Person) (*domain.Person, error){
	log.Printf("+++++++++++++++++++++++++++++++++")
	log.Printf("- repository.AddPerson - person : ", person)

	item, err := dynamodbattribute.MarshalMap(person)
	if err != nil {
		log.Print("erro :", err) 
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
		log.Printf("erro :", err) 
		return nil, erro.ErrInsert
	}

	_, err = r.client.TransactWriteItems(transaction)
	if err != nil {
		log.Printf("erro :", err) 
		return nil, erro.ErrInsert
	}

	return &person , nil
}

func (r *PersonRepository) ListPerson() (*[]domain.Person, error){
	log.Printf("+++++++++++++++++++++++++++++++++")
	log.Printf("- repository.ListPerson -")

	expr, err := expression.NewBuilder().
							WithFilter(       expression.And(
								expression.Contains(expression.Name("sk"), "PERSON"),
								expression.Contains(expression.Name("id"), "PERSON"),
								),).
							Build()
	if err != nil {
		log.Printf("erro :", err) 
		return nil, erro.ErrPreparedQuery
	}


	key := &dynamodb.ScanInput{
		TableName:                 	r.tableName,
		ExpressionAttributeNames:  	expr.Names(),
		ExpressionAttributeValues: 	expr.Values(),
		FilterExpression:    		expr.Filter(),
	}

	result, err := r.client.Scan(key)
	if err != nil {
		log.Printf("Erro(ErrList):", err) 
		return nil, erro.ErrList
	}
	log.Printf("result => ", result)

	persons := []domain.Person{}
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &persons)
    if err != nil {
		log.Printf("erro(ErrUnmarshal) :", err) 
		return nil, erro.ErrUnmarshal
    }

	if len(persons) == 0 {
		return nil, erro.ErrNotFound
	} else {
		return &persons, nil
	}
}

func (r *PersonRepository) GetPerson(id string) (*domain.Person, error){
	log.Printf("+++++++++++++++++++++++++++++++++")
	log.Printf("- repository.GetPerson - id : ", id)

	var keyCond expression.KeyConditionBuilder

	keyCond = expression.KeyAnd(
		expression.Key("id").Equal(expression.Value(id)),
		expression.Key("sk").BeginsWith(id),
	)

	expr, err := expression.NewBuilder().
							WithKeyCondition(keyCond).
							Build()
	if err != nil {
		log.Printf("erro :", err) 
		return nil, erro.ErrPreparedQuery
	}

	key := &dynamodb.QueryInput{
								TableName:                 r.tableName,
								ExpressionAttributeNames:  expr.Names(),
								ExpressionAttributeValues: expr.Values(),
								KeyConditionExpression:    expr.KeyCondition(),
	}

	result, err := r.client.Query(key)
	if err != nil {
		log.Printf("erro :", err) 
		return nil, erro.ErrQuery
	}

	person := []domain.Person{}
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &person)
    if err != nil {
		log.Printf("erro :", err) 
		return nil, erro.ErrUnmarshal
    }

	if len(person) == 0 {
		return nil, erro.ErrNotFound
	} else {
		return &person[0], nil
	}
}