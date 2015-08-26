// Parking Backend
//
// DB storage.
//
// 2015

package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage struct {
	Conn *sql.DB
}

// Init opens a PostgreSQL connection with the given connectionString.
func (s *Storage) Init(connectionString string) (*sql.DB, error) {
	dbase, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	s.Conn = dbase

	return dbase, s.Conn.Ping()
}
