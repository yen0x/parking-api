package model

import (
	"github.com/pborman/uuid"
)

type User struct {
	Uid       uuid.UUID // UUID
	Email     string
	Firstname string
	Lastname  string
	Gender    string
	Phone     string
	Address   string
}
