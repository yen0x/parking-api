// Parking Backend - API
// Auth check for a logged in user.
// Note that this auth adapter call the log adapter
// on successfull token.
// 2015

package api

import (
	"net/http"

	"bitbucket.org/remeh/parking/runtime"
)

type AuthAdapter struct {
	Runtime *runtime.Runtime
	handler http.Handler
}

const (
	COOKIE_TOKEN_KEY = "t"
)

func (a AuthAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// gets the cookie in the request

	cookie, err := r.Cookie(COOKIE_TOKEN_KEY)
	if err != nil {
		w.WriteHeader(403)
		return
	}

	// ensure that this session exists

	session := a.Runtime.SessionStorage.Get(cookie.Value)
	if len(session.User.Uid) == 0 {
		w.WriteHeader(403)
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
