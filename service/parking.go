// Parking Backend - Service
//
// All methods around parking features
//
// 2015

package service

import (
	"bitbucket.org/remeh/parking/runtime"
	"bitbucket.org/remeh/parking/model"
	"github.com/pborman/uuid"
	"time"
)

func CreateParking(rt *runtime.Runtime, address, description, price string, ) (uuid.UUID, error) {
	if rt == nil {
		return []byte{}, nil
	}

	pDAO := rt.Storage.ParkingDAO
	now := time.Now()

    parking := model.Parking{
		Uid:          uuid.Parse("0"),
		UserId:       uuid.Parse("1"),
		Description:  description,
		Address:      address,
		Latitude:     0,
		Longitude:    0,
		DailyPrice:   price,
		CreationTime: now,
		LastUpdate:   now,
	}

	pDAO.Insert()


}
