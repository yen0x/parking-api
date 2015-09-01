package api

import (
	"bitbucket.org/remeh/parking/runtime"
	"io/ioutil"
	"net/http"
)

type CreateParking struct {
	Runtime *runtime.Runtime
}

type CreateParkingBody struct {
	Address     string `json"address"`
	Description string `json"description"`
	Price       int    `json"price"`
}

func (c CreateParking) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		Error(err)
		w.WriteHeader(500)
		return
	}
}
