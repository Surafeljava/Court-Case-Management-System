package service

import (
	entity "github.com/Surafeljava/Court-Case-Management-System/Entity"
	user "github.com/Surafeljava/Court-Case-Management-System/SearchUse"
)

// JudgeSearchService implements menu.UserService interface
type JudgeSearchService struct {
	judgeSearchRepo user.JudgeSearchRepository
}

// NewJudgeSearchService  returns a new UserService object
func NewJudgeSearchService(judgeSearchRepository user.JudgeSearchRepository) user.JudgeSearchRepository {
	return &JudgeSearchService{judgeSearchRepo: judgeSearchRepository}
}

// Judges returns all stored application users
func (us *JudgeSearchService) Judges() ([]entity.Judge, []error) {
	judges, errs := us.judgeSearchRepo.Judges()
	if len(errs) > 0 {
		return nil, errs
	}
	return judges, errs
}

// Judge retrieves an application user by its id
func (us *JudgeSearchService) Judge(judgeID string) (*entity.Judge, []error) {
	usr, errs := us.judgeSearchRepo.Judge(judgeID)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}
