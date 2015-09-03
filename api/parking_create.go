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
	Address     string `json"address"`
	Description string `json"description"`
	Price       string `json"price"`
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
	uuid, err := service.CreateParking(c.Runtime, body.Address, body.Description, body.Price)
	if err != nil {
		Error(err)
		w.WriteHeader(500)
		return
	}

	w.Write(uuid)
}
