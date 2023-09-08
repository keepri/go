package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/keepri/go/internal/config"
	"github.com/keepri/go/internal/server"
)

func main() {
	godotenv.Load()
	config := config.New().LoadFromEnv()

	server := server.Setup(config)
	log.Printf("server starting on port %v...", config.Port)
	log.Fatalln(server.ListenAndServe())
}
