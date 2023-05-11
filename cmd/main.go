package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/takuya-okada-01/heart-note/controller"
	"github.com/takuya-okada-01/heart-note/infrastructure/database"
	repository "github.com/takuya-okada-01/heart-note/infrastructure/repository/user_repository"
	services "github.com/takuya-okada-01/heart-note/services/user_service"
	"github.com/takuya-okada-01/heart-note/session"
)

func main() {

	db := database.Connect()
	ur := repository.NewUserRepository(db)
	us := services.NewUserService(ur)
	uc := controller.NewUserController(us)

	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	v1 := router.Group("/")
	{
		v1.POST("/signup", uc.SignUp)
		v1.POST("/login", uc.Login)
	}

	v2 := router.Group("/user")
	v2.Use(session.SessionCheck())
	{
		v2.GET("/:id", uc.SelectUser)
		v2.PUT("/:id", uc.UpdateUser)
		v2.DELETE("/:id", uc.DeleteUser)
	}
	router.Run()

}
