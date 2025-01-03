package service

import (
	"github.com/gsxhnd/owl/server/db/database"
	"github.com/gsxhnd/owl/server/model"
	"github.com/gsxhnd/owl/utils"
)

type FileService interface {
	CreateFile(model.File) error
	DeleteFiles([]uint) error
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

func (s fileService) CreateFile(data model.File) error {
	datas := make([]model.File, 0)
	datas = append(datas, data)
	return s.db.CreateFiles(datas)
}

func (s fileService) DeleteFiles(ids []uint) error {
	return nil
}

func (s fileService) UpdateAnime(data model.File) error {
	return nil
}

func (s fileService) GetAnimes(p *database.Pagination) ([]model.File, error) {
	return nil, nil
}
