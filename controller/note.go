package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/takuya-okada-01/heart-note/controller/dto"
	"github.com/takuya-okada-01/heart-note/domain"
)

type noteController struct {
	router *gin.RouterGroup
	nu     domain.INoteUseCase
}

func NewNoteController(router *gin.RouterGroup, noteUseCase domain.INoteUseCase) {
	nc := &noteController{router: router, nu: noteUseCase}
	router.POST("/note", nc.InsertNote)
	router.GET("/note/:id", nc.SelectNoteByID)
	router.GET("/note/folder/:folder_id", nc.SelectNoteByFolderID)
	router.PUT("/note/:id", nc.UpdateNote)
	router.DELETE("/note/:id", nc.DeleteNoteByID)
}

func (nc *noteController) InsertNote(ctx *gin.Context) {
	userID, ok := ctx.Keys["user_id"].(string)
	if !ok {
		ctx.JSON(500, gin.H{"message": "user not found"})
		return
	}

	var note dto.NoteRequest
	ctx.BindJSON(&note)

	id, err := nc.nu.InsertNote(&domain.Note{
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

	note, err := nc.nu.SelectNoteByID(userID, id)
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

	notes, err := nc.nu.SelectNoteByFolderID(userID, folderID)
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

	err := nc.nu.UpdateNote(&domain.Note{
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

	err := nc.nu.DeleteNoteByID(userID, id)
	if err != nil {
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "success"})
}
