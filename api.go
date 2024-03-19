package main


import (
	"net/http"
)

type APIServer struct {
	mux *http.ServeMux
	listenAddress string
}

func NewAPIServer(listenAddress string) *APIServer {
	return &APIServer{
		listenAddress: listenAddress,
		mux: http.NewServeMux(),
	}
}

func (api *APIServer) Run() error {
	
	return http.ListenAndServe(api.listenAddress, api.mux)
}


