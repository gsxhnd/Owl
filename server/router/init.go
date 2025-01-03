package router

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gsxhnd/owl/server/handler"
	"github.com/gsxhnd/owl/server/middleware"
	"github.com/gsxhnd/owl/utils"
	"github.com/gsxhnd/owl/web"
)

type Router interface {
	Run() error
}

type router struct {
	cfg    *utils.Config
	app    *fiber.App
	logger utils.Logger
	h      handler.Handler
	m      middleware.Middleware
}

// @title           Owl API
// @version         0.0.1
// @description     This is a sample server celler server.
// @license.name  MIT
// @license.url   https://opensource.org/license/mit
// @host      localhost:8080
// @BasePath  /api/v1
// @externalDocs.description  OpenAPI
func NewRouter(cfg *utils.Config, l utils.Logger, m middleware.Middleware, h handler.Handler) (Router, error) {
	app := fiber.New(fiber.Config{
		EnablePrintRoutes:     cfg.Mode == "dev",
		DisableStartupMessage: cfg.Mode == "prod",
		Prefork:               false,
	})

	return &router{
		cfg:    cfg,
		app:    app,
		logger: l,
		h:      h,
		m:      m,
	}, nil
}

func (r *router) Run() error {
	// r.app.Use(r.m.RequestLog)
	r.app.Get("/ping", r.h.PingHandler.Ping)

	api := r.app.Group("/api/v1")
	// folder api
	api.Post("/folder", r.h.FolderHandler.CreateFolder)
	api.Delete("/folder", r.h.FolderHandler.DeleteFolder)
	api.Put("/folder", r.h.FolderHandler.UpdateFolder)
	api.Get("/folder", r.h.FolderHandler.GetFolders)
	// file api
	api.Delete("/file", r.h.FileHandler.DeleteFiles)
	api.Get("/file", r.h.FileHandler.GetFiles)
	// tag
	api.Post("/tag", r.h.TagHandler.CreateTag)
	api.Delete("/tag", r.h.TagHandler.DeleteTag)
	api.Put("/tag", r.h.TagHandler.UpdateTag)
	api.Get("/tag", r.h.TagHandler.GetTags)
	api.Get("/tag/search", r.h.TagHandler.SearchTags)
	// file tag
	api.Post("/movie_tag", r.h.FileTagHandler.CreateFileTags)
	api.Delete("/movie_tag", r.h.FileTagHandler.DeleteFileTags)

	img := r.app.Group("/api/v1/img")
	img.Get("/movie/:id", r.h.ImageHandler.GetMovieImage)
	img.Get("/actor/:id", r.h.ImageHandler.GetActorImage)

	r.app.Use("/*", filesystem.New(filesystem.Config{
		Root:       http.FS(web.Content),
		PathPrefix: "dist",
		Browse:     true,
	}))

	r.app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	r.logger.Infof("Server actort listening")

	return r.app.Listen(r.cfg.Listen)
}
