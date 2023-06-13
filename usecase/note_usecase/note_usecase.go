package usecase

import (
	"github.com/takuya-okada-01/heart-note/domain"
)

type noteUseCase struct {
	noteRepository domain.INoteRepository
}

func NewNoteUseCase(noteRepository domain.INoteRepository) domain.INoteUseCase {
	return &noteUseCase{noteRepository: noteRepository}
}

func (n *noteUseCase) InsertNote(note *domain.Note) (string, error) {
	return n.noteRepository.InsertNote(note)
}

func (n *noteUseCase) SelectNoteByID(userID string, id string) (domain.Note, error) {
	return n.noteRepository.SelectNoteByID(userID, id)
}

func (n *noteUseCase) SelectNoteByFolderID(userID string, folderID string) ([]domain.Note, error) {
	return n.noteRepository.SelectNoteByFolderID(userID, folderID)
}

func (n *noteUseCase) UpdateNote(note *domain.Note) error {
	return n.noteRepository.UpdateNote(note)
}

func (n *noteUseCase) DeleteNoteByID(userID string, id string) error {
	return n.noteRepository.DeleteNoteByID(userID, id)
}
