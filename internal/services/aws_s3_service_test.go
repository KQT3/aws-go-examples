package services

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_listObjects(t *testing.T) {
	//setup
	awsCredentials := setupAwsCredentials(t)

	//when
	objects := listObjects(awsCredentials)

	//then
	for _, object := range objects.Contents {
		fmt.Println(*object.Key)
	}

	assert.NotNil(t, objects)
}
