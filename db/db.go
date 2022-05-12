package db

import "github.com/intrntsrfr/ghost/model"

type DB interface {
	CreateUpload(u *model.Upload) (*model.Upload, error)
	GetUploadByHash(hash string) *model.Upload
	GetUserUploads(accID int) []*model.Upload
	GetUserByHash(hash string) *model.Account
}
