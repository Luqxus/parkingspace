package main

import "log"

func main() {
	listenAddress := ":3000"
	
//	new mux api server
	server := NewAPIServer(listenAddress)
	
	if err :=  server.Run(); err != nil {
		log.Fatal(err)
	}
}