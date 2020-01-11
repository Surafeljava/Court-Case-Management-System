package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	fmt.Println("Welcome To Court Case Management System")
	db, err := gorm.Open("postgres", "host=localhost  user=postgres dbname=gormdb password=E456yob123 sslmode=disable")
	defer db.Close()

	if err != nil {
		panic(err)
	}
}
