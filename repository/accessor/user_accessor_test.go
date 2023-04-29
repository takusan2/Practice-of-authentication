package accessor

import (
	"fmt"
	"testing"
	"time"

	"github.com/takuya-okada-01/heart-note/repository"
	"github.com/takuya-okada-01/heart-note/repository/entity"
)

func TestUserAccessor(t *testing.T) {
	db := repository.Connect()
	insertUser := entity.User{
		Name:         "test",
		Email:        "hoge@hoge.com",
		PasswordHash: "test",
		Salt:         "test",
		CreatedAt: time.Date(
			2021, 1, 1, 0, 0, 0, 0, time.Local,
		).UTC(),
	}
	userAccessor := NewUserAccessor(db)

	id, err := userAccessor.InsertUser(&insertUser)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(id)
	if err != nil {
		t.Fatal(err)
	}

	user, err := userAccessor.SelectLastUser()
	if err != nil {
		t.Fatal(err)
	}

	want := entity.User{
		ID:           id,
		Name:         insertUser.Name,
		Email:        insertUser.Email,
		PasswordHash: insertUser.PasswordHash,
		Salt:         insertUser.Salt,
		CreatedAt:    insertUser.CreatedAt,
	}

	if user != want {
		t.Errorf("InsertUser == %v, want %v", user, want)
	}
}
