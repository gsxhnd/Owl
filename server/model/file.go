package model

import (
	"time"
)

type File struct {
	Id        uint       `json:"id"`
	Name      string     `json:"name" validate:"required"`
	FolderId  uint       `json:"folder_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
