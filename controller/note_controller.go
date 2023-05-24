package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/takuya-okada-01/heart-note/controller/dto"
	"github.com/takuya-okada-01/heart-note/domain"
	usecase "github.com/takuya-okada-01/heart-note/usecase/note_usecase"
)

type NoteController interface {
	InsertNote(ctx *gin.Context)
	SelectNoteByID(ctx *gin.Context)
	SelectNoteByFolderID(ctx *gin.Context)
	UpdateNote(ctx *gin.Context)
	DeleteNoteByID(ctx *gin.Context)
}

type noteController struct {
	noteService usecase.NoteUseCase
}

func NewNoteController(noteService usecase.NoteUseCase) NoteController {
	return &noteController{noteService: noteService}
}

func (nc *noteController) InsertNote(ctx *gin.Context) {
	userID, ok := ctx.Keys["user_id"].(string)
	if !ok {
		ctx.JSON(500, gin.H{"message": "user not found"})
		return
	}

	var note dto.NoteRequest
	ctx.BindJSON(&note)

	id, err := nc.noteService.InsertNote(&domain.Note{
		Name:     note.Name,
		Content:  note.Content,
		FolderID: note.FolderID,
		UserID:   userID,
	})
	if err != nil {
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"id": id})
}

func (nc *noteController) SelectNoteByID(ctx *gin.Context) {
	userID, ok := ctx.Keys["user_id"].(string)
	if !ok {
		ctx.JSON(500, gin.H{"message": "user not found"})
		return
	}

	id := ctx.Param("id")

	note, err := nc.noteService.SelectNoteByID(userID, id)
	if err != nil {
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(200, note)
}

func (nc *noteController) SelectNoteByFolderID(ctx *gin.Context) {
	userID, ok := ctx.Keys["user_id"].(string)
	if !ok {
		ctx.JSON(500, gin.H{"message": "user not found"})
		return
	}
	folderID := ctx.Query("folderId")
	fmt.Print("folderID", folderID)

	notes, err := nc.noteService.SelectNoteByFolderID(userID, folderID)
	if err != nil {
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(200, notes)
}

func (nc *noteController) UpdateNote(ctx *gin.Context) {
	userID, ok := ctx.Keys["user_id"].(string)
	if !ok {
		ctx.JSON(500, gin.H{"message": "user not found"})
		return
	}

	var note dto.NoteRequest
	id := ctx.Param("id")
	ctx.BindJSON(&note)

	err := nc.noteService.UpdateNote(&domain.Note{
		ID:       id,
		Name:     note.Name,
		Content:  note.Content,
		FolderID: note.FolderID,
		UserID:   userID,
	})
	if err != nil {
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "success"})
}

func (nc *noteController) DeleteNoteByID(ctx *gin.Context) {
	userID, ok := ctx.Keys["user_id"].(string)
	if !ok {
		ctx.JSON(500, gin.H{"message": "user not found"})
		return
	}
	id := ctx.Param("id")

	err := nc.noteService.DeleteNoteByID(userID, id)
	if err != nil {
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "success"})
}
