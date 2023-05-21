package repository_interface

import "github.com/takuya-okada-01/heart-note/domain"

type UserRepository interface {
	InsertUser(user *domain.User) (string, error)
	SelectUser(id string) (domain.User, error)
	SelectUserByEmail(email string) (domain.User, error)
	UpdateUser(user *domain.User) error
	DeleteUser(id string) error
}
