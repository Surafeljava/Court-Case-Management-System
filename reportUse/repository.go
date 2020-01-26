package reportUse

import entity "github.com/Surafeljava/Court-Case-Management-System/Entity"

type ReportRepository interface {
	CreateCourtReport() ([]entity.Case, []entity.Judge, []error)
}
