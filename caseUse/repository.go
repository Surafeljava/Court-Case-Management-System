package caseUse

import "github.com/Surafeljava/Court-Case-Management-System/entity"

type CaseRepository interface {
}

type LoginRepository interface {
	CheckLogin(user entity.User) error
}
