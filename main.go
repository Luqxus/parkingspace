package main

import (
	"log"

	"github.com/luquxSentinel/spacedrive/api"
	"github.com/luquxSentinel/spacedrive/service"
)

func main() {
	listenAddress := ":3000"

	// TODO: new auth service
	authService := service.NewAuthService()
	//	new mux api server
	server := api.New(listenAddress, authService)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
