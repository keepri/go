package main

import (
	"github.com/joho/godotenv"
	"github.com/keepri/go/internal/server"
)

func main() {
	godotenv.Load()
	server.Init()
}
