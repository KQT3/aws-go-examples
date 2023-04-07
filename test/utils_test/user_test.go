package utils_test

import (
	"aws-go-examples/internal/utils"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	// given
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	db := map[string]string{
		"alice": "123",
		"bob":   "456",
	}

	// when
	c.Params = append(c.Params, gin.Param{Key: "name", Value: "alice"})
	utils.GetUser(c, db)

	// then
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"user": "alice", "value": "123"}`, w.Body.String())
}

func Test(t *testing.T) {
	fmt.Println(t)
	daw := utils.Daw()
	fmt.Println("daw: ", daw)
}
