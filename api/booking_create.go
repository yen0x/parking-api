package api

import (
	. "bitbucket.org/remeh/parking/logger"
	"bitbucket.org/remeh/parking/runtime"
	"bitbucket.org/remeh/parking/service"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CreateBooking struct {
	Runtime *runtime.Runtime
}

type CreateBookingBody struct {
	Parking string `json"parking"` //TODO get parking from db
	Start   string `json"start"`
	End     string `json"end"`
	Count   int    `json"count"`
	Email   string `json"email"` // FIXME(remy): won't be useful as soon as we have a session token
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
	body := CreateBookingBody{}
	json.Unmarshal(data, &body)

	fmt.Println(body)

	// FIXME(remy): won't be useful as soon as we have a session token
	user, err := service.GetUser(c.Runtime, body.Email)

	if err != nil {
		Error(err)
		w.WriteHeader(500)
		return
	}

	// Unknown user.

	if len(user.Uid) == 0 {
		w.WriteHeader(400)
		return
	}

	// Checks that the form has been correctly filled.

	if len(body.Start) == 0 ||
		len(body.End) == 0 ||
		len(body.Parking) == 0 ||
		body.Count <= 0 {
		w.WriteHeader(400)
		return

	}

	uuid, err := service.CreateBooking(c.Runtime, user, body.Start, body.End, body.Parking, body.Count)
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
