package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/Surafeljava/Court-Case-Management-System/caseUse/repository"
	"github.com/Surafeljava/Court-Case-Management-System/caseUse/service"
	"github.com/Surafeljava/Court-Case-Management-System/court/handler"
	"github.com/Surafeljava/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	fmt.Println("Welcome To Court Case Management System")

	dbc, err := gorm.Open("postgres", "host=localhost port=5433 user=postgres dbname=courttest2 password=123456")
	defer dbc.Close()

	//TODO: Creating tables on the database
	// dbc.AutoMigrate(&entity.Opponent{})
	// dbc.AutoMigrate(&entity.Case{})
	// dbc.AutoMigrate(&entity.Judge{})
	//dbc.AutoMigrate(&entity.Admin{})
	//dbc.AutoMigrate(&entity.Notification{})

	// ad := entity.Admin{AdminId: "AD1", AdminPwd: "1234"}
	// dbc.Create(&ad)

	if err != nil {
		panic(err)
	}

	tmpl := template.Must(template.ParseGlob("../UI/templates/*"))

	loginRepo := repository.NewLoginRepositoryImpl(dbc)
	loginServ := service.NewLoginServiceImpl(loginRepo)

	caseRepo := repository.NewCaseRepositoryImpl(dbc)
	caseServ := service.NewCaseServiceImpl(caseRepo)

	oppRepo := repository.NewOpponentRepositoryImpl(dbc)
	oppServ := service.NewOpponentServiceImpl(oppRepo)

	adminJudgeRepo := repository.NewJudgeRepositoryImpl(dbc)
	adminJudgeServ := service.NewJudgeServiceImpl(adminJudgeRepo)

	loginHandler := handler.NewLoginHandler(tmpl, loginServ)
	newcaseHandler := handler.NewCaseHandler(tmpl, caseServ)
	opponentHandler := handler.NewOpponentHandler(tmpl, oppServ)
	adminJudgeHandler := handler.NewAdminJudgeHandler(tmpl, adminJudgeServ)

	fs := http.FileServer(http.Dir("../UI/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/login", loginHandler.AuthenticateUser)
	//http.HandleFunc("/login/check", loginHandler.AuthenticateUser)
	http.HandleFunc("/admin/cases/new", newcaseHandler.NewCase)
	http.HandleFunc("/admin/cases/update", newcaseHandler.UpdateCase)
	http.HandleFunc("/admin/cases", newcaseHandler.Cases)
	http.HandleFunc("/admin/opponent/new", opponentHandler.NewOpponent)
	http.HandleFunc("/admin/judge/new", adminJudgeHandler.NewJudge)

	//TODO: notification handlers
	// http.HandleFunc("/admin/notification/new", )
	// http.HandleFunc("/notification", )

	http.ListenAndServe(":8181", nil)

}
