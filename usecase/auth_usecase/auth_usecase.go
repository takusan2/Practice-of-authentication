package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/takuya-okada-01/heart-note/domain"
	"github.com/takuya-okada-01/heart-note/utils"
	"github.com/takuya-okada-01/heart-note/utils/crypto"
)

type authUseCase struct {
	userRepository domain.IUserRepository
}

func NewAuthUseCase(userRepository domain.IUserRepository) domain.IAuthUseCase {
	return &authUseCase{userRepository: userRepository}
}

func (au *authUseCase) SignUpWithEmailAndPassword(ctx *gin.Context, email string, password string) (string, error) {
	var user domain.User
	user.Email = email

	// hash password
	user.Salt = crypto.SecureRandomBase64()
	passwordHash, err := crypto.PasswordEncrypt(user.PasswordHash + user.Salt)
	if err != nil {
		return "", err
	}
	user.PasswordHash = passwordHash

	// insert user
	id, err := au.userRepository.InsertUser(&user)
	if err != nil {
		return "", err
	}

	// generate token
	tokenString, err := utils.GenerateToken(id)
	if err != nil {
		return "", err
	}

	return tokenString, err
}

func (au *authUseCase) LoginWithEmailAndPassword(ctx *gin.Context, email string, password string) (string, error) {
	var user domain.User

	// fetch user by email from entity
	user, err := au.userRepository.SelectUserByEmail(email)
	if err != nil {
		return "", err
	}

	// compare password
	err = crypto.CompareHashAndPassword(user.PasswordHash, password+user.Salt)
	if err != nil {
		return "", err
	}

	// generate token
	tokenString, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return tokenString, err
}

func (au *authUseCase) Logout(ctx *gin.Context) error {
	ctx.SetCookie("AccessToken", "", 0, "/", "", false, true)
	return nil
}
