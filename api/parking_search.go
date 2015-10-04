package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"bitbucket.org/remeh/parking/db/model"
	. "bitbucket.org/remeh/parking/logger"
	"bitbucket.org/remeh/parking/runtime"
	"bitbucket.org/remeh/parking/service"

	"github.com/gorilla/mux"
)

type SearchParking struct {
	Runtime *runtime.Runtime
}

type listParkingEntry struct {
	Uid         string  `json:"uid"`
	Address     string  `json:"address"`
	Zip         string  `json:"zip"`
	City        string  `json:"city"`
	Description string  `json:"description"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Price       int     `json:"price"`
	Currency    string  `json:"currency"`
}

const (
	// Offset on X around the selected pos when looking for a parking (unit: meters)
	OFFSET_X = 500.0
	// Offset on Y aroudn the selected pos when looking for a parking (unit: meters)
	OFFSET_Y = 500.0
)

func (c SearchParking) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// compute a box within which we want to find some parkings
	vars := mux.Vars(r)
	pNELat, pNELon := vars["nelat"], vars["nelon"]
	pSWLat, pSWLon := vars["swlat"], vars["swlon"]
	pStart, pEnd := vars["start"], vars["end"]

	if len(pNELat) == 0 || len(pNELon) == 0 ||
		len(pSWLat) == 0 || len(pSWLon) == 0 {
		w.WriteHeader(400)
		return
	}

	// north eat

	neLon, err := strconv.ParseFloat(pNELon, 64)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	neLat, err := strconv.ParseFloat(pNELat, 64)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	swLon, err := strconv.ParseFloat(pSWLon, 64)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	swLat, err := strconv.ParseFloat(pSWLat, 64)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	start, err := c.parseDate(pStart)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	end, err := c.parseDate(pEnd)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	// compute an area with the given POI
	parkings, err := service.GetParkingsInSurroundingArea(c.Runtime, neLat, neLon, swLat, swLon, start, end)
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

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func (c SearchParking) parseDate(parameter string) (time.Time, error) {
	return time.Parse("2006-01-02", parameter)
}

func (c SearchParking) buildEntries(parkings []model.Parking) []listParkingEntry {
	serialized := make([]listParkingEntry, len(parkings))
	i := 0
	for _, parking := range parkings {
		serialized[i] = c.buildEntry(parking)
		i++
	}
	return serialized
}

func (c SearchParking) buildEntry(parking model.Parking) listParkingEntry {
	return listParkingEntry{
		Uid:         parking.Uid.String(),
		Address:     parking.Address,
		Description: parking.Description,
		Latitude:    parking.Latitude,
		Longitude:   parking.Longitude,
		Zip:         parking.Zip,
		City:        parking.City,
		Price:       parking.DailyPrice,
		Currency:    parking.Currency,
	}
}
