package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"log"
)

type clientDynamodb struct {
	awsConfig      *aws.Config
	dynamodbClient *dynamodb.Client
}

func NewDynamodbConfig() (*clientDynamodb, error) {
	// Create a new config with the credentials provided
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Printf("An error has occurred during the execution time. Error: %s", err.Error())
		return &clientDynamodb{}, err
	}
	return &clientDynamodb{awsConfig: &cfg}, nil
}

func (a *clientDynamodb) NewDynamodbClient() {
	// Create a new Secrets Manager client
	a.dynamodbClient = dynamodb.NewFromConfig(*a.awsConfig)
}

// GetParameter retrieves a parameter from AWS Parameter Store.
func (a *clientSSM) GetParameterDynamodb(parameterName string) (string, error) {
	// Create an input for the GetParameter API
	return "", nil
}
func (a *clientSSM) SetParameterDynamodb(input string) error {
	//setting
	return nil
}
