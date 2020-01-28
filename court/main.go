package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	entity "github.com/Surafeljava/Court-Case-Management-System/Entity"
	usrRepo "github.com/Surafeljava/Court-Case-Management-System/SearchUse/repository"
	usrService "github.com/Surafeljava/Court-Case-Management-System/SearchUse/service"
	aplRepo "github.com/Surafeljava/Court-Case-Management-System/appealUse/repository"
	aplService "github.com/Surafeljava/Court-Case-Management-System/appealUse/service"
	"github.com/Surafeljava/Court-Case-Management-System/caseUse/repository"
	"github.com/Surafeljava/Court-Case-Management-System/caseUse/service"
	repRepo "github.com/Surafeljava/Court-Case-Management-System/reportUse/repository"
	repService "github.com/Surafeljava/Court-Case-Management-System/reportUse/service"
	"github.com/Surafeljava/Court-Case-Management-System/rtoken"

	"github.com/Surafeljava/Court-Case-Management-System/court/handler"
	notificationRepo "github.com/Surafeljava/Court-Case-Management-System/notificationUse/repository"
	notificationServ "github.com/Surafeljava/Court-Case-Management-System/notificationUse/service"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//TODO: Creating database tables
func CreateDBTables(db *gorm.DB) {
	db.CreateTable(&entity.Court{}, &entity.Opponent{}, &entity.Case{}, &entity.Judge{}, &entity.Admin{}, &entity.Notification{}, &entity.Relation{}, &entity.Decision{}, &entity.Witness{}, &entity.Session{})
}

