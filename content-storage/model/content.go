package model

import "time"

type Content struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	FileName   string    `json:"file_name"`
	FileType   string    `json:"file_type"`
	Path       string    `json:"path"`
	UploadedAt time.Time `json:"uploaded_at"`
}
