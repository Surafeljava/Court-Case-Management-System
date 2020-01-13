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

//Appeal ...
func (appealService *AppealService) Appeal(oppNum string) (*entity.Case, *entity.Opponent, *entity.Witness, *entity.Decision, []error) {
	cases, opp, wit, dec, errs := appealService.appealRepo.Appeal(oppNum)

	if len(errs) > 0 {
		return nil, nil, nil, nil, errs
	}

	return cases, opp, wit, dec, errs
}
