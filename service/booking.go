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

	"github.com/mailgun/mailgun-go"
	"github.com/pborman/uuid"
)

const (
	DATE_FORMAT = "2006-01-02"
)

func CreateBooking(rt *runtime.Runtime, user model.User, start, end time.Time, parking model.Parking, count int) (uuid.UUID, error) {
	if rt == nil {
		return []byte{}, nil
	}

	bDAO := rt.Storage.BookingDAO
	uid := uuid.Parse(uuid.New())

	booking := model.Booking{
		Uid:       uid,
		UserId:    user.Uid,
		ParkingId: parking.Uid,
		Start:     start,
		End:       end,
		Count:     count,
	}

	_, err := bDAO.Insert(booking)

	if err == nil {
		mg := mailgun.NewMailgun("sandbox162f60c4aa7a4219b003fe611b60f376.mailgun.org", "key-9e42dce47a4a90a6b6afd88b2be5495e", "")
		m := mg.NewMessage(
			"Dwight Schrute <gege.dert@gmail.com>",                        // From
			"Vous avez une nouvelle réservation",                          // Subject
			"Une nouvelle réservation a été effectué pour votre parking.", // Plain-text body
			"Michael Scott <gege.dert@gmail.com>",                         // Recipients (vararg list)
		)

		_, _, err := mg.Send(m)

		if err != nil {
			return []byte{}, err
		}
	}
	return uid, err
}
