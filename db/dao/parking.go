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
	findInArea *Stmt
	findByUid  *Stmt
}

const (
	PARKING_FIELDS = `"parking"."uid",
				      "parking"."user_id",
				      "parking"."description",
				      "parking"."address",
				      "parking"."zip",
				      "parking"."city",
				      "parking"."latitude",
				      "parking"."longitude",
				      "parking"."daily_price",
					  "parking"."currency",
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
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);
	`); err != nil {
		return err
	}

	if d.findInArea, err = d.db.Prepare(`
		SELECT ` + PARKING_FIELDS + `
		FROM "parking"
		WHERE
			"latitude" <= $1 AND "latitude" >= $3
			AND
			"longitude" >= $4 AND "longitude" <= $2
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

	if d.findByUid, err = d.db.Prepare(`
		SELECT ` + PARKING_FIELDS + `
		FROM "parking"
		WHERE uid = $1
	`); err != nil {
		return err
	}

	return nil
}

func (d *ParkingDAO) FindInArea(topLeftLat, topLeftLon, bottomRightLat, bottomRightLon float64) ([]model.Parking, error) {
	return readParkings(d.findInArea.Query(topLeftLat, topLeftLon, bottomRightLat, bottomRightLon))
}

func (d *ParkingDAO) FindByUser(user model.User) ([]model.Parking, error) {
	return d.FindByUserId(user.Uid)
}
func (d *ParkingDAO) FindByUid(uid uuid.UUID) (model.Parking, error) {
	found := model.Parking{}
	rows, err := d.findByUid.Query(uid.String())
	if rows == nil || err != nil {
		return found, err
	}

	defer rows.Close()
	if !rows.Next() {
		return found, err
	}
	return parkingFromRow(rows)
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
		parking.Zip,
		parking.City,
		parking.Latitude,
		parking.Longitude,
		parking.DailyPrice,
		parking.Currency,
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
		zip,
		city,
		currency string
	var latitude,
		longitude float64
	var dailyPrice int
	var creationTime,
		lastUpdate time.Time

	err := rows.Scan(
		&uid,
		&userId,
		&description,
		&address,
		&zip,
		&city,
		&latitude,
		&longitude,
		&dailyPrice,
		&currency,
		&creationTime,
		&lastUpdate,
	)

	return model.Parking{
		Uid:          uuid.Parse(uid),
		UserId:       uuid.Parse(userId),
		Description:  description,
		Address:      address,
		Zip:          zip,
		City:         city,
		Latitude:     latitude,
		Longitude:    longitude,
		DailyPrice:   dailyPrice,
		Currency:     currency,
		CreationTime: creationTime,
		LastUpdate:   lastUpdate,
	}, err
}
