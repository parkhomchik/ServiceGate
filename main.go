package main

import (
	"net/http"

	

	_ "github.com/jinzhu/gorm/dialects/postgres"
	//"github.com/parkhomchik/ServiceGate/models"
	"github.com/parkhomchik/ServiceGate/models/database"
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	//DataBase
	/*db, err := gorm.Open("postgres", "host=localhost port=5433 user=postgres dbname=servicegate sslmode=disable password=parkhom4ik")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	database.Migrate(db)
*/
	db := database.InitDb()
	database.Migrate(db)
	
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
	db := database.InitDb()
	 

	var json database.Request
	if c.BindJSON(&json) == nil {
		
		//db.NewRecord(&json) // => returns `true` as primary key is blank
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
		c.JSON(http.StatusBadRequest, gin.H{"status": "StatusBadRequest"})
	}

}

func requestv2(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"status": "not implemented"})
}
