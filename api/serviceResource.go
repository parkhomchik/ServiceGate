package main

import (
	"github.com/parkhomchik/ServiceGate/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"strconv"
	//"time"
)

type ServiceResource struct {
	db gorm.DB
}

func (tr *ServiceResource) CreateService(c *gin.Context) {
	var service models.Service

	if c.Bind(&service) != nil {
		c.JSON(400, models.NewError("problem decoding body"))
		return
	}

	//service.Created = int32(time.Now().Unix())

	tr.db.Save(&service)

	c.JSON(201, service)
}

func (tr *ServiceResource) GetAllServices(c *gin.Context) {
	var services []models.Service

	tr.db.Find(&services)

	c.JSON(200, services)
}

func (tr *ServiceResource) GetService(c *gin.Context) {
	id, err := tr.getId(c)
	if err != nil {
		c.JSON(400, models.NewError("problem decoding id sent"))
		return
	}

	var service models.Service

	if tr.db.First(&service, id).RecordNotFound() {
		c.JSON(404, gin.H{"error": "not found"})
	} else {
		c.JSON(200, service)
	}
}

func (tr *ServiceResource) UpdateService(c *gin.Context) {
	id, err := tr.getId(c)
	if err != nil {
		c.JSON(400, models.NewError("problem decoding id sent"))
		return
	}

	var service models.Service

	if c.Bind(&service) != nil {
		c.JSON(400, models.NewError("problem decoding body"))
		return
	}
	service.ID = int32(id)

	var existing models.Service

	if tr.db.First(&existing, id).RecordNotFound() {
		c.JSON(404, models.NewError("not found"))
	} else {
		tr.db.Save(&service)
		c.JSON(200, service)
	}

}
/*
func (tr *ServiceResource) PatchTodo(c *gin.Context) {
	id, err := tr.getId(c)
	if err != nil {
		c.JSON(400, models.NewError("problem decoding id sent"))
		return
	}

	// this is a hack because Gin falsely claims my unmarshalled obj is invalid.
	// recovering from the panic and using my object that already has the json body bound to it.
	var json []api.Patch
	defer func() {
		if r := recover(); r != nil {
			if json[0].Op != "replace" && json[0].Path != "/status" {
				c.JSON(400, api.NewError("PATCH support is limited and can only replace the /status path"))
				return
			}

			var todo api.Todo

			if tr.db.First(&todo, id).RecordNotFound() {
				c.JSON(404, api.NewError("not found"))
			} else {
				todo.Status = json[0].Value

				tr.db.Save(&todo)
				c.JSON(200, todo)
			}
		}
	}()
	c.Bind(&json)
}
*/
func (tr *ServiceResource) DeleteService(c *gin.Context) {
	id, err := tr.getId(c)
	if err != nil {
		c.JSON(400, models.NewError("problem decoding id sent"))
		return
	}

	var service models.Service

	if tr.db.First(&service, id).RecordNotFound() {
		c.JSON(404, models.NewError("not found"))
	} else {
		tr.db.Delete(&service)
		c.Data(204, "application/json", make([]byte, 0))
	}
}

func (tr *ServiceResource) getId(c *gin.Context) (int32, error) {
	idStr := c.Params.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Print(err)
		return 0, err
	}
	return int32(id), nil
}