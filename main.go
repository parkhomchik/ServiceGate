package main

import (
	"net/http"

	"github.com/parkhomchik/ServiceGate/models"
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
	"gopkg.in/gin-gonic/gin.v1"
	"fmt"
	"log"
)

func main() {
	//DataBase
	db, err := gorm.Open("postgres", "host=localhost port=5433 user=postgres dbname=servicegate sslmode=disable password=parkhom4ik")
  	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}


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

	fmt.Println(c.ClientIP())

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
