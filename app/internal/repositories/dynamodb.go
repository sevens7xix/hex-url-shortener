package repositories

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/sevens7xix/hex-url-shortener/app/internal/core/domain"
)

type DynamoDBRepository struct {
	client *dynamodb.Client
}

func NewDynamoDBRepository() *DynamoDBRepository {
	awsRegion := os.Getenv("AWS_REGION")
	awsEndpoint := os.Getenv("AWS_ENDPOINT")

	awsRegion = "us-west-1"
	awsEndpoint = "http://localhost:4566"

	credentials := credentials.StaticCredentialsProvider{
		Value: aws.Credentials{
			AccessKeyID: "123", SecretAccessKey: "xyz",
			Source: "Hard-coded credentials; values are irrelevant for local DynamoDB",
		},
	}

	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if awsEndpoint != "" {
			return aws.Endpoint{
				URL:           awsEndpoint,
				SigningRegion: awsRegion,
			}, nil
		}

		// returning EndpointNotFoundError will allow the service to fallback to its default resolution
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	awsCfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(awsRegion),
		config.WithEndpointResolverWithOptions(customResolver),
		config.WithCredentialsProvider(credentials),
		config.WithClientLogMode(aws.LogRequest|aws.LogRetries))

	if err != nil {
		log.Fatalf("Failed to load SDK Config, %v", err)
	}

	svc := dynamodb.NewFromConfig(awsCfg)

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{{
			AttributeName: aws.String("short"),
			AttributeType: types.ScalarAttributeTypeS,
		}},
		KeySchema: []types.KeySchemaElement{{
			AttributeName: aws.String("short"),
			KeyType:       types.KeyTypeHash,
		}},
		TableName: aws.String("data"),
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
	}

	_, err = svc.CreateTable(context.Background(), input)

	if err != nil {
		log.Printf("Couldn't create table %v. Here's why: %v\n", "data", err)
	} else {
		waiter := dynamodb.NewTableExistsWaiter(svc)
		err = waiter.Wait(context.TODO(), &dynamodb.DescribeTableInput{
			TableName: aws.String("data")}, 5*time.Minute)
		if err != nil {
			log.Printf("Wait for table exists failed. Here's why: %v\n", err)
		}
	}

	return &DynamoDBRepository{
		client: svc,
	}
}

func (repository *DynamoDBRepository) Create(Data domain.Data) error {
	item, err := attributevalue.MarshalMap(Data)

	if err != nil {
		return err
	}

	_, err = repository.client.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String("data"),
		Item:      item,
	})

	if err != nil {
		return err
	}

	return nil
}

func (repository *DynamoDBRepository) Get(shortURL string) (domain.Data, error) {
	data := domain.NewData("", shortURL)
	short, err := attributevalue.Marshal(shortURL)

	if err != nil {
		return domain.Data{}, nil
	}

	key := map[string]types.AttributeValue{"short": short}

	response, err := repository.client.GetItem(context.Background(), &dynamodb.GetItemInput{
		Key:       key,
		TableName: aws.String("data"),
	})

	if err != nil {
		return domain.Data{}, nil
	} else {
		err = attributevalue.UnmarshalMap(response.Item, &data)
		if err != nil {
			log.Printf("Couldn't unmarshal response. Here's why: %v\n", err)
		}
	}

	return data, nil
}
