// Parking Backend
//
// DB storage.
//
// 2015

package db

import (
	"database/sql"

	. "bitbucket.org/remeh/parking/db/dao"

	_ "github.com/lib/pq"
)

type Storage struct {
	Conn *sql.DB

	UserDAO    *UserDAO
	ParkingDAO *ParkingDAO
}

// Init opens a PostgreSQL connection with the given connectionString.
func (s *Storage) Init(connectionString string) (*sql.DB, error) {
	dbase, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	s.Conn = dbase

	// Creates all the DAOs of this storage.
	err = s.createDAOs()
	if err != nil {
		return nil, err
	}

	return dbase, s.Conn.Ping()
}

func (s *Storage) createDAOs() error {
	var err error
	if s.UserDAO, err = NewUserDAO(s.Conn); err != nil {
		return err
	}
	if s.ParkingDAO, err = NewParkingDAO(s.Conn); err != nil {
		return err
	}
	return nil
}
