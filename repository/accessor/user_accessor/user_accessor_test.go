package accessor

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/takuya-okada-01/heart-note/repository/database"
	"github.com/takuya-okada-01/heart-note/repository/database/entity"
)

func TestUserInsertAndSelect(t *testing.T) {
	db := database.Connect()
	userAccessor := NewUserAccessor(db)
	insertUser := entity.User{
		Name:         "test",
		Email:        "hoge@hoge.com",
		PasswordHash: "password",
	}

	id, err := userAccessor.InsertUser(&insertUser)
	if err != nil {
		t.Fatal(err)
	}

	want := id
	user, err := userAccessor.SelectUser(id)
	if err != nil {
		t.Fatal(err)
	}

	if user.ID != want {
		t.Errorf("InsertUser == %s, want %s", user.ID, want)
	}
	userAccessor.DeleteUser(id)
	result, err := json.MarshalIndent(user, "", "  ")
	fmt.Print(string(result))
}

func TestUserUpdate(t *testing.T) {
	db := database.Connect()
	userAccessor := NewUserAccessor(db)
	insertUser := entity.User{
		Name:         "test",
		Email:        "hogehoge@hoge.com",
		PasswordHash: "password",
	}

	id, err := userAccessor.InsertUser(&insertUser)
	if err != nil {
		t.Fatal(err)
	}

	insertUser.ID = id
	insertUser.Name = "test_updated"
	userAccessor.UpdateUser(&insertUser)

	want := "test_updated"
	user, err := userAccessor.SelectUser(id)
	if user.Name != want {
		t.Errorf("InsertUser == %s, want %s", user.Name, want)
	}
	userAccessor.DeleteUser(id)
}

func TestUserDelete(t *testing.T) {
	db := database.Connect()
	userAccessor := NewUserAccessor(db)
	insertUser := entity.User{
		Name:         "test",
		Email:        "hogsfeh@hoge.com",
		PasswordHash: "password",
	}

	id, err := userAccessor.InsertUser(&insertUser)
	if err != nil {
		t.Fatal(err)
	}

	err = userAccessor.DeleteUser(id)
	if err != nil {
		t.Fatal(err)
	}

	user, err := userAccessor.SelectUser(id)
	if err == nil {
		t.Errorf("InsertUser == %s, want %s", user, "nil")
	}
}
