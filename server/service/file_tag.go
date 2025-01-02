package service

import (
	"github.com/gsxhnd/owl/server/db/database"
	"github.com/gsxhnd/owl/server/model"
	"github.com/gsxhnd/owl/utils"
)

type FileTagService interface {
	CreateFileTags(data []model.FileTag) error
	DeleteFileTags(ids []uint) error
}

type fileTagService struct {
	logger utils.Logger
	db     database.Driver
}

func NewFileTagService(l utils.Logger, db database.Driver) FileTagService {
	return &fileTagService{
		logger: l,
		db:     db,
	}
}

func (s *fileTagService) CreateFileTags(movieTags []model.FileTag) error {
	return s.db.CreateMovieTags(movieTags)
}

func (s *fileTagService) DeleteFileTags(ids []uint) error {
	return s.db.DeleteMovieTags(ids)
}
