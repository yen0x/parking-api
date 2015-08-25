// Parking Backend - Runtime
//
// Runtime configuration.
//
// 2015

package runtime

type Config struct {
	// Address to listen to.
	ListenAddr string `envconfig:"ADDR,default=:8080"`
}
