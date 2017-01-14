package database

import "github.com/jinzhu/gorm"

//Client DataBase Table Client
type Client struct {
	gorm.Model
	Name string `gorm:"size:255"` // Default size for string is 255, reset it with this tag
}
