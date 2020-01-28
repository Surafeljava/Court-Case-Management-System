package repository

import (
	entity "github.com/Surafeljava/Court-Case-Management-System/Entity"
	"github.com/jinzhu/gorm"
)

type AdminCourtRepositoryImpl struct {
	conn *gorm.DB
}

func NewAdminCourtRepositoryImpl(Conn *gorm.DB) *AdminCourtRepositoryImpl {
	return &AdminCourtRepositoryImpl{conn: Conn}
}

func (acr *AdminCourtRepositoryImpl) Court() (*entity.Court, []error) {
	court := entity.Court{}

	err := acr.conn.First(&court).GetErrors()
	if len(err) > 0 {
		return nil, err
	}
	return &court, nil
}

func (acr *AdminCourtRepositoryImpl) Admin() (*entity.Admin, []error) {
	admin := entity.Admin{}

	err := acr.conn.First(&admin).GetErrors()
	if len(err) > 0 {
		return nil, err
	}
	return &admin, nil
}

func (acr *AdminCourtRepositoryImpl) CreateCourt(court *entity.Court) (*entity.Court, []error) {
	err := acr.conn.Save(&court).GetErrors()
	if len(err) > 0 {
		return nil, err
	}
	return court, nil
}
func (acr *AdminCourtRepositoryImpl) UpdateCourt(court *entity.Court) (*entity.Court, []error) {
	return nil, nil
}

// DeleteCourt(id int) error
func (acr *AdminCourtRepositoryImpl) CreateAdmin(admin *entity.Admin) (*entity.Admin, []error) {
	err := acr.conn.Save(&admin).GetErrors()
	if len(err) > 0 {
		return nil, err
	}
	return admin, nil
}
