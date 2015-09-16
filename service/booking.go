// Booking Backend - Service
//
// All methods around booking features
//
// 2015

package service

import (
	"time"

	"bitbucket.org/remeh/parking/db/model"
	"bitbucket.org/remeh/parking/runtime"

	"github.com/pborman/uuid"
)

const (
	DATE_FORMAT = "02/01/2006"
)

func CreateBooking(rt *runtime.Runtime, user model.User, startString, endString, parking string, count int) (uuid.UUID, error) {
	if rt == nil {
		return []byte{}, nil
	}

	bDAO := rt.Storage.BookingDAO
	uid := uuid.Parse(uuid.New())
	start, err := time.Parse(DATE_FORMAT, startString)
	if err != nil {
		return []byte{}, err
	}
	end, err := time.Parse(DATE_FORMAT, endString)
	if err != nil {
		return []byte{}, err
	}

	//TODO (jean) check parkingexists
	booking := model.Booking{
		Uid:       uid,
		UserId:    user.Uid,
		ParkingId: user.Uid,
		Start:     start,
		End:       end,
		Count:     count,
	}

	_, err = bDAO.Insert(booking)
	return uid, err
}