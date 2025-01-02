package model

type FileTag struct {
	Id      uint    `json:"id"`
	FileId  uint    `json:"file_id"`
	TagId   uint    `json:"tag_id"`
	TagName *string `json:"tag_name"`
}
