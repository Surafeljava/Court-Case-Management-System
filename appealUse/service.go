package appealUse

import entity "github.com/Surafeljava/Court-Case-Management-System/Entity"

type AppealService interface {
	Appeal(oppNum string) (*entity.Case, *entity.Opponent, *entity.Witness, *entity.Decision, []error)
}
