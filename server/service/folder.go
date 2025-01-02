package service

import (
	"github.com/gsxhnd/owl/server/db/database"
	"github.com/gsxhnd/owl/server/model"
	"github.com/gsxhnd/owl/server/storage"
	"github.com/gsxhnd/owl/utils"
)

type FolderService interface {
	CreateFolder([]model.Folder) error
	DeleteFolder(ids []uint) error
	UpdateFolder(*model.Folder) error
	GetFolders(*database.Pagination) ([]model.Folder, error)
}

type folderService struct {
	logger  utils.Logger
	db      database.Driver
	storage storage.Storage
}

func NewFolderService(l utils.Logger, db database.Driver, s storage.Storage) FolderService {
	return &folderService{
		logger:  l,
		db:      db,
		storage: s,
	}
}

func (s *folderService) CreateFolder(movies []model.Folder) error {
	return s.db.CreateFolder(movies)
}

func (s *folderService) DeleteFolder(ids []uint) error {
	return s.db.DeleteFolders(ids)
}

func (s *folderService) UpdateFolder(m *model.Folder) error {
	return s.db.UpdateFolder(m)
}

func (s *folderService) GetFolders(p *database.Pagination) ([]model.Folder, error) {
	return s.db.GetFolders(p)
}
