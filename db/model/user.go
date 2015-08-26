package model

import (
	"time"

	"github.com/pborman/uuid"
)

type User struct {
	Uid          uuid.UUID // UUID
	Email        string
	Firstname    string
	Lastname     string
	Gender       string // TODO(remy): GenderType enum
	Phone        string
	Address      string
	CreationTime time.Time
	LastUpdate   time.Time
}
