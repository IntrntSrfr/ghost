package model

import "time"

type Upload struct {
	ID        int       `db:"id" json:"id"`
	AccountID int       `db:"account_id" json:"account_id"`
	Created   time.Time `db:"created" json:"created"`
	Filename  string    `db:"filename" json:"filename"`
}
