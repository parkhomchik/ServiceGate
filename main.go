package main

import (
	"net/http"

	"github.com/parkhomchik/ServiceGate/models"

	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	router := gin.Default()

	// version 1
	v1 := router.Group("/v1/service")
	{
		v1.POST("/", request) // INSERT
	}

	// version 2
	v2 := router.Group("/v2/service")
	{
		v2.POST("/", requestv2) // INSERT
	}
	router.Run(":8080")
}

func request(c *gin.Context) {
	var json models.Request
	if c.BindJSON(&json) == nil {
		// check Portal
		//json.PortalName

		//check Service

		//check PortalService

		//check Command

		//check Parameters

		//check created record

		//execute

		// return
		c.JSON(http.StatusOK, json)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "StatusBadRequest"})
	}

}

func requestv2(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"status": "not implemented"})
}
