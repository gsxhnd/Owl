package service

import (
	"github.com/gsxhnd/owl/server/db/database"
	"github.com/gsxhnd/owl/server/model"
	"github.com/gsxhnd/owl/utils"
)

type FileService interface {
	CreateAnimes([]model.File) error
	DeleteAnimes([]uint) error
	UpdateAnime(model.File) error
	GetAnimes(*database.Pagination) ([]model.File, error)
}

type fileService struct {
	logger utils.Logger
	db     database.Driver
}

func NewFileService(l utils.Logger, db database.Driver) FileService {
	return fileService{
		logger: l,
		db:     db,
	}
}

// CreateAnimes implements fileService.
func (s fileService) CreateAnimes(data []model.File) error {
	return s.db.CreateAnimes(data)
}

// DeleteAnimes implements fileService.
func (s fileService) DeleteAnimes(ids []uint) error {
	return s.db.DeleteAnimes(ids)
}

// UpdateAnime implements fileService.
func (s fileService) UpdateAnime(data model.File) error {
	return s.db.UpdateAnime(data)
}

// GetAnimes implements fileService.
func (s fileService) GetAnimes(p *database.Pagination) ([]model.File, error) {
	return s.db.GetAnimes(p)
}
