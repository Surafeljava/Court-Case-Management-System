package appealUse

import "github.com/Surafeljava/Court-Case-Management-System/entity"

type AppealRepositroy interface {
	AppealGet(caseNum string) (*entity.Case, []error)
}

//Comment
