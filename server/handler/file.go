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

type FileHandler interface {
	CreateFile(ctx *fiber.Ctx) error
	DeleteFiles(ctx *fiber.Ctx) error
	GetFile(ctx *fiber.Ctx) error
	GetFiles(ctx *fiber.Ctx) error
}

type fileHandle struct {
	valid  *validator.Validate
	svc    service.FileService
	logger utils.Logger
}

func NewFileHandler(svc service.FileService, v *validator.Validate, l utils.Logger) FileHandler {
	return &fileHandle{
		valid:  v,
		svc:    svc,
		logger: l,
	}
}

// @Summary      Create animes
// @Description  Create animes
// @Tags         anime
// @Produce      json
// @Success      200
// @Router       /anime [post]
func (h *fileHandle) CreateFile(ctx *fiber.Ctx) error {
	var body model.File
	if err := ctx.BodyParser(&body); err != nil {
		h.logger.Errorf(err.Error())
		return ctx.JSON(errno.DecodeError(err))
	}

	if err := h.valid.Var(body, "dive"); err != nil {
		return ctx.JSON(errno.DecodeError(err))
	}

	err := h.svc.CreateFile(body)
	return ctx.JSON(errno.DecodeError(err))
}

// @Summary      Delete animes
// @Description  Delete animes
// @Tags         anime
// @Produce      json
// @Success      200
// @Router       /anime [delete]
func (h *fileHandle) DeleteFiles(ctx *fiber.Ctx) error {
	body := make([]uint, 0)
	if err := ctx.BodyParser(&body); err != nil {
		h.logger.Errorf(err.Error())
		return ctx.JSON(errno.DecodeError(err))
	}

	if err := h.valid.Var(body, "dive"); err != nil {
		return ctx.JSON(errno.DecodeError(err))
	}

	err := h.svc.DeleteFiles(body)
	return ctx.JSON(errno.DecodeError(err))
}

// @Summary      List animes
// @Description  Get animes
// @Tags         anime
// @Produce      json
// @Success      200
// @Router       /anime [get]
func (h *fileHandle) GetFiles(ctx *fiber.Ctx) error {
	p := database.Pagination{
		Limit:  uint64(ctx.QueryInt("page_size", 50)),
		Offset: uint64(ctx.QueryInt("page_size", 50) * ctx.QueryInt("page", 0)),
	}

	data, err := h.svc.GetAnimes(&p)

	return ctx.JSON(errno.DecodeError(err).WithData(data))
}

func (h *fileHandle) GetFile(ctx *fiber.Ctx) error {
	return ctx.Status(200).SendString("pong")
}

func (h *fileHandle) UpdateAnime(ctx *fiber.Ctx) error {
	return ctx.Status(200).SendString("pong")
}
