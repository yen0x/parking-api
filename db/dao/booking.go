// Booking Backend - DAO
//
// Booking DAO.
//
// 2015

package dao

import (
	"bitbucket.org/remeh/parking/db/model"
	. "database/sql"
	"time"

	"github.com/pborman/uuid"
)

type BookingDAO struct {
	db *DB

	insert          *Stmt
	findByUserId    *Stmt
	findByParkingId *Stmt
	deleteByUid     *Stmt
	findByUid       *Stmt
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

	if d.findByUserId, err = d.db.Prepare(`
	   SELECT ` + BOOKING_FIELDS + ` from "booking"
	   where user_id = $1
	   `); err != nil {
		return err
	}

	if d.deleteByUid, err = d.db.Prepare(`
	   DELETE from "booking"
	   where uid = $1
	   `); err != nil {
		return err
	}

	if d.findByUid, err = d.db.Prepare(`
	   SELECT ` + BOOKING_FIELDS + ` from "booking"
	   where uid = $1
	   `); err != nil {
		return err
	}

	if d.findByParkingId, err = d.db.Prepare(`
	   SELECT ` + BOOKING_FIELDS + ` from "booking"
	   where parking_id = $1
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

func (d *BookingDAO) Delete(booking model.Booking) (Result, error) {
	return d.deleteByUid.Exec(booking.Uid.String())
}

func (d *BookingDAO) FindByUserId(uid uuid.UUID) ([]model.Booking, error) {
	return readBookings(d.findByUserId.Query(uid.String()))
}

func (d *BookingDAO) FindByParkingId(uid uuid.UUID) ([]model.Booking, error) {
	return readBookings(d.findByParkingId.Query(uid.String()))
}

func (d *BookingDAO) FindByUid(uid uuid.UUID) (model.Booking, error) {
	found := model.Booking{}
	rows, err := d.findByUid.Query(uid.String())
	if rows == nil || err != nil {
		return found, err
	}

	defer rows.Close()
	if !rows.Next() {
		return found, err
	}
	return bookingFromRow(rows)
}

// readBookings fully reads (and closes) the given rows to return
// the read bookings or an error if something wrong occurred.
func readBookings(rows *Rows, err error) ([]model.Booking, error) {
	result := make([]model.Booking, 0)

	if err != nil || rows == nil {
		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		booking, err := bookingFromRow(rows)
		if err != nil {
			return result, err
		}
		result = append(result, booking)
	}

	return result, err
}

// bookingFromRow reads a booking model from the current row.
func bookingFromRow(rows *Rows) (model.Booking, error) {
	var uid,
		userId,
		parkingId string
	var count int
	var start,
		end time.Time

	err := rows.Scan(
		&uid,
		&userId,
		&parkingId,
		&start,
		&end,
		&count,
	)

	return model.Booking{
		Uid:       uuid.Parse(uid),
		UserId:    uuid.Parse(userId),
		ParkingId: uuid.Parse(parkingId),
		Start:     start,
		End:       end,
		Count:     count,
	}, err
}
