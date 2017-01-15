package database

import "github.com/jinzhu/gorm"

type Service struct {
	gorm.Model
	ServiceName       string
	Active            bool
	ServiceParameters []ServiceParameter
}
