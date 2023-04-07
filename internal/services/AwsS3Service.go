package services

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type AwsCredentials struct {
	AccessKey string
	SecretKey string
}

const S3Bucket = "qt3test"

func listObjects(awsCredentials AwsCredentials) *s3.ListObjectsOutput {
	cred := credentials.NewStaticCredentials(awsCredentials.AccessKey, awsCredentials.SecretKey, "")

	sess := session.Must(session.NewSession(&aws.Config{
		Credentials: cred,
		Region:      aws.String("eu-north-1"),
	}))

	svc := s3.New(sess)

	resp, err := svc.ListObjects(&s3.ListObjectsInput{Bucket: aws.String(S3Bucket)})

	if err != nil {
		fmt.Println("Error: ", err)
	}

	return resp
}
