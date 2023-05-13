package dto

import "time"

type NoteRequest struct {
	Name     string `json:"name"`
	Content  string `json:"content"`
	FolderID string `json:"folder_id"`
	UserID   string `json:"user_id"`
}

type NoteResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Content   string    `json:"content"`
	FolderID  string    `json:"folder_id"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;"`
	UpdatedAt time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;"`
}
