package main

import (
	"github.com/parkhomchik/ServiceGate/models"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/gorm"
)

type Config struct {
	SvcHost    string
	DbUser     string
	DbPassword string
	DbHost     string
	DbName     string
}

type TodoService struct {
}

func (s *TodoService) getDb() (gorm.DB, error) {
	//connectionString := cfg.DbUser + ":" + cfg.DbPassword + "@tcp(" + cfg.DbHost + ":3306)/" + cfg.DbName + "?charset=utf8&parseTime=True"
	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=services sslmode=disable password=parkhom4ik")
 	return *db, nil
}

func (s *TodoService) Migrate(cfg Config) error {
	db, err := s.getDb()
	if err != nil {
		return err
	}
	db.SingularTable(true)

	db.AutoMigrate(&api.Service{})
	return nil
}

func main() {
	s := TodoService{}
	db, err := s.getDb()
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)

	todoResource := &TodoResource{db: db}

	r := gin.Default()

	r.GET("/todo", todoResource.GetAllTodos)
	r.GET("/todo/:id", todoResource.GetTodo)
	r.POST("/todo", todoResource.CreateTodo)
	r.PUT("/todo/:id", todoResource.UpdateTodo)
	r.PATCH("/todo/:id", todoResource.PatchTodo)
	r.DELETE("/todo/:id", todoResource.DeleteTodo)

	r.Run(cfg.SvcHost)

	return nil
}