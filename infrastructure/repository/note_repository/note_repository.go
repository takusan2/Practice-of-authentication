package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/takuya-okada-01/heart-note/infrastructure/database/entity"
)

type NoteRepository interface {
	InsertNote(note *entity.Note) (string, error)
	SelectNoteByID(userID string, id string) (entity.Note, error)
	SelectNoteByFolderID(userID string, folderID int) ([]entity.Note, error)
	UpdateNote(userID string, note *entity.Note) error
	DeleteNoteByID(userID string, id string) error
}

type noteRepository struct {
	db *gorm.DB
}

func NewNoteRepository(db *gorm.DB) NoteRepository {
	return &noteRepository{db: db}
}

func (n *noteRepository) InsertNote(note *entity.Note) (string, error) {
	result := n.db.Table("notes").Create(note)
	if result.Error != nil {
		return "", result.Error
	}
	return note.ID, nil
}

func (n *noteRepository) SelectNoteByID(userID string, id string) (entity.Note, error) {
	var note entity.Note
	result := n.db.Table("notes").Where("user_id = ?", userID).Where("id = ?", id).First(&note)
	if result.Error != nil {
		return note, result.Error
	}

	return note, nil
}

func (n *noteRepository) SelectNoteByFolderID(userID string, folderID int) ([]entity.Note, error) {
	var notes []entity.Note
	result := n.db.Table("notes").Where("user_id = ?", userID).Where("folder_id = ?", folderID).Find(&notes)
	if result.Error != nil {
		return notes, result.Error
	}
	return notes, nil
}

func (n *noteRepository) UpdateNote(userID string, note *entity.Note) error {
	result := n.db.Table("notes").Where("user_id = ?", userID).Where("id = ?", note.ID).Update(note)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (n *noteRepository) DeleteNoteByID(userID string, id string) error {
	result := n.db.Table("notes").Where("user_id = ?", userID).Where("id = ?", id).Delete(&entity.Note{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
