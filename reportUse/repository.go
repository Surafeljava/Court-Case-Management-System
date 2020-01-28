package reportUse

import entity "github.com/Surafeljava/Court-Case-Management-System/Entity"

type ReportRepository interface {
	GetClosedCases() ([]entity.Case, error)
	GetOpenCases() ([]entity.Case, error)
	GetAllCases() ([]entity.Case, error)
	GetCriminalJudges() ([]entity.Judge, error)
	GetCivilJudges() ([]entity.Judge, error)
	GetAllJudges() ([]entity.Judge, error)
	GetAllNotifications() ([]entity.Notification, error)
}
