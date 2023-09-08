package server

import (
	"log"

	"github.com/keepri/go/internal/config"
)

func Init() {
	server_cfg := config.NewServerConfig()
	server := server_cfg.NewServer()
	log.Printf("server starting on port %v...", server_cfg.Port)
	log.Fatalln(server.ListenAndServe())
}
