package appealUse

import entity "github.com/Surafeljava/Court-Case-Management-System/Entity"

type AppealRepositroy interface {
	Appeal(oppNum string) (*entity.Case, *entity.Opponent, *entity.Witness, *entity.Decision, []error)
	RelationForAppeal(oppNum string) (*entity.Relation, []error)
	CaseForAppeal(caseNum string) (*entity.Case, []error)
	WitnessForAppeal(caseNum string) (*entity.Witness, []error)
	DecisionForAppeal(caseNum string) (*entity.Decision, []error)
}
