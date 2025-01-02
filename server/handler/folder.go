package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gsxhnd/owl/server/db/database"
	"github.com/gsxhnd/owl/server/errno"
	"github.com/gsxhnd/owl/server/model"
	"github.com/gsxhnd/owl/server/service"
	"github.com/gsxhnd/owl/server/storage"
	"github.com/gsxhnd/owl/utils"
)

type FolderHandler interface {
	CreateFolder(ctx *fiber.Ctx) error
	DeleteFolder(ctx *fiber.Ctx) error
	UpdateFolder(ctx *fiber.Ctx) error
	GetFolders(ctx *fiber.Ctx) error
}

type folderHandle struct {
	valid   *validator.Validate
	svc     service.FolderService
	logger  utils.Logger
	storage storage.Storage
}

func NewFolderHandler(svc service.FolderService, v *validator.Validate, s storage.Storage, l utils.Logger) FolderHandler {
	return &folderHandle{
		valid:   v,
		svc:     svc,
		logger:  l,
		storage: s,
	}
}

// @Summary      Create folder
// @Description  Create folder
// @Tags         folder
// @Produce      json
// @Success      200  {object}  errno.errno
// @Router       /folder [post]
func (h *folderHandle) CreateFolder(ctx *fiber.Ctx) error {
	var body = make([]model.Folder, 0)
	if err := ctx.BodyParser(&body); err != nil {
		h.logger.Errorf(err.Error())
		return ctx.JSON(errno.DecodeError(err))
	}

	if err := h.valid.Var(body, "dive"); err != nil {
		return ctx.JSON(errno.DecodeError(err))
	}

	err := h.svc.CreateFolder(body)
	return ctx.JSON(errno.DecodeError(err))
}

// @Summary      Delete folder
// @Description  Delete folder
// @Tags         folder
// @Accept       json
// @Produce      json
// @Param        default body []uint true "default"
// @Success      200 {object} errno.errno{data=nil}
// @Router       /folder [delete]
func (h *folderHandle) DeleteFolder(ctx *fiber.Ctx) error {
	var body = make([]uint, 0)
	if err := ctx.BodyParser(&body); err != nil {
		h.logger.Errorf(err.Error())
		return ctx.JSON(errno.DecodeError(err))
	}

	if err := h.valid.Var(body, "dive"); err != nil {
		return ctx.JSON(errno.DecodeError(err))
	}

	err := h.svc.DeleteFolder(body)
	return ctx.JSON(errno.DecodeError(err))
}

// @Summary      Update a folder by id
// @Description  Update a folder info by id
// @Tags         folder
// @Accept       json
// @Produce      json
// @Param        tag body model.Folder true "folder object"
// @Success      200 {object} errno.errno
// @Router       /folder [put]
func (h *folderHandle) UpdateFolder(ctx *fiber.Ctx) error {
	var body model.Folder
	if err := ctx.BodyParser(&body); err != nil {
		h.logger.Errorf(err.Error())
		return ctx.JSON(errno.DecodeError(err))
	}

	if err := h.valid.Struct(body); err != nil {
		h.logger.Errorf(err.Error())
		return ctx.JSON(errno.DecodeError(err))
	}

	if err := h.svc.UpdateFolder(&body); err != nil {
		h.logger.Errorf(err.Error())
		return ctx.JSON(errno.DecodeError(err))
	}

	return ctx.JSON(errno.OK)
}

// @Summary      Get folders
// @Description  Get folders
// @Tags         folder
// @Produce      json
// @Param        page_size query int false "int valid" default(50)
// @Param        page query int false "int valid" default(1)
// @Success      200 {object} errno.errno{data=[]model.Folder}
// @Router       /folder [get]
func (h *folderHandle) GetFolders(ctx *fiber.Ctx) error {
	var p = database.Pagination{
		Limit:  uint64(ctx.QueryInt("page_size", 50)),
		Offset: uint64(ctx.QueryInt("page_size", 50) * (ctx.QueryInt("page", 1) - 1)),
	}

	if err := h.valid.Struct(p); err != nil {
		return ctx.JSON(errno.DecodeError(err))
	}

	data, err := h.svc.GetFolders(&p)

	return ctx.JSON(errno.DecodeError(err).WithData(data))
}

// @Summary      Get movies
// @Description  Get movies
// @Tags         movie
// @Produce      json
// @Param        code path string true "movie code"
// @Success      200 {object} errno.errno{data=model.MovieInfo}
// @Router       /movie/info/:code [get]
// func (h *folderHandle) GetMovieInfo(ctx *fiber.Ctx) error {
// 	code := ctx.Params("code", "")
// 	data, err := h.svc.GetMovieInfo(code)
// 	return ctx.JSON(errno.DecodeError(err).WithData(data))
// }

// @Summary      Search movies
// @Description  Search movies by code
// @Tags         movie
// @Produce      json
// @Param        code query string true "movie code"
// @Success      200 {object} errno.errno{data=[]model.Movie}
// @Router       /movie/search [get]
// func (h *folderHandle) SearchMovies(ctx *fiber.Ctx) error {
// 	data, err := h.svc.SearchMoviesByCode(ctx.Query("code"))
// 	if err != nil {
// 		return ctx.JSON(errno.DecodeError(err))
// 	}
// 	return ctx.JSON(errno.OK.WithData(data))
// }

// @Summary      Upload movie cover
// @Description  Upload movie cover by movie id
// @Tags         movie
// @Produce      json
// @Param        code query string true "movie code"
// @Success      200 {object} errno.errno{data=[]model.Movie}
// @Router       /movie/cover [put]
// func (h *folderHandle) UploadCover(ctx *fiber.Ctx) error {
// 	code := ctx.Params("code", "")

// 	form, err := ctx.MultipartForm()
// 	if err != nil {
// 		return ctx.JSON(errno.DecodeError(err))
// 	}

// 	files := form.File["cover"]
// 	if files == nil || len(files) <= 0 || len(files) > 1 {
// 		return ctx.JSON(errno.DecodeError(errors.New("file is not exist or too many")))
// 	}

// 	file, err := files[0].Open()
// 	if err != nil {
// 		return ctx.JSON(errno.DecodeError(err))
// 	}
// 	defer file.Close()

// 	fileBytes, err := io.ReadAll(file)
// 	if err != nil {
// 		return ctx.JSON(errno.DecodeError(err))
// 	}

// 	err = h.svc.UploadMovieCover(code, files[0].Filename, fileBytes)
// 	return ctx.JSON(errno.DecodeError(err))
// }
