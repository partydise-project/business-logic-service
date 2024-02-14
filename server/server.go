package server

import "github.com/gin-gonic/gin"

func InicializeRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(201, gin.H{
			"Welcome": "Welcome to the business-logic-service",
		})
	})

	return r
}
