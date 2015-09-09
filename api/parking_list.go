package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"bitbucket.org/remeh/parking/db/model"
	. "bitbucket.org/remeh/parking/logger"
	"bitbucket.org/remeh/parking/runtime"
	"bitbucket.org/remeh/parking/service"

	"github.com/gorilla/mux"
)

type ListParking struct {
	Runtime *runtime.Runtime
}

type listParkingEntry struct {
	Uid         string  `json:"uid"`
	Address     string  `json"address"`
	Zip         string  `json:"zip"`
	City        string  `json:"city"`
	Description string  `json"description"`
	Latitude    float64 `json"latitude"`
	Longitude   float64 `json"longitude"`
	Price       string  `json"price"`
}

const (
	// Offset on X around the selected pos when looking for a parking (unit: meters)
	OFFSET_X = 250
	// Offset on Y aroudn the selected pos when looking for a parking (unit: meters)
	OFFSET_Y = 250
)

func (c ListParking) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// compute a box within which we want to find some parkings
	vars := mux.Vars(r)
	pLat, pLon := vars["lat"], vars["lon"]
	if len(pLat) == 0 || len(pLon) == 0 {
		w.WriteHeader(400)
		return
	}

	lon, err := strconv.ParseFloat(pLon, 64)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	lat, err := strconv.ParseFloat(pLat, 64)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	// compute an area with the given POI
	parkings, err := service.GetParkingsInSurroundingArea(c.Runtime, lat, lon, OFFSET_X, OFFSET_Y)
	if err != nil {
		Error(err.Error())
		w.WriteHeader(500)
		return
	}

	// TODO(remy): marshaling the response
	data, err := json.Marshal(c.buildEntries(parkings))
	if err != nil {
		Error(err.Error())
		w.WriteHeader(500)
		return
	}

	w.Write(data)
}

func (c ListParking) buildEntries(parkings []model.Parking) []listParkingEntry {
	serialized := make([]listParkingEntry, len(parkings))
	i := 0
	for _, parking := range parkings {
		serialized[i] = c.buildEntry(parking)
		i++
	}
	return serialized
}

func (c ListParking) buildEntry(parking model.Parking) listParkingEntry {
	return listParkingEntry{
		Uid:         parking.Uid.String(),
		Address:     parking.Address,
		Description: parking.Description,
		Latitude:    parking.Latitude,
		Longitude:   parking.Longitude,
		Zip:         parking.Zip,
		City:        parking.City,
		Price:       parking.DailyPrice,
	}
}
