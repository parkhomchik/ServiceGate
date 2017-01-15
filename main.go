package main

import (
	"net/http"

	// "github.com/jinzhu/gorm/dialects/postgres"
	"fmt"

	"github.com/parkhomchik/ServiceGate/models/database"
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	db := database.InitDb()
	database.Migrate(db)

	router := gin.Default()

	// version 1
	service := router.Group("/service")
	{
		service.POST("/", request) // INSERT
	}

	portal := router.Group("/portal")
	{
		portal.GET("/")
	}

	router.Run(":8080")
}

func request(c *gin.Context) {
	db := database.InitDb()

	var json database.Request
	if c.BindJSON(&json) == nil {

		db.Create(&json)

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
		fmt.Println(c.Request)
		c.JSON(http.StatusBadRequest, gin.H{"status": "StatusBadRequest"})
	}

}

func portalGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
