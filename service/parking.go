// Parking Backend - Service
//
// All methods around parking features
//
// 2015

package service

import (
	"bitbucket.org/remeh/parking/db/model"
	"bitbucket.org/remeh/parking/runtime"
	"github.com/pborman/uuid"
	"time"
)

func CreateParking(rt *runtime.Runtime, address, description, price, user, zip, city string, latitude, longitude float64) (uuid.UUID, error) {
	if rt == nil {
		return []byte{}, nil
	}

	pDAO := rt.Storage.ParkingDAO
	now := time.Now()
	uid := uuid.Parse(uuid.New())

	parking := model.Parking{
		Uid:          uid,
		UserId:       uuid.Parse(user),
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
