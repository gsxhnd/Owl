package handler

import (
	"github.com/google/wire"
)

type Handler struct {
	PingHandler    PingHandler
	FolderHandler  FolderHandler
	FileHandler    FileHandler
	TagHandler     TagHandler
	FileTagHandler FileTagHandler
	LabelHandler   LabelHandler
	ImageHandler   ImageHandler
}

var HandlerSet = wire.NewSet(
	NewPingHandler,
	NewFolderHandler,
	NewFileHandler,
	NewTagHandler,
	NewFileTagHandler,
	NewLabelHandler,
	NewImageHandler,
	wire.Struct(new(Handler), "*"),
)
