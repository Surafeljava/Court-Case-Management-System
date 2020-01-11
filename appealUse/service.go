package appealUse

type AppealService interface {
	AppealGet(caseNum string) (*entity.Case, []error)
}
