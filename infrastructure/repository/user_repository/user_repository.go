package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/takuya-okada-01/heart-note/crypto"
	"github.com/takuya-okada-01/heart-note/domain"
	"github.com/takuya-okada-01/heart-note/domain/repository_interface"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository_interface.UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) InsertUser(user *domain.User) (string, error) {
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
	return user.ID, err
}

func (u *userRepository) SelectUser(id string) (domain.User, error) {
	var user domain.User
	err := u.db.Select("*").Where("id = ?", id).First(&user).Error
	return user, err
}

func (u *userRepository) UpdateUser(user *domain.User) error {
	err := u.db.Model(&user).Where("id = ?", user.ID).Update(user).Error
	return err
}

func (u *userRepository) DeleteUser(id string) error {
	err := u.db.Where("id = ?", id).Delete(&domain.User{}).Error
	return err
}

func (u *userRepository) SelectUserByEmail(email string) (domain.User, error) {
	var user domain.User
	err := u.db.Select("*").Where("email = ?", email).First(&user).Error
	return user, err
}
