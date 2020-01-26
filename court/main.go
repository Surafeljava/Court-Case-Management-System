package main

import (
	"fmt"
	"html/template"
	"net/http"

	usrRepo "github.com/Surafeljava/Court-Case-Management-System/SearchUse/repository"
	usrService "github.com/Surafeljava/Court-Case-Management-System/SearchUse/service"
	aplRepo "github.com/Surafeljava/Court-Case-Management-System/appealUse/repository"
	aplService "github.com/Surafeljava/Court-Case-Management-System/appealUse/service"
	"github.com/Surafeljava/Court-Case-Management-System/caseUse/repository"
	"github.com/Surafeljava/Court-Case-Management-System/caseUse/service"

	"github.com/Surafeljava/Court-Case-Management-System/court/handler"
	notificationRepo "github.com/Surafeljava/Court-Case-Management-System/notificationUse/repository"
	notificationServ "github.com/Surafeljava/Court-Case-Management-System/notificationUse/service"
	"github.com/Surafeljava/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	fmt.Println("Welcome To Court Case Management System")

	dbc, err := gorm.Open("postgres", "host=localhost port=5433 user=postgres dbname=courttest2 password=123456")
	//dbc, err := gorm.Open("postgres", "postgres://postgres:1234@localhost/courttest2?sslmode=disable")
	defer dbc.Close()

	//TODO: Creating tables on the database
	// dbc.AutoMigrate(&entity.Opponent{})
	// dbc.AutoMigrate(&entity.Case{})
	// dbc.AutoMigrate(&entity.Judge{})
	// dbc.AutoMigrate(&entity.Admin{})
	// dbc.AutoMigrate(&entity.Notification{})
	// dbc.AutoMigrate(&entity.Relation{})
	// dbc.AutoMigrate(&entity.Decision{})
	// dbc.AutoMigrate(&entity.Witness{})

	// hasher := md5.New()
	// hasher.Write([]byte("1234"))
	// pwdnew := hex.EncodeToString(hasher.Sum(nil))

	// ad := entity.Admin{AdminId: "AD1", AdminPwd: pwdnew}
	// dbc.Create(&ad)

	if err != nil {
		panic(err)
	}

	// wit := entity.Witness{CaseNum: "CS1", WitnessDoc: "Witness sample doc", WitnessType: "type"}
	// dbc.Create(&wit)

	// des := entity.Decision{CaseNum: "CS1", DecisionDate: time.Now(), Decision: "The final decistion", DecisionDesc: "Decision description kjksahdjk"}
	// dbc.Create(&des)

	// rel := entity.Relation{CaseNum: "CS1", PlId: "OP1", AcId: "OP2"}
	// dbc.Create(&rel)

	// tmpl := template.Must(template.ParseGlob("../UI/templates/*"))

	//Login repository and Service Creation
	loginRepo := repository.NewLoginRepositoryImpl(dbc)
	loginServ := service.NewLoginServiceImpl(loginRepo)

	//Case repository and Service Creation
	caseRepo := repository.NewCaseRepositoryImpl(dbc)
	caseServ := service.NewCaseServiceImpl(caseRepo)

	//Opponent repository and Service Creation
	oppRepo := repository.NewOpponentRepositoryImpl(dbc)
	oppServ := service.NewOpponentServiceImpl(oppRepo)

	//Judge repository and Service Creation
	adminJudgeRepo := repository.NewJudgeRepositoryImpl(dbc)
	adminJudgeServ := service.NewJudgeServiceImpl(adminJudgeRepo)

	loginHandler := handler.NewLoginHandler(tmpl, loginServ)
	newcaseHandler := handler.NewCaseHandler(tmpl, caseServ)
	opponentHandler := handler.NewOpponentHandler(tmpl, oppServ)
	adminJudgeHandler := handler.NewAdminJudgeHandler(tmpl, adminJudgeServ)

	//Searching
	//Case_Search
	caseSearchRepo := usrRepo.NewCaseSearchGormRepo(dbc)
	caseSearchService := usrService.NewCaseSearchService(caseSearchRepo)
	caseSearchHandler := handler.NewCaseSearchHandler(caseSearchService)

	//Judge_Search
	judgeSearchRepo := usrRepo.NewJudgeSearchGormRepo(dbc)
	judgeSearchService := usrService.NewJudgeSearchService(judgeSearchRepo)
	judgeSearchHandler := handler.NewJudgeSearchHandler(judgeSearchService)

	//notification service and Repo
	notificatioRepos := notificationRepo.NewNotificationRepositoryImpl(dbc)
	notificationService := notificationServ.NewNotificationServiceImpl(notificatioRepos)

	//Notification
	adminNotificatHandler := handler.NewNotificationHandler(tmpl, notificationService)
	OppJudgNotificatHandler := handler.NewOppJNotificationHandler(tmpl, notificationService)

	//Appeal
	appealRepo := aplRepo.NewAppealGormRepo(dbc)
	appealService := aplService.NewAppealService(appealRepo)
	oppAppealHandler := handler.NewAppealHandler(appealService)

	fs := http.FileServer(http.Dir("../UI/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/login", loginHandler.AuthenticateUser)
	http.HandleFunc("/admin/cases/new", LoginRequired(UserAuthorized(newcaseHandler.NewCase)))
	http.HandleFunc("/admin/cases/update", newcaseHandler.UpdateCase)
	http.HandleFunc("/admin/cases/delete", newcaseHandler.DeleteCase)
	http.HandleFunc("/admin/cases", newcaseHandler.Cases)
	http.HandleFunc("/admin/opponent/new", opponentHandler.NewOpponent)
	http.HandleFunc("/admin/judge/new", adminJudgeHandler.NewJudge)

	http.HandleFunc("/judge/cases/close", LoginRequired(UserAuthorized(newcaseHandler.CloseCase)))

	//TODO: notification handlers
	// http.HandleFunc("/admin/notification/new", )
	// http.HandleFunc("/notification", )

	//Admin_search
	http.HandleFunc("/v1/adminSearch", adminSearch)

	//Case Search
	http.HandleFunc("/v1/admin/cases", caseSearchHandler.Cases)
	http.HandleFunc("/v1/admin/cases/singlecase", caseSearchHandler.GetSingleCase)

	//Judge Search
	http.HandleFunc("/v1/admin/judges", judgeSearchHandler.Judges)
	http.HandleFunc("/v1/admin/judges/singlejudge", judgeSearchHandler.GetSingleJudge)

	//notification
	http.HandleFunc("/admin/postNotifications", adminNotificatHandler.AdminPostNotification)
	http.HandleFunc("/judge/notifications", OppJudgNotificatHandler.NotificationsJudge)
	http.HandleFunc("/opponent/notifications", OppJudgNotificatHandler.NotificationsOpponent)

	//Appeal
	http.HandleFunc("/admin/oppForAppealTrial", oppAppealHandler.OppTrial)
	http.HandleFunc("/admin/oppAppeal", oppAppealHandler.OppAppeal)

	http.ListenAndServe(":8181", nil)

}

var tmpl = template.Must(template.ParseGlob("../UI/templates/*.html"))

func adminSearch(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "adminSearch.layout", nil)
}

func LoginRequired(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Check if there is a session

		sess, err := r.Cookie("signed_user")

		if err != nil {
			http.Redirect(w, r, "/login", 302)
			return
		}

		//fmt.Printf(sess.Value)
		fmt.Println(sess.Value)
		handler.ServeHTTP(w, r)
	}
}

func UserAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Check the authorization of the user accessing the link
		handler.ServeHTTP(w, r)
	}
}
