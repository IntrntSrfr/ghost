package db

import (
	"github.com/intrntsrfr/ghost/model"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"time"

	_ "github.com/lib/pq"
)

type PsqlDB struct {
	db     *sqlx.DB
	logger *zap.Logger
}

func NewPsqlDB(db *sqlx.DB, logger *zap.Logger) *PsqlDB {
	return &PsqlDB{db, logger}
}

func (p *PsqlDB) CreateUpload(name, ext string, user int) (string, error) {
	_, err := p.db.Exec("INSERT INTO upload(id, account_id, created, filename) VALUES($1, $2, $3, $4)", name, user, time.Now(), name+ext)
	if err != nil {
		return "", err
	}
	return name + ext, nil
}

func (p *PsqlDB) GetUpload() *model.Upload {
	//TODO implement me
	panic("implement me")
}

func (p *PsqlDB) GetUserUploads() []*model.Upload {
	//TODO implement me
	panic("implement me")
}
