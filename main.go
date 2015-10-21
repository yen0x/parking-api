// Parking Backend
//
// Main
//
// 2015

package main

import (
	. "bitbucket.org/remeh/parking/api"
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
	rt.AddApi("/example", LogRoute(rt, Example{rt}))

	rt.AddApi("/user/create", LogRoute(rt, CreateUser{rt}))
	rt.AddApi("/login", LogRoute(rt, Login{rt}))

	rt.AddApi("/parking/create", AuthRoute(rt, CreateParking{rt}))
	rt.AddApi("/parking/list", AuthRoute(rt, ListParking{rt}))
	rt.AddApi("/parking/search/area/{nelat},{nelon}/{swlat},{swlon}/{start}/{end}", LogRoute(rt, SearchParking{rt}))

	rt.AddApi("/booking/create", CreateBooking{rt})
	rt.AddApi("/booking/list", AuthRoute(rt, ListBooking{rt}))
}
