// Parking Backend - API
// Auth check for a logged in user.
// Note that this auth adapter call the log adapter
// on successfull token.
// 2015

package api

import (
	"fmt"
	"net/http"

	"bitbucket.org/remeh/parking/runtime"
)

type AuthAdapter struct {
	Runtime *runtime.Runtime
	handler http.Handler
}

func (a AuthAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// gets the cookie in the request

	cookie, err := r.Cookie(runtime.COOKIE_TOKEN_KEY)
	if err != nil {
		w.WriteHeader(403)
		fmt.Println("could not read cookie")
		return
	}

	// ensure that this session exists

	_, exists := a.Runtime.SessionStorage.Get(cookie.Value)
	if !exists {
		w.WriteHeader(403)
		fmt.Println("could not find session")
		return
	}

	// ok

	LogRoute(a.Runtime, a.handler).ServeHTTP(w, r)
}

func AuthRoute(rt *runtime.Runtime, handler http.Handler) http.Handler {
	return AuthAdapter{
		Runtime: rt,
		handler: handler,
	}
}
