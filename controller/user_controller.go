package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/takuya-okada-01/heart-note/controller/dto"
	"github.com/takuya-okada-01/heart-note/infrastructure/database/entity"
	services "github.com/takuya-okada-01/heart-note/services/user_service"
)

type UserController interface {
	SignUp(ctx *gin.Context)
	Login(ctx *gin.Context)
	SelectUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

type userController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return &userController{userService: userService}
}

func (uc *userController) SignUp(ctx *gin.Context) {
	var user dto.UserRequest
	ctx.BindJSON(&user)

	id, err := uc.userService.SignUpWithEmailAndPassword(user.Email, user.PasswordHash)
	if err != nil {
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"id": id})
}

func (uc *userController) Login(ctx *gin.Context) {

	var user dto.UserRequest
	ctx.BindJSON(&user)

	id, err := uc.userService.LoginWithEmailAndPassword(user.Email, user.PasswordHash)
	if err != nil {
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}

	session := sessions.Default(ctx)
	session.Set("UserId", id)
	session.Save()

	ctx.JSON(200, gin.H{"id": id})
}

func (uc *userController) SelectUser(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := uc.userService.SelectUser(id)

	responoseUser := dto.UserResonse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	if err != nil {
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"user": responoseUser})
}

func (uc *userController) UpdateUser(ctx *gin.Context) {
	var user dto.UserRequest
	id := ctx.Param("id")
	ctx.BindJSON(&user)

	err := uc.userService.UpdateUser(
		&entity.User{
			ID:           id,
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
	id := ctx.Param("id")

	err := uc.userService.DeleteUser(id)
	if err != nil {
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "success"})
}
