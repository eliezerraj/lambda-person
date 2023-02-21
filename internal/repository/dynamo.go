package repository

import(
	"log"
	"os"

	"lambda-person/internal/erro"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type PersonRepository struct {
	client 		dynamodbiface.DynamoDBAPI
	tableName   *string
}

func NewPersonRepository(tableName string) (*PersonRepository, error){
	log.Printf("----------------------------")
	log.Print("- repository.NewPersonRepository tableName: ", tableName) 

	region := os.Getenv("AWS_REGION")
    awsSession, err := session.NewSession(&aws.Config{
        Region: aws.String(region)},
    )
	if err != nil {
		return nil, erro.ErrOpenDatabase
	}

	return &PersonRepository {
		client: dynamodb.New(awsSession),
		tableName: aws.String(tableName),
	}, nil
}
