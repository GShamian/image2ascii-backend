package main

import (
	"log"

	"github.com/GShamian/image2ascii-backend"
	"github.com/GShamian/image2ascii-backend/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(image2ascii.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
