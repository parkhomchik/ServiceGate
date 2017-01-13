package main

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	router := gin.Default()

	// version 1
	v1 := router.Group("/v1/service")
	{
		v1.POST("/", request) // INSERT
	}

<<<<<<< HEAD
	// version 2
	v2 := router.Group("/v2/service")
=======
	// Simple group: v2
	v2 := router.Group("/v2/services")
>>>>>>> a1a8ed669e0805368dbca2e6684af3a0108e8df8
	{
		v2.POST("/", requestv2) // INSERT
	}
	router.Run(":8080")
}

func request(c *gin.Context) {
	//var json Request
	if c.BindJSON(&json) == nil {
	}
	c.JSON(http.StatusOK, json)
}

func requestv2(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}
