package services

import (
	"aws-go-examples/internal/config"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const tableName = "user_images"

func getItemFromDynamoDB(userId string, awsCredentials config.AwsCredentials) (*dynamodb.GetItemOutput, error) {
	dynamoDBService := createClient(awsCredentials)

	key := map[string]*dynamodb.AttributeValue{
		"userId": {
			S: aws.String(userId),
		},
	}

	input := &dynamodb.GetItemInput{
		Key:       key,
		TableName: aws.String(tableName),
	}

	return dynamoDBService.GetItem(input)
}

func getItemFiltered(userId string, imagesCollectionId string, imageIndex string, awsCredentials config.AwsCredentials) (*dynamodb.QueryOutput, error) {
	dynamoDBService := createClient(awsCredentials)

	expressionAttributeValues := map[string]*dynamodb.AttributeValue{
		":userId": {
			S: aws.String(userId),
		},
		":imagesCollectionId": {
			S: aws.String(imagesCollectionId),
		},
	}

	queryRequest := &dynamodb.QueryInput{
		TableName:                 aws.String(tableName),
		KeyConditionExpression:    aws.String("userId = :userId"),
		ExpressionAttributeValues: expressionAttributeValues,
		FilterExpression:          aws.String(fmt.Sprintf("contains(imagesCollection[%s].imagesCollectionId, :imagesCollectionId)", imageIndex)),
		ProjectionExpression:      aws.String(fmt.Sprintf("imagesCollection[%s]", imageIndex)),
	}

	return dynamoDBService.Query(queryRequest)
}

func createClient(awsCredentials config.AwsCredentials) *dynamodb.DynamoDB {
	cred := credentials.NewStaticCredentials(awsCredentials.AccessKey, awsCredentials.SecretKey, "")

	newSession := session.Must(session.NewSession(&aws.Config{
		Credentials: cred,
		Region:      aws.String("us-east-1"),
	}))

	dynamoDBService := dynamodb.New(newSession)
	return dynamoDBService
}
