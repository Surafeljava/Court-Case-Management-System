package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	fmt.Println("Welcome To Court Case Management System")
	db, err := gorm.Open("postgres", "host=localhost port=5433 user=postgres dbname=gormdb password=123456")
	defer db.Close()

	if err != nil {
		panic(err)
	}
}
