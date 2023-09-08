package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/keepri/go/internal/api"
	"github.com/keepri/go/internal/api/handlers"
	"github.com/keepri/go/internal/config"
)

func Setup(c *config.Config) *http.Server {
	port := fmt.Sprint(c.Port)
	router := setupRouter(c)
	router.Mount("/v1", setupV1())

	return &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}
}

func setupRouter(c *config.Config) chi.Router {
	port := fmt.Sprint(c.Port)
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://127.0.0.1:" + port},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	if c.Logger {
		router.Use(middleware.Logger)
	}

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ow, hi Mark!"))
	})

	staticDir := http.Dir("static")
	staticFileServer := http.FileServer(staticDir)
	http.Handle("/static/", http.StripPrefix("/static/", staticFileServer))

	return router
}

func setupV1() chi.Router {
	v1 := chi.NewRouter()

	v1.Route("/api", api.Route)

	v1.Get("/healthz", handlers.Healthz)
	v1.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	return v1
}
