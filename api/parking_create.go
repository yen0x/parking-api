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

type CreateParking struct {
	Runtime *runtime.Runtime
}

type CreateParkingBody struct {
	Address     string  `json"address"`
	Zip         string  `json"zip"`
	City        string  `json"city"`
	Description string  `json"description"`
	Latitude    float64 `json"latitude"`
	Longitude   float64 `json"longitude"`
	Price       string  `json"price"`
	Email       string  `json"email"` // FIXME(remy): won't be useful as soon as we have a session token
}

type CreateParkingResp struct {
	Uid string `json:"uid"`
}

func (c CreateParking) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		Error(err)
		w.WriteHeader(500)
		return
	}
	body := CreateParkingBody{}
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

	if len(body.Address) == 0 ||
		len(body.Zip) == 0 ||
		len(body.City) == 0 ||
		len(body.Description) == 0 ||
		len(body.Price) == 0 {
		w.WriteHeader(400)
		return

	}

	uuid, err := service.CreateParking(c.Runtime, user, body.Address, body.Description, body.Price, body.Zip, body.City, body.Latitude, body.Longitude)
	if err != nil {
		Error(err)
		w.WriteHeader(500)
		return
	}

	resp := CreateParkingResp{
		Uid: uuid.String(),
	}
	data, err = json.Marshal(resp)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
