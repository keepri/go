package handlers

import (
	"net/http"

	"github.com/keepri/go/internal/config"
	"github.com/keepri/go/internal/utils/response"
)

type AuthHandler struct {
	DBService *config.DBConfig
}

func NewAuthHandler(db_service *config.DBConfig) *AuthHandler {
	return &AuthHandler{
		DBService: db_service,
	}
}

func (ah *AuthHandler) Login(w http.ResponseWriter, _ *http.Request) {
	response.WithErr(w, "not implemented", 501)
}

func (ah *AuthHandler) Logout(w http.ResponseWriter, _ *http.Request) {
	response.WithErr(w, "not implemented", 501)
}
