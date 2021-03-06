// Parking Backend - Model
//
// Photo model.
//
// 2015

package model

import (
	"github.com/pborman/uuid"
)

type Photo struct {
	Uid       uuid.UUID
	ParkingId uuid.UUID
	Mimetype  string
	Data      []byte
}
