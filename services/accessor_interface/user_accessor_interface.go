package accessor_interface

import (
	"github.com/takuya-okada-01/heart-note/repository/database/entity"
)

type UserAccessor interface {
	InsertUser(user *entity.User) (string, error)
	SelectUser(id string) (entity.User, error)
	UpdateUser(user *entity.User) error
	DeleteUser(id string) error
	// SelectLastUser() (entity.User, error)
}
