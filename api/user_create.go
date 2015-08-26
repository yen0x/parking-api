// Parking Backend - API
//
// Controller to create an user.
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

type CreateUser struct {
	Runtime *runtime.Runtime
}

type CreateUserBody struct {
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type CreateUserResp struct {
	Uid       string `json:"uid"`
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func (c CreateUser) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Reads the body.

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Error(err)
		w.WriteHeader(500)
		return
	}

	body := CreateUserBody{}
	json.Unmarshal(data, &body)

	// Parameters check.

	if len(body.Email) == 0 || len(body.Firstname) == 0 || len(body.Lastname) == 0 {
		w.WriteHeader(400)
		return
	}

	// Checks that the email isn't already used
	existing, err := service.UserExists(c.Runtime, body.Email)
	if err != nil {
		Error(err)
		w.WriteHeader(500)
		return
	}

	if existing {
		w.WriteHeader(409)
		return
	}

	// Creates and stores the user.

	uuid, err := service.CreateUser(c.Runtime, body.Email, body.Firstname, body.Lastname)
	if err != nil {
		Error(err)
		w.WriteHeader(500)
		return
	}

	// Sends the response.

	resp := CreateUserResp{
		Uid:       uuid.String(),
		Email:     body.Email,
		Firstname: body.Firstname,
		Lastname:  body.Lastname,
	}
	data, err = json.Marshal(resp)
	if err != nil {
		Error(err)
		w.WriteHeader(500)
		return
	}

	w.Write(data)
}
