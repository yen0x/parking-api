package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	. "bitbucket.org/remeh/parking/logger"
	"bitbucket.org/remeh/parking/runtime"
	"bitbucket.org/remeh/parking/service"

	"github.com/pborman/uuid"
)

type CreateBooking struct {
	Runtime *runtime.Runtime
}

type CreateBookingBody struct {
	Parking string `json"parking"`
	Start   string `json"start"`
	End     string `json"end"`
	Count   int    `json"count"`
}

type CreateBookingResp struct {
	Uid string `json:"uid"`
}

func (c CreateBooking) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		Error(err)
		w.WriteHeader(500)
		return
	}

	session, exists := c.Runtime.SessionStorage.GetFromRequest(r)
	if !exists {
		w.WriteHeader(403)
		return
	}

	body := CreateBookingBody{}
	json.Unmarshal(data, &body)

	// Checks that the form has been correctly filled.

	start, err := time.Parse(service.DATE_FORMAT, body.Start)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	end, err := time.Parse(service.DATE_FORMAT, body.End)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	if len(body.Parking) == 0 ||
		body.Count <= 0 {
		w.WriteHeader(400)
		return

	}

	// checks that the parking actually exists.
	parkingUid := uuid.Parse(body.Parking)
	parkingExists, err := service.ParkingExists(c.Runtime, parkingUid)

	if err != nil {
		Error(err)
		w.WriteHeader(500)
		return
	} else if !parkingExists {
		w.WriteHeader(400)
		return
	}

	uuid, err := service.CreateBooking(c.Runtime, session.User, start, end, parkingUid, body.Count)
	if err != nil {
		Error(err)
		w.WriteHeader(500)
		return
	}

	resp := CreateBookingResp{
		Uid: uuid.String(),
	}
	data, err = json.Marshal(resp)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
