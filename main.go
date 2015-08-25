// Parking Backend
//
// Main.
//
// 2015

package main

import (
	. "bitbucket.org/parking/logger"
	"bitbucket.org/parking/runtime"

	"github.com/vrischmann/envconfig"
)

func main() {
	Info("Starting the server.")
	config, err := readConfig()
	if err != nil {
		Error(err.Error())
		return
	}

	Info("Listening on", config.ListenAddr)

	server := runtime.Server{
		Config: config,
	}

	server.Start()
	return
}

// readConfig reads in the environment var
// the configuration to start the runtime.
func readConfig() (runtime.Config, error) {
	config := runtime.Config{}
	err := envconfig.Init(&config)
	return config, err
}
