// Parking Backend - Runtime
//
// Logger helpers.
//
// 2015

package runtime

import (
	"log"
)

func Error(data ...interface{}) {
	log.Println("Error", data)
}

func Warning(data ...interface{}) {
	log.Println("Warning", data)
}

func Info(data ...interface{}) {
	log.Println("Info", data)
}
