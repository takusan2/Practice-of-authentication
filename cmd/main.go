package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/takuya-okada-01/heart-note/controller"
	"github.com/takuya-okada-01/heart-note/infrastructure/database"
	note_repository "github.com/takuya-okada-01/heart-note/infrastructure/repository/note_repository"
	user_repository "github.com/takuya-okada-01/heart-note/infrastructure/repository/user_repository"

	"github.com/takuya-okada-01/heart-note/mysession"
	note_usecase "github.com/takuya-okada-01/heart-note/usecase/note_usecase"
	user_usecase "github.com/takuya-okada-01/heart-note/usecase/user_usecase"
)

func main() {

	db := database.Connect()
	ur := user_repository.NewUserRepository(db)
	uu := user_usecase.NewUserUseCase(ur)
	uc := controller.NewUserController(uu)

	nr := note_repository.NewNoteRepository(db)
	nu := note_usecase.NewNoteUseCase(nr)
	nc := controller.NewNoteController(nu)

	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	v1 := router.Group("/")
	{
		v1.POST("/signup", uc.SignUp)
		v1.POST("/login", uc.Login)
	}

	v2 := router.Group("/user")
	v2.Use(mysession.SessionCheck())
	{
		v2.GET("/", uc.SelectUser)
		v2.PUT("/", uc.UpdateUser)
		v2.DELETE("/:id", uc.DeleteUser)
	}
	v3 := router.Group("/note")
	v3.Use(mysession.SessionCheck())
	{
		v3.POST("/", nc.InsertNote)
		v3.GET("/:id", nc.SelectNoteByID)
		v3.PUT("/:id", nc.UpdateNote)
		v3.DELETE("/:id", nc.DeleteNoteByID)
	}

	v4 := router.Group("/folders")
	v4.Use(mysession.SessionCheck())
	{
		v4.GET("/", nc.SelectNoteByFolderID)
	}

	router.Run()

}
