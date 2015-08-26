// Parking Backend - Runtime
//
// Listening server.
//
// 2015

package runtime

import (
	"net/http"

	"bitbucket.org/remeh/parking/db"
	. "bitbucket.org/remeh/parking/logger"

	"github.com/gorilla/mux"
	"github.com/vrischmann/envconfig"
)

type Runtime struct {
	Config Config

	router *mux.Router

	Storage db.Storage
}

func NewRuntime() *Runtime {
	rt := &Runtime{}
	rt.router = mux.NewRouter()
	rt.readConfig()
	return rt
}

// Starts listening.
func (r *Runtime) Start() error {
	// Opens the database connection.
	Info("Opening the database connection.")
	_, err := r.Storage.Init(r.Config.ConnString)
	if err != nil {
		return err
	}

	// Prepares the router serving the static pages and assets.
	r.prepareStaticRouter()

	// Handles static routes
	http.Handle("/", r.router)

	// Starts listening.
	err = http.ListenAndServe(r.Config.ListenAddr, nil)
	return err
}

func (r *Runtime) AddApi(pattern string, handler http.Handler) {
	r.router.PathPrefix("/api").Subrouter().Handle(pattern, handler)
}

func (r *Runtime) prepareStaticRouter() {
	// Add the final route, the static assets and pages.
	r.router.PathPrefix("/").Handler(http.FileServer(http.Dir(r.Config.PublicDir)))
	Info("Serving static from directory", r.Config.PublicDir)
}

// readConfig reads in the environment var
// the configuration to start the runtime.
func (r *Runtime) readConfig() {
	err := envconfig.Init(&r.Config)
	if err != nil {
		Error("While reading the configuration:", err.Error())
	}
}
