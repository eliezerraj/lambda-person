package repository

import(
	"log"
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

func (r *PersonRepository) AddPersonAddress(person domain.Person, adresses []domain.Address) (*domain.PersonAddress, error){
	log.Printf("+++++++++++++++++++++++++++++++++")
	//log.Printf("- repository.AddPersonAddress - person : ", person)
	//log.Printf("- repository.AddPersonAddress - adresses : ", adresses)

	for _, item_address := range adresses {

		concat_sk := fmt.Sprintf("ADDRESS:%s", item_address.ID)

		item_to_marshall := PersonAddressRecord{
			ID: person.ID,
			SK: concat_sk,

			Address: item_address,
		}
		//log.Printf("- repository.AddPerson - item_to_marshall : ", item_to_marshall)
		
		item, err := dynamodbattribute.MarshalMap(item_to_marshall)
		if err != nil {
			log.Print("erro :", err) 
			return nil, erro.ErrUnmarshal
		}

		transactItems := []*dynamodb.TransactWriteItem{}
		transactItems = append(transactItems, &dynamodb.TransactWriteItem{Put: &dynamodb.Put{
			TableName: r.tableName,
			Item:      item,
		}})
	
		transaction := &dynamodb.TransactWriteItemsInput{TransactItems: transactItems}
		if err := transaction.Validate(); err != nil {
			log.Print("erro :", err) 
			return nil, erro.ErrInsert
		}
	
		_, err = r.client.TransactWriteItems(transaction)
		if err != nil {
			log.Print("erro :", err) 
			return nil, erro.ErrInsert
		}
	}

	personAddress := domain.PersonAddress{person, adresses}
	//log.Printf("- repository.AddPersonAddress - personAddress : ", personAddress)

	return &personAddress , nil
}

func (r *PersonRepository) ListPersonAddress() (*[]domain.PersonAddress, error){
	log.Printf("+++++++++++++++++++++++++++++++++")
	log.Printf("- repository.ListPersonAddress -")

	key := &dynamodb.ScanInput{
		TableName:	r.tableName,
	}

	result, err := r.client.Scan(key)
	if err != nil {
		log.Printf("Erro(ErrList):", err) 
		return nil, erro.ErrList
	}
	log.Printf("result => ", result)

	personAddress := []domain.PersonAddress{}
	
	return &personAddress, nil
}

func (r *PersonRepository) QueryPersonAddress(person domain.Person) (*[]domain.PersonAddress, error){
	log.Printf("+++++++++++++++++++++++++++++++++")
	log.Printf("- repository.QueryPersonAddress -")

	var keyCond expression.KeyConditionBuilder
	sk := "ADDRESS"
	keyCond = expression.KeyAnd(
		expression.Key("id").Equal(expression.Value(person.ID)),
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

	result, err := r.client.Query(key)
	if err != nil {
		log.Printf("Erro(ErrList):", err) 
		return nil, erro.ErrList
	}
	log.Printf("result => ", result)

	personAddress := []domain.PersonAddress{}
	
	return &personAddress, nil
}