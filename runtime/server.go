// Parking Backend - Runtime
//
// Listening server.
//
// 2015

package runtime

import (
	"net/http"

	. "bitbucket.org/parking/logger"

	"github.com/gorilla/mux"
)

type Server struct {
	Config Config
}

// Starts listening.
func (s Server) Start() {
	// Prepares the router.
	router := prepareRouter()
	http.Handle("/", router)

	// Starts listening.
	err := http.ListenAndServe(s.Config.ListenAddr, nil)
	if err != nil {
		Error(err.Error())
	}
}

// prepareRouter creates the router to use
// to answer http requests.
func prepareRouter() *mux.Router {
	r := mux.NewRouter()

	return r
}
