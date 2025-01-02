package storage

import (
	"bytes"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"os"
	"path"

	"github.com/gsxhnd/owl/utils"
)

type localStorage struct {
	path   string
	logger utils.Logger
}

func NewLocalStorage(cfg *utils.Config, l utils.Logger) (Storage, error) {
	var sPath = path.Join(cfg.DataPath, cfg.Storage.Path)

	// if err := utils.MakeDir(path.Join(sPath, "actor")); err != nil {
	// 	return nil, err
	// }

	// if err := utils.MakeDir(path.Join(sPath, "movie")); err != nil {
	// 	return nil, err
	// }

	for i := 0; i <= 255; i++ {
		hex := fmt.Sprintf("%02x", i)
		if err := os.MkdirAll(path.Join(sPath, "movie", hex), os.ModePerm); err != nil {
			fmt.Println(err)
		}
		if err := os.MkdirAll(path.Join(sPath, "actor", hex), os.ModePerm); err != nil {
			fmt.Println(err)
		}
	}

	return &localStorage{
		path: sPath,
	}, nil
}

func (s *localStorage) Ping() error {
	return nil
}

func (s *localStorage) GetImage(cover string, id uint, filename string) ([]byte, string, error) {
	var filepath = path.Join(s.path, cover, "1.jpeg")
	file, err := os.Open(filepath)
	if err != nil {
		return nil, "", err
	}
	defer file.Close()

	var buf bytes.Buffer
	var tee = io.TeeReader(file, &buf)

	_, f, err := image.Decode(tee)
	if err != nil {
		return nil, "", err
	}

	buff, _ := io.ReadAll(&buf)

	return buff, f, nil
}

func (s *localStorage) SaveImage(data []byte, cover string, id uint, filename string) error {
	hex := fmt.Sprintf("%02x", id&0xff)

	if err := os.WriteFile(path.Join(s.path, cover, hex, filename), data, 0644); err != nil {
		s.logger.Errorf("Local save image error: %s", err.Error())
		return err
	}
	return nil
}
