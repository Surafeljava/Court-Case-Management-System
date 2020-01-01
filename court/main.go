package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"github.com/Surafeljava/Court-Case-Management-System/caseUse/repository"
	"github.com/Surafeljava/Court-Case-Management-System/caseUse/service"
	"github.com/Surafeljava/Court-Case-Management-System/court/handler"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	fmt.Println("Welcome To Court Case Management System")
	//db, err := gorm.Open("postgres", "host=localhost port=5433 user=postgres dbname=gormdb password=123456")
	// defer db.Close()
	dbconn, err := sql.Open("postgres", "host=localhost port=5433 user=postgres dbname=gormdb password=123456")
	defer dbconn.Close()

	if err != nil {
		panic(err)
	}

	tmpl := template.Must(template.ParseGlob("../UI/templates/*"))

	loginRepo := repository.NewLoginRepositoryImpl(dbconn)
	loginServ := service.NewLoginServiceImpl(loginRepo)

	loginHandler := handler.NewLoginHandler(tmpl, loginServ)

	newcaseHandler := handler.NewCaseHandler(tmpl, loginServ)

	fs := http.FileServer(http.Dir("../UI/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/login", loginHandler.UserLoginCheck)
	http.HandleFunc("/admin/newcase", newcaseHandler.NewCase)

	http.ListenAndServe(":8181", nil)

}
