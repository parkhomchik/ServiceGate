package main

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1/services")
	{
		v1.GET("/", test)
		v1.GET("/{id}", test)
		v1.POST("/", test)
		v1.PUT("/{id}", test)
		v1.DELETE("/{id}")
	}

	// Simple group: v2
	v2 := router.Group("/v2/services")
	{
		v2.POST("/login", test)
		v2.POST("/submit", test)
		v2.POST("/read", test)
	}
	router.Run(":8080")
}

func test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"handlerName": "xxx"})
}
