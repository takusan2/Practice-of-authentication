package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Note struct {
	ID        string    `gorm:"type:varchar(36);primary_key;"`
	Name      string    `gorm:"type:varchar(255);not null;"`
	Content   string    `gorm:"type:text;"`
	FolderID  string    `gorm:"type:varchar(36);not null;foreignkey:folders(id);"`
	UserID    string    `gorm:"type:varchar(36);not null;foreignkey:users(id);"`
	CreatedAt time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;"`
	UpdatedAt time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;"`
}

func (n *Note) BeforeCreate(tx *gorm.DB) (err error) {
	n.ID = uuid.New().String()
	return
}

type INoteRepository interface {
	InsertNote(note *Note) (string, error)
	SelectNoteByID(userID string, id string) (Note, error)
	SelectNoteByFolderID(userID string, folderID string) ([]Note, error)
	UpdateNote(note *Note) error
	DeleteNoteByID(userID string, id string) error
}

type INoteUseCase interface {
	InsertNote(note *Note) (string, error)
	SelectNoteByID(userID string, id string) (Note, error)
	SelectNoteByFolderID(userID string, folderID string) ([]Note, error)
	UpdateNote(note *Note) error
	DeleteNoteByID(userID string, id string) error
}
