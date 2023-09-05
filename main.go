package main

import (
	"log"
	"net/http"
	"os"

	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	PORT := env("PORT", true)
	router := main_router(PORT)

	router.Route("/api", api_router)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request %s %s%s\n", r.Method, r.Host, r.URL.Path)
		w.Write([]byte("Hello, world!"))
	})

	server := http.Server{
		Handler: router,
		Addr:    ":" + PORT,
	}

	log.Printf("server starting on port %v", PORT)
	log.Fatalln(server.ListenAndServe())
}

func main_router(PORT string) (router *chi.Mux) {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://127.0.0.1:" + PORT},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	return r
}

func api_router(r chi.Router) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request %s %s%s\n", r.Method, r.Host, r.URL.Path)
		w.Write([]byte("Hello, world!"))
	})
}

func env(name string, is_required bool) string {
	env := os.Getenv(name)
	if env == "" && is_required {
		log.Fatalf("%v environment variable not set\n", name)
	}
	return env
}
