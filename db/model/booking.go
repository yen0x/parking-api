// Booking Backend - Model
//
// Booking model.
//
// 2015

package model

import (
	"time"

	"github.com/pborman/uuid"
)

type Booking struct {
	Uid       uuid.UUID
	UserId    uuid.UUID
	ParkingId uuid.UUID
	Start     time.Time
	End       time.Time
	Count     int
}
