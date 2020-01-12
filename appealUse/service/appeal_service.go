package service

import "github.com/Surafeljava/Court-Case-Management-System/appealUse"

type AppealServiceImpl struct {
	appealRepo appealUse.AppealRepositroy
}

func NewAppealServiceImpl(notf appealUse.AppealRepositroy) appealUse.AppealRepositroy {
	return &AppealServiceImpl{notfRepo: notf}

}

func (ar *AppealServiceImpl) AppealGet(caseNum string) (*entity.Case, []error){
	return nil, nil
}
