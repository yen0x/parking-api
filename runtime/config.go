// Parking Backend - Runtime
//
// Runtime configuration.
//
// 2015

package runtime

type Config struct {
	// Address to listen to.
	ListenAddr string `envconfig:"ADDR,default=:8080"`
	// Public directory with pages and assets.
	PublicDir string `envconfig:"PUBLIC,default=public/"`
	// Connection string
	ConnString string `envconfig:"CONN,default=host=/var/run/postgresql sslmode=disable user=parking dbname=parking password=parking"`
}
