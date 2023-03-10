package repository

import(
	"os"

	"github.com/rs/zerolog/log"
	"lambda-person/internal/erro"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var childLogger = log.With().Str("repository", "PersonRepository").Logger()

type PersonRepository struct {
	client 		dynamodbiface.DynamoDBAPI
	tableName   *string
}

func NewPersonRepository(tableName string) (*PersonRepository, error){
	childLogger.Debug().Msg("*** NewPersonRepository")

	region := os.Getenv("AWS_REGION")
    awsSession, err := session.NewSession(&aws.Config{
        Region: aws.String(region)},
    )
	if err != nil {
		childLogger.Error().Err(err).Msg("error message")
		return nil, erro.ErrOpenDatabase
	}

	return &PersonRepository {
		client: dynamodb.New(awsSession),
		tableName: aws.String(tableName),
	}, nil
}
