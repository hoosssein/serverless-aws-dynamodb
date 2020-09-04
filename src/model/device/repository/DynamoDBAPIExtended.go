package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"os"
)

type DynamoDBAPIExtended interface {
	CreateSession() dynamodbiface.DynamoDBAPI
	MarshalMap(in interface{}) (map[string]*dynamodb.AttributeValue, error)
	UnmarshalMap(m map[string]*dynamodb.AttributeValue, out interface{}) error
}

type DynamoDBAPIExtendedImpl struct{}

func (d DynamoDBAPIExtendedImpl) CreateSession() dynamodbiface.DynamoDBAPI {
	var endpoint *string = nil
	if os.Getenv("AWS_ENDPOINT") != "" {
		endpoint = aws.String(os.Getenv("AWS_ENDPOINT"))
	}
	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String(os.Getenv("AWS_REGION")),
		Endpoint: endpoint,
	}))
	return dynamodb.New(sess)
}

func (d DynamoDBAPIExtendedImpl) MarshalMap(in interface{}) (map[string]*dynamodb.AttributeValue, error) {
	return dynamodbattribute.MarshalMap(in)
}

func (d DynamoDBAPIExtendedImpl) UnmarshalMap(m map[string]*dynamodb.AttributeValue, out interface{}) error {
	return dynamodbattribute.UnmarshalMap(m, out)
}

func NewDynamoDBAPIExtended() DynamoDBAPIExtended {
	return DynamoDBAPIExtendedImpl{}
}
