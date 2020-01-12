package searchUse

import entity "github.com/Surafeljava/Court-Case-Management-System/Entity"

//CaseSearchRepository ...
type CaseSearchRepository interface {
	Cases() ([]entity.Case, []error)
	Case(caseNum string) (*entity.Case, []error)
}

//JudgeSearchRepository ...
type JudgeSearchRepository interface {
	Judges() ([]entity.Judge, []error)
	Judge(judgeID string) (*entity.Judge, []error)
}
