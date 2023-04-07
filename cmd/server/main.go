package main

import (
	"aws-go-examples/internal/utils"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	db := map[string]string{"john": "123", "jane": "456"}

	r.GET("/user/:name", func(c *gin.Context) {
		utils.GetUser(c, db)
	})
	return r
}

func main() {
	r := setupRouter()

	r.Run(":8000")
}
