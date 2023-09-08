package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/keepri/go/internal/api/handlers"
)

func ApiRoute(r chi.Router) {
	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", handlers.Login)
		r.Post("/logout", handlers.Logout)
		r.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
			w.Write([]byte("Hello, auth!"))
		})
	})
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, api!"))
	})
}
