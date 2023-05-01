package repository

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/takuya-okada-01/heart-note/repository/database"
	"github.com/takuya-okada-01/heart-note/repository/database/entity"
)

func TestUserInsertAndSelect(t *testing.T) {
	db := database.Connect()
	userRepository := NewUserRepository(db)
	insertUser := entity.User{
		Name:         "test",
		Email:        "hoge@hoge.com",
		PasswordHash: "password",
	}

	id, err := userRepository.InsertUser(&insertUser)
	if err != nil {
		t.Fatal(err)
	}

	want := id
	user, err := userRepository.SelectUser(id)
	if err != nil {
		t.Fatal(err)
	}

	if user.ID != want {
		t.Errorf("InsertUser == %s, want %s", user.ID, want)
	}
	userRepository.DeleteUser(id)
	result, err := json.MarshalIndent(user, "", "  ")
	fmt.Print(string(result))
}

func TestUserUpdate(t *testing.T) {
	db := database.Connect()
	userRepository := NewUserRepository(db)
	insertUser := entity.User{
		Name:         "test",
		Email:        "hogehoge@hoge.com",
		PasswordHash: "password",
	}

	id, err := userRepository.InsertUser(&insertUser)
	if err != nil {
		t.Fatal(err)
	}

	insertUser.ID = id
	insertUser.Name = "test_updated"
	userRepository.UpdateUser(&insertUser)

	want := "test_updated"
	user, err := userRepository.SelectUser(id)
	if user.Name != want {
		t.Errorf("InsertUser == %s, want %s", user.Name, want)
	}
	userRepository.DeleteUser(id)
}

func TestUserDelete(t *testing.T) {
	db := database.Connect()
	userRepository := NewUserRepository(db)
	insertUser := entity.User{
		Name:         "test",
		Email:        "hogsfeh@hoge.com",
		PasswordHash: "password",
	}

	id, err := userRepository.InsertUser(&insertUser)
	if err != nil {
		t.Fatal(err)
	}

	err = userRepository.DeleteUser(id)
	if err != nil {
		t.Fatal(err)
	}

	user, err := userRepository.SelectUser(id)
	if err == nil {
		t.Errorf("InsertUser == %s, want %s", user, "nil")
	}
}

func TestSelectByEmail(t *testing.T) {
	db := database.Connect()
	userRepository := NewUserRepository(db)
	insertUser := entity.User{
		Name:         "test",
		Email:        "hogehoge@hoge.com",
		PasswordHash: "password",
	}

	_, err := userRepository.InsertUser(&insertUser)
	if err != nil {
		t.Fatal(err)
	}

	want := insertUser.Email
	user, err := userRepository.SelectUserByEmail(
		insertUser.Email,
	)
	if err != nil {
		t.Fatal(err)
	}

	if user.Email != want {
		t.Errorf("InsertUser == %s, want %s", user.Email, want)
	}

	userRepository.DeleteUser(user.ID)
}
