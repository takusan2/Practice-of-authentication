package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/takuya-okada-01/heart-note/controller/dto"
	"github.com/takuya-okada-01/heart-note/domain"
)

type userController struct {
	userUseCase domain.IUserUseCase
}

func NewUserController(router gin.RouterGroup, userUseCase domain.IUserUseCase) {
	uc := &userController{userUseCase: userUseCase}
	router.PUT("/user", uc.UpdateUser)
	router.DELETE("/user", uc.DeleteUser)
}

func (uc *userController) UpdateUser(ctx *gin.Context) {
	userID, ok := ctx.Keys["user_id"].(string)
	if !ok {
		ctx.JSON(500, gin.H{"message": "user not found"})
		return
	}

	var user dto.UserRequest
	ctx.BindJSON(&user)

	err := uc.userUseCase.UpdateUser(
		ctx,
		userID,
		user.Name,
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

	err := uc.userUseCase.DeleteUser(ctx, userID)
	if err != nil {
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "success"})
}
