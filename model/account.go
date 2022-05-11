package model

type Account struct {
	ID  int    `db:"id" json:"id"`
	Key string `db:"key_hash" json:"-"`
}
