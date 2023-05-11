package entity

import (
	"time"
)

type Folder struct {
	ID        int       `gorm:"type:int;primary_key;auto_increment;"`
	Name      string    `gorm:"type:varchar(255);not null;"`
	ParentID  int       `gorm:"type:int;unsigned;not null;default:0,foreignkey:folders(id);"`
	UserID    string    `gorm:"type:varchar(36);foreignkey:users(id);"`
	CreatedAt time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;"`
	UpdatedAt time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;"`
}
