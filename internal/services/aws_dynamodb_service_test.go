package services

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getItem(t *testing.T) {
	//setup
	awsCredentials := setupAwsCredentials(t)
	testUserId := "ec8e3d53-5b88-438e-8187-4d4aacdebb04"

	//when
	objects, err := getItemFromDynamoDB(testUserId, awsCredentials)

	//then
	fmt.Println(objects)
	fmt.Println(err)

	assert.NotNil(t, objects)
}

func Test_getItemFromDynamoDB(t *testing.T) {
	//setup
	awsCredentials := setupAwsCredentials(t)
	testUserId := "ec8e3d53-5b88-438e-8187-4d4aacdebb04"
	imageIndex := "0"
	testImagesCollectionId0 := "5ec5d9a2-3104-4472-89ad-fc45bf4ade51"

	//when
	filtered, err := getItemFiltered(testUserId, testImagesCollectionId0, imageIndex, awsCredentials)

	//then
	fmt.Println(filtered)
	fmt.Println(err)

	assert.NotNil(t, filtered)
}
