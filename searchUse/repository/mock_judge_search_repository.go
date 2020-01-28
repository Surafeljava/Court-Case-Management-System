package repository

import (
	"errors"

	entity "github.com/Surafeljava/Court-Case-Management-System/Entity"
	user "github.com/Surafeljava/Court-Case-Management-System/SearchUse"
	"github.com/Surafeljava/gorm"
)

// MockJudgeSearchRepo implements the menu.CategoryRepository interface
type MockJudgeSearchRepo struct {
	conn *gorm.DB
}

// NewMockJudgeSearchRepo will create a new object of MockCaseSearchRepo
func NewMockJudgeSearchRepo(db *gorm.DB) user.JudgeSearchRepository {
	return &MockJudgeSearchRepo{conn: db}
}

// Judges returns all fake Judges
func (mJudRepo *MockJudgeSearchRepo) Judges() ([]entity.Judge, []error) {
	judges := []entity.Judge{entity.JudgeMock}
	return judges, nil
}

// Judge retrieves a fake Judge with judgeID "JU1"
func (mJudRepo *MockJudgeSearchRepo) Judge(judgeID string) (*entity.Judge, []error) {
	judge := entity.JudgeMock
	if judgeID == "JU1" {
		return &judge, nil
	}
	return nil, []error{errors.New("Not found")}
}
