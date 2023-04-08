package services

import (
	"aws-go-examples/internal/config"
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func setupAwsCredentials(t *testing.T) config.AwsCredentials {
	err := godotenv.Load("../../.env.test")
	if err != nil {
		t.Fatal("Error loading .env.test file")
	}
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	awsCredentials := config.AwsCredentials{
		AccessKey: accessKey,
		SecretKey: secretKey,
	}
	return awsCredentials
}
