package api

import (
	"bitbucket.org/remeh/parking/db/model"
	. "bitbucket.org/remeh/parking/logger"
	"bitbucket.org/remeh/parking/runtime"
	"bitbucket.org/remeh/parking/service"

	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/pborman/uuid"
	"net/http"
)

type CalParking struct {
	Runtime *runtime.Runtime
}

func (c CalParking) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	session, exists := c.Runtime.SessionStorage.GetFromRequest(r)
	if !exists {
		w.WriteHeader(403)
		return
	}

	vars := mux.Vars(r)
	parkingUid := vars["parkinguid"]

	parking, err := service.GetParking(c.Runtime, uuid.Parse(parkingUid))
	if !uuid.Equal(parking.UserId, session.User.Uid) {
		w.WriteHeader(403)
		return
	}

	bookings, err := service.GetBookingsByParking(c.Runtime, parking)
	if err != nil {
		Error(err)
		w.WriteHeader(500)
		return
	}

	bookingresps, err := c.serializeBookings(c.Runtime, bookings)
	if err != nil {
		Error(err)
		w.WriteHeader(500)
		return
	}

	data, err := json.Marshal(bookingresps)
	if err != nil {
		Error(err)
		w.WriteHeader(500)
		return
	}

	w.Write(data)
}

func (c CalParking) serializeBookings(rt *runtime.Runtime, bookings []model.Booking) ([]BookingResp, error) {
	rv := make([]BookingResp, len(bookings))
	for i, b := range bookings {
		parking, err := service.GetParkingByUid(rt, b.ParkingId)
		if err != nil {
			return nil, err
		}
		rv[i] = c.serializeBooking(b, parking)
	}
	return rv, nil
}

func (c CalParking) serializeBooking(booking model.Booking, parking model.Parking) BookingResp {
	return BookingResp{
		Uid:        booking.Uid.String(),
		Start:      booking.Start.String(),
		End:        booking.End.String(),
		Count:      booking.Count,
		ParkingUid: parking.Uid.String(),
		Address:    parking.Address,
		City:       parking.City,
	}
}
