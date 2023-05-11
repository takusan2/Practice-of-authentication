package repository

import (
	"fmt"
	"testing"

	"github.com/takuya-okada-01/heart-note/infrastructure/database"
	"github.com/takuya-okada-01/heart-note/infrastructure/database/entity"
)

func TestNoteInsert(t *testing.T) {
	db := database.Connect()
	defer db.Close()

	note := &entity.Note{
		Name:     "test",
		Content:  "test",
		FolderID: 0,
		UserID:   "test",
	}

	nr := NewNoteRepository(db)
	_, err := nr.InsertNote(note)
	if err != nil {
		t.Errorf("InsertNote() == %v, want %v", err, "nil")
	}

	err = nr.DeleteNoteByID(note.UserID, note.ID)
	if err != nil {
		t.Errorf("DeleteNoteByID() == %v, want %v", err, "nil")
	}
}

func TestNoteSelect(t *testing.T) {
	db := database.Connect()
	defer db.Close()

	note := &entity.Note{
		Name:     "test",
		Content:  "test",
		UserID:   "test",
		FolderID: 0,
	}

	nr := NewNoteRepository(db)

	id, err := nr.InsertNote(note)
	if err != nil {
		t.Errorf("InsertNote() == %v, want %v", err, "nil")
	}

	testNote, err := nr.SelectNoteByID(note.UserID, id)
	fmt.Println(testNote)
	if err != nil {
		t.Errorf("SelectNoteByID() == %v, want %v", err, "nil")
	}

	notes, err := nr.SelectNoteByFolderID(note.UserID, 0)
	fmt.Println(notes)
	if err != nil {
		t.Errorf("SelectNoteByFolderID() == %v, want %v", err, "nil")
	}

	err = nr.DeleteNoteByID(note.UserID, id)
	if err != nil {
		t.Errorf("DeleteNoteByID() == %v, want %v", err, "nil")
	}
}

func TestNoteUpdate(t *testing.T) {
	db := database.Connect()
	defer db.Close()

	note := &entity.Note{
		Name:     "test",
		Content:  "test",
		UserID:   "test",
		FolderID: 0,
	}

	nr := NewNoteRepository(db)

	id, err := nr.InsertNote(note)
	if err != nil {
		t.Errorf("InsertNote() == %v, want %v", err, "nil")
	}

	testNote, err := nr.SelectNoteByID(note.UserID, id)
	fmt.Println(testNote)
	if err != nil {
		t.Errorf("SelectNoteByID() == %v, want %v", err, "nil")
	}

	testNote.Name = "test2"
	err = nr.UpdateNote(note.UserID, &testNote)
	if err != nil {
		t.Errorf("UpdateNote() == %v, want %v", err, "nil")
	}

	err = nr.DeleteNoteByID(note.UserID, id)
	if err != nil {
		t.Errorf("DeleteNoteByID() == %v, want %v", err, "nil")
	}
}

func TestNoteDelete(t *testing.T) {
	db := database.Connect()
	defer db.Close()

	note := &entity.Note{
		Name:     "shoud_be_deleted",
		Content:  "test",
		UserID:   "test",
		FolderID: 0,
	}

	nr := NewNoteRepository(db)

	id, err := nr.InsertNote(note)
	if err != nil {
		t.Errorf("InsertNote() == %v, want %v", err, "nil")
	}

	err = nr.DeleteNoteByID(note.UserID, id)
	if err != nil {
		t.Errorf("DeleteNoteByID() == %v, want %v", err, "nil")
	}
}
