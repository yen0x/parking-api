package service

import (
	"time"

	"bitbucket.org/remeh/parking/db/model"
	"bitbucket.org/remeh/parking/runtime"

	"github.com/pborman/uuid"
)

// TODO(remy): comment
func CreateUser(rt *runtime.Runtime, email, firstname, lastname string) (uuid.UUID, error) {
	if rt == nil {
		return []byte{}, nil
	}

	uDAO := rt.Storage.UserDAO
	now := time.Now()

	uid := uuid.Parse(uuid.New())
	user := model.User{
		Uid:          uid,
		Email:        email,
		Firstname:    firstname,
		Lastname:     lastname,
		CreationTime: now,
		LastUpdate:   now,
	}
	_, err := uDAO.Insert(user)

	return uid, err
}

// UserExists returns whether or not an user already uses
// the given email.
func UserExists(rt *runtime.Runtime, email string) (bool, error) {
	uDAO := rt.Storage.UserDAO
	user, err := uDAO.FindByEmail(email)
	return len(user.Uid) > 0, err
}
