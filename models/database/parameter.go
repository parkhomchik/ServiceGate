package database

import "github.com/jinzhu/gorm"

type Parameter struct {
	gorm.Model
	Name  string
	Value string
}
