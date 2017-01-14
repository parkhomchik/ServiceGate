package database

import "github.com/jinzhu/gorm"
import "fmt"

//Migrate создаем и изменяем таблицы
func Migrate(db *gorm.DB) {
	fmt.Println("Migrate Client")
	db.AutoMigrate(&Client{}, &Request{}, &Parameter{})
}
