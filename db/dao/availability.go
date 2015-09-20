//
// Availability DAO.
//
// 2015

package dao

import (
	"bitbucket.org/remeh/parking/db/model"
	. "database/sql"
)

type AvailabilityDAO struct {
	db *DB

	insert *Stmt
}

const (
	AVAILABILITY_FIELDS = `"availability"."parking_uid",
					"availability"."start",
					"availability"."end"`
)

func NewAvailabilityDAO(db *DB) (*AvailabilityDAO, error) {
	dao := &AvailabilityDAO{
		db: db,
	}
	err := dao.initStmt()
	return dao, err
}

func (d *AvailabilityDAO) initStmt() error {
	var err error

	if d.insert, err = d.db.Prepare(`
	   INSERT INTO "availability"
	   (` + insertFields("availability", AVAILABILITY_FIELDS) + `)
	   VALUES ($1, $2, $3);
	   `); err != nil {
		return err
	}
	return nil
}

func (d *AvailabilityDAO) Insert(availability model.Availability) (Result, error) {
	return d.insert.Exec(
		availability.ParkingUid.String(),
		availability.Start,
		availability.End)
}
