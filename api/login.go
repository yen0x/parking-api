// Parking Backend - API
//
// Login route.
//
// 2015

package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	. "bitbucket.org/remeh/parking/logger"
	"bitbucket.org/remeh/parking/runtime"
	"bitbucket.org/remeh/parking/service"
)

type Login struct {
	Runtime *runtime.Runtime
}

type loginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResp struct {
	Uid       string `json:"uid"`
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Token     string `json:"token"`
}

func (c Login) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Reads the body.
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Error(err)
		w.WriteHeader(500)
		return
	}

	var body loginBody
	json.Unmarshal(data, &body)

	// parameters check.

	if len(body.Email) == 0 || len(body.Password) == 0 {
		w.WriteHeader(400)
		return
	}

	success, user, err := service.CheckUserPassword(c.Runtime, body.Email, body.Password)
	if err != nil {
		Error(err.Error())
		w.WriteHeader(500)
		return
	}

	// if login didn't succeed

	if !success {
		w.WriteHeader(403)
		return
	}

	// generates a new session for this user.

	session := c.Runtime.SessionStorage.New(user)

	// return some info to the browser

	resp := loginResp{
		Uid:       user.Uid.String(),
		Email:     user.Email,
		Firstname: user.Firstname,
	}

	// sets the cookie
	cookie := &http.Cookie{
		Name:  runtime.COOKIE_TOKEN_KEY,
		Value: session.Token,
	}
	http.SetCookie(w, cookie)

	data, err = json.Marshal(resp)
	if err != nil {
		Error(err.Error())
		w.WriteHeader(500)
		return
	}

	w.Write(data)
}
