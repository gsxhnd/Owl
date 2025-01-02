package storage

import (
	"errors"

	"github.com/google/wire"
	"github.com/gsxhnd/owl/utils"
)

type Storage interface {
	Ping() error
	GetImage(cover string, id uint, filename string) ([]byte, string, error)
	SaveImage(data []byte, cover string, id uint, filename string) error
}

func NewStorage(cfg *utils.Config, l utils.Logger) (Storage, error) {
	if cfg.Storage.Type == "minio" {
		return NewMinioStorage(cfg)
	}

	if cfg.Storage.Type == "local" {
		return NewLocalStorage(cfg, l)
	}

	return nil, errors.New("no storage type")
}

var StorageSet = wire.NewSet(NewStorage)