func main() {
	fmt.Println("Welcome To Court Case Management System")

	csrfSignKey := []byte(rtoken.GenerateRandomID(32))
	tmpl := template.Must(template.ParseGlob("../UI/templates/*"))

	dbc, err := gorm.Open("postgres", "host=localhost port=5433 user=postgres dbname=courttest2 password=123456")

	//Creating Database Tables
	//CreateDBTables(dbc)
	dbc.AutoMigrate(&entity.Court{})

	defer dbc.Close()

	if err != nil {
		panic(err)
	}

	//Login repository and Service Creation
	loginRepo := repository.NewLoginRepositoryImpl(dbc)
	loginServ := service.NewLoginServiceImpl(loginRepo)

	//Case repository and Service Creation
	caseRepo := repository.NewCaseRepositoryImpl(dbc)
	caseServ := service.NewCaseServiceImpl(caseRepo)

	//Opponent repository and Service Creation
	oppRepo := repository.NewOpponentRepositoryImpl(dbc)
	oppServ := service.NewOpponentServiceImpl(oppRepo)

	//Session Repository and Service Creation
	sessRepo := repository.NewSessionGormRepo(dbc)
	sessServ := service.NewSessionService(sessRepo)

	//Judge repository and Service Creation
	adminJudgeRepo := repository.NewJudgeRepositoryImpl(dbc)
	adminJudgeServ := service.NewJudgeServiceImpl(adminJudgeRepo)

	sess := configSess()
	loginHandler := handler.NewLoginHandler(tmpl, loginServ, sessServ, sess, csrfSignKey)

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

	newcaseHandler := handler.NewCaseHandler(tmpl, caseServ, notificationService)
	opponentHandler := handler.NewOpponentHandler(tmpl, oppServ)
	adminJudgeHandler := handler.NewAdminJudgeHandler(tmpl, adminJudgeServ)

	//Notification
	adminNotificatHandler := handler.NewNotificationHandler(tmpl, notificationService)
	OppJudgNotificatHandler := handler.NewOppJNotificationHandler(tmpl, notificationService)
	judgeNotificationHandler := handler.NewJudgeNotificationHandler(tmpl, notificationService)

	//Report
	reportRepo := repRepo.NewReportGormRepo(dbc)
	reportServ := repService.NewReportServiceImpl(reportRepo)
	reportHandle := handler.NewReportHandler(tmpl, reportServ)

	//Court Create
	courtRepo := repository.NewAdminCourtRepositoryImpl(dbc)
	courtService := service.NewAdminCourtServiceImpl(courtRepo)
	courtHandle := handler.NewAdminCourtHandler(tmpl, courtService)

	//Appeal
	appealRepo := aplRepo.NewAppealGormRepo(dbc)
	appealService := aplService.NewAppealService(appealRepo)
	oppAppealHandler := handler.NewAppealHandler(appealService)

	fs := http.FileServer(http.Dir("../UI/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/login", loginHandler.AuthenticateUser)
	http.HandleFunc("/logout", loginHandler.Logout)
	http.HandleFunc("/admin/cases/new", loginHandler.AuthenticatedUser(newcaseHandler.NewCase))
	http.HandleFunc("/admin/cases/update", loginHandler.AuthenticatedUser(newcaseHandler.UpdateCase))
	http.HandleFunc("/admin/cases/delete", loginHandler.AuthenticatedUser(newcaseHandler.DeleteCase))
	http.HandleFunc("/admin/cases", loginHandler.AuthenticatedUser(newcaseHandler.Cases))
	http.HandleFunc("/admin/opponent/new", loginHandler.AuthenticatedUser(opponentHandler.NewOpponent))
	http.HandleFunc("/admin/judge/new", loginHandler.AuthenticatedUser(adminJudgeHandler.NewJudge))
	http.HandleFunc("/search/case", newcaseHandler.SearchCaseInfo)

	http.HandleFunc("/judge/cases/close", loginHandler.AuthenticatedUser(newcaseHandler.CloseCase))
	http.HandleFunc("/user/changepwd", loginHandler.AuthenticatedUser(loginHandler.ChangePassword))

	//TODO: notification handlers
	// http.HandleFunc("/admin/notification/new", )
	// http.HandleFunc("/notification", )

	//Admin Report and Statistics
	http.HandleFunc("/admin/report", loginHandler.AuthenticatedUser(reportHandle.GetStatistics))

	//Court and Admin Create
	http.HandleFunc("/courtcreate", courtHandle.CreateCourt)
	http.HandleFunc("/admincreate", courtHandle.AdminCreate)

	//Admin_search
	http.HandleFunc("/v1/adminSearch", adminSearch)

	//Case Search
	http.HandleFunc("/v1/admin/cases", caseSearchHandler.Cases)
	http.HandleFunc("/v1/admin/cases/singlecase", caseSearchHandler.GetSingleCase)

	//Judge Search
	http.HandleFunc("/v1/admin/judges", judgeSearchHandler.Judges)
	http.HandleFunc("/v1/admin/judges/singlejudge", judgeSearchHandler.GetSingleJudge)

	//notification
	http.HandleFunc("/admin/notifications/delete", adminNotificatHandler.AdminDeleteNotification)
	http.HandleFunc("/admin/notifications", adminNotificatHandler.AdminNotifications)
	http.HandleFunc("/admin/notifications/update", adminNotificatHandler.AdminUpdateNotification)
	http.HandleFunc("/admin/notifications/postnotification", adminNotificatHandler.AdminPostNotification)

	http.HandleFunc("/judge/notifications", judgeNotificationHandler.NotificationsJudge)
	http.HandleFunc("/opponent/notifications", OppJudgNotificatHandler.NotificationsOpponent)

	http.HandleFunc("/judge/notifications/delete", judgeNotificationHandler.DeleteJudgeNotification)
	http.HandleFunc("/opponent/notifications/delete", OppJudgNotificatHandler.DeleteOpponentNotification)
	http.HandleFunc("/judge/notifications/update", judgeNotificationHandler.SingleNotificationJudge)
	http.HandleFunc("/opponent/notifications/update", OppJudgNotificatHandler.SingleNotificationOpponent)

	//Appeal
	http.HandleFunc("/admin/oppForAppealTrial", oppAppealHandler.OppTrial)
	http.HandleFunc("/admin/oppAppeal", oppAppealHandler.OppAppeal)

	http.ListenAndServe(":8181", nil)

}

var tmpl = template.Must(template.ParseGlob("../UI/templates/*.html"))

func adminSearch(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "adminSearch.layout", nil)
}

func configSess() *entity.Session {
	tokenExpires := time.Now().Add(time.Minute * 30).Unix()
	sessionID := rtoken.GenerateRandomID(32)
	signingString, err := rtoken.GenerateRandomString(32)
	if err != nil {
		panic(err)
	}
	signingKey := []byte(signingString)

	return &entity.Session{
		Expires:    tokenExpires,
		SigningKey: signingKey,
		UUID:       sessionID,
	}
}
