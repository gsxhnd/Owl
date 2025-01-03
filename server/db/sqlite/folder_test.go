package sqlite

import (
	"testing"

	"github.com/gsxhnd/owl/server/model"
	"github.com/stretchr/testify/assert"
)

func Test_sqliteDB_CreateMovies(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, err := getMockDB()
			assert.Nil(t, err)
			list := make([]model.Folder, 0)
			data := model.Folder{}
			list = append(list, data)

			db.CreateFolders(list)
		})
	}
}

func Test_sqliteDB_DeleteMovies(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, err := getMockDB()
			assert.Nil(t, err)

			db.DeleteFolders([]uint{3, 4, 5})
		})
	}
}

func Test_sqliteDB_GetMovies(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, err := getMockDB()
			assert.Nil(t, err)

			db.GetFolders(nil)
		})
	}
}
