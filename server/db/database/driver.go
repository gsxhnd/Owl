package database

import "github.com/gsxhnd/owl/server/model"

type Driver interface {
	Ping() error
	Migrate() error
	CreateFolders([]model.Folder) error
	DeleteFolders([]uint) error
	UpdateFolder(*model.Folder) error
	GetFolders(*Pagination, ...string) ([]model.Folder, error)
	// CreateFile() error
	// DeleteFiles() error
	// CreateActors([]model.Actor) error
	// DeleteActors([]uint) error
	// UpdateActor(actor *model.Actor) error
	// GetActors() ([]model.Actor, error)
	// SearchActorByName(string) ([]model.Actor, error)
	// CreateTags([]model.Tag) error
	// DeleteTags([]uint) error
	// UpdateTag(tag *model.Tag) error
	// GetTags() ([]model.Tag, error)
	// SearchTagsByName(name string) ([]model.Tag, error)
	// CreateMovieActors(movieActors []model.MovieActor) error
	// DeleteMovieActors(ids []uint) error
	// UpdateMovieActor(model.MovieActor) error
	// GetMovieActors() ([]model.MovieActor, error)
	// GetMovieActorsByMovieId(id uint) ([]model.MovieActor, error)
	// CreateMovieTags([]model.FileTag) error
	// DeleteMovieTags(ids []uint) error
	// GetMovieTagByMovieId(movieId uint) ([]model.FileTag, error)
	// UpdateMovieTag(model.FileTag) error
	// CreateAnimes([]model.File) error
	// DeleteAnimes([]uint) error
	// UpdateAnime(model.File) error
	// GetAnimes(*Pagination) ([]model.File, error)
}
