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
	DATE_FORMAT = "2006-01-02"
)

func CreateBooking(rt *runtime.Runtime, user model.User, start, end time.Time, parking uuid.UUID, count int) (uuid.UUID, error) {
	if rt == nil {
		return []byte{}, nil
	}

	bDAO := rt.Storage.BookingDAO
	uid := uuid.Parse(uuid.New())

	booking := model.Booking{
		Uid:       uid,
		UserId:    user.Uid,
		ParkingId: parking,
		Start:     start,
		End:       end,
		Count:     count,
	}

	_, err := bDAO.Insert(booking)
	return uid, err
}
