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

func CreateParking(rt *runtime.Runtime, address, description, price string) (uuid.UUID, error) {
	if rt == nil {
		return []byte{}, nil
	}

	pDAO := rt.Storage.ParkingDAO
	now := time.Now()
	uid := uuid.Parse(uuid.New())

	parking := model.Parking{
		Uid:          uid,
		UserId:       uuid.Parse("1"),
		Description:  description,
		Address:      address,
		Latitude:     0.0,
		Longitude:    0.0,
		DailyPrice:   price,
		CreationTime: now,
		LastUpdate:   now,
	}

	_, err := pDAO.Insert(parking)
	return uid, err
}
