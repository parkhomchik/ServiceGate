package database

import "github.com/jinzhu/gorm"
import "fmt"

//Migrate создаем и изменяем таблицы
func Migrate(db *gorm.DB) {
	fmt.Println("Migrate Client")
	db.AutoMigrate(&Client{}, &Request{}, &Parameter{})
}

func InitDb() *gorm.DB {
    // Openning file
    db, err := gorm.Open("postgres", "host=localhost port=5433 user=postgres dbname=servicegate sslmode=disable password=parkhom4ik")
    db.LogMode(true)
    // Error
    if err != nil {
        panic(err)
    }
    // Creating the table
    /*if !db.HasTable(&Client{}) {
        db.CreateTable(&Client{})
    }

	if !db.HasTable(&Request{}) {
		db.CreateTable(&Request{})
	}

	if !db.HasTable(&Parameter{}) {
		db.CreateTable(&Parameter{})
	}*/

    return db
}