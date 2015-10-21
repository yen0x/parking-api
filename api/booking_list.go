package api

import (
	"bitbucket.org/remeh/parking/db/model"
	. "bitbucket.org/remeh/parking/logger"
	"bitbucket.org/remeh/parking/runtime"
	"bitbucket.org/remeh/parking/service"

	"encoding/json"
	"net/http"
)

type ListBooking struct {
	Runtime *runtime.Runtime
}

type BookingResp struct {
	Uid        string `json:"uid"`
	Start      string `json:"start"`
	End        string `json:"end"`
	Count      int    `json:"count"`
	ParkingUid string `json:"parkinguid"`
	Address    string `json:"address"`
	City       string `json:"city"`
}

func (c ListBooking) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	session, exists := c.Runtime.SessionStorage.GetFromRequest(r)
	if !exists {
		w.WriteHeader(403)
		return
	}

	bookings, err := service.GetBookings(c.Runtime, session.User)
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

func (c ListBooking) serializeBookings(rt *runtime.Runtime, bookings []model.Booking) ([]BookingResp, error) {
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

func (c ListBooking) serializeBooking(booking model.Booking, parking model.Parking) BookingResp {
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
