package entity

import (
	"time"
)

// CaseMock  Food Case
var CaseMock = Case{
	ID:            1,
	CaseNum:       "CS1",
	CaseTitle:     "Murder",
	CaseDesc:      "Killing two people",
	CaseStatus:    "Open",
	CaseType:      "criminal",
	CaseCreation:  time.Time{},
	CaseCourtDate: time.Time{},
	CaseJudge:     "JD1",
}

// AdminMock mocks  Admin
var AdminMock = Admin{
	ID:       1,
	AdminId:  "AD1",
	AdminPwd: "1234",
}

// UserTypeMock mocks  Decision
var UserTypeMock = UserType{
	UsrId:  "US1",
	UsrPwd: "1234",
}

// RelationMock mocks  Decision
var RelationMock = Relation{
	ID:      1,
	CaseNum: "CS1",
	PlId:    "PL1",
	AcId:    "AC1",
}

// DecisionMock mocks  Decision
var DecisionMock = Decision{
	ID:           1,
	CaseNum:      "CS1",
	DecisionDate: time.Time{},
	Decision:     "Must be in jail",
	DecisionDesc: "document",
}

// WitnessMock mocks  Witness
var WitnessMock = Witness{
	ID:          1,
	CaseNum:     "CS1",
	WitnessDoc:  "Duresa",
	WitnessType: "person",
}

// JudgeMock mocks  Judge
var JudgeMock = Judge{
	ID:           1,
	JudgeId:      "JU1",
	JudgePwd:     "1234",
	JudgeName:    "Tesfaye",
	JudgeGender:  "male",
	JudgeAddress: "Addis Ababa",
	JudgePhone:   "0903054480",
	JudgeType:    "criminal",
	JudgePhoto:   "photo.png",
}

// NotificationMock mocks  Notification
var NotificationMock = Notification{
	ID:             1,
	NotDescription: "Team meeting",
	NotTitle:       "Notice",
	NotLevel:       "all",
	NotDate:        time.Time{},
}

// OpponentMock mocks  Opponent
var OpponentMock = Opponent{
	ID:         1,
	OppId:      "OP1",
	OppPwd:     "1234",
	OppType:    "criminal",
	OppName:    "Haylu",
	OppGender:  "Male",
	OppBD:      time.Time{},
	OppAddress: "Addis Ababa",
	OppPhone:   "0909090909",
	OppPhoto:   "photoOpponent",
}

// SuccessMessageMock mocks  SuccessMessage
var SuccessMessageMock = SuccessMessage{
	Status:  "OK",
	Message: "messaage",
}

// SessionMock mocks  Session
var SessionMock = Session{
	ID:         1,
	UUID:       "AD1",
	Expires:    30,
	SigningKey: []byte{},
}
