package handlers

import "net/http"

type HealthzHandler struct{}

func NewHealthzHandler() *HealthzHandler {
	return &HealthzHandler{}
}

func (hh *HealthzHandler) Healthz(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("cool"))
}
