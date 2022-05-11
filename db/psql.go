package db

import (
	"github.com/intrntsrfr/ghost/model"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	_ "github.com/lib/pq"
)

type PsqlDB struct {
	db     *sqlx.DB
	logger *zap.Logger
}

func NewPsqlDB(db *sqlx.DB, logger *zap.Logger) *PsqlDB {
	return &PsqlDB{db, logger}
}

func (p *PsqlDB) CreateUpload(u *model.Upload) (*model.Upload, error) {
	err := p.db.QueryRow("INSERT INTO upload(hash, extension, created, account_id) VALUES($1, $2, $3, $4) RETURNING id", u.Hash, u.Extension, u.Created, u.AccountID).Scan(&u.ID)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (p *PsqlDB) GetUpload() *model.Upload {
	//TODO implement me
	panic("implement me")
}

func (p *PsqlDB) GetUploadByHash(hash string) *model.Upload {
	u := &model.Upload{}
	err := p.db.Get(u, "SELECT * FROM upload WHERE hash=$1", hash)
	if err != nil {
		return nil
	}
	return u
}

func (p *PsqlDB) GetUserUploads() []*model.Upload {
	//TODO implement me
	panic("implement me")
}
