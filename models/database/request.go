package database

import "github.com/jinzhu/gorm"

//Request table
type Request struct {
	gorm.Model
	//ID int64
	ExternalID  int64
	ServiceName string
	CommandName string
	Parameters  []Parameter
}

func (r *Request) Create(db *gorm.DB) {

}

func (r *Request) Update(db *gorm.DB, newRequest Request) {

}

func (r *Request) Delete(db *gorm.DB) {

}
