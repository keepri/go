package handlers

import "net/http"

func Healthz(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("cool"))
}
