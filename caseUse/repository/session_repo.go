package repository

import (
	entity "github.com/Surafeljava/Court-Case-Management-System/Entity"
	"github.com/Surafeljava/Court-Case-Management-System/caseUse"
	"github.com/jinzhu/gorm"
)

// SessionGormRepo implements user.SessionRepository interface
type SessionGormRepo struct {
	conn *gorm.DB
}

// NewSessionGormRepo  returns a new SessionGormRepo object
func NewSessionGormRepo(db *gorm.DB) caseUse.SessionRepository {
	return &SessionGormRepo{conn: db}
}

// Session returns a given stored session
func (sr *SessionGormRepo) Session(sessionID string) (*entity.Session, []error) {
	session := entity.Session{}
	errs := sr.conn.Find(&session, "uuid=?", sessionID).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &session, errs
}

// StoreSession stores a given session
func (sr *SessionGormRepo) StoreSession(session *entity.Session) (*entity.Session, []error) {
	sess := session
	errs := sr.conn.Save(sess).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return sess, errs
}

// DeleteSession deletes a given session
func (sr *SessionGormRepo) DeleteSession(sessionID string) (*entity.Session, []error) {
	sess, errs := sr.Session(sessionID)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = sr.conn.Delete(sess, sess.ID).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return sess, errs
}
