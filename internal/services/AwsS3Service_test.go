package services

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_listObjects(t *testing.T) {
	//given
	err := godotenv.Load("../../.env.test")
	if err != nil {
		t.Fatal("Error loading .env.test file")
	}
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	awsCredentials := AwsCredentials{
		AccessKey: accessKey,
		SecretKey: secretKey,
	}

	//when
	objects := listObjects(awsCredentials)

	//then
	for _, object := range objects.Contents {
		fmt.Println(*object.Key)
	}

	assert.NotNil(t, objects)
}
