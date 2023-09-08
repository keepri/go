package handlers

import (
	"net/http"

	"github.com/keepri/go/internal/auth"
	"github.com/keepri/go/internal/config"
	"github.com/keepri/go/internal/models"
	"github.com/keepri/go/internal/utils/response"
)

type UserHandler struct {
	DBService *config.DBConfig
}

func NewUserHandler(db_service *config.DBConfig) *UserHandler {
	return &UserHandler{
		DBService: db_service,
	}
}

func (ah *UserHandler) User(w http.ResponseWriter, r *http.Request) {
	api_key, err := auth.GetApiKey(&r.Header)
	if err != nil {
		response.WithErr(w, err.Error(), 400)
	}

	user, err := ah.DBService.DB.UserByApiKey(r.Context(), api_key)
	if err != nil {
		response.WithErr(w, err.Error(), 404)
	}

	response.WithJson(w, models.UserModel(&user), 200)
}
