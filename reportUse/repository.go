package reportUse

import "github.com/Surafeljava/Court-Case-Management-System/entity"

type ReportRepository interface {
	CreateCourtReport() ([]entity.Case, []entity.Judge, []error)
}
