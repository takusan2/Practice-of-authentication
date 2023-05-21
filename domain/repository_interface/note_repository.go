package repository_interface

import "github.com/takuya-okada-01/heart-note/domain"

type NoteRepository interface {
	InsertNote(note *domain.Note) (string, error)
	SelectNoteByID(userID string, id string) (domain.Note, error)
	SelectNoteByFolderID(userID string, folderID string) ([]domain.Note, error)
	UpdateNote(note *domain.Note) error
	DeleteNoteByID(userID string, id string) error
}
