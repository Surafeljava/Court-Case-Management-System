package caseUse

import entity "github.com/Surafeljava/Court-Case-Management-System/Entity"

//CaseService ...
type CaseService interface {
	Cases() ([]entity.Case, error)
	Case(id int) (*entity.Case, []error)
	CreateCase(casedoc *entity.Case) []error
	UpdateCase(casedoc *entity.Case) (*entity.Case, []error)
	CloseCase(casedoc entity.Case) error
	ExtendCase(casedoc entity.Case) error
	DeleteCase(id int) error
}

//OpponentService ...
type OpponentService interface {
	CreateOpponent(opp *entity.Opponent) (*entity.Opponent, []error)
}

//JudgeService ..
type JudgeService interface {
	CreateJudge(judge *entity.Judge) (*entity.Judge, []error)
}

//LoginService ...
type LoginService interface {
	CheckLogin(user *entity.UserType) (*entity.UserType, []error)
	CheckAdmin(id string, pwd string) (*entity.Admin, []error)
	CheckJudge(id string, pwd string) (*entity.Judge, []error)
	CheckOpponent(id string, pwd string) (*entity.Opponent, []error)
}

//CaseSearchService ...
type CaseSearchService interface {
	Cases() ([]entity.Case, []error)
	Case(id uint) (*entity.Case, []error)
}
