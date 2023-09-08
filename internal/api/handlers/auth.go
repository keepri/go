package handlers

import (
	"net/http"

	"github.com/keepri/go/internal/utils/response"
)

func Login(w http.ResponseWriter, r *http.Request) {
	response.WithErr(w, "not implemented", 501)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	response.WithErr(w, "not implemented", 501)
}
