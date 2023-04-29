package accessor

import (
	"context"

	"github.com/takuya-okada-01/heart-note/repository/entity"
	"github.com/uptrace/bun"
)

type UserAccessor interface {
	InsertUser(user *entity.User) (int64, error)
	SelectUser(id int) error
	UpdateUser(user *entity.User) error
	DeleteUser(id int) error
	SelectLastUser() (entity.User, error)
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

func (u *userAccessor) InsertUser(user *entity.User) (int64, error) {
	result, err := u.db.NewInsert().Model(user).Exec(context.Background())
	id, err := result.LastInsertId()
	return id, err
}

func (u *userAccessor) SelectUser(id int) error {
	var user entity.User
	err := u.db.NewSelect().Model(&user).Where("id = ?", id).Scan(context.Background())
	return err
}

func (u *userAccessor) UpdateUser(user *entity.User) error {
	_, err := u.db.NewUpdate().Model(user).Where("id = ?", user.ID).Exec(context.Background())
	return err
}

func (u *userAccessor) DeleteUser(id int) error {
	_, err := u.db.NewDelete().Model(&entity.User{}).Where("id = ?", id).Exec(context.Background())
	return err
}
