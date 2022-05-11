package db

import "github.com/intrntsrfr/ghost/model"

type DB interface {
	CreateUpload() *model.Upload
	GetUpload() *model.Upload
	GetUserUploads() []*model.Upload
}
