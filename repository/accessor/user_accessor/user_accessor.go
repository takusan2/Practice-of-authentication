package accessor

import (
	"context"

	"github.com/takuya-okada-01/heart-note/crypto"
	"github.com/takuya-okada-01/heart-note/repository/entity"
	"github.com/uptrace/bun"
)

type UserAccessor interface {
	InsertUser(user *entity.User) (int64, error)
	SelectUser(id int64) (entity.User, error)
	UpdateUser(user *entity.User) error
	DeleteUser(id int64) error
	SelectLastUser() (entity.User, error)
	LoginWithEmailAndPassword(email, password string) (int64, error)
	SignUpWithEmailAndPassword(email, password string) (int64, error)
}

type userAccessor struct {
	db *bun.DB
}

func NewUserAccessor(db *bun.DB) UserAccessor {
	return &userAccessor{db: db}
}

func (u *userAccessor) SelectLastUser() (entity.User, error) {
	var user entity.User
	err := u.db.NewSelect().Model(&user).Order("id DESC").Limit(1).Scan(context.Background())
	return user, err
}

func (u *userAccessor) SelectUser(id int64) (entity.User, error) {
	var user entity.User
	err := u.db.NewSelect().Model(&user).Where("id = ?", id).Scan(context.Background())
	return user, err
}

func (u *userAccessor) LoginWithEmailAndPassword(email, password string) (int64, error) {
	var user entity.User
	err := u.db.NewSelect().Model(&user).Where("email = ?", email).Scan(context.Background())
	if err != nil {
		return 0, err
	}
	err = crypto.CompareHashAndPassword(user.PasswordHash, password)
	if err != nil {
		return 0, err
	}
	return user.ID, err
}

func (u *userAccessor) InsertUser(user *entity.User) (int64, error) {
	var err error
	user.PasswordHash, err = crypto.PasswordEncrypt(user.PasswordHash)
	if err != nil {
		return 0, err
	}

	result, err := u.db.NewInsert().Model(user).Exec(context.Background())
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	return id, err
}

func (u *userAccessor) SignUpWithEmailAndPassword(email, password string) (int64, error) {
	var user entity.User
	user.Email = email
	user.PasswordHash = password

	id, err := u.InsertUser(&user)
	if err != nil {
		return 0, err
	}
	return id, err
}

func (u *userAccessor) UpdateUser(user *entity.User) error {
	_, err := u.db.NewUpdate().Model(user).Where("id = ?", user.ID).Exec(context.Background())
	return err
}

func (u *userAccessor) DeleteUser(id int64) error {
	_, err := u.db.NewDelete().Model(&entity.User{}).Where("id = ?", id).Exec(context.Background())
	return err
}
