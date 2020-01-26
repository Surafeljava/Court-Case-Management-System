package reportUse

import entity "github.com/Surafeljava/Court-Case-Management-System/Entity"

type ReportService interface {
	CreateCourtReport() ([]entity.Case, []entity.Judge, []error)
}
