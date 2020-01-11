package entity

// Case ...
type Case struct {
	ID            uint
	CaseTitle     string `gorm:"type:varchar(255);not null"`
	CaseDesc      string `gorm:"type:varchar(255);not null"`
	CaseStatus    string `gorm:"type:varchar(255);not null"`
	CaseType      string `gorm:"type:varchar(255);not null"`
	CaseCreation  string `gorm:"type:varchar(255);not null"`
	CaseCourtDate string `gorm:"type:varchar(255);not null"`
	CaseJudge     string `gorm:"type:varchar(255);not null"`
}

//Judge ...
type Judge struct {
	ID           uint
	JudgeName    string `gorm:"type:varchar(255);not null"`
	JudgePwd     string `gorm:"type:varchar(255);not null"`
	JudgeGender  string `gorm:"type:varchar(255);not null"`
	JudgeAddress string `gorm:"type:varchar(255);not null"`
	JudgePhone   string `gorm:"type:varchar(255);not null"`
	JudgeType    string `gorm:"type:varchar(255);not null"`
	JudgePhoto   string `gorm:"type:varchar(255);not null"`
}

//Opponent ...
type Opponent struct {
	opponentID   uint
	opponentName string `gorm:"type:varchar(255);not null"`
}
