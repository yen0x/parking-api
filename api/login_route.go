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

	// return some info to the browser

	// TODO(remy): generate a token and store them either in base or ram

	resp := loginResp{
		Uid:       user.Uid.String(),
		Email:     user.Email,
		Firstname: user.Firstname,
	}

	data, err = json.Marshal(resp)
	if err != nil {
		Error(err.Error())
		w.WriteHeader(500)
		return
	}

	w.Write(data)
}
