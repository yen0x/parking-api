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

	// FIXME(remy): are float64 enough ?
	Latitude  float64
	Longitude float64

	CreationTime time.Time
	LastUpdate   time.Time
}
