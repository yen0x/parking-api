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

type Runtime struct {
	Config Config
	// TODO(remy): storage connection.
}

// Starts listening.
func (s Runtime) Start() error {
	// Prepares the router.
	s.prepareAPIRouter()

	s.prepareStaticRouter()

	// Starts listening.
	err := http.ListenAndServe(s.Config.ListenAddr, nil)
	return err
}

// prepareRouter creates the router to use
// to answer http requests.
func (s Runtime) prepareAPIRouter() {
	router := mux.NewRouter()

	http.Handle("/api", router)
}

func (s Runtime) prepareStaticRouter() {
	// Add the final route, the static assets and pages.
	router := mux.NewRouter()
	router.PathPrefix("/").Handler(http.FileServer(http.Dir(s.Config.PublicDir)))
	http.Handle("/", router)
	Info("Serving static from directory", s.Config.PublicDir)
}
