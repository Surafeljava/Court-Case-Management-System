package service

import (
	"github.com/danieltefera/Project_Search/entity"
	"github.com/danieltefera/Project_Search/user"
)

// CaseSearchService implements menu.UserService interface
type CaseSearchService struct {
	caseSearchRepo user.CaseSearchRepository
}

// NewCaseSearchService  returns a new UserService object
func NewCaseSearchService(caseSearchRepository user.CaseSearchRepository) user.CaseSearchService {
	return &CaseSearchService{caseSearchRepo: caseSearchRepository}
}

// Cases returns all stored application users
func (us *CaseSearchService) Cases() ([]entity.Case, []error) {
	usrs, errs := us.caseSearchRepo.Cases()
	if len(errs) > 0 {
		return nil, errs
	}
	return usrs, errs
}

// Case retrieves an application user by its id
func (us *CaseSearchService) Case(id uint) (*entity.Case, []error) {
	usr, errs := us.caseSearchRepo.Case(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}
