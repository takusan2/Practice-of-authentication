package services

import (
	"github.com/takuya-okada-01/heart-note/crypto"
	"github.com/takuya-okada-01/heart-note/repository/database/entity"
	"github.com/takuya-okada-01/heart-note/services/repository_interface"
)

type userService struct {
	userRepository repository_interface.UserRepository
}

func NewUserService(userRepository repository_interface.UserRepository) *userService {
	return &userService{userRepository: userRepository}
}

func (u *userService) InsertUser(user *entity.User) (string, error) {
	return u.userRepository.InsertUser(user)
}

func (u *userService) SelectUser(id string) (entity.User, error) {
	return u.userRepository.SelectUser(id)
}

func (u *userService) UpdateUser(user *entity.User) error {
	return u.userRepository.UpdateUser(user)
}

func (u *userService) DeleteUser(id string) error {
	return u.userRepository.DeleteUser(id)
}

func (u *userService) LoginWithEmailAndPassword(email, password string) (string, error) {
	var user entity.User

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

func (u *userService) SignUpWithEmailAndPassword(email, password string) (string, error) {
	var user entity.User
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
