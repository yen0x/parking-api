// Parking Backend
//
// Main.
//
// 2015

package main

import (
	"bitbucket.org/parking/api"
	. "bitbucket.org/parking/logger"
	"bitbucket.org/parking/runtime"
)

func main() {
	Info("Starting the runtime.")

	rt := runtime.NewRuntime()

	declareApiRoutes(rt)

	Info("Listening on", rt.Config.ListenAddr)

	err := rt.Start()
	if err != nil {
		Error(err.Error())
	}
}

func declareApiRoutes(rt *runtime.Runtime) {
	rt.AddApi("/example", api.Example{rt})
}
