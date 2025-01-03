package service

import (
	"github.com/gsxhnd/owl/server/db/database"
	"github.com/gsxhnd/owl/server/storage"
	"github.com/gsxhnd/owl/utils"
)

type PingService interface {
	Ping() (*pingResp, error)
}

type pingService struct {
	logger  utils.Logger
	db      database.Driver
	storage storage.Storage
}

func NewPingService(l utils.Logger, db database.Driver, s storage.Storage) PingService {
	return &pingService{
		logger:  l,
		db:      db,
		storage: s,
	}
}

type pingResp struct {
	DBVersion string `json:"db_version"`
}

func (p *pingService) Ping() (*pingResp, error) {
	if err := p.db.Ping(); err != nil {
		p.logger.Errorf(err.Error())
		return nil, err
	}

	if err := p.storage.Ping(); err != nil {
		p.logger.Errorf(err.Error())
		return nil, err
	}

	version, err := p.db.Version()
	if err != nil {
		p.logger.Errorf(err.Error())
		return nil, err
	}

	return &pingResp{
		DBVersion: version,
	}, nil
}
