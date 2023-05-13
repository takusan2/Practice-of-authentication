package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/takuya-okada-01/heart-note/domain"
)

type NoteRepository interface {
	InsertNote(note *domain.Note) (string, error)
	SelectNoteByID(userID string, id string) (domain.Note, error)
	SelectNoteByFolderID(userID string, folderID string) ([]domain.Note, error)
	UpdateNote(note *domain.Note) error
	DeleteNoteByID(userID string, id string) error
}

type noteRepository struct {
	db *gorm.DB
}

func NewNoteRepository(db *gorm.DB) NoteRepository {
	return &noteRepository{db: db}
}

func (n *noteRepository) InsertNote(note *domain.Note) (string, error) {
	result := n.db.Table("notes").Create(note)
	if result.Error != nil {
		return "", result.Error
	}
	return note.ID, nil
}

func (n *noteRepository) SelectNoteByID(userID string, id string) (domain.Note, error) {
	var note domain.Note
	result := n.db.Table("notes").Where("user_id = ?", userID).Where("id = ?", id).First(&note)
	if result.Error != nil {
		return note, result.Error
	}

	return note, nil
}

func (n *noteRepository) SelectNoteByFolderID(userID string, folderID string) ([]domain.Note, error) {
	var notes []domain.Note
	result := n.db.Table("notes").Where("user_id = ?", userID).Where("folder_id = ?", folderID).Find(&notes)
	if result.Error != nil {
		return notes, result.Error
	}
	return notes, nil
}

func (n *noteRepository) UpdateNote(note *domain.Note) error {
	result := n.db.Table("notes").Where("user_id = ?", note.UserID).Where("id = ?", note.ID).Update(note)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (n *noteRepository) DeleteNoteByID(userID string, id string) error {
	result := n.db.Table("notes").Where("user_id = ?", userID).Where("id = ?", id).Delete(&domain.Note{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
