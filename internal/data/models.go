package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Todo  TodoModel
	Users UserModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Todo:  TodoModel{DB: db},
		Users: UserModel{DB: db},
	}
}
