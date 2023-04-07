package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context, db map[string]string) {
	user := c.Param("name")
	value, ok := db[user]
	if ok {
		c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
	}
}
