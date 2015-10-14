// Booking Backend - Service
//
// All methods around booking features
//
// 2015

package service

import (
	"bitbucket.org/remeh/parking/db/model"
	"bitbucket.org/remeh/parking/runtime"
	"fmt"
	"time"

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
	//fetching parking owner data
	owner, err := GetUserByParking(rt, parking.Uid)
	if err != nil {
		return []byte{}, err
	}

	_, err = bDAO.Insert(booking)

	if err == nil {
		mg := mailgun.NewMailgun("sandbox162f60c4aa7a4219b003fe611b60f376.mailgun.org", "key-9e42dce47a4a90a6b6afd88b2be5495e", "")
		m := mg.NewMessage(
			"Parking Particuliers <contact@parking-particuliers.com>",     // From
			"Vous avez une nouvelle réservation",                          // Subject
			"Une nouvelle réservation a été effectué pour votre parking.", // Plain-text body
			owner.ToMailString(),                                          // Recipients (vararg list)
		)
		m2 := mg.NewMessage(
			"Parking Particuliers <contact@parking-particuliers.com>", // From
			"Confirmaiton de réservation",                             // Subject
			"Nous vous confirmons la réservation faites à DATA",       // Plain-text body
			user.ToMailString(),                                       // Recipients (vararg list)
		)

		_, _, err := mg.Send(m)
		_, _, err2 := mg.Send(m2)

		if err != nil {
			return []byte{}, err
		}
		if err2 != nil {
			return []byte{}, err2
		}
		fmt.Println("Emails sent ")
	}
	return uid, err
}

func GetBookings(rt *runtime.Runtime, user model.User) ([]model.Booking, error) {
	bDAO := rt.Storage.BookingDAO
	bookings, err := bDAO.FindByUserId(user.Uid)
	return bookings, err
}
