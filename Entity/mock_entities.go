package entity

import "time"

// CaseryMock mocks Food Menu Category
var CaseryMock = Case{
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
