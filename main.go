package main

import (
	"github.com/joho/godotenv"
	"github.com/keepri/go/internal/server"
)

type Test struct {
	Hello string
}

func main() {
	godotenv.Load()
	server.Init()
}
