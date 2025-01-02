//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package di

import (
	"github.com/google/wire"
	"github.com/gsxhnd/owl/server/db"
	"github.com/gsxhnd/owl/server/handler"
	"github.com/gsxhnd/owl/server/middleware"
	"github.com/gsxhnd/owl/server/router"
	"github.com/gsxhnd/owl/server/service"
	"github.com/gsxhnd/owl/server/storage"
	"github.com/gsxhnd/owl/utils"
)

func InitApp() (*Application, error) {
	wire.Build(
		utils.UtilsSet,
		NewApplication,
		router.NewRouter,
		middleware.NewMiddleware,
		handler.HandlerSet,
		service.ServiceSet,
		storage.StorageSet,
		db.DBSet,
	)
	return &Application{}, nil
}
