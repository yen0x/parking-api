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

func CreateParking(rt *runtime.Runtime, user model.User, address, description, zip, city, currency string, price int, latitude, longitude float64) (uuid.UUID, error) {
	if rt == nil {
		return []byte{}, nil
	}

	// inserts the parking

	pDAO := rt.Storage.ParkingDAO
	aDAO := rt.Storage.AvailabilityDAO

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
		Currency:     currency,
		CreationTime: now,
		LastUpdate:   now,
	}

	_, err := pDAO.Insert(parking)
	if err != nil {
		return []byte(""), err
	}

	// then, atm, creates a global availability

	past, future := timeLimits()
	avail := model.Availability{
		ParkingUid: uid,
		Start:      past,
		End:        future,
	}
	_, err = aDAO.Insert(avail)

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
func GetParkingsInSurroundingArea(rt *runtime.Runtime, neLat, neLon, swLat, swLon float64, start, end time.Time) ([]model.Parking, error) {
	pDAO := rt.Storage.ParkingDAO
	return pDAO.FindInArea(neLat, neLon, swLat, swLon, start, end)
}

// ParkingExistsForUser returns whether or not a parking exists for the user
func ParkingExistsForUser(rt *runtime.Runtime, parkingUid uuid.UUID, user model.User) (bool, error) {
	pDAO := rt.Storage.ParkingDAO
	parking, err := pDAO.FindByUid(parkingUid)
	return parking.UserId.String() == user.Uid.String(), err
}

// ParkingExists returns whether or not a parking exists
func ParkingExists(rt *runtime.Runtime, parkingUid uuid.UUID) (bool, error) {
	pDAO := rt.Storage.ParkingDAO
	parking, err := pDAO.FindByUid(parkingUid)
	return len(parking.UserId) > 0, err
}

// Get Parking based on uuid
func GetParking(rt *runtime.Runtime, parkingUid uuid.UUID) (model.Parking, error) {
	pDAO := rt.Storage.ParkingDAO
	parking, err := pDAO.FindByUid(parkingUid)
	return parking, err
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

	// reconvert
	dLat := offLat * 180 / math.Pi
	dLon := offLon * 180 / math.Pi

	return lat + dLat, lon - dLon, lat - dLat, lon + dLat
}

func timeLimits() (time.Time, time.Time) {
	past, _ := time.Parse(DATE_FORMAT, "1970-01-01")
	future, _ := time.Parse(DATE_FORMAT, "2200-01-01")
	return past, future
}
