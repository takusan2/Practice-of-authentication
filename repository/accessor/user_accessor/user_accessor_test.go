package accessor

import (
	"context"
	"fmt"
	"testing"

	"github.com/takuya-okada-01/heart-note/repository/database"
	"github.com/takuya-okada-01/heart-note/repository/database/entity"
)

func TestUserAccessor(t *testing.T) {
	db := database.Connect()
	insertUser := entity.User{
		Name:         "test",
		Email:        "hoge@hoge.com",
		PasswordHash: "test",
	}

	userAccessor := NewUserAccessor(db)
	err := db.ResetModel(context.Background(), (*entity.User)(nil))
	if err != nil {
		t.Fatal(err)
	}

	id, err := userAccessor.InsertUser(&insertUser)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(id)
	if err != nil {
		t.Fatal(err)
	}

	want := insertUser.ID

	user, err := userAccessor.SelectLastUser()
	if err != nil {
		t.Fatal(err)
	}

	if user.ID != want {
		t.Errorf("InsertUser == %d, want %d", user.ID, want)
	}
	fmt.Printf("user: %+v\n", user)
}

func TestAuthentication(t *testing.T) {
	db := database.Connect()
	insertUser := entity.User{
		Name:         "test",
		Email:        "hoge@hoge.com",
		PasswordHash: "test",
	}

	err := db.ResetModel(context.Background(), (*entity.User)(nil))
	if err != nil {
		t.Fatal(err)
	}

	userAccessor := NewUserAccessor(db)

	want, err := userAccessor.SignUpWithEmailAndPassword(insertUser.Email, insertUser.PasswordHash)
	if err != nil {
		t.Fatal(err)
	}

	id, err := userAccessor.LoginWithEmailAndPassword(
		insertUser.Email, insertUser.PasswordHash,
	)
	if err != nil {
		t.Fatal(err)
	}

	if id != want {
		t.Errorf("InsertUser.ID == %d, want %d", id, want)
	}
}
