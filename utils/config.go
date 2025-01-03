package utils

import (
	"path"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	Mode           string `env:"MODE" envDefault:"dev"`
	Listen         string `env:"LISTEN" envDefault:":8080"`
	DataPath       string `env:"DATA_PATH" envDefault:"./data/.owl"`
	Log            LogConfig
	Storage        StorageConfig
	DatabaseConfig DatabaseConfig
}

type LogConfig struct {
	Path       string `env:"LOG_PATH" envDefault:"./log"`
	FileName   string `env:"LOG_FILE_NAME" envDefault:"./owl.log"`
	Level      string `env:"LOG_LEVEL" envDefault:"debug"`
	MaxBackups int    `env:"LOG_MAX_BACKUPS" envDefault:"10"`
	MaxAge     int    `env:"LOG_MAX_AGE" envDefault:"7"`
}

type StorageConfig struct {
	Type       string `env:"STORAGE_TYPE" envDefault:"local"`
	Path       string `env:"STORAGE_PATH" envDefault:"./data/cover"`
	Endpoint   string `env:"STORAGE_ENDPOINT" envDefault:"localhost:9000"`
	BucketName string `env:"STORAGE_BUCKET_NAME" envDefault:"owl"`
	AccessKey  string `env:"STORAGE_ACCESS_KEY"`
	SecretKey  string `env:"STORAGE_SECRET_KEY"`
}

type DatabaseConfig struct {
	Type string `env:"DB_TYPE" envDefault:"sqlite3"`
	Path string `env:"DB_PATH" envDefault:"owl.db"`
}

func NewConfig() (*Config, error) {
	var c Config

	opts := env.Options{
		Prefix: "JAHA_",
	}

	if err := env.ParseWithOptions(&c, opts); err != nil {
		return nil, err
	}

	if err := MakeDir(path.Join(c.DataPath)); err != nil {
		return nil, err
	}

	if err := MakeDir(path.Join(c.DataPath, c.Log.Path)); err != nil {
		return nil, err
	}

	return &c, nil
}
