// Parking Backend
//
// DB storage.
//
// 2015

package db

import (
	"database/sql"

	. "bitbucket.org/parking/db/dao"

	_ "github.com/lib/pq"
)

type Storage struct {
	Conn *sql.DB

	UserDAO *UserDAO
}

// Init opens a PostgreSQL connection with the given connectionString.
func (s *Storage) Init(connectionString string) (*sql.DB, error) {
	dbase, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	s.Conn = dbase

	// Creates all the DAOs of this storage.
	s.createDAOs()

	return dbase, s.Conn.Ping()
}

func (s *Storage) createDAOs() {
	s.UserDAO = NewUserDAO(s.Conn)
}
