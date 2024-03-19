package main

import (
	"github.com/luquxSentinel/spacedrive/service"
	"log"
)

func main() {
	listenAddress := ":3000"
	

	// TODO: new auth service
	authService := service.NewAuthService()
	//	new mux api server
	server := NewAPIServer(listenAddress, authService)
	
	if err :=  server.Run(); err != nil {
		log.Fatal(err)
	}
}