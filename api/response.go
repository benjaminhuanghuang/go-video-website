package main

import (
	"encoding/json"
	"io"
	"net/http"

	"./defs"
)

func sendErrorRessponse(w http.ResponseWriter, errResp defs.ErrorResponse) {
	w.WriteHeader(errResp.HttpSC)

	resStr, _ := json.Marshal(&errResp.Error) // to json string
	io.WriteString(w, string(resStr))
}

func sendNormalRessponse(w http.ResponseWriter, resp string, sc int) {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}
