package usecase

import (
	"github.com/takuya-okada-01/heart-note/crypto"
	"github.com/takuya-okada-01/heart-note/domain"
	"github.com/takuya-okada-01/heart-note/domain/repository_interface"
)

type UserUseCase interface {
	InsertUser(user *domain.User) (string, error)
	SelectUser(id string) (domain.User, error)
	UpdateUser(user *domain.User) error
	DeleteUser(id string) error
	SignUpWithEmailAndPassword(email, password string) (string, error)
	LoginWithEmailAndPassword(email, password string) (string, error)
}

type userUseCase struct {
	userRepository repository_interface.UserRepository
}

func NewUserUseCase(userRepository repository_interface.UserRepository) UserUseCase {
	return &userUseCase{userRepository: userRepository}
}

func (u *userUseCase) InsertUser(user *domain.User) (string, error) {
	return u.userRepository.InsertUser(user)
}

func (u *userUseCase) SelectUser(id string) (domain.User, error) {
	return u.userRepository.SelectUser(id)
}

func (u *userUseCase) UpdateUser(user *domain.User) error {
	return u.userRepository.UpdateUser(user)
}

func (u *userUseCase) DeleteUser(id string) error {
	return u.userRepository.DeleteUser(id)
}

func (u *userUseCase) LoginWithEmailAndPassword(email, password string) (string, error) {
	var user domain.User

	user, err := u.userRepository.SelectUserByEmail(email)
	if err != nil {
		return "", err
	}

	err = crypto.CompareHashAndPassword(user.PasswordHash, password+user.Salt)
	if err != nil {
		return "", err
	}

	return user.ID, err
}

func (u *userUseCase) SignUpWithEmailAndPassword(email, password string) (string, error) {
	var user domain.User
	user.Email = email
	user.PasswordHash = password
	user.Salt = crypto.SecureRandomBase64()
	crypto.PasswordEncrypt(user.PasswordHash + user.Salt)

	id, err := u.userRepository.InsertUser(&user)
	if err != nil {
		return "", err
	}

	return id, err
}
