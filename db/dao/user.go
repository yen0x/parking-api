package dao

import (
	"database/sql"
)

type UserDAO struct {
	db *sql.DB
}

func NewUserDAO(db *sql.DB) (*UserDAO, error) {
	dao := &UserDAO{
		db: db,
	}
	err := dao.initStmt()
	return dao, err
}

func (d *UserDAO) initStmt() error {
	return nil
}
