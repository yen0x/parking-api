// Booking Backend - DAO
//
// Booking DAO.
//
// 2015

package dao

import (
	"bitbucket.org/remeh/parking/db/model"
	. "database/sql"

	//"github.com/pborman/uuid"
)

type BookingDAO struct {
	db *DB

	insert *Stmt
}

const (
	BOOKING_FIELDS = `"booking"."uid",
					"booking"."user_id",
					"booking"."parking_id",
					"booking"."start",
					"booking"."end",
					"booking"."count"`
)

func NewBookingDAO(db *DB) (*BookingDAO, error) {
	dao := &BookingDAO{
		db: db,
	}
	err := dao.initStmt()
	return dao, err
}

func (d *BookingDAO) initStmt() error {
	var err error

	if d.insert, err = d.db.Prepare(`
	   INSERT INTO "booking"
	   (` + insertFields("booking", BOOKING_FIELDS) + `)
	   VALUES ($1, $2, $3, $4, $5, $6);
	   `); err != nil {
		return err
	}
	return nil
}

func (d *BookingDAO) Insert(booking model.Booking) (Result, error) {
	if len(booking.Uid) == 0 {
		return nil, nil
	}

	return d.insert.Exec(
		booking.Uid.String(),
		booking.UserId.String(),
		booking.ParkingId.String(),
		booking.Start,
		booking.End,
		booking.Count,
	)
}
