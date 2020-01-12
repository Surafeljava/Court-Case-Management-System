package caseUse

import entity "github.com/Surafeljava/Court-Case-Management-System/Entity"

//CaseRepository ...
type CaseRepository interface {
	Cases() ([]entity.Case, error)
	JudgeCases(juid string) ([]entity.Case, error)
	CaseJudges(case_type string) ([]entity.Judge, error)
	Case(id int) (*entity.Case, []error)
	CreateCase(casedoc *entity.Case) []error
	UpdateCase(casedoc *entity.Case) (*entity.Case, []error)
	CloseCase(casedoc string, decision *entity.Decision) []error
	ExtendCase(casedoc *entity.Case) []error
	DeleteCase(id int) []error
}

//OpponentRepository ...
type OpponentRepository interface {
	Opponents() ([]entity.Opponent, error)
	Opponent(id int) (*entity.Opponent, []error)
	CreateOpponent(opp *entity.Opponent) (*entity.Opponent, []error)
}

//JudgeRepository ...
type JudgeRepository interface {
	Judges() ([]entity.Judge, error)
	Judge(id int) (*entity.Judge, []error)
	CreateJudge(judge *entity.Judge) (*entity.Judge, []error)
	UpdateCase(judge *entity.Judge) (*entity.Judge, []error)
	DeleteCase(id int) error
}

//LoginRepository ...
type LoginRepository interface {
	CheckLogin(user *entity.UserType) (*entity.UserType, []error)
	CheckAdmin(id string, pwd string) (*entity.Admin, []error)
	CheckJudge(id string, pwd string) (*entity.Judge, []error)
	CheckOpponent(id string, pwd string) (*entity.Opponent, []error)
}

//CaseSearchRepository ...
type CaseSearchRepository interface {
	Cases() ([]entity.Case, []error)
	Case(id uint) (*entity.Case, []error)
}
