package services

import (
	"testing"

	"github.com/takuya-okada-01/heart-note/repository/database"
	repository "github.com/takuya-okada-01/heart-note/repository/user_repository"
)

func TestSignUpWithEmailandPassword(t *testing.T) {
	db := database.Connect()
	userRepository := repository.NewUserRepository(db)
	userService := NewUserService(userRepository)

	email := "hoge@hoge.com"
	password := "password"

	id, err := userService.SignUpWithEmailAndPassword(email, password)
	if err != nil {
		t.Fatal(err)
	}

	user, err := userService.SelectUser(id)
	if err != nil {
		t.Fatal(err)
	}

	if user.Email != email {
		t.Fatal("email is not equal: want ", email, " got ", user.Email)
	}

	userService.DeleteUser(id)

}
func TestLoginWithEmailAndPassword(t *testing.T) {
	db := database.Connect()
	userRepository := repository.NewUserRepository(db)
	userService := NewUserService(userRepository)

	email := "hoge@hoge.com"
	password := "password"

	want, err := userService.SignUpWithEmailAndPassword(email, password)
	if err != nil {
		t.Fatal(err)
	}

	id, err := userService.LoginWithEmailAndPassword(email, password)
	if err != nil {
		t.Fatal(err)
	}

	if id != want {
		t.Fatal("id is not equal: want ", want, " got ", id)
	}

	userService.DeleteUser(id)
}
