package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"./dbops"
	"./defs"
	"./session"
)

// CreateUser http handler
func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)

	ubody := &defs.UserCredential{}
	// json -> UserCredential
	if err := json.Unmarshal(res, ubody); err != nil {
		sendErrorRessponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	if err := dbops.AddUserCredential(ubody.Username, ubody.Pwd); err != nil {
		sendErrorRessponse(w, defs.ErrorDBError)
		return
	}

	id := session.GenerateNewSessionId(ubody.Username)
	su := &defs.SignedUp{Success: true, SessionId: id}

	if resp, err := json.Marshal(su); err != nil {
		sendErrorRessponse(w, defs.ErrorInternalFaults)
		return
	}
	sendNormalRessponse(w, string(resp), 201)

}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userName := p.ByName("user_name")

	io.WriteString(w, userName)

}
