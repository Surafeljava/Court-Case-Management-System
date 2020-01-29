package entity

import (
	"time"

	"github.com/Surafeljava/gorm"
)

type Admin struct {
	ID       uint
	AdminId  string `gorm:"type:varchar(50);not null"`
	AdminPwd string `gorm:"type:varchar(50);not null"`
}

type UserType struct {
	gorm.Model
	UsrId  string
	UsrPwd string
}

type Case struct {
	ID            uint
	CaseNum       string `gorm:"type:varchar(50);not null"`
	CaseTitle     string `gorm:"type:varchar(50);not null"`
	CaseDesc      string `gorm:"type:varchar(50);not null"`
	CaseStatus    string `gorm:"type:varchar(50);not null"`
	CaseType      string `gorm:"type:varchar(50);not null"`
	CaseCreation  time.Time
	CaseCourtDate time.Time
	CaseJudge     string `gorm:"type:varchar(50);not null"`
}

type CaseInfo struct {
	CaseTitle  string    `json:"case_title" gorm:"type:varchar(255)"`
	CaseStatus string    `json:"case_status" gorm:"type:varchar(255)"`
	CourtDate  time.Time `json:"court_date" gorm:"type:varchar(255)"`
}

type Relation struct {
	ID      uint
	CaseNum string `gorm:"type:varchar(255);not null"`
	PlId    string `gorm:"type:varchar(255);not null"`
	AcId    string `gorm:"type:varchar(255);not null"`
	//JuId    string `gorm:"type:varchar(255);not null"`
}

type Decision struct {
	ID           uint
	CaseNum      string `gorm:"type:varchar(255);not null"`
	DecisionDate time.Time
	Decision     string `gorm:"type:varchar(255);not null"`
	DecisionDesc string `gorm:"type:varchar(255);not null"`
}

type Witness struct {
	ID          uint
	CaseNum     string `gorm:"type:varchar(255);not null"`
	WitnessDoc  string `gorm:"type:varchar(255);not null"`
	WitnessType string `gorm:"type:varchar(255);not null"`
}

type Judge struct {
	ID           uint
	JudgeId      string `gorm:"type:varchar(50);not null"`
	JudgePwd     string `gorm:"type:varchar(50);not null"`
	JudgeName    string `gorm:"type:varchar(50);not null"`
	JudgeGender  string `gorm:"type:varchar(50);not null"`
	JudgeAddress string `gorm:"type:varchar(50);not null"`
	JudgePhone   string `gorm:"type:varchar(50);not null"`
	JudgeType    string `gorm:"type:varchar(50);not null"`
	JudgePhoto   string `gorm:"type:varchar(50);not null"`
}

type Notification struct {
	ID             uint
	NotDescription string `gorm:"type:varchar(255);not null"`
	NotTitle       string `gorm:"type:varchar(255);not null"`
	NotLevel       string `gorm:"type:varchar(50);not null"`
	NotDate        time.Time
}

type Opponent struct {
	ID         uint
	OppId      string `gorm:"type:varchar(50);not null"`
	OppPwd     string `gorm:"type:varchar(50);not null"`
	OppType    string `gorm:"type:varchar(50);not null"`
	OppName    string `gorm:"type:varchar(50);not null"`
	OppGender  string `gorm:"type:varchar(50);not null"`
	OppBD      time.Time
	OppAddress string `gorm:"type:varchar(50);not null"`
	OppPhone   string `gorm:"type:varchar(50);not null"`
	OppPhoto   string `gorm:"type:varchar(50);not null"`
}

type SuccessMessage struct {
	Status  string
	Message string
}

//TODO: unfinished session work...
type Session struct {
	ID         uint
	UUID       string `gorm:"type:varchar(255);not null"`
	Expires    int64  `gorm:"type:varchar(255);not null"`
	SigningKey []byte `gorm:"type:varchar(255);not null"`
}

type Messg struct {
	UserID       string
	UserPwd      string
	UserName     string
	AddtionalMsg string
}

type Court struct {
	ID           uint
	CourtName    string `gorm:"type:varchar(255);not null"`
	CourtLevel   string `gorm:"type:varchar(255);not null"`
	CourtAddress string `gorm:"type:varchar(255);not null"`
	CourtPhone   string `gorm:"type:varchar(255);not null"`
}

type AppealForm struct {
	CaseNum          string
	CaseCreationDate time.Time
	CaseTitle        string
	CaseDesc         string
	OppName          string
	OppGender        string
	OppAddress       string
	OppPhone         string
	WitDocm          string
	WitTy            string
	Decision         string
	DecDate          time.Time
	DacDesc          string
}
