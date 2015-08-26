// Parking Backend - Model
//
// Review model.
//
// 2015

package model

import (
	"time"

	"github.com/pborman/uuid"
)

type Review struct {
	Uid          uuid.UUID
	CreationTime time.Time
	// Author of the review.
	AuthorId uuid.UUID
	// User criticized
	UserId uuid.UUID
	// [0;5]
	Note int
	// Actual review text
	Text string
}
