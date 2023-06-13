package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/takuya-okada-01/heart-note/domain"
)

type userUseCase struct {
	userRepository domain.IUserRepository
}

func NewUserUseCase(userRepository domain.IUserRepository) domain.IUserUseCase {
	return &userUseCase{userRepository: userRepository}
}

func (u *userUseCase) UpdateUser(ctx *gin.Context, id string, name string) error {
	user, err := u.userRepository.SelectUser(id)
	if err != nil {
		return err
	}
	user.Name = name
	return u.userRepository.UpdateUser(&user)
}

func (u *userUseCase) DeleteUser(ctx *gin.Context, id string) error {
	return u.userRepository.DeleteUser(id)
}
