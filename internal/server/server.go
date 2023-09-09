package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/keepri/go/internal/config"
	"github.com/keepri/go/internal/handlers"
)

func Init() {
	router := chi.NewRouter()
	v1 := chi.NewRouter()
	cfg := config.NewServerConfig()
	db_service := config.NewDBConfig()
	user_handler := handlers.NewUserHandler(db_service)
	auth_handler := handlers.NewAuthHandler(db_service)
	healthz_handler := handlers.NewHealthzHandler()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{fmt.Sprintf("http://127.0.0.1:%d", cfg.Port)},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	if cfg.Logger {
		router.Use(middleware.Logger)
	}

	// /v1 routes
	v1.Route("/", func(r chi.Router) {

		r.Get("/healthz", healthz_handler.Healthz)
		r.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
			w.Write([]byte("Hello, v1!"))
		})

		// /v1/api routes
		r.Route("/api", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, _ *http.Request) {
				w.Write([]byte("Hello, api!"))
			})

			// /v1/api/users routes
			r.Route("/users", func(r chi.Router) {
				r.Get("/", user_handler.User)
			})

			// /v1/api/auth routes
			r.Route("/auth", func(r chi.Router) {
				r.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
					w.Write([]byte("Hello, auth!"))
				})
				r.Post("/login", auth_handler.Login)
				r.Post("/logout", auth_handler.Logout)
			})
		})
	})

	router.Mount("/v1", v1)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ow, hi Mark!"))
	})

	staticDir := http.Dir("static")
	staticFileServer := http.FileServer(staticDir)
	router.Handle("/static/*", http.StripPrefix("/static/", staticFileServer))

	server := cfg.NewServer(router)
	log.Printf("server starting on port %v...", cfg.Port)
	log.Fatalln(server.ListenAndServe())
}
