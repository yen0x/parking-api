package api

import (
	. "bitbucket.org/remeh/parking/logger"
	"bitbucket.org/remeh/parking/runtime"
	"bitbucket.org/remeh/parking/service"

	"errors"
	"github.com/gorilla/mux"
	"github.com/pborman/uuid"
	"net/http"
)

type DeleteParking struct {
	Runtime *runtime.Runtime
}

//TODO jean: check there aint bookings on the parking

func (c DeleteParking) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["uid"]
	session, exists := c.Runtime.SessionStorage.GetFromRequest(r)
	if !exists {
		w.WriteHeader(403)
		return
	}

	parking, err := service.GetParkingByUid(c.Runtime, uuid.Parse(uid))
	if err != nil {
		Error(err)
		w.WriteHeader(500)
		return
	}
	if parking.Uid == nil {
		Error(errors.New("Parking does not exist"))
		w.WriteHeader(400)
		w.Write(([]byte)("Unknown parking"))
		return
	} else if !uuid.Equal(session.User.Uid, parking.UserId) {
		w.WriteHeader(403)
		w.Write(([]byte)("Unauthorized operation"))
		return
	}

	_, err = service.DeleteParking(c.Runtime, parking)
	if err != nil {
		Error(err)
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
	w.Write(([]byte)("Deleted parking " + uid))
}
