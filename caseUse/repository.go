package caseUse

import "github.com/Surafeljava/Court-Case-Management-System/entity"

type CaseRepository interface {
	CreateCase(cased entity.Case) error
}

type LoginRepository interface {
	CheckLogin(user entity.User) error
}

type CaseSearchRepository interface {
	Cases() ([]entity.Case, []error)
	Case(id uint) (*entity.Case, []error)
}
