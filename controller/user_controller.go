package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/takuya-okada-01/heart-note/controller/dto"
	"github.com/takuya-okada-01/heart-note/domain"
	"github.com/takuya-okada-01/heart-note/session_manager"
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
	userService    usecase.UserUseCase
	sessionManager session_manager.SessionManager
}

func NewUserController(userService usecase.UserUseCase, sessionManager session_manager.SessionManager) UserController {
	return &userController{userService: userService, sessionManager: sessionManager}
}

func (uc *userController) SignUp(ctx *gin.Context) {
	var user dto.UserRequest
	ctx.BindJSON(&user)

	id, err := uc.userService.SignUpWithEmailAndPassword(user.Email, user.PasswordHash)
	if err != nil {
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}
	if err := uc.sessionManager.SetSession(ctx); err != nil {
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

	if err := uc.sessionManager.SetSession(ctx); err != nil {
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"id": id})
}

func (uc *userController) SelectUser(ctx *gin.Context) {
	session := sessions.Default(ctx)
	userID := session.Get("UserId").(string)

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

	ctx.JSON(200, gin.H{"user": responoseUser})
}

func (uc *userController) UpdateUser(ctx *gin.Context) {
	session := sessions.Default(ctx)
	userID := session.Get("UserId").(string)

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
	session := sessions.Default(ctx)
	userID := session.Get("UserId").(string)

	err := uc.userService.DeleteUser(userID)
	if err != nil {
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "success"})
}
