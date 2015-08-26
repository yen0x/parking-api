package api

import (
	"net/http"

	"bitbucket.org/remeh/parking/runtime"
)

type Example struct {
	Runtime *runtime.Runtime
}

func (c Example) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("example route"))
}
