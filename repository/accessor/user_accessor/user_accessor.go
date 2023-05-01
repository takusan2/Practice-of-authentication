package accessor

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/takuya-okada-01/heart-note/crypto"
	"github.com/takuya-okada-01/heart-note/repository/database/entity"
	"github.com/takuya-okada-01/heart-note/services/accessor_interface"
)

type userAccessor struct {
	db *gorm.DB
}

func NewUserAccessor(db *gorm.DB) accessor_interface.UserAccessor {
	return &userAccessor{db: db}
}

func (u *userAccessor) InsertUser(user *entity.User) (string, error) {
	var err error
	user.Salt = crypto.SecureRandomBase64()
	user.PasswordHash, err = crypto.PasswordEncrypt(user.PasswordHash + user.Salt)
	if err != nil {
		return "", err
	}

	result := u.db.Table("users").Create(user)
	err = result.Error
	if err != nil {
		return "", err
	}

	fmt.Print("id: ", user.ID, "\n")

	return user.ID, err
}

func (u *userAccessor) SelectUser(id string) (entity.User, error) {
	var user entity.User
	err := u.db.Select("*").Where("id = ?", id).First(&user).Error
	return user, err
}

func (u *userAccessor) UpdateUser(user *entity.User) error {
	err := u.db.Model(&user).Where("id = ?", user.ID).Update(user).Error
	return err
}

func (u *userAccessor) DeleteUser(id string) error {
	err := u.db.Where("id = ?", id).Delete(&entity.User{}).Error
	return err
}

// func (u *userAccessor) LoginWithEmailAndPassword(email, password string) (string, error) {
// 	var user entity.User
// 	err := u.db.NewSelect().Model(&user).Where("email = ?", email).Scan(context.Background())
// 	if err != nil {
// 		return 0, err
// 	}
// 	err = crypto.CompareHashAndPassword(user.PasswordHash, password)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return user.ID, err
// }

// func (u *userAccessor) SignUpWithEmailAndPassword(email, password string) (int64, error) {
// 	var user entity.User
// 	user.Email = email
// 	user.PasswordHash = password

// 	id, err := u.InsertUser(&user)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return id, err
// }
