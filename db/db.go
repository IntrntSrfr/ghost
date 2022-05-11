package db

import "github.com/intrntsrfr/ghost/model"

type DB interface {
	CreateUpload(u *model.Upload) (*model.Upload, error)
	GetUpload() *model.Upload
	GetUploadByHash(hash string) *model.Upload
	GetUserUploads() []*model.Upload
}
