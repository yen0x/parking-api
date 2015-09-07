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

	// FIXME(remy): maybe we should create a special type for money ?
	DailyPrice string

	CreationTime time.Time
	LastUpdate   time.Time
}
