// Parking Backend - DAO
//
// Parking DAO.
//
// 2015

package dao

import (
	. "database/sql"
	"time"

	"bitbucket.org/remeh/parking/db/model"

	"github.com/pborman/uuid"
)

type ParkingDAO struct {
	db *DB

	insert     *Stmt
	findByUser *Stmt
}

const (
	PARKING_FIELDS = `"parking"."uid",
				      "parking"."user_id",
				      "parking"."description",
				      "parking"."address",
				      "parking"."latitude",
				      "parking"."longitude",
				      "parking"."daily_price",
				      "parking"."creation_time",
				      "parking"."last_update"`
)

func NewParkingDAO(db *DB) (*ParkingDAO, error) {
	dao := &ParkingDAO{
		db: db,
	}
	err := dao.initStmt()
	return dao, err
}

func (d *ParkingDAO) initStmt() error {
	var err error

	if d.insert, err = d.db.Prepare(`
		INSERT INTO "parking"
		(` + insertFields("parking", PARKING_FIELDS) + `)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);
	`); err != nil {
		return err
	}

	if d.findByUser, err = d.db.Prepare(`
		SELECT ` + PARKING_FIELDS + `
		FROM "parking"
		WHERE user_id = $1
	`); err != nil {
		return err
	}

	return nil
}

func (d *ParkingDAO) FindByUser(user model.User) ([]model.Parking, error) {
	return d.FindByUserId(user.Uid)
}
func (d *ParkingDAO) FindByUserId(uid uuid.UUID) ([]model.Parking, error) {
	return readParkings(d.findByUser.Query(uid.String()))
}

func (d *ParkingDAO) Insert(parking model.Parking) (Result, error) {
	if len(parking.Uid) == 0 {
		return nil, nil
	}

	return d.insert.Exec(
		parking.Uid.String(),
		parking.UserId.String(),
		parking.Description,
		parking.Address,
		parking.Latitude,
		parking.Longitude,
		parking.DailyPrice,
		parking.CreationTime,
		parking.LastUpdate,
	)
}

// readParkings fully reads (and closes) the given rows to return
// the read parkings or an error if something wrong occurred.
func readParkings(rows *Rows, err error) ([]model.Parking, error) {
	result := make([]model.Parking, 0)

	if err != nil || rows == nil {
		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		parking, err := parkingFromRow(rows)
		if err != nil {
			return result, err
		}
		result = append(result, parking)
	}

	return result, err
}

// parkingFromRow reads an parking model from the current row.
func parkingFromRow(rows *Rows) (model.Parking, error) {
	var uid,
		userId,
		description,
		address,
		dailyPrice string
	var latitude,
		longitude float64
	var creationTime,
		lastUpdate time.Time

	err := rows.Scan(
		&uid,
		&userId,
		&description,
		&address,
		&latitude,
		&longitude,
		&dailyPrice,
		&creationTime,
		&lastUpdate,
	)

	return model.Parking{
		Uid:          uuid.Parse(uid),
		UserId:       uuid.Parse(userId),
		Description:  description,
		Address:      address,
		Latitude:     latitude,
		Longitude:    longitude,
		DailyPrice:   dailyPrice,
		CreationTime: creationTime,
		LastUpdate:   lastUpdate,
	}, err
}
