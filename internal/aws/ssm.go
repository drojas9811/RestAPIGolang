package aws

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"log"
)

type clientSSM struct {
	awsConfig *aws.Config
	ssmClient *ssm.Client
}

func NewSSMConfig() (*clientSSM, error) {
	// Create a new config with the credentials provided
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Printf("An error has occurred during the execution time. Error: %s", err.Error())
		return &clientSSM{}, err
	}
	client:= ssm.NewFromConfig(cfg)
	return &clientSSM{awsConfig: &cfg, ssmClient: client}, nil
}
/*
func (a *clientSSM) NewSSMClient() {
	// Create a new Secrets Manager client
	a.ssmClient = ssm.NewFromConfig(*a.awsConfig)
}
*/

// GetParameter retrieves a parameter from AWS Parameter Store.
func (a *clientSSM) GetParameterSSM(parameterName string) (string, error) {
	// Create an input for the GetParameter API
	value := true
	input := &ssm.GetParameterInput{
		Name:           aws.String(parameterName),
		WithDecryption: &value,
	}
	// Retrieve the parameter value from Parameter Store
	result, err := a.ssmClient.GetParameter(context.TODO(), input)
	if err != nil {
		log.Printf("An error has occurred during the execution time. Error: %s", err)
		return "", errors.New("error")
	}
	return *result.Parameter.Value, nil
}
func (a *clientSSM) SetParameterSSM(input string) error {
	//setting
	return nil
}
