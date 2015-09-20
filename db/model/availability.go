// Booking Backend - Model
//
// Parking availability model.
//
// 2015

package model

import (
	"time"

	"github.com/pborman/uuid"
)

type Availability struct {
	ParkingUid uuid.UUID
	Start      time.Time
	End        time.Time
}
