// Parking Backend - API
// Adapter to log a request on a route.
// 2015

package api

import (
	"fmt"
	"net/http"

	. "bitbucket.org/remeh/parking/logger"
	"bitbucket.org/remeh/parking/runtime"
)

type LogAdapter struct {
	Runtime *runtime.Runtime
	handler http.Handler
}

func (a LogAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// then propagate to the next handler if no 403 has been returned.
	sWriter := &StatusWriter{w, 200}
	a.handler.ServeHTTP(sWriter, r)

	Info(fmt.Sprintf("HIT - %s %s %s referer[%s] user-agent[%s] addr[%s] code[%d]", r.Method, r.URL.String(), r.Proto, r.Referer(), r.UserAgent(), r.RemoteAddr, sWriter.Status))

}

// LogRoute creates a route which will log the route access.
func LogRoute(rt *runtime.Runtime, handler http.Handler) http.Handler {
	return LogAdapter{
		Runtime: rt,
		handler: handler,
	}
}

type StatusWriter struct {
	http.ResponseWriter
	Status int
}

func (w *StatusWriter) WriteHeader(code int) {
	w.Status = code
	w.ResponseWriter.WriteHeader(code)
}
