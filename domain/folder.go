package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Folder struct {
	ID        string    `gorm:"type:varchar(36);primary_key;"`
	Name      string    `gorm:"type:varchar(255);not null;"`
	ParentID  int       `gorm:"type:int;unsigned;not null;default:0,foreignkey:folders(id);"`
	UserID    string    `gorm:"type:varchar(36);foreignkey:users(id);"`
	CreatedAt time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;"`
	UpdatedAt time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;"`
}

func (n Folder) BeforeCreate(tx *gorm.DB) (err error) {
	n.ID = uuid.New().String()
	return
}
