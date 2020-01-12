package appealUse

type AppealRepositroy interface {
	AppealGet(caseNum string) (*entity.Case, []error)

}
//Comment
