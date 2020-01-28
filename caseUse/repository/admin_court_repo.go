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

func (acr *AdminCourtRepositoryImpl) CreateCourt(court *entity.Court) (*entity.Court, []error) {
	return nil, nil
}
func (acr *AdminCourtRepositoryImpl) UpdateCourt(court *entity.Court) (*entity.Court, []error) {
	return nil, nil
}

// DeleteCourt(id int) error
func (acr *AdminCourtRepositoryImpl) CreateAdmin(admin *entity.Admin) (*entity.Admin, []error) {
	return nil, nil
}
