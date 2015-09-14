// Parking Backend - Model
//
// Parking model.
//
// 2015

package model

import (
	"time"

	"github.com/pborman/uuid"
)

type Parking struct {
	Uid uuid.UUID
	// User ID of the owner.
	UserId uuid.UUID

	Description string
	Address     string
	Zip         string
	City        string

	// FIXME(remy): are float64 enough ?
	Latitude  float64
	Longitude float64

	DailyPrice int
	Currency   string

	CreationTime time.Time
	LastUpdate   time.Time
}
