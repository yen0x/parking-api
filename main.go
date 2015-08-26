// Parking Backend
//
// Main
//
// 2015

package main

import (
	"bitbucket.org/remeh/parking/api"
	. "bitbucket.org/remeh/parking/logger"
	"bitbucket.org/remeh/parking/runtime"
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

	rt.AddApi("/user/create", api.CreateUser{rt})
}
