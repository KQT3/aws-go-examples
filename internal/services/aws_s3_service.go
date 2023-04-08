package services

import (
	"aws-go-examples/internal/config"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const S3Bucket = "qt3test"

func listObjects(awsCredentials config.AwsCredentials) *s3.ListObjectsOutput {
	cred := credentials.NewStaticCredentials(awsCredentials.AccessKey, awsCredentials.SecretKey, "")

	newSession := session.Must(session.NewSession(&aws.Config{
		Credentials: cred,
		Region:      aws.String("eu-north-1"),
	}))

	s3Service := s3.New(newSession)

	response, err := s3Service.ListObjects(&s3.ListObjectsInput{Bucket: aws.String(S3Bucket)})

	if err != nil {
		fmt.Println("Error: ", err)
	}

	return response
}
