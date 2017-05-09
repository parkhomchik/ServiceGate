package main

import (
	"github.com/parkhomchik/ServiceGate/models"
	"github.com/parkhomchik/ServiceGate/db"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/gorm"
)

type TodoService struct {
}

func (s *TodoService) getDb() (gorm.DB, error) {
	db, err := gorm.Open("postgres", "host=localhost port=5433 user=postgres dbname=services sslmode=disable password=parkhom4ik")
 	return *db, err
}

func (s *TodoService) Migrate() error {
	db, err := s.getDb()
	if err != nil {
		return err
	}
	db.SingularTable(true)

	db.AutoMigrate(&models.Service{})
	return nil
}

func main() {
	s := TodoService{}
	err := s.Migrate()
	db, err := s.getDb()
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)

	serviceResource := DataBaseManager.ServiceResource{Db: db}
	//serviceResource := ServiceResource{db : db}

	r := gin.Default()

	r.GET("/service", serviceResource.GetAllServices)
	r.GET("/service/:id", serviceResource.GetService)
	r.POST("/service", serviceResource.CreateService)
	r.PUT("/service/:id", serviceResource.UpdateService)
	//r.PATCH("/service/:id", serviceResource.PatchService)
	r.DELETE("/service/:id", serviceResource.DeleteService)

	r.Run(":8080")	
}