package caseUse

import "github.com/Surafeljava/Court-Case-Management-System/entity"

//CaseService ...
type CaseService interface {
}

//LoginService ...
type LoginService interface {
	CheckLogin(user entity.User) error
}

//CaseSearchService ...
type CaseSearchService interface {
	Cases() ([]entity.Case, []error)
	Case(id uint) (*entity.Case, []error)
}

//JudgeSearchService ...
type JudgeSearchService interface {
	Judges() ([]entity.Judge, []error)
	Judge(id uint) (*entity.Judge, []error)
}
