package db

import (
	"database/sql"
	"github.com/intrntsrfr/ghost/model"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	_ "github.com/lib/pq"
)

type PsqlDB struct {
	*sqlx.DB
	logger *zap.Logger
}

func NewPsqlDB(db *sqlx.DB, logger *zap.Logger) *PsqlDB {
	return &PsqlDB{db, logger}
}

func (p *PsqlDB) CreateUpload(u *model.Upload) (*model.Upload, error) {
	err := p.QueryRow("INSERT INTO upload(hash, extension, created, account_id) VALUES($1, $2, $3, $4) RETURNING id", u.Hash, u.Extension, u.Created, u.AccountID).Scan(&u.ID)
	if err != nil {
		p.logger.Error("could not create upload", zap.Error(err))
		return nil, err
	}
	return u, nil
}

func (p *PsqlDB) GetUploadByHash(hash string) *model.Upload {
	u := &model.Upload{}
	err := p.Get(u, "SELECT * FROM upload WHERE hash=$1", hash)
	if err != nil {
		if err != sql.ErrNoRows {
			p.logger.Error("could not fetch upload", zap.Error(err))
		}
		return nil
	}
	return u
}

func (p *PsqlDB) GetUserUploads(accID int) []*model.Upload {
	var uploads []*model.Upload
	err := p.Select(&uploads, "SELECT * FROM upload WHERE account_id=$1", accID)
	if err != nil && err != sql.ErrNoRows {
		p.logger.Error("could not fetch user uploads", zap.Error(err))
		return nil
	}
	return uploads
}

func (p *PsqlDB) GetUserByHash(hash string) *model.Account {
	a := &model.Account{}
	err := p.Get(a, "SELECT * FROM account WHERE key_hash=$1", hash)
	if err != nil {
		if err != sql.ErrNoRows {
			p.logger.Error("could not fetch user", zap.Error(err))
		}
		return nil
	}
	return a
}
