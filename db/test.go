package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/gorm"	
)

func main() {
	_, err := gorm.Open("postgres", "host=localhost port=5433 user=postgres dbname=services sslmode=disable password=password")
 	if err != nil {
		 fmt.Println(err)
	 }
}