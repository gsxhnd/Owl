package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gsxhnd/owl/server/db/database"
	"github.com/gsxhnd/owl/server/errno"
	"github.com/gsxhnd/owl/server/model"
	"github.com/gsxhnd/owl/server/service"
	"github.com/gsxhnd/owl/utils"
)

type LabelHandler interface {
	CreateLabels(ctx *fiber.Ctx) error
	DeleteLabels(ctx *fiber.Ctx) error
	// UpdateActor(ctx *fiber.Ctx) error
	GetLabels(ctx *fiber.Ctx) error
}

type labelHandle struct {
	valid  *validator.Validate
	svc    service.LabelService
	logger utils.Logger
}

func NewLabelHandler(svc service.LabelService, v *validator.Validate, l utils.Logger) LabelHandler {
	return &labelHandle{
		valid:  v,
		svc:    svc,
		logger: l,
	}
}

// CreateActors implements ActorHandler.
// @Summary      Create label
// @Description  Create actor
// @Tags         actor
// @Produce      json
// @Success      200
// @Router       /actor [post]
func (h *labelHandle) CreateLabels(ctx *fiber.Ctx) error {
	var body = make([]model.Label, 0)
	if err := ctx.BodyParser(&body); err != nil {
		h.logger.Errorf(err.Error())
		return ctx.JSON(errno.DecodeError(err))
	}

	if err := h.valid.Var(body, "dive"); err != nil {
		return ctx.JSON(errno.DecodeError(err))
	}

	err := h.svc.CreateLabels(body)
	return ctx.JSON(errno.DecodeError(err))
}

// DeleteActors implements ActorHandler.
// @Summary      Delete actors
// @Description  Delete actor
// @Tags         actor
// @Produce      json
// @Success      200
// @Router       /actor [delete]
func (h *labelHandle) DeleteLabels(ctx *fiber.Ctx) error {
	var body = make([]uint, 0)
	if err := ctx.BodyParser(&body); err != nil {
		h.logger.Errorf(err.Error())
		return ctx.JSON(errno.DecodeError(err))
	}

	if err := h.valid.Var(body, "dive"); err != nil {
		return ctx.JSON(errno.DecodeError(err))
	}

	err := h.svc.DeleteLabels(body)
	return ctx.JSON(errno.DecodeError(err))
}

// UpdateActor implements ActorHandler.
// @Summary      Update a actor by id
// @Description  Update a actor by id
// @Tags         actor
// @Accept       json
// @Produce      json
// @Param        tag body model.Actor true "Actor object"
// @Success      200 {object} errno.errno
// @Router       /actor [put]
func (h *labelHandle) UpdateActor(ctx *fiber.Ctx) error {
	var body model.Label
	if err := ctx.BodyParser(&body); err != nil {
		h.logger.Errorf(err.Error())
		return ctx.JSON(errno.DecodeError(err))
	}

	if err := h.valid.Struct(body); err != nil {
		h.logger.Errorf(err.Error())
		return ctx.JSON(errno.DecodeError(err))
	}
	h.svc.UpdateLabels(&body)

	return ctx.JSON(errno.OK)
}

// GetActors implements ActorHandler.
// @Summary      Get actors
// @Description  Get actors List
// @Tags         actor
// @Produce      json
// @Success      200  {object}   errno.errno{data=[]model.Actor}
// @Router       /actor [get]
func (h *labelHandle) GetLabels(ctx *fiber.Ctx) error {
	var p = database.Pagination{
		Limit:  uint64(ctx.QueryInt("page_size", 50)),
		Offset: uint64(ctx.QueryInt("page_size", 50) * ctx.QueryInt("page", 0)),
	}

	data, err := h.svc.GetLabels(&p)

	return ctx.JSON(errno.DecodeError(err).WithData(data))
}

// GetActor implements ActorHandler.
// TODO: get actor movies
func (h *labelHandle) GetActor(ctx *fiber.Ctx) error {
	panic("unimplemented")
}
