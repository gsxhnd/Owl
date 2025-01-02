package db

import (
	"path"

	"github.com/google/wire"
	"github.com/gsxhnd/owl/server/db/database"
	"github.com/gsxhnd/owl/server/db/sqlite"
	"github.com/gsxhnd/owl/utils"
)

func NewDatabase(cfg *utils.Config, logger utils.Logger) (database.Driver, error) {
	var driver database.Driver
	var err error

	var dbPath = path.Join(cfg.DataPath, cfg.DatabaseConfig.Path)

	logger.Debugf("jaha using local database, path: %s", dbPath)
	driver, err = sqlite.NewSqliteDB(dbPath, logger)
	if err != nil {
		return nil, err
	}

	if err := driver.Migrate(); err != nil {
		return nil, err
	}
	return driver, nil
}

var DBSet = wire.NewSet(NewDatabase)
