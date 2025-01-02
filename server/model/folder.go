package model

import (
	"time"
)

type Folder struct {
	Id        uint       `json:"id"`
	Name      string     `json:"name"`
	Pid       uint       `json:"pid"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
