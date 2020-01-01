package caseUse

import "github.com/Surafeljava/Court-Case-Management-System/entity"

type CaseService interface {
}

type LoginService interface {
	CheckLogin(user entity.User) error
}
