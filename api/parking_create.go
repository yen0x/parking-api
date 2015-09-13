package api

import (
	. "bitbucket.org/remeh/parking/logger"
	"bitbucket.org/remeh/parking/runtime"
	"bitbucket.org/remeh/parking/service"

	"encoding/json"
	"io/ioutil"
	"net/http"
)

type CreateParking struct {
	Runtime *runtime.Runtime
}

type CreateParkingBody struct {
	Address     string  `json:"address"`
	Zip         string  `json:"zip"`
	City        string  `json:"city"`
	Description string  `json:"description"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Price       int     `json:"price"`
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
	err = json.Unmarshal(data, &body)

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

	// Checks that the form has been correctly filled.

	if len(body.Address) == 0 ||
		len(body.Zip) == 0 ||
		len(body.City) == 0 ||
		len(body.Description) == 0 ||
		body.Price == 0 {
		w.WriteHeader(400)
		return
	}

	uuid, err := service.CreateParking(c.Runtime, session.User, body.Address, body.Description, body.Zip, body.City, "EUR", body.Price, body.Latitude, body.Longitude)
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
