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
	dataPath      string
	thumbnailPath string
	logger        utils.Logger
}

func NewLocalStorage(cfg *utils.Config, l utils.Logger) (Storage, error) {
	dataPath := path.Join(cfg.DataPath)
	thumbnailPath := path.Join(cfg.DataPath, ".owl")

	for i := 0; i <= 255; i++ {
		hex := fmt.Sprintf("%02x", i)
		if err := os.MkdirAll(path.Join(thumbnailPath, "thumbnail", hex), os.ModePerm); err != nil {
			l.Errorw("mkdir thumbnail path failed", "error", err)
			return nil, err
		}
	}

	return &localStorage{
		dataPath:      dataPath,
		thumbnailPath: thumbnailPath,
	}, nil
}

func (s *localStorage) Ping() error {
	return nil
}

func (s *localStorage) GetImage(cover string, id uint, filename string) ([]byte, string, error) {
	filepath := path.Join(s.dataPath, cover, "1.jpeg")
	file, err := os.Open(filepath)
	if err != nil {
		return nil, "", err
	}
	defer file.Close()

	var buf bytes.Buffer
	tee := io.TeeReader(file, &buf)

	_, f, err := image.Decode(tee)
	if err != nil {
		return nil, "", err
	}

	buff, _ := io.ReadAll(&buf)

	return buff, f, nil
}

func (s *localStorage) SaveImage(data []byte, cover string, id uint, filename string) error {
	hex := fmt.Sprintf("%02x", id&0xff)

	if err := os.WriteFile(path.Join(s.dataPath, cover, hex, filename), data, 0644); err != nil {
		s.logger.Errorf("Local save image error: %s", err.Error())
		return err
	}
	return nil
}
