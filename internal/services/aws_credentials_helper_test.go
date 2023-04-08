package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_setupAwsCredentials(t *testing.T) {
	credentials := setupAwsCredentials(t)
	assert.NotEmpty(t, credentials.AccessKey, "AccessKey is empty")
	assert.NotEmpty(t, credentials.SecretKey, "SecretKey is empty")
}
