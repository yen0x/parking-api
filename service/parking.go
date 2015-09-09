// Parking Backend - Service
//
// All methods around parking features
//
// 2015

package service

import (
	"math"
	"time"

	"bitbucket.org/remeh/parking/db/model"
	"bitbucket.org/remeh/parking/runtime"

	"github.com/pborman/uuid"
)

func CreateParking(rt *runtime.Runtime, user model.User, address, description, price, zip, city string, latitude, longitude float64) (uuid.UUID, error) {
	if rt == nil {
		return []byte{}, nil
	}

	pDAO := rt.Storage.ParkingDAO
	now := time.Now()
	uid := uuid.Parse(uuid.New())

	parking := model.Parking{
		Uid:          uid,
		UserId:       user.Uid,
		Description:  description,
		Address:      address,
		Zip:          zip,
		City:         city,
		Latitude:     latitude,
		Longitude:    longitude,
		DailyPrice:   price,
		CreationTime: now,
		LastUpdate:   now,
	}

	_, err := pDAO.Insert(parking)
	return uid, err
}

// GetParkings returns the many parkings of the given user.
func GetParkings(rt *runtime.Runtime, user model.User) ([]model.Parking, error) {
	pDAO := rt.Storage.ParkingDAO
	parkings, err := pDAO.FindByUser(user)
	return parkings, err
}

// GetParkingsInSurroundingArea returns the parkings find in the area around
// the given lat-lon point.
func GetParkingsInSurroundingArea(rt *runtime.Runtime, lat, lon, width, height float64) ([]model.Parking, error) {
	pDAO := rt.Storage.ParkingDAO

	// compute an area with the given POI
	topLeftLat, topLeftLon, bottomRightLat, bottomRightLon := computeArea(lat, lon, width, height)

	return pDAO.FindInArea(topLeftLat, topLeftLon, bottomRightLat, bottomRightLon)
}

// computeArea takes the given POI on the map and computes a square
// around this point of width/height.
//
// width and height parameters are in meters such as the final result is
//
// <--- width (m) --->
// A-----------------x ^              A (first, second returned values)
// x                 x |              B (third, fourth returned values)
// x    (lat/lon)    x |
// x        o        x height (m)
// x                 x |
// x                 x |
// x-----------------B v
//
// NOTE(remy): these are approximative computating, see :
// http://gis.stackexchange.com/questions/2951/algorithm-for-offsetting-a-latitude-longitude-by-some-amount-of-meters
func computeArea(lat, lon, width, height float64) (float64, float64, float64, float64) {
	var earthRadius float64 = 6378137

	// offsets
	offLat := width / earthRadius
	offLon := height / (earthRadius * math.Cos(math.Pi*lat/float64(180)))

	return lat - offLat, lon - offLon, lat + offLat, lon + offLat
}
