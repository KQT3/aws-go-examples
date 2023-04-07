package utils_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"aws-go-examples/internal/utils"
)

func TestGetUser(t *testing.T) {
	// Set up Gin context and mock database
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	db := map[string]string{
		"john": "123",
		"jane": "456",
	}

	// Test case 1: User exists in database
	c.Params = append(c.Params, gin.Param{Key: "name", Value: "john"})
	utils.GetUser(c, db)
	assert.Equal(t, http.StatusOK, c.Writer.Status())
	assert.JSONEq(t, `{"user": "john", "value": "123"}`, c.Writer.Body.String())

	// Test case 2: User does not exist in database
	c.Params = append(c.Params, gin.Param{Key: "name", Value: "bob"})
	utils.GetUser(c, db)
	assert.Equal(t, http.StatusOK, c.Writer.Status())
	assert.JSONEq(t, `{"user": "bob", "status": "no value"}`, c.Writer.Body.String())
}
