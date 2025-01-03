package service

import (
	"github.com/gsxhnd/owl/server/db/database"
	"github.com/gsxhnd/owl/server/model"
	"github.com/gsxhnd/owl/utils"
)

type LabelService interface {
	CreateLabels([]model.Label) error
	DeleteLabels([]uint) error
	UpdateLabels(*model.Label) error
	GetLabels(*database.Pagination) ([]model.Label, error)
}

type labelService struct {
	logger utils.Logger
	db     database.Driver
}

func NewLabelService(l utils.Logger, db database.Driver) LabelService {
	return &labelService{
		logger: l,
		db:     db,
	}
}

func (s labelService) CreateLabels(actors []model.Label) error {
	return nil
}

func (s labelService) DeleteLabels(ids []uint) error {
	return nil
}

func (s labelService) UpdateLabels(actor *model.Label) error {
	return nil
}

func (s labelService) GetLabels(p *database.Pagination) ([]model.Label, error) {
	return nil, nil
}
