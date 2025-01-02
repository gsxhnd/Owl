package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gsxhnd/owl/server/errno"
	"github.com/gsxhnd/owl/server/model"
	"github.com/gsxhnd/owl/server/service"
	"github.com/gsxhnd/owl/utils"
)

type FileTagHandler interface {
	CreateFileTags(ctx *fiber.Ctx) error
	DeleteFileTags(ctx *fiber.Ctx) error
}

type fileTagHandle struct {
	valid  *validator.Validate
	svc    service.FileTagService
	logger utils.Logger
}

func NewFileTagHandler(svc service.FileTagService, v *validator.Validate, l utils.Logger) FileTagHandler {
	return &fileTagHandle{
		valid:  v,
		svc:    svc,
		logger: l,
	}
}

// @Summary      Create movie tags
// @Description  Create movie tags
// @Tags         movie_tag
// @Produce      json
// @Param        default body []model.MovieTag true "default"
// @Success      200  {object}   errno.errno
// @Router       /movie_tag [post]
func (h *fileTagHandle) CreateFileTags(ctx *fiber.Ctx) error {
	var body = make([]model.FileTag, 0)
	if err := ctx.BodyParser(&body); err != nil {
		h.logger.Errorf(err.Error())
		return ctx.JSON(errno.DecodeError(err))
	}

	if err := h.valid.Var(body, "dive"); err != nil {
		return ctx.JSON(errno.DecodeError(err))
	}

	err := h.svc.CreateFileTags(body)
	return ctx.JSON(errno.DecodeError(err))
}

// @Summary      Delete movie tags
// @Description  Delete movie tags
// @Tags         movie_tag
// @Produce      json
// @Param        default body []uint true "default"
// @Success      200 {object} errno.errno{data=nil}
// @Router       /movie_tag [delete]
func (h *fileTagHandle) DeleteFileTags(ctx *fiber.Ctx) error {
	var body = make([]uint, 0)
	if err := ctx.BodyParser(&body); err != nil {
		h.logger.Errorf(err.Error())
		return ctx.JSON(errno.DecodeError(err))
	}

	if err := h.valid.Var(body, "dive"); err != nil {
		return ctx.JSON(errno.DecodeError(err))
	}

	err := h.svc.DeleteFileTags(body)
	return ctx.JSON(errno.DecodeError(err))
}
