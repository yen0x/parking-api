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
	User        string  `json"user"`
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
	//	latitude, err := strconv.ParseFloat(body.Latitude, 64)
	//	longitude, err := strconv.ParseFloat(body.Longitude, 64)

	uuid, err := service.CreateParking(c.Runtime, body.Address, body.Description, body.Price, body.User, body.Zip, body.City, body.Latitude, body.Longitude)
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
