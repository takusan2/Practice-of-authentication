package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/takuya-okada-01/heart-note/config"
	"github.com/takuya-okada-01/heart-note/controller"
	"github.com/takuya-okada-01/heart-note/infrastructure/database"
	note_repository "github.com/takuya-okada-01/heart-note/infrastructure/repository/note_repository"
	user_repository "github.com/takuya-okada-01/heart-note/infrastructure/repository/user_repository"
	"github.com/takuya-okada-01/heart-note/middleware"
	auth_usecase "github.com/takuya-okada-01/heart-note/usecase/auth_usecase"
	note_usecase "github.com/takuya-okada-01/heart-note/usecase/note_usecase"
)

func main() {
	godotenv.Load(config.ProjectRootPath + "/.env")

	db := database.Connect()
	defer db.Close()
	defer fmt.Print("db closed\n")

	router := gin.Default()
	store := cookie.NewStore([]byte(os.Getenv("SECRET_KEY")))
	router.Use(sessions.Sessions("AccessToken", store))

	ur := user_repository.NewUserRepository(db)
	// uu := user_usecase.NewUserUseCase(ur)
	// uc := controller.NewUserController(uu)

	au := auth_usecase.NewAuthUseCase(ur)
	v1 := router.Group("/")
	controller.NewAuthController(v1, au)

	// v2 := router.Group("/user")
	// v2.Use(middleware.VerifyToken())
	// v2.Use(middleware.SessionCheck())
	// {
	// 	v2.GET("/", uc.SelectUser)
	// 	v2.PUT("/", uc.UpdateUser)
	// 	v2.DELETE("/:id", uc.DeleteUser)
	// }
	nr := note_repository.NewNoteRepository(db)
	nu := note_usecase.NewNoteUseCase(nr)
	v3 := router.Group("/note")
	v3.Use(middleware.VerifyToken())
	v3.Use(middleware.SessionCheck())
	controller.NewNoteController(v3, nu)

	router.Run()

}
