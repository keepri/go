package config

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/keepri/go/internal/api"
	"github.com/keepri/go/internal/api/handlers"
	"github.com/keepri/go/internal/utils/env"
)

type ServerConfig struct {
	Port   int
	Debug  bool
	Logger bool
}

func NewServerConfig() *ServerConfig {
	return &ServerConfig{
		Port:   env.Get("port", true).Atoi(),
		Debug:  len(env.Get("debug", false).Val) > 0,
		Logger: len(env.Get("logger", false).Val) > 0,
	}
}

func (c *ServerConfig) NewServer() *http.Server {
	return &http.Server{
		Handler: c.MainRouter(),
		Addr:    fmt.Sprintf(":%d", c.Port),
	}
}

func (c *ServerConfig) MainRouter() chi.Router {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{fmt.Sprintf("http://127.0.0.1:%d", c.Port)},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	if c.Logger {
		router.Use(middleware.Logger)
	}

	router.Mount("/v1", c.V1Router())
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ow, hi Mark!"))
	})

	staticDir := http.Dir("static")
	staticFileServer := http.FileServer(staticDir)
	http.Handle("/static/", http.StripPrefix("/static/", staticFileServer))

	return router
}

func (c *ServerConfig) V1Router() chi.Router {
	v1 := chi.NewRouter()

	v1.Route("/api", api.ApiRoute)

	v1.Get("/healthz", handlers.Healthz)
	v1.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	return v1
}
