package appealUse

import "github.com/Surafeljava/Court-Case-Management-System/entity"

type AppealService interface {
	AppealGet(caseNum string) (*entity.Case, []error)
}
