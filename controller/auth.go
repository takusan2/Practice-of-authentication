package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/takuya-okada-01/heart-note/domain"
)

type authController struct {
	authUseCase domain.IAuthUseCase
}

func NewAuthController(router *gin.RouterGroup, authUseCase domain.IAuthUseCase) {
	ac := &authController{authUseCase: authUseCase}
	router.POST("/signup", ac.SignUpWithEmailAndPassword)
	router.POST("/login", ac.LoginWithEmailAndPassword)
	router.POST("/logout", ac.Logout)
}

func (ac *authController) SignUpWithEmailAndPassword(ctx *gin.Context) {
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")

	tokenString, err := ac.authUseCase.SignUpWithEmailAndPassword(ctx, email, password)
	if err != nil {
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}

	ctx.SetCookie("AccessToken", tokenString, 3600*24*30, "/", "localhost", false, true)
	ctx.JSON(200, gin.H{"id": tokenString})
}

func (ac *authController) LoginWithEmailAndPassword(ctx *gin.Context) {
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")

	tokenString, err := ac.authUseCase.LoginWithEmailAndPassword(ctx, email, password)
	if err != nil {
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}

	ctx.SetCookie("AccessToken", tokenString, 3600*24*30, "/", "localhost", false, true)
	ctx.JSON(200, gin.H{"id": tokenString})
}

func (ac *authController) Logout(ctx *gin.Context) {
	ac.authUseCase.Logout(ctx)
	ctx.JSON(200, gin.H{"message": "logout"})
}
