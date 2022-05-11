package model

import (
	"time"
)

type Upload struct {
	ID        int       `db:"id" json:"id"`
	Hash      string    `db:"hash" json:"hash"`
	Extension string    `db:"extension" json:"extension"`
	Created   time.Time `db:"created" json:"created"`
	AccountID int       `db:"account_id" json:"account_id"`
}

func (u Upload) String() string {
	return u.Hash + u.Extension
}
