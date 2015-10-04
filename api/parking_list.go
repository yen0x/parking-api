package api

import (
	"bitbucket.org/remeh/parking/db/model"
	. "bitbucket.org/remeh/parking/logger"
	"bitbucket.org/remeh/parking/runtime"
	"bitbucket.org/remeh/parking/service"

	"encoding/json"
	"net/http"
)

type ListParking struct {
	Runtime *runtime.Runtime
}

type ParkingResp struct {
	Uid     string
	Address string
	City    string
}

func (c ListParking) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	session, exists := c.Runtime.SessionStorage.GetFromRequest(r)
	if !exists {
		w.WriteHeader(403)
		return
	}

	parkings, err := service.GetParkings(c.Runtime, session.User)
	if err != nil {
		Error(err)
		w.WriteHeader(500)
		return
	}

	data, err := json.Marshal(c.serializeParkings(parkings))
	if err != nil {
		Error(err)
		w.WriteHeader(500)
		return
	}

	w.Write(data)
}

func (c ListParking) serializeParkings(parkings []model.Parking) []ParkingResp {
	rv := make([]ParkingResp, len(parkings))
	for i, p := range parkings {
		rv[i] = c.serializeParking(p)
	}
	return rv
}

func (c ListParking) serializeParking(parking model.Parking) ParkingResp {
	return ParkingResp{
		Uid:     parking.Uid.String(),
		Address: parking.Address,
		City:    parking.City,
	}
}
