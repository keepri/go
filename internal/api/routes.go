package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/keepri/go/internal/api/handlers"
)

func Route(r chi.Router) {
	r.Route("/auth", authRoute)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, api!"))
	})
}

func authRoute(r chi.Router) {
	r.Post("/login", handlers.Login)
	r.Post("/logout", handlers.Logout)
	r.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("Hello, auth!"))
	})
}
