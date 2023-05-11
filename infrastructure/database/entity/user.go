package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type User struct {
	ID           string    `gorm:"type:varchar(36);primary_key;"`
	Name         string    `gorm:"type:varchar(255);not null;"`
	Email        string    `gorm:"type:varchar(255);not null;unique;"`
	PasswordHash string    `gorm:"type:varchar(255);not null;"`
	Salt         string    `gorm:"type:varchar(255);not null;"`
	CreatedAt    time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;"`
	UpdatedAt    time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return
}
