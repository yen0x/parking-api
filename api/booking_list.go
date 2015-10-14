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

	data, err := json.Marshal(c.serializeBookings(bookings))
	if err != nil {
		Error(err)
		w.WriteHeader(500)
		return
	}

	w.Write(data)
}

func (c ListBooking) serializeBookings(bookings []model.Booking) []BookingResp {
	rv := make([]BookingResp, len(bookings))
	for i, p := range bookings {
		rv[i] = c.serializeBooking(p)
	}
	return rv
}

func (c ListBooking) serializeBooking(booking model.Booking) BookingResp {
	return BookingResp{
		Uid:   booking.Uid.String(),
		Start: booking.Start.String(),
		End:   booking.End.String(),
		Count: booking.Count,
	}
}
