package searchUse

import entity "github.com/Surafeljava/Court-Case-Management-System/Entity"

//CaseSearchService ...
type CaseSearchService interface {
	Cases() ([]entity.Case, []error)
	Case(caseNum string) (*entity.Case, []error)
}

//JudgeSearchService ...
type JudgeSearchService interface {
	Judges() ([]entity.Judge, []error)
	Judge(judgeID string) (*entity.Judge, []error)
}
