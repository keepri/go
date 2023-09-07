package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/keepri/rss/internal/config"
	"github.com/keepri/rss/internal/server"
)

func main() {
	godotenv.Load()
	config := config.New().LoadFromEnv()

	server := server.Setup(config)
	log.Printf("server starting on port %v...", config.Port)
	log.Fatalln(server.ListenAndServe())
}
