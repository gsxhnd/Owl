package service

import "github.com/google/wire"

var ServiceSet = wire.NewSet(
	NewPingService,
	NewFolderService,
	NewFileService,
	NewTagService,
	NewFileTagService,
	NewLabelService,
)
