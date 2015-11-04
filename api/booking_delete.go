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

type DeleteBooking struct {
	Runtime *runtime.Runtime
}

func (c DeleteBooking) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["uid"]
	session, exists := c.Runtime.SessionStorage.GetFromRequest(r)
	if !exists {
		w.WriteHeader(403)
		return
	}

	booking, err := service.FindBookingByUid(c.Runtime, uuid.Parse(uid))
	if err != nil {
		Error(err)
		w.WriteHeader(500)
		return
	}
	if booking.Uid == nil {
		Error(errors.New("Booking does not exist"))
		w.WriteHeader(400)
		w.Write(([]byte)("Unknown booking"))
		return
	} else if !uuid.Equal(session.User.Uid, booking.UserId) {
		w.WriteHeader(403)
		w.Write(([]byte)("Unauthorized operation"))
		return
	}

	_, err = service.DeleteBooking(c.Runtime, booking)
	if err != nil {
		Error(err)
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
	w.Write(([]byte)("Deleted booking " + uid))
}
