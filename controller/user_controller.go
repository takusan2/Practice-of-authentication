package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/takuya-okada-01/heart-note/controller/dto"
	"github.com/takuya-okada-01/heart-note/domain"
	usecase "github.com/takuya-okada-01/heart-note/usecase/user_usecase"
)

type UserController interface {
	SignUp(ctx *gin.Context)
	Login(ctx *gin.Context)
	SelectUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

type userController struct {
	userService usecase.UserUseCase
}

func NewUserController(userService usecase.UserUseCase) UserController {
	return &userController{userService: userService}
}

func (uc *userController) SignUp(ctx *gin.Context) {
	var user dto.UserRequest
	ctx.BindJSON(&user)

	tokenString, err := uc.userService.SignUpWithEmailAndPassword(user.Email, user.PasswordHash)
	if err != nil {
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}

	// cookieにトークンをセット
	ctx.SetCookie("SessionID", tokenString, 3600*24*30, "/", "localhost", false, true)
	ctx.JSON(200, gin.H{"id": tokenString})
}

func (uc *userController) Login(ctx *gin.Context) {
	var user dto.UserRequest
	ctx.BindJSON(&user)

	tokenString, err := uc.userService.LoginWithEmailAndPassword(user.Email, user.PasswordHash)

	if err != nil {
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}

	// cookieにトークンをセット
	ctx.SetCookie("SessionID", tokenString, 3600*24*30, "/", "localhost", false, true)
	ctx.JSON(200, gin.H{"token": tokenString})
}

func (uc *userController) SelectUser(ctx *gin.Context) {
	userID, ok := ctx.Keys["user_id"].(string)
	if !ok {
		ctx.JSON(500, gin.H{"message": "user not found"})
		return
	}

	user, err := uc.userService.SelectUser(userID)
	responoseUser := dto.UserResonse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	if err != nil {
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(200, responoseUser)
}

func (uc *userController) UpdateUser(ctx *gin.Context) {
	userID, ok := ctx.Keys["user_id"].(string)
	if !ok {
		ctx.JSON(500, gin.H{"message": "user not found"})
		return
	}

	var user dto.UserRequest
	ctx.BindJSON(&user)

	err := uc.userService.UpdateUser(
		&domain.User{
			ID:           userID,
			Name:         user.Name,
			Email:        user.Email,
			PasswordHash: user.PasswordHash,
		},
	)
	if err != nil {
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "success"})
}

func (uc *userController) DeleteUser(ctx *gin.Context) {
	userID, ok := ctx.Keys["user_id"].(string)
	if !ok {
		ctx.JSON(500, gin.H{"message": "user not found"})
		return
	}

	err := uc.userService.DeleteUser(userID)
	if err != nil {
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "success"})
}
