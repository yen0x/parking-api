// Parking Backend - Model
//
// User model.
//
// 2015

package model

import (
	"time"

	"github.com/pborman/uuid"
)

type User struct {
	Uid          uuid.UUID
	Email        string
	Firstname    string
	Lastname     string
	Password     string
	Salt         string
	Gender       string // TODO(remy): GenderType enum
	Phone        string
	Address      string
	CreationTime time.Time
	LastUpdate   time.Time
}

func (u *User) ToMailString() string {
	return u.Firstname + " " + u.Lastname + "<" + u.Email + ">"
}
