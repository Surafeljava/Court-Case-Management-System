package service

import (
	entity "github.com/Surafeljava/Court-Case-Management-System/Entity"
	appealUse "github.com/Surafeljava/Court-Case-Management-System/appealUse"
)

//AppealService ...
type AppealService struct {
	appealRepo appealUse.AppealRepositroy
}

//NewAppealService ...
func NewAppealService(appealRepository appealUse.AppealRepositroy) appealUse.AppealService {
	return &AppealService{appealRepo: appealRepository}

}
