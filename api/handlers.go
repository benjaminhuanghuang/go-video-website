package main

import (
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// CreateUser http handler
func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "Create User ")
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userName := p.ByName("user_name")

	io.WriteString(w, userName)

}
