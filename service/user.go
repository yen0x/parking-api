// Parking Backend - Service
//
// All methods around users features.
//
// 2015

package service

import (
	"time"

	"bitbucket.org/remeh/parking/db/model"
	"bitbucket.org/remeh/parking/runtime"

	"github.com/pborman/uuid"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser tries to store a new user into the database. It creates the
// uid of the users and returns it.
func CreateUser(rt *runtime.Runtime, email, firstname, password string) (uuid.UUID, error) {
	if rt == nil {
		return []byte{}, nil
	}

	uDAO := rt.Storage.UserDAO
	now := time.Now()

	uid := uuid.Parse(uuid.New())

	// Crypts and creates salt.
	cryptedPassword, err := cryptPassword(password)
	if err != nil {
		return uid, err
	}

	_, err = uDAO.Create(uid.String(), email, firstname, cryptedPassword, now)

	return uid, err
}

// TODO(remy): comment.
func CheckUserPassword(rt *runtime.Runtime, email, password string) (bool, model.User, error) {
	user, err := GetUser(rt, email)
	if err != nil {
		return false, user, err
	}

	crypted, err := rt.Storage.UserDAO.GetCryptedPassword(user)
	if err != nil {
		return false, user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(crypted), []byte(password))
	return err == nil, user, nil
}

// UserExists returns whether or not an user already uses
// the given email.
func UserExists(rt *runtime.Runtime, email string) (bool, error) {
	uDAO := rt.Storage.UserDAO
	user, err := uDAO.FindByEmail(email)
	return len(user.Uid) > 0, err
}

// GetUser returns an user from the database by an email.
func GetUser(rt *runtime.Runtime, email string) (model.User, error) {
	uDAO := rt.Storage.UserDAO
	return uDAO.FindByEmail(email)
}

func cryptPassword(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(b), err
}

func GetUserByParking(rt *runtime.Runtime, parkingId uuid.UUID) (model.User, error) {
	uDAO := rt.Storage.UserDAO
	user, err := uDAO.FindByParking(parkingId)
	return user, err
}
